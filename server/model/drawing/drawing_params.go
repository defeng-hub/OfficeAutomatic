package drawing

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"time"
)

// Param 项目的详细参数之一
type Param struct {
	global.GVA_MODEL
	ProjectId     uint    `json:"projectId"`
	Project       Project `gorm:"foreignKey:ProjectId" json:"project"`
	X             int     `json:"x"`
	Y             int     `json:"y"`
	Text          string  `json:"text"`
	AutoIncrement int     `json:"autoIncrement"` //是否自增  0不自增 1自增填充, 选用1就不会填充本身的Text
	//画笔
	BrushId uint  `json:"brushId" `
	Brush   Brush `gorm:"foreignKey:BrushId" json:"brush"`
}

func (Param) TableName() string {
	return "drawing_params"
}

// Increment 自增参数
type Increment struct {
	ID        uint      `gorm:"primarykey" json:"ID"` // 主键ID
	ParamId   uint      `gorm:"primarykey" json:"paramId"`
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间
	WinNum    int       `json:"winNum"`
}

func (Increment) TableName() string {
	return "drawing_increment"
}
