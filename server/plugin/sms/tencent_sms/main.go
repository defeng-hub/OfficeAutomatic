package sms

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/plugin/sms/tencent_sms/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/plugin/sms/tencent_sms/router"
	"github.com/gin-gonic/gin"
)

type tencentSmsPlugin struct {
}

func CreateTencentSmsPlug(SecretId, SecretKey, SdkAppId, SignName string) *tencentSmsPlugin {
	global.GlobalConfig.SecretId = SecretId
	global.GlobalConfig.SecretKey = SecretKey
	global.GlobalConfig.SdkAppId = SdkAppId
	global.GlobalConfig.SignName = SignName
	return &tencentSmsPlugin{}
}

func (*tencentSmsPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitTencentSmsRouter(group)
}

func (*tencentSmsPlugin) RouterPath() string {
	return "tencentSms"
}
