import service from '@/utils/request'

// UpdateTemplatesApi
// @Tags      Txyun
// @Summary   更新短信模板列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200
// @Router    /txyun/sms/UpdateTemplates [post]
export const UpdateTemplates = () => {
  return service({
    url: '/txyun/sms/UpdateTemplates',
    method: 'post',
  })
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
export const GetSmsList = (params) => {
    return service({
      url: '/txyun/sms/GetSmsList',
      method: 'get',
      params
    })
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
export const SendSms = (data) => {
  return service({
    url: '/txyun/sms/SendSms',
    method: 'post',
    data
  })
}

// AddSmsProject
// @Tags      Txyun
// @Summary   添加sms项目
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     smsmodel.AddSmsProjectReq  true  "项目名, 备注, 模板ID"
// @Success   200   {object}  response.Response{data=object,msg=string}  "获取全部sms项目"
// @Router    /txyun/sms/AddSmsProject [post]
export const AddSmsProject = (data) => {
  return service({
    url: '/txyun/sms/AddSmsProject',
    method: 'post',
    data
  })
}

// GetAllSmsProject
// @Tags      Txyun
// @Summary   获取全部sms项目
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Success   200   {object}  response.Response{data=object,msg=string}  "获取全部sms项目"
// @Router    /txyun/sms/GetAllSmsProject [post]
export const GetAllSmsProject = (data) => {
  return service({
    url: '/txyun/sms/GetAllSmsProject',
    method: 'post',
    data
  })
}


// DelSmsProject
// @Tags      Txyun
// @Summary   删除sms项目
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body     smsmodel.DelSmsProjectReq  true  "模板ID"
// @Success   200   {object}  response.Response{data=object,msg=string}  "获取sms项目"
// @Router    /txyun/sms/DelSmsProject [post]
export const DelSmsProject = (data) => {
  return service({
    url: '/txyun/sms/DelSmsProject',
    method: 'post',
    data
  })
}