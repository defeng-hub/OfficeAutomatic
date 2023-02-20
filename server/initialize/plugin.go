package initialize

import (
	"fmt"
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/middleware"
	"github.com/defeng-hub/ByOfficeAutomatic/server/plugin/email"
	"github.com/defeng-hub/ByOfficeAutomatic/server/plugin/tencent_sms"
	"github.com/defeng-hub/ByOfficeAutomatic/server/utils/plugin"
	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		fmt.Println(Plugin[i].RouterPath())
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(Router *gin.Engine) {
	PublicGroup := Router.Group("")
	fmt.Println("无鉴权插件安装==》", PublicGroup)

	PrivateGroup := Router.Group("")
	fmt.Println("鉴权插件安装==》", PrivateGroup)
	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//  添加跟角色挂钩权限的插件 示例 本地示例模式于在线仓库模式注意上方的import 可以自行切换 效果相同
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.Nickname,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))

	// 腾讯云 发送信息组件
	PluginInit(PrivateGroup, tencent_sms.CreateTencentSmsPlug("短信的SecretId", "短信的SecretKey", "短信的 SdkAppId", "短信的 SignName"))

}
