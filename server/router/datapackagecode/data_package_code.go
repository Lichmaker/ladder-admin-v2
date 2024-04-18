package datapackagecode

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataPackageCodeRouter struct {
}

// InitDataPackageCodeRouter 初始化 dataPackageCode表 路由信息
func (s *DataPackageCodeRouter) InitDataPackageCodeRouter(Router *gin.RouterGroup) {
	dataPackageCodeRouter := Router.Group("dataPackageCode").Use(middleware.OperationRecord())
	dataPackageCodeRouterWithoutRecord := Router.Group("dataPackageCode")
	var dataPackageCodeApi = v1.ApiGroupApp.DatapackagecodeApiGroup.DataPackageCodeApi
	{
		dataPackageCodeRouter.POST("createDataPackageCode", dataPackageCodeApi.CreateDataPackageCode)             // 新建dataPackageCode表
		dataPackageCodeRouter.DELETE("deleteDataPackageCode", dataPackageCodeApi.DeleteDataPackageCode)           // 删除dataPackageCode表
		dataPackageCodeRouter.DELETE("deleteDataPackageCodeByIds", dataPackageCodeApi.DeleteDataPackageCodeByIds) // 批量删除dataPackageCode表
		dataPackageCodeRouter.PUT("updateDataPackageCode", dataPackageCodeApi.UpdateDataPackageCode)              // 更新dataPackageCode表
		dataPackageCodeRouter.POST("batchCreateDataPackageCode", dataPackageCodeApi.BatchCreateDataPackageCode)
		dataPackageCodeRouter.POST("exchange", dataPackageCodeApi.ExchangeDataPackageCode)
		dataPackageCodeRouter.POST("consumePackage", dataPackageCodeApi.ConsumePackage)

	}
	{
		dataPackageCodeRouterWithoutRecord.GET("findDataPackageCode", dataPackageCodeApi.FindDataPackageCode)       // 根据ID获取dataPackageCode表
		dataPackageCodeRouterWithoutRecord.GET("getDataPackageCodeList", dataPackageCodeApi.GetDataPackageCodeList) // 获取dataPackageCode表列表
		dataPackageCodeRouterWithoutRecord.GET("getMyDataPackageCodeList", dataPackageCodeApi.GetMyDataPackageCodeList)
	}
}
