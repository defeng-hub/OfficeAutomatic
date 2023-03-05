package router

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/router/example"
	"github.com/defeng-hub/ByOfficeAutomatic/server/router/renshiguanli"
	"github.com/defeng-hub/ByOfficeAutomatic/server/router/system"
	"github.com/defeng-hub/ByOfficeAutomatic/server/router/txyun"
)

type RouterGroup struct {
	System       system.RouterGroup
	Example      example.RouterGroup
	Txyun        txyun.RouterGroup
	Renshiguanli renshiguanli.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
