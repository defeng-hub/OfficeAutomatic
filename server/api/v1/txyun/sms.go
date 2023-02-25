package txyun

import (
	"github.com/defeng-hub/ByOfficeAutomatic/server/global"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/request"
	"github.com/defeng-hub/ByOfficeAutomatic/server/model/common/response"
	smsmodel "github.com/defeng-hub/ByOfficeAutomatic/server/model/txyun/sms"
	"github.com/defeng-hub/ByOfficeAutomatic/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strings"
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

// SendSms
// @Tags      Txyun
// @Summary   发送短信
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     smsmodel.SendSmsRequest  true  "模板ID, 手机号列表, 模板"
// @Success   200   {object}  response.Response{data=object,msg=string}  "发送短信"
// @Router    /txyun/sms/SendSms [post]
func (e *SmsHandler) SendSms(c *gin.Context) {
	req := smsmodel.SendSmsRequest{}
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 去除不满足手机号格式的
	split := strings.Split(req.TplPhones, "\n")
	var phones []string
	var errorphones []string
	for _, v := range split {
		if utils.CheckMobile(v) {
			phones = append(phones, v)
		} else {
			errorphones = append(errorphones, v)
		}
	}

	err = TxyunService.SendSms(req.TemplateId, phones, req.TplParams)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("发送成功", c)

}

// GetAllSmsProject
// @Tags      Txyun
// @Summary   获取全部sms项目
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=object,msg=string}  "获取全部sms项目"
// @Router    /txyun/sms/GetAllSmsProject [post]
func (e *SmsHandler) GetAllSmsProject(c *gin.Context) {
	List, total, err := TxyunService.AllSmsProject()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(gin.H{
		"list":  List,
		"total": total,
	}, "获取成功", c)
}

// AddSmsProject
// @Tags      Txyun
// @Summary   添加sms项目
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     smsmodel.AddSmsProjectReq  true  "项目名, 备注, 模板ID"
// @Success   200   {object}  response.Response{data=object,msg=string}  "获取sms项目"
// @Router    /txyun/sms/AddSmsProject [post]
func (e *SmsHandler) AddSmsProject(c *gin.Context) {
	var req smsmodel.AddSmsProjectReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	project, err := TxyunService.AddSmsProject(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(project, "新增短信项目成功", c)
}

// DelSmsProject
// @Tags      Txyun
// @Summary   删除sms项目
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     request.GetById  true  "模板ID"
// @Success   200   {object}  response.Response{data=object,msg=string}  "获取sms项目"
// @Router    /txyun/sms/DelSmsProject [post]
func (e *SmsHandler) DelSmsProject(c *gin.Context) {
	var req request.GetById
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	project, err := TxyunService.DelSmsProject(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(project, "删除短信项目成功", c)
}

// SmsProjectRows
// @Tags      Txyun
// @Summary   通过sms项目获取全部数据
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     smsmodel.SmsProjectRowsPageReq  true  "请求参数"
// @Success   200   {object}  response.Response{data=object,msg=string}  "通过sms项目Id获取全部数据"
// @Router    /txyun/sms/SmsProjectRows [post]
func (e *SmsHandler) SmsProjectRows(c *gin.Context) {
	var req smsmodel.SmsProjectRowsPageReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rows, total, err := TxyunService.SmsProjectRows(&req)
	if err != nil {
		return
	}
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     rows,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, "获取成功", c)
}

// DelSmsProjectRow
// @Tags      Txyun
// @Summary   删除sms项目成员
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     request.GetById  true  "请求参数"
// @Success   200   {object}  response.Response{data=object,msg=string}  "删除sms项目成员"
// @Router    /txyun/sms/DelSmsProjectRow [post]
func (e *SmsHandler) DelSmsProjectRow(c *gin.Context) {
	var req request.GetById
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	rows, err := TxyunService.DelSmsProjectRow(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(rows, "删除成功", c)
}

// AddSmsProjectRow
// @Tags      Txyun
// @Summary   添加sms项目成员
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     smsmodel.AddSmsProjectRowReq  true  "请求参数"
// @Success   200   {object}  response.Response{data=object,msg=string}  "添加sms项目成员"
// @Router    /txyun/sms/AddSmsProjectRow [post]
func (e *SmsHandler) AddSmsProjectRow(c *gin.Context) {
	var req smsmodel.AddSmsProjectRowsReq
	err := c.Bind(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	split := strings.Split(req.TplPhones, "\n")
	var rows []*smsmodel.SmsProjectRow
	for _, v := range split {
		if utils.CheckMobile(v) {
			mod := smsmodel.SmsProjectRow{
				SmsProjectId: req.SmsProjectId,
				Phone:        v,
				Param1:       req.Param1,
				Param2:       req.Param2,
				Param3:       req.Param3,
				Param4:       req.Param4,
				Param5:       req.Param5,
				Param6:       req.Param6,
				Param7:       req.Param7,
				Param8:       req.Param8,
				Param9:       req.Param9,
			}
			rows = append(rows, &mod)
		}
	}

	row, err := TxyunService.AddSmsProjectRow(rows)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(row, "添加成功", c)
}
