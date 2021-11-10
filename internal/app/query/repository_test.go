package query

import (
	"testing"
	"time"

	"github.com/romberli/das/config"
	"github.com/romberli/go-util/constant"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestQueryRepositoryAll(t *testing.T) {
	TestMySQLRepo_GetByServiceNames(t)
	TestMySQLRepo_GetByDBName(t)
	TestMySQLRepo_GetBySQLID(t)
	TestClickhouseRepo_GetByDBName(t)
	TestClickhouseRepo_GetByServiceNames(t)
	TestClickhouseRepo_GetBySQLID(t)
}

func init() {
	viper.Set(config.DBMonitorMySQLUserKey, config.DefaultDBMonitorMySQLUser)
	viper.Set(config.DBMonitorMySQLPassKey, config.DefaultDBMonitorMySQLPass)
}

func TestMySQLRepo_GetByServiceNames(t *testing.T) {
	asst := assert.New(t)

	q := NewQuerierWithGlobal(NewConfigWithDefault())
	monitorSystem, err := q.getMonitorSystemByMySQLServerID(testMySQLServerID)

	monitorRepo, err := q.getMonitorRepo(monitorSystem)
	mr := monitorRepo
	queries, err := mr.GetByServiceNames([]string{testServiceName})
	asst.Nil(err, "test TestMySQLRepo_GetByServiceNames() Failed")
	asst.NotZero(len(queries), "test TestMySQLRepo_GetByServiceNames() Failed")
}

func TestMySQLRepo_GetByDBName(t *testing.T) {
	asst := assert.New(t)

	q := NewQuerierWithGlobal(NewConfigWithDefault())
	monitorSystem, err := q.getMonitorSystemByMySQLServerID(testMySQLServerID)

	mr, err := q.getMonitorRepo(monitorSystem)
	qu, err := mr.GetByDBName(testServiceName, testDBName)
	asst.Equal(nil, err, "test MySQLRepo_GetByDBName() Failed")
	asst.Equal(true, qu != nil, "test MySQLRepo_GetByDBName() Failed")
}

func TestMySQLRepo_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	q := NewQuerierWithGlobal(NewConfigWithDefault())
	q.config.startTime = time.Now().Add(-constant.Week)

	q.config.endTime = time.Now()
	monitorSystem, err := q.getMonitorSystemByMySQLServerID(testMySQLServerID)

	monitorRepo, err := q.getMonitorRepo(monitorSystem)

	mr := monitorRepo
	query, err := mr.GetBySQLID(testServiceName, testSQLID)
	asst.Equal(nil, err, "test MySQLRepo_GetBySQLID() Failed")
	asst.NotNil(query, "test MySQLRepo_GetBySQLID() Failed")
}

func TestClickhouseRepo_GetByDBName(t *testing.T) {
	asst := assert.New(t)

	q := NewQuerierWithGlobal(NewConfigWithDefault())
	monitorSystem, err := q.getMonitorSystemByMySQLServerID(testMySQLServerID)

	monitorRepo, err := q.getMonitorRepo(monitorSystem)

	cr := monitorRepo
	qu, err := cr.GetByDBName(testServiceName, testDBName)

	asst.Equal(nil, err, "test ClickhouseRepo_GetByDBName() Failed")
	asst.Equal(true, qu != nil, "test ClickhouseRepo_GetByDBName() Failed")

}

func TestClickhouseRepo_GetByServiceNames(t *testing.T) {
	asst := assert.New(t)

	q := NewQuerierWithGlobal(NewConfigWithDefault())
	monitorSystem, err := q.getMonitorSystemByMySQLServerID(testMySQLServerID)

	monitorRepo, err := q.getMonitorRepo(monitorSystem)

	cr := monitorRepo
	qu, err := cr.GetByServiceNames([]string{testServiceName})
	asst.Equal(nil, err, "test ClickhouseRepo_GetByServiceNames() Failed")
	asst.Equal(true, qu != nil, "test ClickhouseRepo_GetByServiceNames() Failed")
}

func TestClickhouseRepo_GetBySQLID(t *testing.T) {
	asst := assert.New(t)

	q := NewQuerierWithGlobal(NewConfigWithDefault())
	q.config.startTime = time.Date(2021, 11, 1, 1, 1, 1, 1, time.Local)
	q.config.endTime = time.Date(2021, 11, 30, 1, 1, 1, 1, time.Local)
	monitorSystem, err := q.getMonitorSystemByMySQLServerID(testMySQLServerID)
	monitorRepo, err := q.getMonitorRepo(monitorSystem)

	cr := monitorRepo
	qu, err := cr.GetBySQLID(testServiceName, testSQLID)
	asst.Equal(nil, err, "test ClickhouseRepo_GetBySQLID() Failed")
	asst.Equal(true, qu != nil, "test ClickhouseRepo_GetBySQLID() Failed")
}
