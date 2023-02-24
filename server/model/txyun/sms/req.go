package sms

type SendSmsRequest struct {
	TemplateId int      `json:"template_id,omitempty"`
	TplPhones  string   `json:"tpl_phones,omitempty"`
	TplParams  []string `json:"tpl_params,omitempty"`
}
