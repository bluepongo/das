package metadata

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/metadata"
	"github.com/romberli/das/pkg/message"
	msgmeta "github.com/romberli/das/pkg/message/metadata"
	"github.com/romberli/das/pkg/resp"
	"github.com/romberli/go-util/common"
	"github.com/romberli/go-util/constant"
	"github.com/romberli/log"
)

const (
	userIDJSON      = "id"
	userNameJSON    = "user_name"
	employeeIDJSON  = "employee_id"
	accountNameJSON = "account_name"
	loginNameJSON   = "login_name"
	emailJSON       = "email"
	telephoneJSON   = "telephone"
	mobileJSON      = "mobile"

	userAppsStruct               = "Apps"
	userDBsStruct                = "DBs"
	userMiddlewareClustersStruct = "MiddlewareClusters"
	userMySQLClustersStruct      = "MySQLClusters"
	userMySQLServersStruct       = "MySQLServers"
	userNameStruct               = "LoginName"
	departmentNameStruct         = "DepartmentName"
	employeeIDStruct             = "EmployeeID"
	accountNameStruct            = "AccountName"
	emailStruct                  = "Email"
	telephoneStruct              = "Telephone"
	roleStruct                   = "Role"
	mobileStruct                 = "Mobile"
)

// @Tags 	user
// @Summary get all users
// @Accept	application/json
// @Param	token body string true "token"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router  /api/v1/metadata/user [get]

func GetUser(c *gin.Context) {
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get entities
	err := s.GetAll()
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUserAll, err)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// responseF
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUserAll, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUserAll)
}

// @Tags 	user
// @Summary get user by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id path int true "user id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router 	/api/v1/metadata/user/get/:id [get]
func GetUserByID(c *gin.Context) {
	// get param
	idStr := c.Param(envIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, envIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUserByID, errors.Trace(err), id)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err = s.GetByID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetUserByID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetUserByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetUserByID, id)
}

// @Tags 	user
// @Summary get user by name
// @Accept	application/json
// @Param	token body string true "token"
// @Param	user_name path string true "user name"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router 	/api/v1/metadata/user/user-name/:user_name [get]
func GetByUserName(c *gin.Context) {
	// get param
	userName := c.Param(userNameJSON)
	if userName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, userNameJSON)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByUserName(userName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetByUserName, err, userName)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetByUserName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetByUserName, userName)
}

// @Tags 	user
// @Summary get user by employee id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	employee_id path string true "employee id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router 	/api/v1/metadata/user/employee-id/:employee_id [get]
func GetUserByEmployeeID(c *gin.Context) {
	// get param
	employeeID := c.Param(employeeIDJSON)
	if employeeID == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, employeeIDStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByEmployeeID(employeeID)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetEmployeeID, err, employeeID)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetEmployeeID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetEmployeeID, employeeID)
}

// @Tags 	user
// @Summary get user by account name
// @Accept	application/json
// @Param	token body string true "token"
// @Param	account_name path string true "account name"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router  /api/v1/metadata/user/account-name/:account_name [get]
func GetUserByAccountName(c *gin.Context) {
	// get param
	accountName := c.Param(accountNameJSON)
	if accountName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, accountNameStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByAccountName(accountName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAccountName, err, accountName)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAccountName, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAccountName, accountName)
}

// @Tags 	user
// @Summary get user by loginName
// @Accept	application/json
// @Param	token body string true "token"
// @Param	login_name path string true "login name"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router  /api/v1/metadata/user/login-name/:login_name [get]
func GetByAccountNameOrEmployeeID(c *gin.Context) {
	// get param
	loginName := c.Param(loginNameJSON)
	if loginName == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, loginNameJSON)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByAccountNameOrEmployeeID(loginName)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetByAccountNameOrEmployeeID, err, loginName)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetByAccountNameOrEmployeeID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetByAccountNameOrEmployeeID, loginName)
}

// @Tags 	user
// @Summary get user by email
// @Accept	application/json
// @Param	token body string true "token"
// @Param	email path string true "email"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router  /api/v1/metadata/user/email/:email [get]
func GetUserByEmail(c *gin.Context) {
	// get param
	email := c.Param(emailJSON)
	if email == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, emailStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByEmail(email)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetEmail, err, email)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetEmail, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetEmail, email)
}

// @Tags 	user
// @Summary get user by telephone
// @Accept	application/json
// @Param	token body string true "token"
// @Param	telephone path string true "telephone"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router  /api/v1/metadata/user/telephone/:telephone [get]
func GetUserByTelephone(c *gin.Context) {
	// get param
	telephone := c.Param(telephoneJSON)
	if telephone == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, telephoneStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByTelephone(telephone)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetTelephone, err, telephone)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetTelephone, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetTelephone, telephone)
}

// @Tags 	user
// @Summary get user by mobile
// @Accept	application/json
// @Param	token body string true "token"
// @Param	mobile path string true "mobile"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router  /api/v1/metadata/user/mobile/:mobile [get]
func GetUserByMobile(c *gin.Context) {
	// get param
	mobile := c.Param(mobileJSON)
	if mobile == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, mobileStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get UserRepo
	err := s.GetByMobile(mobile)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMobile, err, mobile)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMobile, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMobile, mobile)
}

// @Tags 	user
// @Summary get apps by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id path int true "user id"
// @Produce application/json
// @Success 200 {string} string "{"apps": [ {"id": 2,"app_name": "app2","level": 2,"del_flag": 0,"create_time": "2021-10-25T09:46:04.275796+08:00","last_update_time": "2021-12-21T09:15:47.688546+08:00"}]}"
// @Router 	/api/v1/metadata/user/app/:id [get]
func GetAppsByUserID(c *gin.Context) {
	// get param
	idStr := c.Param(userIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, userIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get entity
	err = s.GetAppsByUserID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAppsByUserID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(userAppsStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAppsByUserID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAppsByUserID, id)

}

// @Tags 	user
// @Summary get dbs by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id path int true "user id"
// @Produce application/json
// @Success 200 {string} string "{"dbs": [{"id": 1,"db_name": "db2","cluster_id": 3,"cluster_type": 1,"env_id": 1,"del_flag": 0,"create_time": "2022-01-04T15:08:33.418288+08:00","last_update_time": "2022-01-25T16:17:26.284761+08:00"},}]}"
// @Router 	/api/v1/metadata/user/db/:id [get]
func GetDBsByUserID(c *gin.Context) {
	// get param
	idStr := c.Param(userIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, userIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get entity
	err = s.GetDBsByUserID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetDBsByUserID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(userDBsStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetDBsByUserID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetDBsByUserID, id)

}

// @Tags 	user
// @Summary get middleware clusters by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id path int true "user id"
// @Produce application/json
// @Success 200 {string} string "{"middleware_clusters": [{"id": 1,"cluster_name": "middleware-cluster-1","env_id": 1,"del_flag": 0,"create_time": "2021-11-09T18:06:57.917596+08:00","last_update_time": "2021-11-18T15:39:52.927116+08:00"}]}"
// @Router 	/api/v1/metadata/user/middleware-cluster/:id [get]
func GetMiddlewareClustersByUserID(c *gin.Context) {
	// get param
	idStr := c.Param(userIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, userIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get entity
	err = s.GetMiddlewareClustersByUserID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMiddlewareClustersByUserID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(userMiddlewareClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMiddlewareClustersByUserID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMiddlewareClustersByUserID, id)

}

// @Tags 	user
// @Summary get mysql clusters by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id path int true "user id"
// @Produce application/json
// @Success 200 {string} string "{"mysql_clusters": [{"id": 1,"cluster_name": "mysql-cluster-pmm2","middleware_cluster_id": 0,"env_id": 1"monitor_system_id": 1,"del_flag": 0,"last_update_time": "2021-12-21T09:16:10.750725+08:00","create_time": "2021-09-02T09:02:22.346672+08:00",},]}"
// @Router 	/api/v1/metadata/user/mysql-cluster/:id [get]
func GetMySQLClustersByUserID(c *gin.Context) {
	// get param
	idStr := c.Param(userIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, userIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get entity
	err = s.GetMySQLClustersByUserID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetMySQLClustersByUserID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(userMySQLClustersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetMySQLClustersByUserID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetMySQLClustersByUserID, id)

}

// @Tags 	user
// @Summary get all mysql servers by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id path int true "user id"
// @Produce application/json
// @Success	200 {string} string "{"mysql_servers":[{"port_num":3306,"create_time":"2021-09-02T11:16:06.561525+08:00","last_update_time":"2022-03-01T08:19:09.779365+08:00","cluster_id":1,"server_name":"192-168-10-219","service_name":"192-168-10-219:3306","host_ip":"192.168.10.219","id":1,"deployment_type":1,"version":"5.7","del_flag":0}]}"
// @Router 	/api/v1/metadata/user/all-mysql-server/:id [get]
func GetAllMySQLServersByUserID(c *gin.Context) {
	// get param
	idStr := c.Param(userIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, userIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// get entity
	err = s.GetAllMySQLServersByUserID(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataGetAllMySQLServersByUserID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.MarshalWithFields(userMySQLServersStruct)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataGetAllMySQLServersByUserID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataGetAllMySQLServersByUserID, id)

}

// @Tags 	user
// @Summary add a new user
// @Accept	application/json
// @Param	token body string true "token"
// @Param	user_name 		body string true "user name"
// @Param	department_name body string true "department name"
// @Param	employee_id 	body string true "employee id"
// @Param	telephone 		body string true "user telephone"
// @Param	role 			body int 	true "user role"
// @Param	account_name 	body string true "account name"
// @Param	email 			body string true "email"
// @Param	mobile 			body string true "mobile"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router 	/api/v1/metadata/user [post]
func AddUser(c *gin.Context) {
	var fields map[string]interface{}

	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.UserInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, ok := fields[userNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, userNameStruct)
		return
	}
	_, ok = fields[emailStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, emailStruct)
		return
	}
	_, ok = fields[departmentNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, departmentNameStruct)
		return
	}
	// _, ok = fields[employeeIDStruct]
	// if !ok {
	// 	resp.ResponseNOK(c, message.ErrFieldNotExists, employeeIDStruct)
	// 	return
	// }
	_, ok = fields[accountNameStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, accountNameStruct)
		return
	}
	_, ok = fields[roleStruct]
	if !ok {
		resp.ResponseNOK(c, message.ErrFieldNotExists, roleStruct)
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// insert into middleware
	err = s.Create(fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataAddUser, err, fields[userNameStruct])
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataAddUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataAddUser, fields[userNameStruct])
}

// @Tags 	user
// @Summary update user by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id				path int	true	"user id"
// @Param	user_name 		body string true 	"user name"
// @Param	department_name body string true 	"department name"
// @Param	employee_id 	body string true 	"employee id"
// @Param	telephone 		body string true 	"user telephone"
// @Param	role 			body int 	true 	"user role"
// @Param	account_name 	body string true 	"account name"
// @Param	email 			body string true 	"email"
// @Param	mobile 			body string true 	"mobile"
// @Param 	del_flag 		body int	false	"delete flag"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router 	/api/v1/metadata/user/update/:id [post]
func UpdateUserByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(envIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, envIDJSON)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}
	// unmarshal data
	fields, err = common.UnmarshalToMapWithStructTag(data, &metadata.UserInfo{}, constant.DefaultMiddlewareTag)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, err)
		return
	}
	_, userNameExists := fields[userNameStruct]
	_, departmentNameExists := fields[departmentNameStruct]
	_, employeeIDExists := fields[employeeIDStruct]
	_, accountNameExists := fields[accountNameStruct]
	_, emailExists := fields[emailStruct]
	_, mobileExists := fields[mobileStruct]
	_, telephoneExists := fields[telephoneStruct]
	_, roleExists := fields[roleStruct]
	_, delFlagExists := fields[envDelFlagStruct]
	if !userNameExists && !departmentNameExists && !employeeIDExists && !accountNameExists && !emailExists && !telephoneExists && !roleExists && !delFlagExists && !mobileExists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, fmt.Sprintf("%s and %s", userNameStruct, envDelFlagStruct))
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// update UserRepo
	err = s.Update(id, fields)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataUpdateUser, err)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err, id)
		return
	}
	// resp
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataUpdateUser, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.DebugMetadataUpdateUser, id)
}

// @Tags 	user
// @Summary delete user by id
// @Accept	application/json
// @Param	token body string true "token"
// @Param	id path int true "user id"
// @Produce application/json
// @Success 200 {string} string "{"users": [{"id": 18,"employee_id": "21213434","account_name": "kf-Tom","mobile": "18088888888","role": 2,"user_name": "Tom","department_name": "kf","email": "test@test.com.cn","telephone": "02188888888","del_flag": 0,"create_time": "2022-03-07T15:56:32.277857+08:00","last_update_time": "2022-03-07T15:56:32.277857+08:00"}]}"
// @Router 	/api/v1/metadata/user/delete/:id [get]
func DeleteUserByID(c *gin.Context) {
	var fields map[string]interface{}

	// get params
	idStr := c.Param(envIDJSON)
	if idStr == constant.EmptyString {
		resp.ResponseNOK(c, message.ErrFieldNotExists, envIDJSON)
		return
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.ResponseNOK(c, message.ErrTypeConversion, errors.Trace(err))
		return
	}
	// init service
	s := metadata.NewUserServiceWithDefault()
	// update entities
	err = s.Delete(id)
	if err != nil {
		resp.ResponseNOK(c, msgmeta.ErrMetadataDeleteUserByID, err, id)
		return
	}
	// marshal service
	jsonBytes, err := s.Marshal()
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, err)
		return
	}
	// response
	jsonStr := string(jsonBytes)
	log.Debug(message.NewMessage(msgmeta.DebugMetadataDeleteUserByID, jsonStr).Error())
	resp.ResponseOK(c, jsonStr, msgmeta.InfoMetadataDeleteUserByID, fields[userNameStruct])
}
