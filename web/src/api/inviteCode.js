import service from '@/utils/request'

// @Tags InviteCode
// @Summary 创建inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InviteCode true "创建inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /inviteCode/createInviteCode [post]
export const createInviteCode = (data) => {
  return service({
    url: '/inviteCode/createInviteCode',
    method: 'post',
    data
  })
}

// @Tags InviteCode
// @Summary 删除inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InviteCode true "删除inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inviteCode/deleteInviteCode [delete]
export const deleteInviteCode = (params) => {
  return service({
    url: '/inviteCode/deleteInviteCode',
    method: 'delete',
    params
  })
}

// @Tags InviteCode
// @Summary 批量删除inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inviteCode/deleteInviteCode [delete]
export const deleteInviteCodeByIds = (params) => {
  return service({
    url: '/inviteCode/deleteInviteCodeByIds',
    method: 'delete',
    params
  })
}

export const deleteMyInviteCodeByIds = (params) => {
  return service({
    url: '/inviteCode/deleteMyInviteCodeByIds',
    method: 'delete',
    params
  })
}

// @Tags InviteCode
// @Summary 更新inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.InviteCode true "更新inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inviteCode/updateInviteCode [put]
export const updateInviteCode = (data) => {
  return service({
    url: '/inviteCode/updateInviteCode',
    method: 'put',
    data
  })
}

// @Tags InviteCode
// @Summary 用id查询inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.InviteCode true "用id查询inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inviteCode/findInviteCode [get]
export const findInviteCode = (params) => {
  return service({
    url: '/inviteCode/findInviteCode',
    method: 'get',
    params
  })
}

// @Tags InviteCode
// @Summary 分页获取inviteCode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取inviteCode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inviteCode/getInviteCodeList [get]
export const getInviteCodeList = (params) => {
  return service({
    url: '/inviteCode/getInviteCodeList',
    method: 'get',
    params
  })
}

export const getMyInviteCodeList = (params) => {
  return service({
    url: '/inviteCode/getMyInviteCodeList',
    method: 'get',
    params
  })
}

