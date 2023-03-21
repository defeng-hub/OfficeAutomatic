package email

import (
	"fmt"
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	EmailEngine "github.com/defeng-hub/Go-Email-Template"
)

type LeaveTongzhiShenheren struct{}
type LeaveTongzhiShenherenOption struct {
	name      string // 审核人姓名
	username  string // 发起人姓名
	userphone string // 发起人手机号
	url       string
}

func (w *LeaveTongzhiShenheren) DefaultOption(name string, username string, userphone string, url string) LeaveTongzhiShenherenOption {
	return LeaveTongzhiShenherenOption{
		name:      name,
		username:  username,
		userphone: userphone,
		url:       url,
	}
}

func (w *LeaveTongzhiShenheren) email(op LeaveTongzhiShenherenOption) EmailEngine.Email {
	return EmailEngine.Email{
		Body: EmailEngine.Body{
			Name: op.name,
			Intros: []string{
				fmt.Sprintf("您收到了用户: %s(%s) 的请假申请, 来处理一下把", op.username, op.userphone),
			},
			Actions: []EmailEngine.Action{
				{
					Instructions: "点击按钮查看详情",
					Button: EmailEngine.Button{
						Text: "去处理",
						Link: op.url,
					},
				},
			},
			Outros: []string{
				"需要帮助，或者存在其他问题? 只要回复这封邮件，我们很乐意帮忙。",
			},
		},
	}
}

func (w *LeaveTongzhiShenheren) Send(op LeaveTongzhiShenherenOption, subject string, to string) {
	email := w.email(op)
	err := EmailEngine.SetDefaultEmailValues(&email)
	if err != nil {
		return
	}
	html, err := hermes.GenerateHTML(email)
	if err != nil {
		return
	}
	text, err := hermes.GeneratePlainText(email)
	if err != nil {
		return
	}
	Auth.Server = global.GVA_CONFIG.Email.Host
	Auth.Port = global.GVA_CONFIG.Email.Port
	Auth.SMTPUser = global.GVA_CONFIG.Email.From       //用户名
	Auth.SMTPPassword = global.GVA_CONFIG.Email.Secret //密码
	Auth.SenderEmail = global.GVA_CONFIG.Email.From    //发件人地址
	Auth.SenderIdentity = global.GVA_CONFIG.Email.Nickname

	err = EmailEngine.Send(Auth, EmailEngine.EmailSendOptions{
		To:      to,
		Subject: subject}, html, text)

	if err != nil {
		return
	}
}
