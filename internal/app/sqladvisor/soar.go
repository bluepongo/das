package sqladvisor

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/romberli/das/config"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/internal/dependency/sqladvisor"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/linux"
	"github.com/romberli/go-util/middleware/sql/parser"
	"github.com/spf13/viper"
)

const (
	logExp               = `^\d{4}/\d{2}/\d{2} \d{2}:\d{2}:\d{2}\.\d{3}`
	errExp               = `\[1;31m\[E]|\[1;31m\[F]`
	defaultLogMessageLen = 3
)

var _ sqladvisor.Advisor = (*DefaultAdvisor)(nil)

type DefaultAdvisor struct {
	parser     *parser.Parser
	soarBin    string
	configFile string
}

// NewDefaultAdvisor returns a new *DefaultAdvisor
func NewDefaultAdvisor(soarBin, configFile string) *DefaultAdvisor {
	return newDefaultAdvisor(soarBin, configFile)
}

// NewDefaultAdvisorWithDefault returns a new *DefaultAdvisor with default value
func NewDefaultAdvisorWithDefault() *DefaultAdvisor {
	soarBin := viper.GetString(config.SQLAdvisorSoarBinKey)
	configFile := viper.GetString(config.SQLAdvisorSoarConfigKey)

	return newDefaultAdvisor(soarBin, configFile)
}

// newDefaultAdvisor returns a new *DefaultAdvisor
func newDefaultAdvisor(soarBin, configFile string) *DefaultAdvisor {
	return &DefaultAdvisor{
		parser:     parser.NewParserWithDefault(),
		soarBin:    soarBin,
		configFile: configFile,
	}
}

// GetParser returns the parser
func (da *DefaultAdvisor) GetParser() *parser.Parser {
	return da.parser
}

// GetFingerprint returns the fingerprint of the sql text
func (da *DefaultAdvisor) GetFingerprint(sqlText string) string {
	return da.parser.GetFingerprint(sqlText)
}

// GetSQLID returns the identity of the sql text
func (da *DefaultAdvisor) GetSQLID(sqlText string) string {
	return da.parser.GetSQLID(sqlText)
}

// Advise parses the sql text and returns the tuning advice,
// note that only the first sql statement in the sql text will be advised
func (da *DefaultAdvisor) Advise(dbID int, sqlText string) (string, string, error) {
	return da.adviseWithDefault(dbID, sqlText)
}

// advise parses the sql text and returns the tuning advice,
// note that only the first sql statement in the sql text will be advised
func (da *DefaultAdvisor) adviseWithDefault(dbID int, sqlText string) (string, string, error) {
	user := viper.GetString(config.DBSoarMySQLUserKey)
	pass := viper.GetString(config.DBSoarMySQLPassKey)

	return da.advise(dbID, sqlText, user, pass)
}

// advise parses the sql text and returns the tuning advice,
// note that only the first sql statement in the sql text will be advised
func (da *DefaultAdvisor) advise(dbID int, sqlText, user, pass string) (string, string, error) {
	dsn, err := da.getOnlineDSN(dbID, user, pass)
	if err != nil {
		return constant.EmptyString, constant.EmptyString, err
	}

	command := fmt.Sprintf(`%s -config=%s -online-dsn=%s -query="%s"`, da.soarBin, da.configFile, dsn, sqlText)

	result, err := linux.ExecuteCommand(command)
	if err != nil {
		return constant.EmptyString, constant.EmptyString, err
	}

	return da.parseResult(result)
}

// getOnlineDSN returns the online dsn which will be used by soar
func (da *DefaultAdvisor) getOnlineDSNWithDefault(dbID int) (string, error) {
	user := viper.GetString(config.DBSoarMySQLUserKey)
	pass := viper.GetString(config.DBSoarMySQLPassKey)

	return da.getOnlineDSN(dbID, user, pass)
}

func (da *DefaultAdvisor) getOnlineDSN(dbID int, user, pass string) (string, error) {
	// get db service
	dbService := metadata.NewDBServiceWithDefault()
	err := dbService.GetByID(dbID)
	if err != nil {
		return constant.EmptyString, nil
	}
	// get db
	db := dbService.DBs[constant.ZeroInt]
	clusterID := db.GetClusterID()
	dbName := db.GetDBName()
	// get mysql server service
	mysqlServerService := metadata.NewMySQLServerServiceWithDefault()
	err = mysqlServerService.GetByClusterID(clusterID)
	if err != nil {
		return constant.EmptyString, err
	}

	mysqlServers := mysqlServerService.GetMySQLServers()
	if len(mysqlServers) == constant.ZeroInt {
		return constant.EmptyString, errors.New(fmt.Sprintf("could not find mysql server of the database. db id: %d", dbID))
	}
	// get mysql server
	mysqlServer := mysqlServerService.GetMySQLServers()[constant.ZeroInt]
	hostIP := mysqlServer.GetHostIP()
	portNum := mysqlServer.GetPortNum()

	return fmt.Sprintf("%s:%s@%s:%d/%s", user, pass, hostIP, portNum, dbName), nil
}

func (da *DefaultAdvisor) getDBSoarMySQLUser() string {
	return viper.GetString(config.DBSoarMySQLUserKey)
}

// parseResult parses result, it will split the advice information and the log information
func (da *DefaultAdvisor) parseResult(result string) (string, string, error) {
	var (
		advice  string
		message string
		errMsg  string
	)

	isLogMsg := true
	logExpression, err := regexp.Compile(logExp)
	if err != nil {
		return constant.EmptyString, constant.EmptyString, err
	}
	errExpression, err := regexp.Compile(errExp)
	if err != nil {
		return constant.EmptyString, constant.EmptyString, err
	}

	lines := strings.Split(result, constant.CRLFString)
	for _, line := range lines {
		if isLogMsg {
			isLogMsg = logExpression.Match([]byte(line))
		}

		if isLogMsg {
			message += line + constant.CRLFString
			stringList := strings.Split(line, constant.SpaceString)
			if len(stringList) >= defaultLogMessageLen {
				if errExpression.Match([]byte(stringList[2])) {
					errMsg += line + constant.CRLFString
				}
			}

			continue
		}

		advice += line
	}

	if errMsg != constant.EmptyString {
		return advice, message, errors.New(fmt.Sprintf("parse result failed. error:\n%s", errMsg))
	}

	return advice, message, nil
}
