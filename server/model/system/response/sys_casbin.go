package response

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
