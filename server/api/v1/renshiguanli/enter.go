package renshiguanli

import "github.com/defeng-hub/ByOfficeAutomatic/server/service"

type ApiGroup struct {
	LeaveApi
}

var (
	leaveService = service.ServiceGroupApp.RenshiguanliServiceGroup.LeaveService
)
