package email

import (
	"fmt"
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	EmailEngine "github.com/defeng-hub/Go-Email-Template"
)

type ResetUser struct{}
type ResetUserOption struct {
	name     string
	username string
	phone    string
	url      string
}

func (w *ResetUser) DefaultOption(name string, username string, url string) ResetUserOption {
	return ResetUserOption{
		name:     name,
		username: username,
		url:      url,
	}
}

func (w *ResetUser) Email(op ResetUserOption) EmailEngine.Email {
	return EmailEngine.Email{
		Body: EmailEngine.Body{
			Name: op.name,
			Intros: []string{
				fmt.Sprintf("您的账户: %s 密码已经被重置, 默认为123456。请及时登录查看并修改", op.username),
			},
			Actions: []EmailEngine.Action{
				{
					Instructions: "点击按钮查看详情",
					Button: EmailEngine.Button{
						Text: "确认您的帐户",
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

func (w *ResetUser) Send(op ResetUserOption, subject string, to string) {
	email := w.Email(op)
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
