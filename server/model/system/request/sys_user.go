package request

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/system"
)

// Register User register structure
type Register struct {
	Username     string `json:"userName" example:"用户名"`
	Password     string `json:"passWord" example:"密码"`
	NickName     string `json:"nickName" example:"昵称"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
	// 自主添加的
	Sex                 int    `json:"sex" example:"性别 0未选择 1男 2女"`
	Address             string `json:"address" example:"通信地址"`
	Wno                 string `json:"wno" example:"职工号"`
	UserTeachingGradeID uint   `json:"userTeachingGradeID" example:"教学等级ID"`
	JoinCompanyTime     string `json:"joinCompanyTime" example:"加入公司时间"`
	JoinWorkTime        string `json:"joinWorkTime" example:"参加工作时间"`
	Desc0               string `json:"desc0" example:"本职工作单位/职务"`
	Resume              string `json:"resume" example:"个人简历"`
	Desc1               string `json:"desc1" example:"教师技能等级/职务变动情况记录"`
	Desc2               string `json:"desc2" example:"本职工作变动情况记录"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色ID
}

// Modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                                                           // 主键ID
	NickName     string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	Phone        string                `json:"phone"  gorm:"comment:用户手机号"`                                                          // 用户手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                                                                // 角色ID
	Email        string                `json:"email"  gorm:"comment:用户邮箱"`                                                           // 用户邮箱
	HeaderImg    string                `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	SideMode     string                `json:"sideMode"  gorm:"comment:用户侧边主题"`                                                      // 用户侧边主题
	Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                                                           //冻结用户
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`

	// 自主添加的
	Sex                 int    `json:"sex" example:"性别 0未选择 1男 2女"`
	Address             string `json:"address" example:"通信地址"`
	Wno                 string `json:"wno" example:"职工号"`
	UserTeachingGradeID uint   `json:"userTeachingGradeID" example:"教学等级ID"`
	JoinCompanyTime     string `json:"joinCompanyTime" example:"加入公司时间"`
	JoinWorkTime        string `json:"joinWorkTime" example:"参加工作时间"`
	Desc0               string `json:"desc0" example:"本职工作单位/职务"`
	Resume              string `json:"resume" example:"个人简历"`
	Desc1               string `json:"desc1" example:"教师技能等级/职务变动情况记录"`
	Desc2               string `json:"desc2" example:"本职工作变动情况记录"`
}
