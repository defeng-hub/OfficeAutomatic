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
// @Router    /renshi/CreateLeaveForm [post]
func (e *LeaveApi) CreateLeaveForm(c *gin.Context) {
	var leaveform renshiguanli.LeaveForm
	err := c.ShouldBindJSON(&leaveform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(leaveform, utils.EnteringLeaveFormVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	leaveform.UserId = utils.GetUserID(c)

	//err = LeaveService.CreateExaCustomer(customer)
	err = leaveService.EnteringLeaveForm(&leaveform)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteLeaveForm
// @Tags      Renshiguanli
// @Summary   删除请假表单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById  true  "请求内容"
// @Success   200   {object}  response.Response{msg=string}  "响应内容"
// @Router    /renshi/DeleteLeaveForm [post]
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

// ChangeLeaveApproval
// @Tags      Renshiguanli
// @Summary   更改审核状态
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      renshiguanli.LeaveForm  true  "请求内容"
// @Success   200   {object}  response.Response{msg=string}  "响应内容"
// @Router    /renshi/ChangeLeaveApproval [post]
func (e *LeaveApi) ChangeLeaveApproval(c *gin.Context) {
	var leaveform renshiguanli.LeaveForm
	err := c.ShouldBindJSON(&leaveform)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(leaveform, utils.EnteringLeaveFormVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = leaveService.ChangeLeaveApproval(&leaveform)
	if err != nil {
		global.GVA_LOG.Error("更改审核权限失败!", zap.Error(err))
		response.FailWithMessage("更改审核权限失败", c)
		return
	}
	response.OkWithMessage("更改审核权限成功", c)
}
