import service from '@/utils/request'

export const CreateLeaveForm = (data) => {
    return service({
      url: '/renshi/leave/CreateLeaveForm',
      method: 'post',
      data
    })
}

export const GetMyselfLeaves = (params) => {
  return service({
    url: '/renshi/leave/GetMyselfLeaves',
    method: 'get',
    params
  })
}

export const GetDaichuliLeaves = (params) => {
  return service({
    url: '/renshi/leave/GetDaichuliLeaves',
    method: 'get',
    params
  })
}

// DeleteLeaveForm
// @Tags      Renshiguanli
// @Summary   删除请假表单
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.GetById  true  "请求内容"
// @Success   200   {object}  response.Response{msg=string}  "响应内容"
// @Router    /renshi/leave/DeleteLeaveForm [post]
export const DeleteLeaveForm = (data) => {
  return service({
    url: '/renshi/leave/DeleteLeaveForm',
    method: 'post',
    data
  })
}