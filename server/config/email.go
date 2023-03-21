package config

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // 收件人:多个以英文逗号分隔
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // 收件人
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址
	IsSSL    bool   `mapstructure:"is-ssl" json:"is-ssl" yaml:"is-ssl"`       // 是否SSL
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // 密钥
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // 昵称
}

type EmailSendSetting struct {
	//User 用户相关
	// 1、新增用户时 发送邮件通知该用户
	UserRegister bool `mapstructure:"user-register" json:"user-register" yaml:"user-register"`
	// 2、重置用户密码时 发送邮件通知该用户
	UserReset bool `mapstructure:"user-reset" json:"user-reset" yaml:"user-reset"`

	//leave 请假相关
	// 一、用户申请请假时 发送邮件通知审核人
	LeaveInformShenheren bool `mapstructure:"leave-inform-shenheren" json:"leave-inform-shenheren" yaml:"leave-inform-shenheren"`
	// 二、请假审核完毕后 发送邮件通知申请人
	LeaveInformShenqingren bool `mapstructure:"leave-inform-shenqingren" json:"leave-inform-shenqingren" yaml:"leave-inform-shenqingren"`
}
