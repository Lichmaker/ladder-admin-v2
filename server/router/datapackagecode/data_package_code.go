package datapackagecode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type LaDataPackageCodeRouter struct {
}

// InitLaDataPackageCodeRouter 初始化 dataPackageCode表 路由信息
func (s *LaDataPackageCodeRouter) InitLaDataPackageCodeRouter(Router *gin.RouterGroup) {
	dataPackageCodeRouter := Router.Group("dataPackageCode").Use(middleware.OperationRecord())
	dataPackageCodeRouterWithoutRecord := Router.Group("dataPackageCode")
	var dataPackageCodeApi = v1.ApiGroupApp.DatapackagecodeApiGroup.LaDataPackageCodeApi
	{
		dataPackageCodeRouter.POST("createLaDataPackageCode", dataPackageCodeApi.CreateLaDataPackageCode)   // 新建dataPackageCode表
		dataPackageCodeRouter.DELETE("deleteLaDataPackageCode", dataPackageCodeApi.DeleteLaDataPackageCode) // 删除dataPackageCode表
		dataPackageCodeRouter.DELETE("deleteLaDataPackageCodeByIds", dataPackageCodeApi.DeleteLaDataPackageCodeByIds) // 批量删除dataPackageCode表
		dataPackageCodeRouter.PUT("updateLaDataPackageCode", dataPackageCodeApi.UpdateLaDataPackageCode)    // 更新dataPackageCode表
	}
	{
		dataPackageCodeRouterWithoutRecord.GET("findLaDataPackageCode", dataPackageCodeApi.FindLaDataPackageCode)        // 根据ID获取dataPackageCode表
		dataPackageCodeRouterWithoutRecord.GET("getLaDataPackageCodeList", dataPackageCodeApi.GetLaDataPackageCodeList)  // 获取dataPackageCode表列表
	}
}
