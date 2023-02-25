package sms

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"strings"
)

// SmsTemplate Sms模板表
type SmsTemplate struct {
	global.GVA_MODEL
	TemplateId      int64  `json:"template_id" gorm:"comment:模板ID。"`
	International   int64  `json:"international" gorm:"comment:是否国际/港澳台短信，其中0表示国内短信，1表示国际/港澳台短信。"`
	StatusCode      int64  `json:"status_code" gorm:"comment:申请模板状态，其中0表示审核通过，1表示审核中，-1表示审核未通过或审核失败。"`
	ReviewReply     string `json:"review_reply" gorm:"comment:审核回复，审核人员审核后给出的回复，通常是审核未通过的原因。"`
	TemplateName    string `json:"template_name" gorm:"comment:模板名称。"`
	CreateTime      int64  `json:"create_time" gorm:"comment:提交审核时间，UNIX 时间戳（单位：秒）。"`
	TemplateContent string `json:"template_content" gorm:"comment:模板内容。"`
}

func (SmsTemplate) TableName() string {
	return "txyun_sms_template"
}

func (e SmsTemplate) GetParamsNum() int {
	ls := e.TemplateContent
	if strings.Index(ls, "{1}") == -1 {
		return 0
	} else if strings.Index(ls, "{2}") == -1 {
		return 1
	} else if strings.Index(ls, "{3}") == -1 {
		return 2
	} else if strings.Index(ls, "{4}") == -1 {
		return 3
	} else if strings.Index(ls, "{5}") == -1 {
		return 4
	} else if strings.Index(ls, "{6}") == -1 {
		return 5
	} else if strings.Index(ls, "{7}") == -1 {
		return 6
	} else if strings.Index(ls, "{8}") == -1 {
		return 7
	} else if strings.Index(ls, "{9}") == -1 {
		return 8
	} else {
		return 9
	}
}

//  SMS项目---SMS项目---SMS项目

// SmsProject SMS项目
type SmsProject struct {
	global.GVA_MODEL
	ProjectName string `json:"project_name" gorm:"comment:项目名"`
	Comment     string `json:"comment" gorm:"comment:备注"`
	TemplateId  int64  `json:"template_id" gorm:"comment:模板ID"`
	// references 是关联到外表的键
	SmsTemplate SmsTemplate `json:"sms_template" gorm:"foreignKey:TemplateId;references:TemplateId;"`
	RunNum      int         `json:"run_num" gorm:"comment:运行次数"`
	RowNum      int         `json:"row_num" gorm:"comment:条数"`
	ParamNum    int         `json:"param_num" gorm:"comment:参数个数"`
}

func (SmsProject) TableName() string {
	return "txyun_sms_project"
}

type SmsProjectRow struct {
	global.GVA_MODEL
	SmsProjectId uint       `json:"sms_project_id"`
	SmsProject   SmsProject `json:"sms_project" gorm:"foreignKey:SmsProjectId"`
	Phone        string     `json:"phone"`
	Param1       string     `json:"param1"`
	Param2       string     `json:"param2"`
	Param3       string     `json:"param3"`
	Param4       string     `json:"param4"`
	Param5       string     `json:"param5"`
	Param6       string     `json:"param6"`
	Param7       string     `json:"param7"`
	Param8       string     `json:"param8"`
	Param9       string     `json:"param9"`
}

func (SmsProjectRow) TableName() string {
	return "txyun_sms_project_row"
}
