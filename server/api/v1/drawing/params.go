package drawing

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/request"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/response"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/drawing"
	"github.com/defeng-hub/ByOfficeAutomatic/server/utils"
	"github.com/gin-gonic/gin"
)

type ParamApi struct{}

func (e *ParamApi) GetParamsByProjectId(c *gin.Context) {
	var cin request.GetById
	err := c.ShouldBindJSON(&cin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验入参
	err = utils.Verify(cin, utils.Rules{"ID": {utils.NotEmpty()}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	params, err := paramService.GetParamsByProjectId(cin.Uint())
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(params, "获取成功", c)
	return
}

func (e *ParamApi) CreateParam(c *gin.Context) {
	var cin drawing.ReqParam
	err := c.ShouldBindJSON(&cin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验入参
	err = utils.Verify(cin, utils.Rules{"Title": {utils.NotEmpty()}, "ProjectId": {utils.NotEmpty()}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = paramService.CreateParam(cin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
	return

}

func (e *ParamApi) ChangeParamBranch(c *gin.Context) {
	var cin drawing.ReqChangeParamBranch
	err := c.ShouldBindJSON(&cin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验入参
	err = utils.Verify(cin, utils.Rules{"ParamId": {utils.NotEmpty()}, "BranchId": {utils.NotEmpty()}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = paramService.ChangeParamBranch(cin.ParamId, cin.BranchId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更改成功", c)
	return
}

func (e *ParamApi) ChangeParam(c *gin.Context) {
	var cin drawing.ReqParam
	err := c.ShouldBindJSON(&cin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = paramService.ChangeParam(cin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更改成功", c)
	return
}

func (e *ParamApi) DeleteParam(c *gin.Context) {
	var cin request.GetById
	err := c.ShouldBindJSON(&cin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验入参
	err = utils.Verify(cin, utils.Rules{"ID": {utils.NotEmpty()}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = paramService.DeleteParam(cin.Uint())
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
	return
}

func (e *ParamApi) CreateImage(c *gin.Context) {
	var cin request.GetById
	err := c.ShouldBindJSON(&cin)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验入参
	err = utils.Verify(cin, utils.Rules{"ID": {utils.NotEmpty()}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = paramService.CreateImage(cin.Uint())
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
	return
}
