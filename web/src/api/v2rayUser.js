import service from '@/utils/request'

// @Tags V2rayUser
// @Summary 创建V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.V2rayUser true "创建V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v2rayUser/createV2rayUser [post]
export const createV2rayUser = (data) => {
  return service({
    url: '/v2rayUser/createV2rayUser',
    method: 'post',
    data
  })
}

// @Tags V2rayUser
// @Summary 删除V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.V2rayUser true "删除V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v2rayUser/deleteV2rayUser [delete]
export const deleteV2rayUser = (data) => {
  return service({
    url: '/v2rayUser/deleteV2rayUser',
    method: 'delete',
    data
  })
}

// @Tags V2rayUser
// @Summary 删除V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v2rayUser/deleteV2rayUser [delete]
export const deleteV2rayUserByIds = (data) => {
  return service({
    url: '/v2rayUser/deleteV2rayUserByIds',
    method: 'delete',
    data
  })
}

// @Tags V2rayUser
// @Summary 更新V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.V2rayUser true "更新V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v2rayUser/updateV2rayUser [put]
export const updateV2rayUser = (data) => {
  return service({
    url: '/v2rayUser/updateV2rayUser',
    method: 'put',
    data
  })
}

// @Tags V2rayUser
// @Summary 用id查询V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.V2rayUser true "用id查询V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /v2rayUser/findV2rayUser [get]
export const findV2rayUser = (params) => {
  return service({
    url: '/v2rayUser/findV2rayUser',
    method: 'get',
    params
  })
}

// @Tags V2rayUser
// @Summary 分页获取V2rayUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取V2rayUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v2rayUser/getV2rayUserList [get]
export const getV2rayUserList = (params) => {
  return service({
    url: '/v2rayUser/getV2rayUserList',
    method: 'get',
    params
  })
}
