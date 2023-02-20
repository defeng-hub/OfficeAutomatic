package initialize

import (
	_ "github.com/defeng-hub/ByOfficeAutomatic/server/source/example"
	_ "github.com/defeng-hub/ByOfficeAutomatic/server/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
