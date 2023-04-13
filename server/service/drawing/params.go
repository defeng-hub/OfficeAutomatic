package drawing

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	. "github.com/defeng-hub/ByOfficeAutomatic/server/model/drawing"
	"github.com/defeng-hub/ByOfficeAutomatic/server/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"os"
	"strconv"
	"strings"
)

type ParamService struct{}

func (e *ParamService) GetParamsByProjectId(id uint) (params []*Param, err error) {
	err = global.GVA_DB.Preload("Brush").Preload("Increment").Where("project_id = ?", id).Find(&params).Error
	return
}

func (e *ParamService) CreateParam(reqparam ReqParam) error {
	if reqparam.AutoIncrementBool == true { //	自增
		return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
			param := &Param{
				ProjectId:     reqparam.ProjectId,
				Title:         reqparam.Title,
				X:             reqparam.X,
				Y:             reqparam.Y,
				Text:          "",
				AutoIncrement: 1,
				BrushId:       1,
			}
			if res := tx.Create(param); res.Error != nil || res.RowsAffected == 0 {
				// 返回任何错误都会回滚事务
				return res.Error
			}

			increment := Increment{
				ParamId: param.ID,
				WinNum:  reqparam.WinNum,
			}
			if err := tx.Create(&increment).Error; err != nil {
				// 返回任何错误都会回滚事务
				return err
			}

			return nil
		})
	} else { // 不自增
		return global.GVA_DB.Create(&Param{
			ProjectId:     reqparam.ProjectId,
			Title:         reqparam.Title,
			X:             reqparam.X,
			Y:             reqparam.Y,
			Text:          reqparam.Text,
			AutoIncrement: 0,
			BrushId:       1,
		}).Error
	}

}

func (e *ParamService) ChangeParamBranch(paramId uint, branchId uint) (err error) {
	if err = global.GVA_DB.First(&Brush{}, branchId).Error; err != nil {
		return err
	}
	if err = global.GVA_DB.First(&Param{}, paramId).Update("brush_id", branchId).Error; err != nil {
		return err
	}
	return nil
}

func (e *ParamService) ChangeParam(reqparam ReqParam) error {
	if reqparam.AutoIncrementBool == true { //	自增
		//目前只会改 x y坐标
		return global.GVA_DB.First(&Param{}, reqparam.ID).Updates(
			map[string]interface{}{"x": reqparam.X, "y": reqparam.Y}).Error
	} else {
		return global.GVA_DB.First(&Param{}, reqparam.ID).Updates(
			map[string]interface{}{"text": reqparam.Text, "x": reqparam.X, "y": reqparam.Y}).Error
	}
}

func (e *ParamService) DeleteParam(id uint) error {
	return global.GVA_DB.Unscoped().Delete(&Param{}, id).Error
}

// CreateImage 传入一个项目ID, 返回一个图片
// 绘图, 巨牛逼!!!
func (e *ParamService) CreateImage(id uint) error {
	project := &Project{}
	if err := global.GVA_DB.First(project, id).Error; err != nil {
		return err
	}

	params, err := e.GetParamsByProjectId(project.ID)
	if err != nil {
		return err
	}
	// 产生info参数
	var infos []*DrawTextInfo

	tx, iserror := global.GVA_DB.Begin(), true //开启事务
	defer func() {
		if iserror {
			tx.Rollback()
		}
	}()

	uuid := ""
	for _, obj := range params {

		var text string
		if obj.AutoIncrement == 0 { //0不自增 1自增填充,
			text = obj.Text
		} else {
			text = obj.Text + strconv.Itoa(obj.Increment.WinNum)
			uuid = text //选用 最后一个自动递增参数 作为uuid!!!
			tx.Model(obj.Increment).Update("win_num", obj.Increment.WinNum+1)
		}
		//rgba(255, 206, 102, 1)
		//
		rgba, err := string2rgba(obj.Brush.FontColor)
		if err != nil || len(rgba) != 4 {
			return errors.Wrap(err, "画笔颜色异常")
		}
		// 创建画笔
		brush, err := NewTextBrush(obj.Brush.Path, float64(obj.Brush.FontSize), NewFontColor(rgba[0], rgba[1], rgba[2], rgba[3]))
		if err != nil {
			return errors.Wrap(err, "创建画笔失败")
		}
		info := &DrawTextInfo{
			Text:   text,
			X:      obj.X,
			Y:      obj.Y,
			brunch: brush,
		}

		infos = append(infos, info)
	}
	//fmt.Println(infos)
	bytes, err := os.ReadFile(project.ImgSrc)
	if err != nil {
		return errors.Wrap(err, "打开项目模板图失败")
	}

	path := utils.FileName("uploads/drawing/", project.Title+".png")
	err = DrawStringOnImageAndSave(bytes, infos, path)
	if err != nil {
		return err
	}

	//图片入库 drawing_image_db
	tx.Create(&ImageDB{
		ProjectId: project.ID,
		Uuid:      uuid,
		Path:      path,
	})
	// 事务提交
	err = tx.Commit().Error
	if err != nil {
		return err
	}

	iserror = false // 没有任何错误, 不用Rollback
	return nil
}

//func (e *ParamService)
func string2rgba(s string) ([]uint8, error) {
	s = strings.TrimLeft(s, "rgba(")
	s = strings.TrimRight(s, ")")
	ls := strings.Split(s, ",")

	var res []uint8
	for idx, re := range ls {

		score, err := strconv.ParseFloat(strings.Trim(re, " "), 64)

		if err != nil {
			return nil, err
		}
		if idx == 3 {
			res = append(res, uint8(score*255))
		} else {
			res = append(res, uint8(score))
		}
	}

	return res, nil
}
