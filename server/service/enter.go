package service

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/service/example"
	"github.com/defeng-hub/ByOfficeAutomatic/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
