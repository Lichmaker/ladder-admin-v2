package datastat

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DataSummaryRouter struct {
}

// InitDataSummaryRouter 初始化 DataSummary 路由信息
func (s *DataSummaryRouter) InitDataSummaryRouter(Router *gin.RouterGroup) {
	dataSummaryRouter := Router.Group("dataSummary").Use(middleware.OperationRecord())
	dataSummaryRouterWithoutRecord := Router.Group("dataSummary")
	var dataSummaryApi = v1.ApiGroupApp.DatastatApiGroup.DataSummaryApi
	{
		dataSummaryRouter.POST("createDataSummary", dataSummaryApi.CreateDataSummary)             // 新建DataSummary
		dataSummaryRouter.DELETE("deleteDataSummary", dataSummaryApi.DeleteDataSummary)           // 删除DataSummary
		dataSummaryRouter.DELETE("deleteDataSummaryByIds", dataSummaryApi.DeleteDataSummaryByIds) // 批量删除DataSummary
		dataSummaryRouter.PUT("updateDataSummary", dataSummaryApi.UpdateDataSummary)              // 更新DataSummary
	}
	{
		dataSummaryRouterWithoutRecord.GET("findDataSummary", dataSummaryApi.FindDataSummary)           // 根据ID获取DataSummary
		dataSummaryRouterWithoutRecord.GET("getDataSummaryList", dataSummaryApi.GetDataSummaryList)     // 获取DataSummary列表
		dataSummaryRouterWithoutRecord.GET("getMyDataSummaryList", dataSummaryApi.GetMyDataSummaryList) // 获取登录用户的DataSummary列表
	}
}
