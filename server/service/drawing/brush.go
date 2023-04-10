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
	if brush.Path == "" || brush.FontSize == 0 || brush.FontColor == "" {
		return errors.New("新增画笔入参不正确")
	}
	return global.GVA_DB.Create(brush).Error
}
func (e *ProjectService) DeleteBrush(id uint) error {
	if id == 0 {
		return errors.New("删除画笔入参不正确")
	}
	return global.GVA_DB.Delete(&Brush{}, id).Error
}

func (e *ProjectService) AllProject() (list []*Project, err error) {
	err = global.GVA_DB.Find(&list).Error
	return
}
func (e *ProjectService) CreateProject(project *Project) error {
	if project.BrushId == 0 || project.Title == "" {
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
