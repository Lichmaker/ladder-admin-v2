import service from '@/utils/request'

// @Tags DataPackageCode
// @Summary 创建dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataPackageCode true "创建dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dataPackageCode/createDataPackageCode [post]
export const createDataPackageCode = (data) => {
  return service({
    url: '/dataPackageCode/createDataPackageCode',
    method: 'post',
    data
  })
}

// @Tags DataPackageCode
// @Summary 删除dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataPackageCode true "删除dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataPackageCode/deleteDataPackageCode [delete]
export const deleteDataPackageCode = (params) => {
  return service({
    url: '/dataPackageCode/deleteDataPackageCode',
    method: 'delete',
    params
  })
}

// @Tags DataPackageCode
// @Summary 批量删除dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataPackageCode/deleteDataPackageCode [delete]
export const deleteDataPackageCodeByIds = (params) => {
  return service({
    url: '/dataPackageCode/deleteDataPackageCodeByIds',
    method: 'delete',
    params
  })
}

// @Tags DataPackageCode
// @Summary 更新dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.DataPackageCode true "更新dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataPackageCode/updateDataPackageCode [put]
export const updateDataPackageCode = (data) => {
  return service({
    url: '/dataPackageCode/updateDataPackageCode',
    method: 'put',
    data
  })
}

// @Tags DataPackageCode
// @Summary 用id查询dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.DataPackageCode true "用id查询dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataPackageCode/findDataPackageCode [get]
export const findDataPackageCode = (params) => {
  return service({
    url: '/dataPackageCode/findDataPackageCode',
    method: 'get',
    params
  })
}

// @Tags DataPackageCode
// @Summary 分页获取dataPackageCode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取dataPackageCode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataPackageCode/getDataPackageCodeList [get]
export const getDataPackageCodeList = (params) => {
  return service({
    url: '/dataPackageCode/getDataPackageCodeList',
    method: 'get',
    params
  })
}
