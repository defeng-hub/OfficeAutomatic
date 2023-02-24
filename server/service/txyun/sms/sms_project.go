package sms

import (
	"fmt"
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	smsmodel "github.com/defeng-hub/ByOfficeAutomatic/server/model/txyun/sms"
)

func (e *TencentSmsService) AllSmsProject() (list interface{}, total int64, err error) {
	var Projects []smsmodel.SmsProject
	db := global.GVA_DB.Model(smsmodel.SmsProject{})

	err = db.Count(&total).Error
	if err != nil {
		return Projects, 0, err
	} else {
		err = db.Find(&Projects).Error
		if err != nil {
			return Projects, 0, err
		}
	}
	return Projects, total, err
}

func (e *TencentSmsService) AddSmsProject(req *smsmodel.AddSmsProjectReq) (*smsmodel.SmsProject, error) {
	db := global.GVA_DB.Model(smsmodel.SmsTemplate{})
	template := smsmodel.SmsTemplate{}
	tx := db.Find(&template, "template_id = ?", req.TemplateId)
	if tx.RowsAffected != 1 {
		return nil, fmt.Errorf("短信模板不存在")
	}

	//	找到了这个模板, 那么就制造smsProject
	db = global.GVA_DB.Model(smsmodel.SmsProject{})
	project := smsmodel.SmsProject{
		ProjectName: req.ProjectName,
		Comment:     req.Comment,
		TemplateId:  template.TemplateId,
		RunNum:      0,
		RowNum:      0,
		ParamNum:    template.GetParamsNum(),
	}
	err := db.Create(&project).Error

	if err != nil {
		return nil, err
	} else {
		return &project, nil
	}
}
