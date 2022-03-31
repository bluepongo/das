package sqladvisor

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/pingcap/errors"
	"github.com/romberli/das/internal/app/sqladvisor"
	"github.com/romberli/das/pkg/message"
	msgadvisor "github.com/romberli/das/pkg/message/sqladvisor"
	"github.com/romberli/das/pkg/resp"
	utilsqladvisor "github.com/romberli/das/pkg/util/sqladvisor"
)

const (
	sqlTextJSON     = "sql_text"
	fingerprintJSON = "fingerprint"
	sqlIDJSON       = "sql_id"
	dbIDJSON        = "db_id"
)

// @Tags sqladvisor
// @Summary get sql fingerprint
// @Accept  application/json
// @Param	sql_text body string true "sql text"
// @Produce application/json
// @Success 200 {string} string "{"fingerprint":"select * from a","sql_text":"select * from a;"}"
// @Router	/api/v1/sqladvisor/fingerprint [get]
func GetFingerprint(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	sqlText, exists := dataMap[sqlTextJSON]
	if !exists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, sqlTextJSON)
		return
	}
	// init service
	service := sqladvisor.NewServiceWithDefault()
	// get fingerprint
	fingerprint := service.GetFingerprint(sqlText)
	respData := map[string]string{sqlTextJSON: sqlText, fingerprintJSON: fingerprint}
	respMessage, err := json.Marshal(respData)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, errors.Trace(err))
		return
	}

	resp.ResponseOK(c, string(respMessage), msgadvisor.InfoSQLAdvisorGetFingerprint, sqlTextJSON, fingerprint)
}

// @Tags sqladvisor
// @Summary get sql id
// @Accept  application/json
// @Param	sql_text body string true "sql text"
// @Produce application/json
// @Success 200 {string} string "{"sql_id":"EE56B94E867DC9D5","sql_text":"select * from a;"}"
// @Router	/api/v1/sqladvisor/sql-id [get]
func GetSQLID(c *gin.Context) {
	// get data
	data, err := c.GetRawData()
	if err != nil {
		resp.ResponseNOK(c, message.ErrGetRawData, errors.Trace(err))
		return
	}

	dataMap := make(map[string]string)
	err = json.Unmarshal(data, &dataMap)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}

	sqlText, exists := dataMap[sqlTextJSON]
	if !exists {
		resp.ResponseNOK(c, message.ErrFieldNotExists, sqlTextJSON)
		return
	}
	// init service
	service := sqladvisor.NewServiceWithDefault()
	// get sql id
	sqlID := service.GetSQLID(sqlText)
	respData := map[string]string{sqlTextJSON: sqlText, sqlIDJSON: sqlID}
	respMessage, err := json.Marshal(respData)
	if err != nil {
		resp.ResponseNOK(c, message.ErrMarshalData, errors.Trace(err))
		return
	}

	resp.ResponseOK(c, string(respMessage), msgadvisor.InfoSQLAdvisorGetSQLID, sqlTextJSON, sqlID)
}

// @Tags sqladvisor
// @Summary get advice
// @Accept  application/json
// @Param	db_id		body int	true "db id"
// @Param	sql_text	body string true "sql text"
// @Produce application/json
// @Success 200 {string} string "[ {  "ID": "B95017DB61875675",  "Fingerprint": "select * from t_meta_db_info where create_time<?",  "Score": 85,  "Sample": "select * from t_meta_db_info where create_time<'2021-01-01'",  "Explain": null,  "HeuristicRules": [    {      "Item": "COL.001",      "Severity": "L1",      "Summary": "不建议使用 SELECT * 类型查询",      "Content": "当表结构变更时，使用 * 通配符选择所有列将导致查询的含义和行为会发生更改，可能导致查询返回更多的数据。",      "Case": "select * from tbl where id=1",      "Position": 0    }  ],  "IndexRules": [    {      "Item": "IDX.001",      "Severity": "L2",      "Summary": "为das库的t_meta_db_info表添加索引",      "Content": "为列create_time添加索引; 由于未开启数据采样，各列在索引中的顺序需要自行调整。",      "Case": "ALTER TABLE `das`.`t_meta_db_info` add index `idx_create_time` (`create_time`) ; ",      "Position": 0    }  ],  "Tables": [    "`soar`.`t_meta_db_info`"  ]} ]"
// @Router	/api/v1/sqladvisor/advise [post]
func Advise(c *gin.Context) {
	var rd *utilsqladvisor.SQLInfo
	// get data
	err := c.ShouldBindJSON(&rd)
	if err != nil {
		resp.ResponseNOK(c, message.ErrUnmarshalRawData, errors.Trace(err))
		return
	}
	// init service
	service := sqladvisor.NewServiceWithDefault()
	advice, err := service.Advise(rd.GetDBID(), rd.GetSQLText())
	if err != nil {
		resp.ResponseNOK(c, msgadvisor.ErrSQLAdvisorAdvice, err, rd.GetDBID(), rd.GetSQLText())
		return
	}

	resp.ResponseOK(c, advice, msgadvisor.InfoSQLAdvisorAdvice, rd.GetDBID(), rd.GetSQLText(), advice)
}
