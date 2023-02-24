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
		err = db.Preload("SmsTemplate").Find(&Projects).Error
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

func (e *TencentSmsService) DelSmsProject(req *smsmodel.SmsProjectIdReq) (*smsmodel.SmsProject, error) {
	db := global.GVA_DB.Model(smsmodel.SmsProject{})
	project := smsmodel.SmsProject{}
	tx := db.Where("id = ?", req.Id).Find(&project)
	if tx.RowsAffected == 0 {
		return nil, fmt.Errorf("模板ID不存在")
	}

	tx = db.Where("id = ?", req.Id).Delete(&smsmodel.SmsProject{})
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &project, nil
}

// SmsProjectRows 下边是project row 的service
// SmsProjectRows 通过项目ID拿到项目的全部行数据
func (e *TencentSmsService) SmsProjectRows(req *smsmodel.SmsProjectIdReq) (
	[]*smsmodel.SmsProjectRow, error) {

	return nil, nil
}

func (e *TencentSmsService) DelSmsProjectRow(req *smsmodel.SmsProjectRowIdReq) (
	*smsmodel.SmsProjectRow, error) {
	return nil, nil
}

func (e *TencentSmsService) AddSmsProjectRow(req *smsmodel.AddSmsProjectRowReq) (
	*smsmodel.SmsProjectRow, error) {
	row := smsmodel.SmsProjectRow{
		SmsProjectID: req.SmsProjectID,
		Phone:        req.Phone,
		Param1:       req.Param1,
		Param2:       req.Param2,
		Param3:       req.Param3,
		Param4:       req.Param4,
		Param5:       req.Param5,
		Param6:       req.Param6,
		Param7:       req.Param7,
		Param8:       req.Param8,
		Param9:       req.Param9,
	}
	db := global.GVA_DB.Model(smsmodel.SmsProjectRow{})
	tx := db.Create(&row)
	if tx.Error != nil && tx.RowsAffected != 0 {
		return nil, tx.Error
	}
	return &row, nil
}

func (e *TencentSmsService) UpdateSmsProjectRow(req *smsmodel.SmsProjectRow) (
	*smsmodel.SmsProjectRow, error) {
	return nil, nil
}
