package drawing

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/request"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/response"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/drawing"
	"github.com/defeng-hub/ByOfficeAutomatic/server/utils"
	"github.com/gin-gonic/gin"
)

type ProjectApi struct{}

func (e *ProjectApi) GetAllProject(c *gin.Context) {
	project, err := projectService.AllProject()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(project, "获取成功", c)
}

func (e *ProjectApi) GetProjectById(c *gin.Context) {
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
	obj, err := projectService.GetProjectById(cin.Uint())
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(obj, "获取成功", c)
}

func (e *ProjectApi) GetAllBranch(c *gin.Context) {
	project, err := projectService.AllBrush()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(project, "获取成功", c)
}

func (e *ProjectApi) CreateBranch(c *gin.Context) {
	var brush drawing.Brush
	err := c.ShouldBindJSON(&brush)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验入参
	err = utils.Verify(brush, utils.Rules{"Name": {utils.NotEmpty()}, "Path": {utils.NotEmpty()}, "FontSize": {utils.NotEmpty()}, "FontColor": {utils.NotEmpty()}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectService.CreateBrush(&brush)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

func (e *ProjectApi) CreateProject(c *gin.Context) {
	var project drawing.Project
	err := c.ShouldBindJSON(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验入参
	err = utils.Verify(project, utils.Rules{"Title": {utils.NotEmpty()}, "ImgSrc": {utils.NotEmpty()}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectService.CreateProject(&project)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("添加成功", c)
}

// DeleteProject 传入 json id值
func (e *ProjectApi) DeleteProject(c *gin.Context) {
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
	err = projectService.DeleteProject(cin.Uint())
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (e *ProjectApi) DeleteBranch(c *gin.Context) {
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
	if cin.Uint() == 1 {
		response.FailWithMessage("系统保留笔刷, 不可删除", c)
		return
	}

	err = projectService.DeleteBrush(cin.Uint())
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

func (e *ProjectApi) ChangeBrush(c *gin.Context) {
	var brush drawing.Brush
	err := c.ShouldBindJSON(&brush)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 校验入参
	err = utils.Verify(brush, utils.Rules{"Name": {utils.NotEmpty()}, "Path": {utils.NotEmpty()}, "FontSize": {utils.NotEmpty()}, "FontColor": {utils.NotEmpty()}})
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = projectService.ChangeBrush(&brush)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}
