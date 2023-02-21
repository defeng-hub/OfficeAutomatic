package txyun

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/service/txyun/oss"
	"github.com/defeng-hub/ByOfficeAutomatic/server/service/txyun/sms"
)

//import "github.com/defeng-hub/ByOfficeAutomatic/server/service/txyun/sms"

type ServiceGroup struct {
	sms.TencentSmsService
	oss.TencentOssService
}
