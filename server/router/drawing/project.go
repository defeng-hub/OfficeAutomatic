package drawing

import (
	v1 "github.com/defeng-hub/ByOfficeAutomatic/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct{}

func (ProjectRouter) Init(Router *gin.RouterGroup) {
	R := Router.Group("drawing/")
	Api := v1.ApiGroupApp.DrawingApiGroup.ProjectApi
	paramApi := v1.ApiGroupApp.DrawingApiGroup.ParamApi
	{
		R.GET("project/GetAllProject", Api.GetAllProject)
		R.POST("project/GetProjectById", Api.GetProjectById)
		R.POST("project/CreateProject", Api.CreateProject)
		R.POST("project/CreateBranch", Api.CreateBranch)
		R.POST("project/DeleteBranch", Api.DeleteBranch)
		R.GET("project/GetAllBranch", Api.GetAllBranch)
	}
	{
		// param
		R.POST("project/GetParamsByProjectId", paramApi.GetParamsByProjectId)
		R.POST("project/CreateParam", paramApi.CreateParam)
		R.POST("project/ChangeParamBranch", paramApi.ChangeParamBranch)
		R.POST("project/ChangeParam", paramApi.ChangeParam)
		R.POST("project/DeleteParam", paramApi.DeleteParam)
		R.POST("project/CreateImage", paramApi.CreateImage)
	}
}
