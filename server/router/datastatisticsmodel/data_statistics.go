package datastatisticsmodel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataStatisticsRouter struct {
}

// InitDataStatisticsRouter 初始化 dataStatistics表 路由信息
func (s *DataStatisticsRouter) InitDataStatisticsRouter(Router *gin.RouterGroup) {
	dataStatisticsRouter := Router.Group("dataStatistics").Use(middleware.OperationRecord())
	dataStatisticsRouterWithoutRecord := Router.Group("dataStatistics")
	var dataStatisticsApi = v1.ApiGroupApp.DatastatisticsmodelApiGroup.DataStatisticsApi
	{
		dataStatisticsRouter.POST("createDataStatistics", dataStatisticsApi.CreateDataStatistics)   // 新建dataStatistics表
		dataStatisticsRouter.DELETE("deleteDataStatistics", dataStatisticsApi.DeleteDataStatistics) // 删除dataStatistics表
		dataStatisticsRouter.DELETE("deleteDataStatisticsByIds", dataStatisticsApi.DeleteDataStatisticsByIds) // 批量删除dataStatistics表
		dataStatisticsRouter.PUT("updateDataStatistics", dataStatisticsApi.UpdateDataStatistics)    // 更新dataStatistics表
	}
	{
		dataStatisticsRouterWithoutRecord.GET("findDataStatistics", dataStatisticsApi.FindDataStatistics)        // 根据ID获取dataStatistics表
		dataStatisticsRouterWithoutRecord.GET("getDataStatisticsList", dataStatisticsApi.GetDataStatisticsList)  // 获取dataStatistics表列表
	}
}
