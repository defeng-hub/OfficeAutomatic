package renshiguanli

import "github.com/defeng-hub/ByOfficeAutomatic/server/service"

type ApiGroup struct {
	LeaveApi
}

var (
	userService  = service.ServiceGroupApp.SystemServiceGroup.UserService
	leaveService = service.ServiceGroupApp.RenshiguanliServiceGroup.LeaveService

	emailLeaveTongzhiShenheren   = service.ServiceGroupApp.EmailServiceGroup.LeaveTongzhiShenheren   // 一个实例
	emailLeaveTongzhiShenqingren = service.ServiceGroupApp.EmailServiceGroup.LeaveTongzhiShenqingren // 一个实例
)
