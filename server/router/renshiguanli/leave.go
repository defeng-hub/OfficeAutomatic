package renshiguanli

import (
	v1 "github.com/defeng-hub/ByOfficeAutomatic/server/api/v1"
	"github.com/defeng-hub/ByOfficeAutomatic/server/middleware"
	"github.com/gin-gonic/gin"
)

type LeaveRouter struct{}

func (LeaveRouter) Init(Router *gin.RouterGroup) {
	R := Router.Group("renshi/").Use(middleware.OperationRecord())
	Api := v1.ApiGroupApp.RenshiguanliApiGroup.LeaveApi
	{
		R.POST("leave/CreateLeaveForm", Api.CreateLeaveForm)
		R.GET("leave/GetMyselfLeaves", Api.GetMyselfLeaves)
		R.GET("leave/GetDaichuliLeaves", Api.GetDaichuliLeaves)
		R.POST("leave/DeleteLeaveForm", Api.DeleteLeaveForm)
	}
}
