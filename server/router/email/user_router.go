package email

import (
	v1 "github.com/defeng-hub/ByOfficeAutomatic/server/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (UserRouter) Init(Router *gin.RouterGroup) {
	//.Use(middleware.OperationRecord()) 假如操作日志
	R := Router.Group("email/user/")
	Api := v1.ApiGroupApp.EmailApiGroup.UserEmailApi
	{
		R.GET("RegisterSuccess", Api.RegisterSuccess)
	}
}
