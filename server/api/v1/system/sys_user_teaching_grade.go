package system

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/request"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/response"
	model "github.com/defeng-hub/ByOfficeAutomatic/server/model/system"
	"github.com/gin-gonic/gin"
)

// GetAllUserTeachingGrade
// @Tags     Base
// @Summary  获取全部 教学技能等级
// @Produce  application/json
// @Success  200   {object}  response.Response{data=map[string]interface{},msg=string}  "返回全部 教学技能等级"
// @Router   /user/GetAllUserTeachingGrade [get]
func (b *BaseApi) GetAllUserTeachingGrade(c *gin.Context) {
	var grades []model.UserTeachingGrade

	err := global.GVA_DB.Order("shunxv").Find(&grades).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(grades, "获取成功", c)
}

// DeleteUserTeachingGrade
// @Tags     Base
// @Summary  根据ID删除 教学技能等级
// @Produce  application/json
// @Param    data  body      request.GetById               true  "ID"
// @Success  200   {object}  response.Response{data=map[string]interface{},msg=string}  "响应内容"
// @Router   /user/DeleteUserTeachingGrade [post]
func (b *BaseApi) DeleteUserTeachingGrade(c *gin.Context) {
	var ID request.GetById
	err := c.ShouldBindJSON(&ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = global.GVA_DB.Delete(&model.UserTeachingGrade{
		GVA_MODEL: global.GVA_MODEL{ID: ID.Uint()},
	}).Error
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// ChangeUserTeachingGrade
// @Tags     Base
// @Summary  添加教学等级
// @Produce  application/json
// @Param    data  body      model.UserTeachingGrade               true  "请求内容"
// @Success  200   {object}  response.Response{data=map[string]interface{},msg=string}  "响应内容"
// @Router   /user/ChangeUserTeachingGrade [post]
func (b *BaseApi) ChangeUserTeachingGrade(c *gin.Context) {
	var grade model.UserTeachingGrade
	err := c.ShouldBindJSON(&grade)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return

	}
	if grade.ID == 0 {
		err = global.GVA_DB.Create(&grade).Error
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithMessage("添加成功", c)
	} else {
		//Updates 更新非0 行
		err = global.GVA_DB.Model(&grade).Updates(model.UserTeachingGrade{
			Name:   grade.Name,
			Shunxv: grade.Shunxv,
			Wage:   grade.Wage,
		}).Error
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		response.OkWithMessage("修改成功", c)
	}
}
