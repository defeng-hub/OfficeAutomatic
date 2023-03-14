package renshiguanli

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/request"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/response"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/renshiguanli"
	"github.com/defeng-hub/ByOfficeAutomatic/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LeaveApi struct{}

// CreateLeaveForm
// @Tags      Renshiguanli
// @Summary   创建请假表单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      renshiguanli.LeaveForm  true  "请求内容"
// @Success   200   {object}  response.Response{msg=string}  "响应内容"
// @Router    /renshi/leave/CreateLeaveForm [post]
func (e *LeaveApi) CreateLeaveForm(c *gin.Context) {
	var leaveform renshiguanli.SubmitLeaveForm
	err := c.ShouldBindJSON(&leaveform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	leaveform.UserId = utils.GetUserID(c)

	err = utils.Verify(leaveform, utils.EnteringLeaveFormVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	//err = LeaveService.CreateExaCustomer(customer)
	err = leaveService.EnteringLeaveForm(&renshiguanli.LeaveForm{
		GVA_MODEL:     leaveform.GVA_MODEL,
		UserId:        leaveform.UserId,
		LeaveType:     renshiguanli.LeaveType(leaveform.LeaveType),
		LeaveContent:  leaveform.LeaveContent,
		BeginTime:     leaveform.BeginTime,
		EndTime:       leaveform.EndTime,
		Image:         leaveform.Image,
		Approval:      0,
		ShenpiUserID:  leaveform.ShenpiUserID,
		ShenpiUser2ID: leaveform.ShenpiUser2ID,
	})
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// GetMyselfLeaves
// @Tags      Renshiguanli
// @Summary   分页获取请假表单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取列表,返回包括列表,总数,页码,每页数量"
// @Router    /renshi/leave/GetMyselfLeaves [get]
func (e *LeaveApi) GetMyselfLeaves(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// page 和 pagesize不能为空
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	leaves, total, err := leaveService.SelectUserLeaves(utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     leaves,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)

}

// DeleteLeaveForm
// @Tags      Renshiguanli
// @Summary   删除请假表单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById  true  "请求内容"
// @Success   200   {object}  response.Response{msg=string}  "响应内容"
// @Router    /renshi/leave/DeleteLeaveForm [post]
func (e *LeaveApi) DeleteLeaveForm(c *gin.Context) {
	var req request.GetById
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = leaveService.DeleteLeaveForm(req)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// GetDaichuliLeaves
// @Tags      Renshiguanli
// @Summary   分页获取 待处理的请假表单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取列表,返回包括列表,总数,页码,每页数量"
// @Router    /renshi/leave/GetDaichuliLeaves [get]
func (e *LeaveApi) GetDaichuliLeaves(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// page 和 pagesize不能为空
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	leaves, total, err := leaveService.SelectDaichuliLeaves(utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     leaves,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)

}

// GetYichuliLeaves
// @Tags      Renshiguanli
// @Summary   分页获取 已处理的请假表单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取列表,返回包括列表,总数,页码,每页数量"
// @Router    /renshi/leave/GetYichuliLeaves [get]
func (e *LeaveApi) GetYichuliLeaves(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// page 和 pagesize不能为空
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	leaves, total, err := leaveService.SelectYichuliLeaves(utils.GetUserID(c), pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     leaves,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)

}

// ChangeLeaveApproval 未完成
// @Tags      Renshiguanli
// @Summary   更改审核状态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      renshiguanli.LeaveForm  true  "请求内容"
// @Success   200   {object}  response.Response{msg=string}  "响应内容"
// @Router    /renshi/leave/ChangeLeaveApproval [post]
func (e *LeaveApi) ChangeLeaveApproval(c *gin.Context) {
	var leaveform renshiguanli.LeaveForm
	err := c.ShouldBindJSON(&leaveform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	approval := leaveform.Approval
	global.GVA_DB.First(&leaveform)

	// 更改审核状态
	leaveform.Approval = approval

	err = utils.Verify(leaveform, utils.EnteringLeaveFormVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetUserID(c)

	// 判断审核人是否是提交者
	if leaveform.ShenpiUserID == uid || leaveform.ShenpiUser2ID == uid {
		err = leaveService.ChangeLeaveApproval(&leaveform)
		if err != nil {
			global.GVA_LOG.Error("更改审核权限失败!", zap.Error(err))
			response.FailWithMessage("更改审核权限失败", c)
			return
		}
		response.OkWithMessage("更改审核权限成功", c)
	} else {
		response.FailWithMessage("审核人不正确", c)
	}
}
