package drawing

import "github.com/defeng-hub/ByOfficeAutomatic/server/service"

type ApiGroup struct {
	ProjectApi
	ParamApi
}

var (
	projectService = service.ServiceGroupApp.DrawingServiceGroup.ProjectService
	paramService   = service.ServiceGroupApp.DrawingServiceGroup.ParamService
)
