package system

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
