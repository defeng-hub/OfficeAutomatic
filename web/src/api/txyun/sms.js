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


export const GetSmsList = (params) => {
    return service({
      url: '/txyun/sms/GetSmsList',
      method: 'get',
      params
    })
  }