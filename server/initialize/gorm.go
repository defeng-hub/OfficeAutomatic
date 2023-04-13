package initialize

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/drawing"
	"os"

	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm 初始化数据库并产生数据库全局变量
// Author SliverHorn
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	default:
		return GormMysql()
	}
}

// RegisterTables 注册数据库表专用
// Author SliverHorn
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		//// 系统模块表
		//system.SysApi{},
		//system.SysUser{},
		//system.UserTeachingGrade{}, //教学等级表
		//system.SysBaseMenu{},
		//system.JwtBlacklist{},
		//system.SysAuthority{},
		//system.SysDictionary{},
		//system.SysOperationRecord{},
		//system.SysAutoCodeHistory{},
		//system.SysDictionaryDetail{},
		//system.SysBaseMenuParameter{},
		//system.SysBaseMenuBtn{},
		//system.SysAuthorityBtn{},
		//system.SysAutoCode{},
		//
		//// 示例模块表
		//example.ExaFile{},
		//example.ExaCustomer{},
		//example.ExaFileChunk{},
		//example.ExaFileUploadAndDownload{},
		//
		//// 腾讯云
		//sms.SmsTemplate{},
		//sms.SmsProject{},
		//sms.SmsProjectRow{},
		//
		//// 请假表
		//renshiguanli.LeaveForm{},

		// 画图表
		//drawing.Project{},
		//drawing.Brush{},
		drawing.ImageDB{},
		//drawing.Param{},
		//drawing.Increment{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
