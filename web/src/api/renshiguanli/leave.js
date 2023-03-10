import service from '@/utils/request'

export const CreateLeaveForm = (data) => {
    return service({
      url: '/renshi/leave/CreateLeaveForm',
      method: 'post',
      data
    })
}