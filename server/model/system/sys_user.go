package system

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID                uuid.UUID         `json:"uuid" gorm:"index;comment:用户UUID"`                                                     // 用户UUID
	Username            string            `json:"userName" gorm:"index;comment:用户登录名"`                                                  // 用户登录名
	Password            string            `json:"-"  gorm:"comment:用户登录密码"`                                                             // 用户登录密码
	NickName            string            `json:"nickName" gorm:"default:系统用户;comment:用户昵称"`                                            // 用户昵称
	SideMode            string            `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                          // 用户侧边主题
	HeaderImg           string            `json:"headerImg" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor           string            `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                           // 基础颜色
	ActiveColor         string            `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                      // 活跃颜色
	AuthorityId         uint              `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                        // 用户角色ID
	Authority           SysAuthority      `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities         []SysAuthority    `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone               string            `json:"phone"  gorm:"comment:用户手机号"`                     // 用户手机号
	Email               string            `json:"email"  gorm:"comment:用户邮箱"`                      // 用户邮箱
	Enable              int               `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户是否被冻结 1正常 2冻结
	Sex                 int               `json:"sex" gorm:"default:0;commit:性别 0未选择 1男 2女"`
	Address             string            `json:"address" gorm:"commit:通信地址"`
	Wno                 string            `json:"wno" gorm:"commit:职工号"`
	UserTeachingGradeID uint              `json:"userTeachingGradeID" gorm:"default:1;comment:教学等级ID"`
	UserTeachingGrade   UserTeachingGrade `json:"userTeachingGrade" gorm:"foreignKey:UserTeachingGradeID;comment:教学等级"`
	JoinCompanyTime     string            `json:"joinCompanyTime" gorm:"commit:加入公司时间"`
	JoinWorkTime        string            `json:"joinWorkTime" gorm:"commit:参加工作时间"`
	Desc0               string            `json:"desc0" gorm:"size:200;commit:本职工作单位/职务"`
	Resume              string            `json:"resume" gorm:"size:3000;commit:个人简历"`
	Desc1               string            `json:"desc1" gorm:"size:3000;commit:教师技能等级/职务变动情况记录"`
	Desc2               string            `json:"desc2" gorm:"size:3000;commit:本职工作变动情况记录"`
}

func (SysUser) TableName() string {
	return "sys_users"
}

// UserTeachingGrade 教学等级表
type UserTeachingGrade struct {
	global.GVA_MODEL
	Name string  `json:"title" gorm:"commit:教学等级"`
	Wage float64 `json:"wage" gorm:"commit:时薪"`
}

func (UserTeachingGrade) TableName() string {
	return "sys_user_teaching_grade"
}
