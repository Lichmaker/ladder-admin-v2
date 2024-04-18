import service from '@/utils/request'

// @Tags LaDataPackageCode
// @Summary 创建dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LaDataPackageCode true "创建dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dataPackageCode/createLaDataPackageCode [post]
export const createLaDataPackageCode = (data) => {
  return service({
    url: '/dataPackageCode/createLaDataPackageCode',
    method: 'post',
    data
  })
}

// @Tags LaDataPackageCode
// @Summary 删除dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LaDataPackageCode true "删除dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataPackageCode/deleteLaDataPackageCode [delete]
export const deleteLaDataPackageCode = (params) => {
  return service({
    url: '/dataPackageCode/deleteLaDataPackageCode',
    method: 'delete',
    params
  })
}

// @Tags LaDataPackageCode
// @Summary 批量删除dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataPackageCode/deleteLaDataPackageCode [delete]
export const deleteLaDataPackageCodeByIds = (params) => {
  return service({
    url: '/dataPackageCode/deleteLaDataPackageCodeByIds',
    method: 'delete',
    params
  })
}

// @Tags LaDataPackageCode
// @Summary 更新dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.LaDataPackageCode true "更新dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataPackageCode/updateLaDataPackageCode [put]
export const updateLaDataPackageCode = (data) => {
  return service({
    url: '/dataPackageCode/updateLaDataPackageCode',
    method: 'put',
    data
  })
}

// @Tags LaDataPackageCode
// @Summary 用id查询dataPackageCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.LaDataPackageCode true "用id查询dataPackageCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataPackageCode/findLaDataPackageCode [get]
export const findLaDataPackageCode = (params) => {
  return service({
    url: '/dataPackageCode/findLaDataPackageCode',
    method: 'get',
    params
  })
}

// @Tags LaDataPackageCode
// @Summary 分页获取dataPackageCode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取dataPackageCode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataPackageCode/getLaDataPackageCodeList [get]
export const getLaDataPackageCodeList = (params) => {
  return service({
    url: '/dataPackageCode/getLaDataPackageCodeList',
    method: 'get',
    params
  })
}
