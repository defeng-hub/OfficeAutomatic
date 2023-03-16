package renshiguanli

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/system"
	"time"
)

type ApprovalType int32

const (
	ApprovalType_daishenpi = iota //待审批
	ApprovalType_tongguo          //通过
	ApprovalType_jujue            //拒绝
)

type LeaveType int32

const (
	LeaveType_shijia     = iota //事假
	LeaveType_tiaoxiu           //调休
	LeaveType_bingjia           //病假
	LeaveType_nianjia           //年假
	LeaveType_chanjia           //产假
	LeaveType_peichanjia        //陪产假
	LeaveType_hunjia            //婚假
	LeaveType_lijia             //例假
	LeaveType_sangjia           //丧假
	LeaveType_burujia           //哺乳假
)

type LeaveForm struct {
	global.GVA_MODEL
	UserId       uint           `json:"userId" gorm:"commit:请假人ID"`
	User         system.SysUser `gorm:"foreignKey:UserId"`
	LeaveType    LeaveType      `json:"leaveType" gorm:"commit:请假类型"`
	LeaveContent string         `json:"leaveContent" gorm:"commit:请假事由"`

	//开始时间必须小于结束时间
	//请假时长必须大于等于2h
	//会自动计算请假时长
	BeginTime time.Time    `json:"beginTime" gorm:"commit:请假时间"`
	EndTime   time.Time    `json:"endTime" gorm:"commit:结束时间"`
	Image     string       `json:"image" gorm:"commit:附加图片"`
	Approval  ApprovalType `json:"approval" gorm:"commit:审核状态"`
	Daynum    float32      `json:"daynum" gorm:"commit:请假时长"`

	//审批人
	// 默认审批人是部门总管
	// 第二审批人是老板
	ShenpiUserID  uint           `json:"shenpiUserID" gorm:"审批人ID"`
	ShenpiUser    system.SysUser `gorm:"foreignKey:ShenpiUserID"`
	ShenpiUser2ID uint           `json:"shenpiUser2ID" gorm:"审批人2ID"`
	ShenpiUser2   system.SysUser `gorm:"foreignKey:ShenpiUser2ID"`
}

func (LeaveForm) TableName() string {
	return "renshi_leave_form"
}
