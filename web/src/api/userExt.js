import service from '@/utils/request'

// @Tags UserExt
// @Summary 创建userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserExt true "创建userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userExt/createUserExt [post]
export const createUserExt = (data) => {
  return service({
    url: '/userExt/createUserExt',
    method: 'post',
    data
  })
}

// @Tags UserExt
// @Summary 删除userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserExt true "删除userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userExt/deleteUserExt [delete]
export const deleteUserExt = (params) => {
  return service({
    url: '/userExt/deleteUserExt',
    method: 'delete',
    params
  })
}

// @Tags UserExt
// @Summary 批量删除userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userExt/deleteUserExt [delete]
export const deleteUserExtByIds = (params) => {
  return service({
    url: '/userExt/deleteUserExtByIds',
    method: 'delete',
    params
  })
}

// @Tags UserExt
// @Summary 更新userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UserExt true "更新userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userExt/updateUserExt [put]
export const updateUserExt = (data) => {
  return service({
    url: '/userExt/updateUserExt',
    method: 'put',
    data
  })
}

// @Tags UserExt
// @Summary 用id查询userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UserExt true "用id查询userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userExt/findUserExt [get]
export const findUserExt = (params) => {
  return service({
    url: '/userExt/findUserExt',
    method: 'get',
    params
  })
}

// @Tags UserExt
// @Summary 分页获取userExt表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取userExt表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userExt/getUserExtList [get]
export const getUserExtList = (params) => {
  return service({
    url: '/userExt/getUserExtList',
    method: 'get',
    params
  })
}

// @Tags User
// @Summary 分页获取用户列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body modelInterface.PageInfo true "分页获取用户列表"
// @Success 200 {string} json "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /user/getUserList [post]
export const getUserListV2 = (data) => {
  return service({
    url: '/userExt/v2/getList',
    method: 'get',
    params: data,
  })
}

export const createUserV2 = (data) => {
  return service({
    url: '/userExt/v2/create',
    method: 'post',
    data
  })
}

export const updateUserV2 = (data) => {
  return service({
    url: '/userExt/v2/update',
    method: 'put',
    data
  })
}

export const getDashboardInfo = (showLoading) => {
  return service({
    url: '/userExt/v2/dashboard',
    method: 'get',
    donNotShowLoading: !showLoading
  })
}

