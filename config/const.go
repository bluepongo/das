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
package config

import (
	"github.com/romberli/go-util/constant"
)

// global constant
const (
	DefaultCommandName = "das"
	DefaultDaemon      = false
	DefaultBaseDir     = constant.CurrentDir
	// log
	DefaultLogDir    = "./log"
	MinLogMaxSize    = 1
	MaxLogMaxSize    = constant.MaxInt
	MinLogMaxDays    = 1
	MaxLogMaxDays    = constant.MaxInt
	MinLogMaxBackups = 1
	MaxLogMaxBackups = constant.MaxInt
	// server
	DefaultServerAddr         = "0.0.0.0:6090"
	DefaultServerReadTimeout  = 5
	DefaultServerWriteTimeout = 10
	MinServerReadTimeout      = 0
	MaxServerReadTimeout      = 60
	MinServerWriteTimeout     = 1
	MaxServerWriteTimeout     = 60
	DaemonArgTrue             = "--daemon=true"
	DaemonArgFalse            = "--daemon=false"
	// db
	DefaultDBName                  = "das"
	DefaultDBUser                  = "root"
	DefaultDBPass                  = "root"
	MinDBPoolMaxConnections        = 1
	MaxDBPoolMaxConnections        = constant.MaxInt
	MinDBPoolInitConnections       = 1
	MaxDBPoolInitConnections       = constant.MaxInt
	MinDBPoolMaxIdleConnections    = 1
	MaxDBPoolMaxIdleConnections    = constant.MaxInt
	MinDBPoolMaxIdleTime           = 1
	MaxDBPoolMaxIdleTime           = constant.MaxInt
	MinDBPoolKeepAliveInterval     = 1
	MaxDBPoolKeepAliveInterval     = constant.MaxInt
	DefaultDBMonitorPrometheusUser = "admin"
	DefaultDBMonitorPrometheusPass = "admin"
	DefaultDBMonitorClickhouseUser = ""
	DefaultDBMonitorClickhousePass = ""
	DefaultDBMonitorMySQLUser      = "root"
	DefaultDBMonitorMySQLPass      = "root"
	DefaultDBApplicationMySQLUser  = "root"
	DefaultDBApplicationMySQLPass  = "root"
	// privilege
	DefaultPrivilegeEnabled = true
	// metadata
	MetadataUserDeveloperRole          = 1
	MetadataUserDBARole                = 2
	MetadataUserAdminRole              = 3
	MinMetadataTableAnalyzeMinRole     = MetadataUserDeveloperRole
	MaxMetadataTableAnalyzeMinRole     = MetadataUserAdminRole
	DefaultMetadataTableAnalyzeMinRole = MetadataUserDBARole
	// alert
	DefaultAlertSMTPEnabled = true
	AlertSMTPTextFormat     = "text"
	AlertSMTPHTMLFormat     = "html"
	DefaultAlterSMTPFormat  = AlertSMTPTextFormat
	DefaultAlertSMTPURL     = "127.0.0.1:25"
	DefaultAlertSMTPUser    = "root"
	DefaultAlertSMTPPass    = "root"
	DefaultAlertSMTPFrom    = "mail@example.com"
	DefaultAlertHTTPEnabled = false
	DefaultAlertHTTPURL     = "http://127.0.0.1:8080"
	DefaultAlertHTTPConfig  = "{}"
	// healthcheck
	DefaultHealthCheckMaxRange       = 30
	HealthcheckAlertOwnerTypeApp     = "app"
	HealthcheckAlertOwnerTypeDB      = "db"
	HealthcheckAlertOwnerTypeAll     = "all"
	DefaultHealthcheckAlertOwnerType = HealthcheckAlertOwnerTypeAll
	// query
	DefaultQueryMinRowsExamined = 100000
	// sqladvisor
	DefaultSQLAdvisorSoarBin       = "./soar"
	DefaultSQLAdvisorSoarConfig    = "./soar.yaml"
	DefaultSQLAdvisorSoarBlacklist = "./soar.blacklist"
)

// configuration constant
const (
	// config
	ConfKey = "config"
	// daemon
	DaemonKey = "daemon"
	// log
	LogFileKey       = "log.file"
	LogLevelKey      = "log.level"
	LogFormatKey     = "log.format"
	LogMaxSizeKey    = "log.maxSize"
	LogMaxDaysKey    = "log.maxDays"
	LogMaxBackupsKey = "log.maxBackups"
	// server
	ServerAddrKey         = "server.addr"
	ServerPidFileKey      = "server.pidFile"
	ServerReadTimeoutKey  = "server.readTimeout"
	ServerWriteTimeoutKey = "server.writeTimeout"
	// database
	DBDASMySQLAddrKey           = "db.das.mysql.addr"
	DBDASMySQLNameKey           = "db.das.mysql.name"
	DBDASMySQLUserKey           = "db.das.mysql.user"
	DBDASMySQLPassKey           = "db.das.mysql.pass"
	DBPoolMaxConnectionsKey     = "db.pool.maxConnections"
	DBPoolInitConnectionsKey    = "db.pool.initConnections"
	DBPoolMaxIdleConnectionsKey = "db.pool.maxIdleConnections"
	DBPoolMaxIdleTimeKey        = "db.pool.maxIdleTime"
	DBPoolKeepAliveIntervalKey  = "db.pool.keepAliveInterval"
	DBApplicationMySQLUserKey   = "db.application.mysql.user"
	DBApplicationMySQLPassKey   = "db.application.mysql.pass"
	DBMonitorPrometheusUserKey  = "db.monitor.prometheus.user"
	DBMonitorPrometheusPassKey  = "db.monitor.prometheus.pass"
	DBMonitorClickhouseUserKey  = "db.monitor.clickhouse.user"
	DBMonitorClickhousePassKey  = "db.monitor.clickhouse.pass"
	DBMonitorMySQLUserKey       = "db.monitor.mysql.user"
	DBMonitorMySQLPassKey       = "db.monitor.mysql.pass"
	// privilege
	PrivilegeEnabledKey = "privilege.enabled"
	// metadata
	MetadataTableAnalyzeMinRoleKey = "metadata.table.analyze.minRole"
	// alert
	AlertSMTPEnabledKey = "alert.smtp.enabled"
	AlertSMTPFormatKey  = "alert.smtp.format"
	AlertSMTPURLKey     = "alert.smtp.url"
	AlertSMTPUserKey    = "alert.smtp.user"
	AlertSMTPPassKey    = "alert.smtp.pass"
	AlertSMTPFromKey    = "alert.smtp.from"
	AlertHTTPEnabledKey = "alert.http.enabled"
	AlertHTTPURLKey     = "alert.http.url"
	AlertHTTPConfigKey  = "alert.http.config"
	// healthcheck
	HealthcheckMaxRangeKey       = "healthcheck.maxRange"
	HealthcheckAlertOwnerTypeKey = "healthcheck.alert.ownerType"
	// query
	QueryMinRowsExaminedKey = "query.minRowsExamined"
	// sqladvisor
	SQLAdvisorSoarBinKey       = "sqladvisor.soar.bin"
	SQLAdvisorSoarConfigKey    = "sqladvisor.soar.config"
	SQLAdvisorSoarSamplingKey  = "sqladvisor.soar.sampling"
	SQLAdvisorSoarProfilingKey = "sqladvisor.soar.profiling"
	SQLAdvisorSoarTraceKey     = "sqladvisor.soar.trace"
	SQLAdvisorSoarExplainKey   = "sqladvisor.soar.explain"
)
