import service from '@/utils/request'

export const CreateProject = (data) => {
    return service({
      url: '/drawing/project/CreateProject',
      method: 'post',
      data
    })
}

export const GetAllProject = (params) => {
  return service({
    url: '/drawing/project/GetAllProject',
    method: 'get',
    params
  })
}




//  画笔相关
export const GetAllBranch = (params) => {
    return service({
      url: '/drawing/project/GetAllBranch',
      method: 'get',
      params
    })
}
export const CreateBranch = (data) => {
    return service({
      url: '/drawing/project/CreateBranch',
      method: 'post',
      data
    })
}
export const DeleteBranch = (data) => {
    return service({
      url: '/drawing/project/DeleteBranch',
      method: 'post',
      data
    })
}

export const GetProjectById = (data) => {
  return service({
    url: '/drawing/project/GetProjectById',
    method: 'post',
    data
  })
}
export const GetParamsByProjectId = (data) => {
  return service({
    url: '/drawing/project/GetParamsByProjectId',
    method: 'post',
    data
  })
}
export const CreateParam = (data) => {
  return service({
    url: '/drawing/project/CreateParam',
    method: 'post',
    data
  })
}
export const ChangeParamBranch = (data) => {
  return service({
    url: '/drawing/project/ChangeParamBranch',
    method: 'post',
    data
  })
}
export const ChangeParam = (data) => {
  return service({
    url: '/drawing/project/ChangeParam',
    method: 'post',
    data
  })
}
export const DeleteParam = (data) => {
  return service({
    url: '/drawing/project/DeleteParam',
    method: 'post',
    data
  })
}
export const CreateImage = (data) => {
  return service({
    url: '/drawing/project/CreateImage',
    method: 'post',
    data
  })
}