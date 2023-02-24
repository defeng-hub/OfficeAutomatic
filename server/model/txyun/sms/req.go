package sms

type SendSmsRequest struct {
	TemplateId int      `json:"template_id,omitempty"`
	TplPhones  string   `json:"tpl_phones,omitempty"`
	TplParams  []string `json:"tpl_params,omitempty"`
}

type AddSmsProjectReq struct {
	ProjectName string `json:"project_name,omitempty" gorm:"comment:项目名"`
	Comment     string `json:"comment,omitempty" gorm:"comment:备注"`
	TemplateId  uint   `json:"template_id,omitempty" gorm:"comment:模板ID"`
}
