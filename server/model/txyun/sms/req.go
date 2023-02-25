package sms

type SendSmsRequest struct {
	TemplateId int      `json:"template_id,omitempty"`
	TplPhones  string   `json:"tpl_phones,omitempty"`
	TplParams  []string `json:"tpl_params,omitempty"`
}

//SMS Project
type AddSmsProjectReq struct {
	ProjectName string `json:"project_name,omitempty" gorm:"comment:项目名"`
	Comment     string `json:"comment,omitempty" gorm:"comment:备注"`
	TemplateId  uint   `json:"template_id,omitempty" gorm:"comment:模板ID"`
}
type SmsProjectIdReq struct {
	Id uint `json:"id"`
}

type SmsProjectRowsPageReq struct {
	ID       uint   `json:"id"`
	Page     int    `json:"page" form:"page"`          // 页码
	PageSize int    `json:"page_size" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`    //关键字
}

//SMS Project Row
type SmsProjectRowIdReq struct {
	Id uint `json:"id"`
}
type AddSmsProjectRowReq struct {
	SmsProjectId uint   `json:"sms_project_id"`
	Phone        string `json:"phone"`
	Param1       string `json:"param1"`
	Param2       string `json:"param2"`
	Param3       string `json:"param3"`
	Param4       string `json:"param4"`
	Param5       string `json:"param5"`
	Param6       string `json:"param6"`
	Param7       string `json:"param7"`
	Param8       string `json:"param8"`
	Param9       string `json:"param9"`
}
