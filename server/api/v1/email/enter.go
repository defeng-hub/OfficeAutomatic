package email

import "github.com/defeng-hub/ByOfficeAutomatic/server/service"

type ApiGroup struct {
	UserEmailApi   // user模块邮件推送API
	SystemEmailApi // 系统模块邮件推送API
}

var (
	emailRegisterSuccess = service.ServiceGroupApp.EmailServiceGroup.RegisterSuccess // 一个实例
)
