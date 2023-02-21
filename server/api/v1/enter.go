package v1

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/api/v1/example"
	"github.com/defeng-hub/ByOfficeAutomatic/server/api/v1/system"
	"github.com/defeng-hub/ByOfficeAutomatic/server/api/v1/txyun"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	TxyunApiGroup   txyun.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
