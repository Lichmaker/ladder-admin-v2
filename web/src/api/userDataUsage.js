import service from '@/utils/request'

// @Tags UserDataUsage
// @Summary 创建userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDataUsage true "创建userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userDataUsage/createUserDataUsage [post]
export const createUserDataUsage = (data) => {
  return service({
    url: '/userDataUsage/createUserDataUsage',
    method: 'post',
    data
  })
}

// @Tags UserDataUsage
// @Summary 删除userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDataUsage true "删除userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDataUsage/deleteUserDataUsage [delete]
export const deleteUserDataUsage = (params) => {
  return service({
    url: '/userDataUsage/deleteUserDataUsage',
    method: 'delete',
    params
  })
}

// @Tags UserDataUsage
// @Summary 批量删除userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDataUsage/deleteUserDataUsage [delete]
export const deleteUserDataUsageByIds = (params) => {
  return service({
    url: '/userDataUsage/deleteUserDataUsageByIds',
    method: 'delete',
    params
  })
}

// @Tags UserDataUsage
// @Summary 更新userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserDataUsage true "更新userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userDataUsage/updateUserDataUsage [put]
export const updateUserDataUsage = (data) => {
  return service({
    url: '/userDataUsage/updateUserDataUsage',
    method: 'put',
    data
  })
}

// @Tags UserDataUsage
// @Summary 用id查询userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserDataUsage true "用id查询userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userDataUsage/findUserDataUsage [get]
export const findUserDataUsage = (params) => {
  return service({
    url: '/userDataUsage/findUserDataUsage',
    method: 'get',
    params
  })
}

// @Tags UserDataUsage
// @Summary 分页获取userDataUsage表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userDataUsage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDataUsage/getUserDataUsageList [get]
export const getUserDataUsageList = (params) => {
  return service({
    url: '/userDataUsage/getUserDataUsageList',
    method: 'get',
    params
  })
}
