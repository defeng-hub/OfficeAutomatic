package sms

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
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

// SmsProject SMS项目
type SmsProject struct {
	global.GVA_MODEL
	ProjectName string      `json:"project_name,omitempty" gorm:"comment:项目名"`
	Comment     string      `json:"comment,omitempty" gorm:"comment:备注"`
	TemplateId  uint        `json:"template_id,omitempty" gorm:"comment:模板ID"`
	SmsTemplate SmsTemplate `json:"sms_template" gorm:"foreignKey:TemplateId"`
	RunNum      int         `json:"run_num,omitempty" gorm:"comment:运行次数"`
	RowNum      int         `json:"row_num,omitempty" gorm:"comment:条数"`
	ParamNum    int         `json:"param_num,omitempty" gorm:"comment:参数个数"`
}

func (SmsProject) TableName() string {
	return "txyun_sms_project"
}

type SmsProjectRow struct {
	SmsProjectID uint       `json:"sms_project_id,omitempty"`
	SmsProject   SmsProject `json:"sms_project" gorm:"foreignKey:SmsProjectID"`
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
