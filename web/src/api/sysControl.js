import service from '@/utils/request'

export const restartV2Ray = (data) => {
  return service({
    url: '/sysControl/restartV2Ray',
    method: 'post',
    data
  })
}

export const myPlayground = (data) => {
    return service({
      url: '/sysControl/playground',
      method: 'post',
      data
    })
  }
  