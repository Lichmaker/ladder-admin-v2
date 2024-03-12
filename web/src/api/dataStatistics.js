import service from '@/utils/request'

// @Tags DataStatistics
// @Summary 创建dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataStatistics true "创建dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dataStatistics/createDataStatistics [post]
export const createDataStatistics = (data) => {
  return service({
    url: '/dataStatistics/createDataStatistics',
    method: 'post',
    data
  })
}

// @Tags DataStatistics
// @Summary 删除dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataStatistics true "删除dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataStatistics/deleteDataStatistics [delete]
export const deleteDataStatistics = (params) => {
  return service({
    url: '/dataStatistics/deleteDataStatistics',
    method: 'delete',
    params
  })
}

// @Tags DataStatistics
// @Summary 批量删除dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataStatistics/deleteDataStatistics [delete]
export const deleteDataStatisticsByIds = (params) => {
  return service({
    url: '/dataStatistics/deleteDataStatisticsByIds',
    method: 'delete',
    params
  })
}

// @Tags DataStatistics
// @Summary 更新dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataStatistics true "更新dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataStatistics/updateDataStatistics [put]
export const updateDataStatistics = (data) => {
  return service({
    url: '/dataStatistics/updateDataStatistics',
    method: 'put',
    data
  })
}

// @Tags DataStatistics
// @Summary 用id查询dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DataStatistics true "用id查询dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataStatistics/findDataStatistics [get]
export const findDataStatistics = (params) => {
  return service({
    url: '/dataStatistics/findDataStatistics',
    method: 'get',
    params
  })
}

// @Tags DataStatistics
// @Summary 分页获取dataStatistics表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取dataStatistics表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataStatistics/getDataStatisticsList [get]
export const getDataStatisticsList = (params) => {
  return service({
    url: '/dataStatistics/getDataStatisticsList',
    method: 'get',
    params
  })
}
