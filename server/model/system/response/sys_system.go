package response

import "github.com/defeng-hub/ByOfficeAutomatic/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
