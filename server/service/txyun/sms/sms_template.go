package sms

import (
	"fmt"
	"strconv"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"go.uber.org/zap"

	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/request"
	smsmodel "github.com/defeng-hub/ByOfficeAutomatic/server/model/txyun/sms"
)

type TencentSmsService struct{}

// SendSms 发送短信
func (e *TencentSmsService) SendSms(tplId int, phoneNumbers []string, tplParams []string) error {

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(phoneNumbers)

	//app id
	request.SmsSdkAppId = common.StringPtr("1400640189")

	// 签名
	request.SignName = common.StringPtr("高能博远君")

	strId := strconv.Itoa(tplId)
	//模板id
	request.TemplateId = common.StringPtr(strId)

	//模板参数
	request.TemplateParamSet = common.StringPtrs(tplParams)
	response, err := client.SendSms(request)
	if err != nil {
		zap.L().Error("An API error has returned: %s", zap.Error(err), zap.String("response", response.ToJsonString()))
		return err
	}
	return nil
}

// UpdateTemplates 更新数据库的短信列表
func (e *TencentSmsService) UpdateTemplates() error {
	// 实例化一个client选项，可选的，没有特殊需求可以跳过
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	// 实例化要请求产品的client对象,clientProfile是可选的
	client, _ := sms.NewClient(credential, "ap-beijing", cpf)
	// 实例化一个请求对象,每个接口都会对应一个request对象
	request := sms.NewDescribeSmsTemplateListRequest()
	international := uint64(0)
	request.International = &international
	request.Limit = &international
	request.Offset = &international

	// 返回的resp是一个DescribeSmsTemplateListResponse的实例，与请求对象对应
	response, err := client.DescribeSmsTemplateList(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return err
	}
	if err != nil {
		return err
	}

	// 删除全部表
	global.GVA_DB.Where("1 = 1").Delete(&smsmodel.SmsTemplate{})
	var smsSet []smsmodel.SmsTemplate
	for _, obj := range response.Response.DescribeTemplateStatusSet {
		sm := smsmodel.SmsTemplate{
			TemplateId:      int64(*obj.TemplateId),
			International:   int64(*obj.International),
			StatusCode:      int64(*obj.StatusCode),
			ReviewReply:     *obj.ReviewReply,
			TemplateName:    *obj.TemplateName,
			CreateTime:      int64(*obj.CreateTime),
			TemplateContent: *obj.TemplateContent,
		}
		smsSet = append(smsSet, sm)
	}
	global.GVA_DB.Model(&smsmodel.SmsTemplate{}).Create(&smsSet)

	return nil
}

// GetSmsInfoList 分页获取腾讯云信息列表
func (e *TencentSmsService) GetSmsInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(smsmodel.SmsTemplate{})

	var List []smsmodel.SmsTemplate
	err = db.Count(&total).Error
	if err != nil {
		return List, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Find(&List).Error
	}
	return List, total, err
}
