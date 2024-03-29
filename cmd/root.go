/*
Copyright © 2020 Romber Li <romber2001@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pingcap/errors"
	"github.com/romberli/das/config"
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/go-util/middleware/mysql"
	"github.com/romberli/log"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultConfigFileType = "yaml"

var (
	// config
	baseDir string
	cfgFile string
	// daemon
	daemonStr string
	// log
	logFileName   string
	logLevel      string
	logFormat     string
	logMaxSize    int
	logMaxDays    int
	logMaxBackups int
	// server
	serverAddr         string
	serverPid          int
	serverPidFile      string
	serverReadTimeout  int
	serverWriteTimeout int
	// database
	dbDASMySQLAddr           string
	dbDASMySQLName           string
	dbDASMySQLUser           string
	dbDASMySQLPass           string
	dbPoolMaxConnections     int
	dbPoolInitConnections    int
	dbPoolMaxIdleConnections int
	dbPoolMaxIdleTime        int
	dbPoolKeepAliveInterval  int
	dbMonitorPrometheusUser  string
	dbMonitorPrometheusPass  string
	dbMonitorClickhouseUser  string
	dbMonitorClickhousePass  string
	dbMonitorMySQLUser       string
	dbMonitorMySQLPass       string
	dbApplicationMySQLUser   string
	dbApplicationMySQLPass   string
	// privilege
	privilegeEnabledStr string
	// metadata
	metadataTableAnalyzeMinRole int
	// alert
	alertSMTPEnabledStr string
	alertSMTPFormat     string
	alertSMTPURL        string
	alertSMTPUser       string
	alertSMTPPass       string
	alertSMTPFrom       string
	alertHTTPEnabledStr string
	alertHTTPURL        string
	alertHTTPConfig     string
	// healthcheck
	healthcheckMaxRange       int
	healthcheckAlertOwnerType string
	// query
	queryMinRowsExamined int
	// sqladvisor
	sqladvisorSoarBin          string
	sqladvisorSoarConfig       string
	sqladvisorSoarSamplingStr  string
	sqladvisorSoarProfilingStr string
	sqladvisorSoarTraceStr     string
	sqladvisorSoarExplainStr   string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "das",
	Short: "das is short for database autonomy service",
	Long:  `das provides database autonomy service, which includes disk space usage prediction, sql tuning advice, database health check, and so on...`,
	Run: func(cmd *cobra.Command, args []string) {
		// if no subcommand is set, it will print help information.
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				fmt.Println(fmt.Sprintf("%+v", message.NewMessage(message.ErrPrintHelpInfo, errors.Trace(err))))
				os.Exit(constant.DefaultAbnormalExitCode)
			}

			os.Exit(constant.DefaultNormalExitCode)
		}

		// init config
		err := initConfig()
		if err != nil {
			fmt.Println(fmt.Sprintf("%+v", message.NewMessage(message.ErrInitConfig, err)))
			os.Exit(constant.DefaultAbnormalExitCode)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(fmt.Sprintf("%+v", errors.Trace(err)))
		os.Exit(constant.DefaultAbnormalExitCode)
	}
}

func init() {
	// set usage template
	rootCmd.SetUsageTemplate(UsageTemplateWithoutDefault())

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	// config
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", constant.DefaultRandomString, "config file path")
	// daemon
	rootCmd.PersistentFlags().StringVar(&daemonStr, "daemon", constant.DefaultRandomString, fmt.Sprintf("whether run in background as a daemon(default: %s)", constant.FalseString))
	// log
	rootCmd.PersistentFlags().StringVar(&logFileName, "log-file", constant.DefaultRandomString, fmt.Sprintf("specify the log file name(default: %s)", filepath.Join(config.DefaultLogDir, log.DefaultLogFileName)))
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", constant.DefaultRandomString, fmt.Sprintf("specify the log level(default: %s)", log.DefaultLogLevel))
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", constant.DefaultRandomString, fmt.Sprintf("specify the log format(default: %s)", log.DefaultLogFormat))
	rootCmd.PersistentFlags().IntVar(&logMaxSize, "log-max-size", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max size(default: %d, unit: MB)", log.DefaultLogMaxSize))
	rootCmd.PersistentFlags().IntVar(&logMaxDays, "log-max-days", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max days(default: %d)", log.DefaultLogMaxDays))
	rootCmd.PersistentFlags().IntVar(&logMaxBackups, "log-max-backups", constant.DefaultRandomInt, fmt.Sprintf("specify the log file max backups(default: %d)", log.DefaultLogMaxBackups))
	// server
	rootCmd.PersistentFlags().StringVar(&serverAddr, "server-addr", constant.DefaultRandomString, fmt.Sprintf("specify the server port(default: %s)", config.DefaultServerAddr))
	rootCmd.PersistentFlags().StringVar(&serverPidFile, "server-pid-file", constant.DefaultRandomString, fmt.Sprintf("specify the server pid file path(default: %s)", filepath.Join(config.DefaultBaseDir, fmt.Sprintf("%s.pid", config.DefaultCommandName))))
	rootCmd.PersistentFlags().IntVar(&serverReadTimeout, "server-read-timeout", constant.DefaultRandomInt, fmt.Sprintf("specify the read timeout in seconds of http request(default: %d)", config.DefaultServerReadTimeout))
	rootCmd.PersistentFlags().IntVar(&serverWriteTimeout, "server-write-timeout", constant.DefaultRandomInt, fmt.Sprintf("specify the write timeout in seconds of http request(default: %d)", config.DefaultServerWriteTimeout))
	// database
	rootCmd.PersistentFlags().StringVar(&dbDASMySQLAddr, "db-das-mysql-addr", constant.DefaultRandomString, fmt.Sprintf("specify das database address(format: host:port)(default: %s)", fmt.Sprintf("%s:%d", constant.DefaultLocalHostIP, constant.DefaultMySQLPort)))
	rootCmd.PersistentFlags().StringVar(&dbDASMySQLName, "db-das-mysql-name", constant.DefaultRandomString, fmt.Sprintf("specify das database name(default: %s)", config.DefaultDBName))
	rootCmd.PersistentFlags().StringVar(&dbDASMySQLUser, "db-das-mysql-user", constant.DefaultRandomString, fmt.Sprintf("specify das database user name(default: %s)", config.DefaultDBUser))
	rootCmd.PersistentFlags().StringVar(&dbDASMySQLPass, "db-das-mysql-pass", constant.DefaultRandomString, fmt.Sprintf("specify das database user password(default: %s)", config.DefaultDBPass))
	rootCmd.PersistentFlags().IntVar(&dbPoolMaxConnections, "db-pool-max-connections", constant.DefaultRandomInt, fmt.Sprintf("specify max connections of the connection pool(default: %d)", mysql.DefaultMaxConnections))
	rootCmd.PersistentFlags().IntVar(&dbPoolInitConnections, "db-pool-init-connections", constant.DefaultRandomInt, fmt.Sprintf("specify initial connections of the connection pool(default: %d)", mysql.DefaultMaxIdleConnections))
	rootCmd.PersistentFlags().IntVar(&dbPoolMaxIdleConnections, "db-pool-max-idle-connections", constant.DefaultRandomInt, fmt.Sprintf("specify max idle connections of the connection pool(default: %d)", mysql.DefaultMaxIdleConnections))
	rootCmd.PersistentFlags().IntVar(&dbPoolMaxIdleTime, "db-pool-max-idle-time", constant.DefaultRandomInt, fmt.Sprintf("specify max idle time of connections of the connection pool, (default: %d, unit: seconds)", mysql.DefaultMaxIdleTime))
	rootCmd.PersistentFlags().IntVar(&dbPoolKeepAliveInterval, "db-pool-keep-alive-interval", constant.DefaultRandomInt, fmt.Sprintf("specify keep alive interval of connections of the connection pool(default: %d, unit: seconds)", mysql.DefaultKeepAliveInterval))
	rootCmd.PersistentFlags().StringVar(&dbApplicationMySQLUser, "db-application-mysql-user", constant.DefaultRandomString, fmt.Sprintf("specify mysql user name of application(default: %s)", config.DefaultDBUser))
	rootCmd.PersistentFlags().StringVar(&dbApplicationMySQLPass, "db-application-mysql-pass", constant.DefaultRandomString, fmt.Sprintf("specify mysql user password of application(default: %s)", config.DefaultDBPass))
	rootCmd.PersistentFlags().StringVar(&dbMonitorPrometheusUser, "db-monitor-prometheus-user", constant.DefaultRandomString, fmt.Sprintf("specify prometheus user name of monitor system(default: %s)", config.DefaultDBUser))
	rootCmd.PersistentFlags().StringVar(&dbMonitorPrometheusPass, "db-monitor-prometheus-pass", constant.DefaultRandomString, fmt.Sprintf("specify prometheus user password of monitor system(default: %s)", config.DefaultDBPass))
	rootCmd.PersistentFlags().StringVar(&dbMonitorClickhouseUser, "db-monitor-clickhouse-user", constant.DefaultRandomString, fmt.Sprintf("specify clickhouse user name of monitor system(default: %s)", config.DefaultDBUser))
	rootCmd.PersistentFlags().StringVar(&dbMonitorClickhousePass, "db-monitor-clickhouse-pass", constant.DefaultRandomString, fmt.Sprintf("specify clickhouse user password of monitor system(default: %s)", config.DefaultDBPass))
	rootCmd.PersistentFlags().StringVar(&dbMonitorMySQLUser, "db-monitor-mysql-user", constant.DefaultRandomString, fmt.Sprintf("specify mysql user name of monitor system(default: %s)", config.DefaultDBUser))
	rootCmd.PersistentFlags().StringVar(&dbMonitorMySQLPass, "db-monitor-mysql-pass", constant.DefaultRandomString, fmt.Sprintf("specify mysql user password of monitor system(default: %s)", config.DefaultDBPass))
	// privilege
	rootCmd.PersistentFlags().StringVar(&privilegeEnabledStr, "privilege-enabled", constant.DefaultRandomString, fmt.Sprintf("specify if enables privilege module(default: %s)", constant.TrueString))
	// metadata
	rootCmd.PersistentFlags().IntVar(&metadataTableAnalyzeMinRole, "metadata-table-analyze-min-role", constant.DefaultRandomInt, fmt.Sprintf("specify the minimum role which has analyzing table privilege(default: %d)", config.DefaultMetadataTableAnalyzeMinRole))
	// alert
	rootCmd.PersistentFlags().StringVar(&alertSMTPEnabledStr, "alert-smtp-enabled", constant.DefaultRandomString, fmt.Sprintf("specify if enables smtp method(default: %s)", constant.TrueString))
	rootCmd.PersistentFlags().StringVar(&alertSMTPFormat, "alert-smtp-format", constant.DefaultRandomString, fmt.Sprintf("specify the email content format(default: %s)", config.DefaultAlterSMTPFormat))
	rootCmd.PersistentFlags().StringVar(&alertSMTPURL, "alert-smtp-url", constant.DefaultRandomString, fmt.Sprintf("specify the url of the smtp server(default: %s)", config.DefaultAlertSMTPURL))
	rootCmd.PersistentFlags().StringVar(&alertSMTPUser, "alert-smtp-user", constant.DefaultRandomString, fmt.Sprintf("specify the username of the smtp server(default: %s)", config.DefaultAlertSMTPUser))
	rootCmd.PersistentFlags().StringVar(&alertSMTPPass, "alert-smtp-pass", constant.DefaultRandomString, fmt.Sprintf("specify the password of the smtp server(default: %s)", config.DefaultAlertSMTPPass))
	rootCmd.PersistentFlags().StringVar(&alertSMTPFrom, "alert-smtp-from", constant.DefaultRandomString, fmt.Sprintf("specify the from email address(default: %s)", config.DefaultAlertSMTPFrom))
	rootCmd.PersistentFlags().StringVar(&alertHTTPEnabledStr, "alert-http-enabled", constant.DefaultRandomString, fmt.Sprintf("specify if enables http method(default: %s)", constant.FalseString))
	rootCmd.PersistentFlags().StringVar(&alertHTTPURL, "alert-http-url", constant.DefaultRandomString, fmt.Sprintf("specify actual alert api url(default: %s)", config.DefaultAlertHTTPURL))
	rootCmd.PersistentFlags().StringVar(&alertHTTPConfig, "alert-http-config", constant.DefaultRandomString, fmt.Sprintf("specify alert api parameters(default: %s)", config.DefaultAlertHTTPConfig))
	// healthcheck
	rootCmd.PersistentFlags().IntVar(&healthcheckMaxRange, "healthcheck-max-range", constant.DefaultRandomInt, fmt.Sprintf("specify healthcheck maximum range(default: %d)", config.DefaultHealthCheckMaxRange))
	rootCmd.PersistentFlags().StringVar(&healthcheckAlertOwnerType, "healthcheck-alert-owner-type", constant.DefaultRandomString, fmt.Sprintf("specify healthcheck alert owner type(default: %s)", config.DefaultHealthcheckAlertOwnerType))
	// query
	rootCmd.PersistentFlags().IntVar(&queryMinRowsExamined, "query-min-rows-examined", constant.DefaultRandomInt, fmt.Sprintf("specify query min rows examined(default: %d", config.DefaultQueryMinRowsExamined))
	// sqladvisor
	rootCmd.PersistentFlags().StringVar(&sqladvisorSoarBin, "sqladvisor-soar-bin", constant.DefaultRandomString, fmt.Sprintf("specify binary path of soar(default: %s)", config.DefaultSQLAdvisorSoarBin))
	rootCmd.PersistentFlags().StringVar(&sqladvisorSoarConfig, "sqladvisor-soar-config", constant.DefaultRandomString, fmt.Sprintf("specify config file path of soar(default: %s)", config.DefaultSQLAdvisorSoarConfig))
	rootCmd.PersistentFlags().StringVar(&sqladvisorSoarSamplingStr, "sqladvisor-soar-sampling", constant.DefaultRandomString, fmt.Sprintf("specify if enabling sampling for soar(default: %s)", constant.FalseString))
	rootCmd.PersistentFlags().StringVar(&sqladvisorSoarProfilingStr, "sqladvisor-soar-profiling", constant.DefaultRandomString, fmt.Sprintf("specify if enabling profiling for soar(default: %s)", constant.FalseString))
	rootCmd.PersistentFlags().StringVar(&sqladvisorSoarTraceStr, "sqladvisor-soar-trace", constant.DefaultRandomString, fmt.Sprintf("specify if enabling trace for soar(default: %s)", constant.FalseString))
	rootCmd.PersistentFlags().StringVar(&sqladvisorSoarExplainStr, "sqladvisor-soar-explain", constant.DefaultRandomString, fmt.Sprintf("specify if enabling explain for soar(default: %s)", constant.FalseString))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() error {
	var err error

	// init default config
	err = initDefaultConfig()
	if err != nil {
		return message.NewMessage(message.ErrInitDefaultConfig, err)
	}

	// read config with config file
	err = ReadConfigFile()
	if err != nil {
		return message.NewMessage(message.ErrReadConfigFile, err)
	}

	// override config with command line arguments
	err = OverrideConfig()
	if err != nil {
		return message.NewMessage(message.ErrOverrideCommandLineArgs, err)
	}

	// init log
	fileName := viper.GetString(config.LogFileKey)
	level := viper.GetString(config.LogLevelKey)
	format := viper.GetString(config.LogFormatKey)
	maxSize := viper.GetInt(config.LogMaxSizeKey)
	maxDays := viper.GetInt(config.LogMaxDaysKey)
	maxBackups := viper.GetInt(config.LogMaxBackupsKey)

	fileNameAbs := fileName
	isAbs := filepath.IsAbs(fileName)
	if !isAbs {
		fileNameAbs, err = filepath.Abs(fileName)
		if err != nil {
			return message.NewMessage(message.ErrAbsoluteLogFilePath, errors.Trace(err), fileName)
		}
	}
	_, _, err = log.InitFileLogger(fileNameAbs, level, format, maxSize, maxDays, maxBackups)
	if err != nil {
		return message.NewMessage(message.ErrInitLogger, err)
	}

	log.SetDisableDoubleQuotes(true)
	log.SetDisableEscape(true)

	return nil
}

// initDefaultConfig initiate default configuration
func initDefaultConfig() (err error) {
	// get base dir
	baseDir, err = filepath.Abs(config.DefaultBaseDir)
	if err != nil {
		return message.NewMessage(message.ErrBaseDir, errors.Trace(err), config.DefaultCommandName)
	}
	// set default config value
	config.SetDefaultConfig(baseDir)
	err = config.ValidateConfig()
	if err != nil {
		return err
	}

	return nil
}

// ReadConfigFile read configuration from config file, it will override the init configuration
func ReadConfigFile() (err error) {
	if cfgFile != constant.EmptyString && cfgFile != constant.DefaultRandomString {
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType(defaultConfigFileType)
		err = viper.ReadInConfig()
		if err != nil {
			return errors.Trace(err)
		}
		err = config.ValidateConfig()
		if err != nil {
			return message.NewMessage(message.ErrValidateConfig, err)
		}
	}

	return nil
}

// OverrideConfig read configuration from command line interface, it will override the config file configuration
func OverrideConfig() (err error) {
	// override config
	if cfgFile != constant.EmptyString && cfgFile != constant.DefaultRandomString {
		viper.Set(config.ConfKey, cfgFile)
	}

	// override daemon
	if daemonStr != constant.DefaultRandomString {
		daemon, err := cast.ToBoolE(daemonStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.DaemonKey, daemon)
	}

	// override log
	if logFileName != constant.DefaultRandomString {
		viper.Set(config.LogFileKey, logFileName)
	}
	if logLevel != constant.DefaultRandomString {
		logLevel = strings.ToLower(logLevel)
		viper.Set(config.LogLevelKey, logLevel)
	}
	if logFormat != constant.DefaultRandomString {
		logLevel = strings.ToLower(logFormat)
		viper.Set(config.LogFormatKey, logFormat)
	}
	if logMaxSize != constant.DefaultRandomInt {
		viper.Set(config.LogMaxSizeKey, logMaxSize)
	}
	if logMaxDays != constant.DefaultRandomInt {
		viper.Set(config.LogMaxDaysKey, logMaxDays)
	}
	if logMaxBackups != constant.DefaultRandomInt {
		viper.Set(config.LogMaxBackupsKey, logMaxBackups)
	}

	// override server
	if serverAddr != constant.DefaultRandomString {
		viper.Set(config.ServerAddrKey, serverAddr)
	}
	if serverPidFile != constant.DefaultRandomString {
		viper.Set(config.ServerPidFileKey, serverPidFile)
	}
	if serverReadTimeout != constant.DefaultRandomInt {
		viper.Set(config.ServerReadTimeoutKey, serverReadTimeout)
	}
	if serverWriteTimeout != constant.DefaultRandomInt {
		viper.Set(config.ServerWriteTimeoutKey, serverWriteTimeout)
	}

	// override database
	if dbDASMySQLAddr != constant.DefaultRandomString {
		viper.Set(config.DBDASMySQLAddrKey, dbDASMySQLAddr)
	}
	if dbDASMySQLName != constant.DefaultRandomString {
		viper.Set(config.DBDASMySQLNameKey, dbDASMySQLName)
	}
	if dbDASMySQLUser != constant.DefaultRandomString {
		viper.Set(config.DBDASMySQLUserKey, dbDASMySQLUser)
	}
	if dbDASMySQLPass != constant.DefaultRandomString {
		viper.Set(config.DBDASMySQLPassKey, dbDASMySQLPass)
	}
	if dbPoolMaxConnections != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxConnectionsKey, dbPoolMaxConnections)
	}
	if dbPoolInitConnections != constant.DefaultRandomInt {
		viper.Set(config.DBPoolInitConnectionsKey, dbPoolInitConnections)
	}
	if dbPoolMaxIdleConnections != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxIdleConnectionsKey, dbPoolMaxIdleConnections)
	}
	if dbPoolMaxIdleTime != constant.DefaultRandomInt {
		viper.Set(config.DBPoolMaxIdleTimeKey, dbPoolMaxIdleTime)
	}
	if dbPoolKeepAliveInterval != constant.DefaultRandomInt {
		viper.Set(config.DBPoolKeepAliveIntervalKey, dbPoolKeepAliveInterval)
	}
	if dbMonitorPrometheusUser != constant.DefaultRandomString {
		viper.Set(config.DBMonitorPrometheusUserKey, dbMonitorPrometheusUser)
	}
	if dbMonitorPrometheusPass != constant.DefaultRandomString {
		viper.Set(config.DBMonitorPrometheusPassKey, dbMonitorPrometheusPass)
	}
	if dbMonitorClickhouseUser != constant.DefaultRandomString {
		viper.Set(config.DBMonitorClickhouseUserKey, dbMonitorClickhouseUser)
	}
	if dbMonitorClickhousePass != constant.DefaultRandomString {
		viper.Set(config.DBMonitorClickhousePassKey, dbMonitorClickhousePass)
	}
	if dbMonitorMySQLUser != constant.DefaultRandomString {
		viper.Set(config.DBMonitorMySQLUserKey, dbMonitorMySQLUser)
	}
	if dbMonitorMySQLPass != constant.DefaultRandomString {
		viper.Set(config.DBMonitorMySQLPassKey, dbMonitorMySQLPass)
	}
	if dbApplicationMySQLUser != constant.DefaultRandomString {
		viper.Set(config.DBApplicationMySQLUserKey, dbApplicationMySQLUser)
	}
	if dbApplicationMySQLPass != constant.DefaultRandomString {
		viper.Set(config.DBApplicationMySQLPassKey, dbApplicationMySQLPass)
	}

	// override privilege
	if privilegeEnabledStr != constant.DefaultRandomString {
		privilegeEnabled, err := cast.ToBoolE(privilegeEnabledStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.PrivilegeEnabledKey, privilegeEnabled)
	}

	// override metadata
	if metadataTableAnalyzeMinRole != constant.DefaultRandomInt {
		viper.Set(config.MetadataTableAnalyzeMinRoleKey, metadataTableAnalyzeMinRole)
	}

	// override alert
	if alertSMTPEnabledStr != constant.DefaultRandomString {
		alertSMTPEnabled, err := cast.ToBoolE(alertSMTPEnabledStr)
		if err != nil {
			return errors.Trace(err)
		}

		viper.Set(config.AlertSMTPEnabledKey, alertSMTPEnabled)
	}
	if alertSMTPFormat != constant.DefaultRandomString {
		viper.Set(config.AlertSMTPFormatKey, alertSMTPFormat)
	}
	if alertSMTPURL != constant.DefaultRandomString {
		viper.Set(config.AlertSMTPURLKey, alertSMTPURL)
	}
	if alertSMTPUser != constant.DefaultRandomString {
		viper.Set(config.AlertSMTPUserKey, alertSMTPUser)
	}
	if alertSMTPPass != constant.DefaultRandomString {
		viper.Set(config.AlertSMTPPassKey, alertSMTPPass)
	}
	if alertSMTPFrom != constant.DefaultRandomString {
		viper.Set(config.AlertSMTPFromKey, alertSMTPFrom)
	}
	if alertHTTPEnabledStr != constant.DefaultRandomString {
		alertHTTPEnabled, err := cast.ToBoolE(alertHTTPEnabledStr)
		if err != nil {
			return errors.Trace(err)
		}
		viper.Set(config.AlertHTTPEnabledKey, alertHTTPEnabled)
	}
	if alertHTTPURL != constant.DefaultRandomString {
		viper.Set(config.AlertHTTPURLKey, alertHTTPURL)
	}
	if alertHTTPConfig != constant.DefaultRandomString {
		viper.Set(config.AlertHTTPConfigKey, alertHTTPConfig)
	}

	// override healthcheck
	if healthcheckMaxRange != constant.DefaultRandomInt {
		viper.Set(config.HealthcheckMaxRangeKey, healthcheckMaxRange)
	}
	if healthcheckAlertOwnerType != constant.DefaultRandomString {
		viper.Set(config.HealthcheckAlertOwnerTypeKey, healthcheckAlertOwnerType)
	}

	// override query
	if queryMinRowsExamined != constant.DefaultRandomInt {
		viper.Set(config.QueryMinRowsExaminedKey, queryMinRowsExamined)
	}

	// override sqladvisor
	if sqladvisorSoarBin != constant.DefaultRandomString {
		viper.Set(config.SQLAdvisorSoarBinKey, sqladvisorSoarBin)
	}
	if sqladvisorSoarConfig != constant.DefaultRandomString {
		viper.Set(config.SQLAdvisorSoarConfigKey, sqladvisorSoarConfig)
	}
	if sqladvisorSoarSamplingStr == constant.TrueString {
		viper.Set(config.SQLAdvisorSoarSamplingKey, true)
	} else {
		viper.Set(config.SQLAdvisorSoarSamplingKey, false)
	}
	if sqladvisorSoarProfilingStr == constant.TrueString {
		viper.Set(config.SQLAdvisorSoarProfilingKey, true)
	} else {
		viper.Set(config.SQLAdvisorSoarProfilingKey, false)
	}
	if sqladvisorSoarTraceStr == constant.TrueString {
		viper.Set(config.SQLAdvisorSoarTraceKey, true)
	} else {
		viper.Set(config.SQLAdvisorSoarTraceKey, false)
	}
	if sqladvisorSoarExplainStr == constant.TrueString {
		viper.Set(config.SQLAdvisorSoarExplainKey, true)
	} else {
		viper.Set(config.SQLAdvisorSoarExplainKey, false)
	}

	// validate configuration
	err = config.ValidateConfig()
	if err != nil {
		return message.NewMessage(message.ErrValidateConfig, err)
	}

	return err
}

// UsageTemplateWithoutDefault returns a usage template which does not contain default part
func UsageTemplateWithoutDefault() string {
	return `Usage:{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}{{if gt (len .Aliases) 0}}

Aliases:
  {{.NameAndAliases}}{{end}}{{if .HasExample}}

Examples:
{{.Example}}{{end}}{{if .HasAvailableSubCommands}}

Available Commands:{{range .Commands}}{{if (or .IsAvailableCommand (eq .Name "help"))}}
  {{rpad .Name .NamePadding }} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableLocalFlags}}

Flags:
{{.LocalFlags.FlagUsagesWithoutDefault | trimTrailingWhitespaces}}{{end}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsagesWithoutDefault | trimTrailingWhitespaces}}{{end}}{{if .HasHelpSubCommands}}

Additional help topics:{{range .Commands}}{{if .IsAdditionalHelpTopicCommand}}
  {{rpad .CommandPath .CommandPathPadding}} {{.Short}}{{end}}{{end}}{{end}}{{if .HasAvailableSubCommands}}

Use "{{.CommandPath}} [command] --help" for more information about a command.{{end}}
`
}
