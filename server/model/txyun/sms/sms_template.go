package sms

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
)

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
