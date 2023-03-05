package service

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/service/example"
	"github.com/defeng-hub/ByOfficeAutomatic/server/service/renshiguanli"
	"github.com/defeng-hub/ByOfficeAutomatic/server/service/system"
	"github.com/defeng-hub/ByOfficeAutomatic/server/service/txyun"
)

type ServiceGroup struct {
	SystemServiceGroup       system.ServiceGroup
	ExampleServiceGroup      example.ServiceGroup
	TxyunServiceGroup        txyun.ServiceGroup
	RenshiguanliServiceGroup renshiguanli.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
