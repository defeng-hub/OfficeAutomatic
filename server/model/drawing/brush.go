package drawing

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
)

// Brush 画笔
type Brush struct {
	global.GVA_MODEL
	Path      string `json:"path"`      //画笔路径(ttf 文件路径)
	FontSize  uint   `json:"fontSize"`  //字体大小
	FontColor string `json:"fontColor"` //字体颜色
}

func (Brush) TableName() string {
	return "drawing_brush"
}

// Project 项目管理
type Project struct {
	global.GVA_MODEL
	Title  string `json:"title"`
	ImgSrc string `json:"imgSrc"`
}

func (Project) TableName() string {
	return "drawing_project"
}
