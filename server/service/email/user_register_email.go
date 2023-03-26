package email

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	EmailEngine "github.com/defeng-hub/Go-Email-Template"
)

type RegisterSuccess struct{}

type Option struct {
	name     string
	username string
	wno      string
	phone    string
	url      string
}

func (w *RegisterSuccess) DefaultOption(name string, username string,
	wno string, phone string, url string) Option {
	return Option{
		name:     name,
		username: username,
		wno:      wno,
		phone:    phone,
		url:      url,
	}
}

func (w *RegisterSuccess) Email(op Option) EmailEngine.Email {
	return EmailEngine.Email{
		Body: EmailEngine.Body{
			Name: op.name,
			Intros: []string{
				"欢迎来到博远天合! 我们很高兴您能加入。",
			},
			Dictionary: []EmailEngine.Entry{
				{Key: "用户名", Value: op.username},
				{Key: "职工号", Value: op.wno},
				{Key: "手机号", Value: op.phone},
			},
			Actions: []EmailEngine.Action{
				{
					Instructions: "To get started with boyuantianhe, please click here:",
					Button: EmailEngine.Button{
						Text: "确认您的帐户",
						Link: op.url,
					},
				},
			},
			Outros: []string{
				"需要帮助，还是有其他问题？ 只要回复这封邮件，我们很乐意帮忙。",
			},
		},
	}
}

func (w *RegisterSuccess) Send(op Option, subject string, to string) {
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
