package email

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	EmailEngine "github.com/defeng-hub/Go-Email-Template"
)

var Auth EmailEngine.EmailSmtpAuthentication
var hermes EmailEngine.Hermes

type ServiceGroup struct {
	RegisterSuccess // 注册用户 发送邮件
	ResetUser       // 重置用户密码  发送邮件
	LeaveTongzhiShenheren
	LeaveTongzhiShenqingren
}

func init() {
	//Auth.
	Auth.Server = global.GVA_CONFIG.Email.Host
	Auth.Port = global.GVA_CONFIG.Email.Port
	Auth.SMTPUser = global.GVA_CONFIG.Email.From       //用户名
	Auth.SMTPPassword = global.GVA_CONFIG.Email.Secret //密码
	Auth.SenderEmail = global.GVA_CONFIG.Email.From    //发件人地址
	Auth.SenderIdentity = global.GVA_CONFIG.Email.Nickname
}

func init() {
	hermes = EmailEngine.Hermes{
		Theme:         new(EmailEngine.Default),
		TextDirection: EmailEngine.TDLeftToRight,
		Product: EmailEngine.Product{
			Name: "博远天合教育",
			Link: "https://www.bythedu.com/",
			//Logo:        "https://pic.bythedu.com/gofile/%E9%80%8F%E6%98%8E_logo_2023-02-12_20-39-27.png",
			Copyright:   "Copyright © 2023 boyuantianhe. All rights reserved.",
			TroubleText: "如果你无法点击 '{ACTION}' 按钮，请复制下面的URL到你的浏览器中访问",
		},
	}
	err := EmailEngine.SetDefaultHermesValues(&hermes)
	if err != nil {
		return
	}

}
