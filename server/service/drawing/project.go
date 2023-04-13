package drawing

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	. "github.com/defeng-hub/ByOfficeAutomatic/server/model/drawing"
	"github.com/pkg/errors"
)

//type BrushService struct{}

type ProjectService struct{}

func (e *ProjectService) AllBrush() (list []*Brush, err error) {
	err = global.GVA_DB.Find(&list).Error
	return
}
func (e *ProjectService) CreateBrush(brush *Brush) error {
	if brush.Name == "" || brush.Path == "" || brush.FontSize == 0 || brush.FontColor == "" {
		return errors.New("新增画笔入参不正确")
	}
	return global.GVA_DB.Create(brush).Error
}
func (e *ProjectService) DeleteBrush(id uint) error {
	if id == 0 {
		return errors.New("删除画笔入参不正确")
	}
	return global.GVA_DB.Unscoped().Delete(&Brush{}, id).Error
}

func (e *ProjectService) ChangeBrush(brush *Brush) error {
	err := global.GVA_DB.First(&Brush{}, brush.ID).Error
	if err != nil {
		return errors.Wrap(err, "找不到当前画笔")
	}

	return global.GVA_DB.First(&Brush{}, brush.ID).Updates(
		map[string]interface{}{"name": brush.Name, "font_size": brush.FontSize, "font_color": brush.FontColor}).Error
}

func (e *ProjectService) AllProject() (list []*Project, err error) {
	err = global.GVA_DB.Find(&list).Error
	return
}
func (e *ProjectService) CreateProject(project *Project) error {
	if project.ImgSrc == "" || project.Title == "" {
		return errors.New("新增项目入参不正确")
	}
	return global.GVA_DB.Create(project).Error
}
func (e *ProjectService) DeleteProject(id uint) error {
	if id == 0 {
		return errors.New("删除项目入参不正确")
	}
	return global.GVA_DB.Delete(&Project{}, id).Error
}

func (e *ProjectService) GetProjectById(id uint) (obj *Project, err error) {
	if id == 0 {
		return nil, errors.New("删除项目入参不正确")
	}
	err = global.GVA_DB.First(&obj, id).Error
	return
}
