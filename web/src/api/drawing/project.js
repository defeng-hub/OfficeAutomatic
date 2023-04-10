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