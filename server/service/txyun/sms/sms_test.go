package sms

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/core"
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/initialize"
	"go.uber.org/zap"
	"testing"
)

func TestFunc(t *testing.T) {
	global.GVA_VP = core.Viper("config.yaml") // 初始化Viper
	initialize.OtherInit()
	global.GVA_LOG = core.Zap() // 初始化zap日志库
	zap.ReplaceGlobals(global.GVA_LOG)
	global.GVA_DB = initialize.Gorm() // gorm连接数据库
	ls := TencentSmsService{}
	ls.GetTemplates()
}
