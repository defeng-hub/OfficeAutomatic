package email

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/response"
	"github.com/gin-gonic/gin"
)

type UserEmailApi struct{}

// RegisterSuccess
// @Tags      Email
// @Summary   注册成功邮件
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{msg=string}  "响应内容"
// @Router    /email/user/RegisterSuccess [get]
func (UserEmailApi) RegisterSuccess(c *gin.Context) {
	option := emailRegisterSuccess.DefaultOption("王德丰", "admin",
		"wdf123", "15555150781", "http://www.baidu.com")
	emailRegisterSuccess.Send(option, "测试测试", "15555150781@163.com")
	response.OkWithMessage("success", c)
}
