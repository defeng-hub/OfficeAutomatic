package renshiguanli

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/system"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestEntering(t *testing.T) {
	leaveform := &LeaveForm{
		GVA_MODEL: global.GVA_MODEL{
			ID:        0,
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
			DeletedAt: gorm.DeletedAt{},
		},
		UserId:        0,
		User:          system.SysUser{},
		LeaveType:     0,
		LeaveContent:  "",
		BeginTime:     time.Now().AddDate(0, 0, -3), // 两天前
		EndTime:       time.Now(),                   // 当前时间
		Image:         "",
		Approval:      0,
		ShenpiUserID:  0,
		ShenpiUser:    system.SysUser{},
		ShenpiUser2ID: 0,
		ShenpiUser2:   system.SysUser{},
	}
	if leaveform.EndTime.Sub(leaveform.BeginTime).Hours() >= 3*24 {
		//键入老板的ID
		leaveform.ShenpiUser2ID = 1
	}
	global.GVA_DB.Create(leaveform)
}
