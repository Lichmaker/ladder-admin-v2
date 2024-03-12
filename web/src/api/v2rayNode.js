import service from '@/utils/request'

// @Tags V2rayNode
// @Summary 创建v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.V2rayNode true "创建v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v2rayNode/createV2rayNode [post]
export const createV2rayNode = (data) => {
  return service({
    url: '/v2rayNode/createV2rayNode',
    method: 'post',
    data
  })
}

// @Tags V2rayNode
// @Summary 删除v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.V2rayNode true "删除v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v2rayNode/deleteV2rayNode [delete]
export const deleteV2rayNode = (params) => {
  return service({
    url: '/v2rayNode/deleteV2rayNode',
    method: 'delete',
    params
  })
}

// @Tags V2rayNode
// @Summary 批量删除v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v2rayNode/deleteV2rayNode [delete]
export const deleteV2rayNodeByIds = (params) => {
  return service({
    url: '/v2rayNode/deleteV2rayNodeByIds',
    method: 'delete',
    params
  })
}

// @Tags V2rayNode
// @Summary 更新v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.V2rayNode true "更新v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v2rayNode/updateV2rayNode [put]
export const updateV2rayNode = (data) => {
  return service({
    url: '/v2rayNode/updateV2rayNode',
    method: 'put',
    data
  })
}

// @Tags V2rayNode
// @Summary 用id查询v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.V2rayNode true "用id查询v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /v2rayNode/findV2rayNode [get]
export const findV2rayNode = (params) => {
  return service({
    url: '/v2rayNode/findV2rayNode',
    method: 'get',
    params
  })
}

// @Tags V2rayNode
// @Summary 分页获取v2rayNode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取v2rayNode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v2rayNode/getV2rayNodeList [get]
export const getV2rayNodeList = (params) => {
  return service({
    url: '/v2rayNode/getV2rayNodeList',
    method: 'get',
    params
  })
}

export const pushV2rayNodeData = (data) => {
  return service({
    url: '/v2rayNode/pushData',
    method: 'post',
    data
  })
}

