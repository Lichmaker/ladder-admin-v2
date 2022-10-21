import service from '@/utils/request'

// @Tags DataSummary
// @Summary 创建DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataSummary true "创建DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataSummary/createDataSummary [post]
export const createDataSummary = (data) => {
  return service({
    url: '/dataSummary/createDataSummary',
    method: 'post',
    data
  })
}

// @Tags DataSummary
// @Summary 删除DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataSummary true "删除DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataSummary/deleteDataSummary [delete]
export const deleteDataSummary = (data) => {
  return service({
    url: '/dataSummary/deleteDataSummary',
    method: 'delete',
    data
  })
}

// @Tags DataSummary
// @Summary 删除DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataSummary/deleteDataSummary [delete]
export const deleteDataSummaryByIds = (data) => {
  return service({
    url: '/dataSummary/deleteDataSummaryByIds',
    method: 'delete',
    data
  })
}

// @Tags DataSummary
// @Summary 更新DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataSummary true "更新DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataSummary/updateDataSummary [put]
export const updateDataSummary = (data) => {
  return service({
    url: '/dataSummary/updateDataSummary',
    method: 'put',
    data
  })
}

// @Tags DataSummary
// @Summary 用id查询DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DataSummary true "用id查询DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataSummary/findDataSummary [get]
export const findDataSummary = (params) => {
  return service({
    url: '/dataSummary/findDataSummary',
    method: 'get',
    params
  })
}

// @Tags DataSummary
// @Summary 分页获取DataSummary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取DataSummary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataSummary/getDataSummaryList [get]
export const getDataSummaryList = (params) => {
  return service({
    url: '/dataSummary/getDataSummaryList',
    method: 'get',
    params
  })
}

export const getMyDataSummaryList = (params) => {
    return service({
      url: '/dataSummary/getMyDataSummaryList',
      method: 'get',
      params
    })
  }
  