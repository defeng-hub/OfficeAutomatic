package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/sms/tencent_sms/global"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20210111"
	"go.uber.org/zap"
)

type TencentSmsService struct{}

// SendSms 发送短信
func (e *TencentSmsService) SendSms(tplId string, phoneNumbers, tplParams []string) error {
	credential := common.NewCredential(
		global.GlobalConfig.SecretId,
		global.GlobalConfig.SecretKey,
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "sms.tencentcloudapi.com"
	client, _ := sms.NewClient(credential, "ap-nanjing", cpf)
	request := sms.NewSendSmsRequest()

	request.PhoneNumberSet = common.StringPtrs(phoneNumbers)
	//app id
	request.SmsSdkAppId = common.StringPtr(global.GlobalConfig.SdkAppId)
	// 签名
	request.SignName = common.StringPtr(global.GlobalConfig.SignName)
	//模板id
	request.TemplateId = common.StringPtr(tplId)
	//模板参数
	request.TemplateParamSet = common.StringPtrs(tplParams)
	response, err := client.SendSms(request)
	if err != nil {
		zap.L().Error("An API error has returned: %s", zap.Error(err), zap.String("response", response.ToJsonString()))
		return err
	}
	return nil
}
