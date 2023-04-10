package drawing

import "github.com/defeng-hub/ByOfficeAutomatic/server/service"

type ApiGroup struct {
	ProjectApi
}

var (
	projectService = service.ServiceGroupApp.DrawingServiceGroup.ProjectService
)
