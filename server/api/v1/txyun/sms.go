package txyun

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/request"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/response"
	"github.com/defeng-hub/ByOfficeAutomatic/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SmsHandler struct{}

// UpdateTemplatesApi
// @Tags      Txyun
// @Summary   更新短信模板列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200
// @Router    /txyun/sms/UpdateTemplates [post]
func (s *SmsHandler) UpdateTemplatesApi(c *gin.Context) {
	err := TxyunService.UpdateTemplates()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetSmsList
// @Tags      Txyun
// @Summary   分页获取短信模板列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo  true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /txyun/sms/GetSmsList [post]
func (e *SmsHandler) GetSmsList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.Bind(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	List, total, err := TxyunService.GetSmsInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     List,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)

}
