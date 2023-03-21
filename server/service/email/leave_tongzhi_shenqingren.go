package email

import (
	"fmt"
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	EmailEngine "github.com/defeng-hub/Go-Email-Template"
)

type LeaveTongzhiShenqingren struct{}
type LeaveTongzhiShenqingrenOption struct {
	name  string // 申请人姓名
	state string // 审核状态
	url   string
}

func (w *LeaveTongzhiShenqingren) DefaultOption(name string, state string, url string) LeaveTongzhiShenqingrenOption {
	return LeaveTongzhiShenqingrenOption{
		name:  name,
		state: state,
		url:   url,
	}
}

func (w *LeaveTongzhiShenqingren) email(op LeaveTongzhiShenqingrenOption) EmailEngine.Email {
	return EmailEngine.Email{
		Body: EmailEngine.Body{
			Name: op.name,
			Intros: []string{
				fmt.Sprintf("您的请假申请已经 <%s> ,请点击下方按钮查看详情", op.state),
			},
			Actions: []EmailEngine.Action{
				{
					Instructions: "点击按钮查看详情",
					Button: EmailEngine.Button{
						Text: "查看详情",
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

func (w *LeaveTongzhiShenqingren) Send(op LeaveTongzhiShenqingrenOption, subject string, to string) {
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
