package drawing

import (
	v1 "github.com/defeng-hub/ByOfficeAutomatic/server/api/v1"
	"github.com/gin-gonic/gin"
)

type ProjectRouter struct{}

func (ProjectRouter) Init(Router *gin.RouterGroup) {
	R := Router.Group("drawing/")
	Api := v1.ApiGroupApp.DrawingApiGroup.ProjectApi
	{
		R.GET("project/GetAllProject", Api.GetAllProject)
		R.GET("project/GetAllBranch", Api.GetAllBranch)
	}
}
