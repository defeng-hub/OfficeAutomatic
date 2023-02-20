package router

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/router/example"
	"github.com/defeng-hub/ByOfficeAutomatic/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
