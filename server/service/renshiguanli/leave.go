package renshiguanli

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/request"
	. "github.com/defeng-hub/ByOfficeAutomatic/server/model/renshiguanli"
)

type LeaveService struct{}

func (e *LeaveService) EnteringLeaveForm(leaveform *LeaveForm) error {
	//EnteringLeaveFormVerify 验证入参了
	if leaveform.EndTime.Sub(leaveform.BeginTime).Hours() >= 3*24 {
		//键入老板的ID
		leaveform.ShenpiUser2ID = 1
	}
	return global.GVA_DB.Create(leaveform).Error
}

func (e *LeaveService) DeleteLeaveForm(req request.GetById) error {
	return global.GVA_DB.Delete(&LeaveForm{}, req.ID).Error
}

func (e *LeaveService) ChangeLeaveApproval(leaveform *LeaveForm) error {
	//	ApprovalType_daishenpi = iota //待审批 0
	//	ApprovalType_tongguo          //通过 1
	//	ApprovalType_jujue            //拒绝 2

	// 只更新approval (审核状态)
	return global.GVA_DB.Model(&leaveform).
		Select("approval").Updates(map[string]interface{}{"approval": leaveform.Approval}).Error
}

func (e *LeaveService) SelectUserLeaves(userId uint) (res []*LeaveForm, err error) {
	err = global.GVA_DB.Where("user_id = ?", userId).Find(&res).Error
	return res, err
}
