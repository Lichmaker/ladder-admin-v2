package v2raynode

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type V2rayNodeRouter struct {
}

// InitV2rayNodeRouter 初始化 v2rayNode表 路由信息
func (s *V2rayNodeRouter) InitV2rayNodeRouter(Router *gin.RouterGroup) {
	v2rayNodeRouter := Router.Group("v2rayNode").Use(middleware.OperationRecord())
	v2rayNodeRouterWithoutRecord := Router.Group("v2rayNode")
	var v2rayNodeApi = v1.ApiGroupApp.V2raynodeApiGroup.V2rayNodeApi
	{
		v2rayNodeRouter.POST("createV2rayNode", v2rayNodeApi.CreateV2rayNode)             // 新建v2rayNode表
		v2rayNodeRouter.DELETE("deleteV2rayNode", v2rayNodeApi.DeleteV2rayNode)           // 删除v2rayNode表
		v2rayNodeRouter.DELETE("deleteV2rayNodeByIds", v2rayNodeApi.DeleteV2rayNodeByIds) // 批量删除v2rayNode表
		v2rayNodeRouter.PUT("updateV2rayNode", v2rayNodeApi.UpdateV2rayNode)              // 更新v2rayNode表
		v2rayNodeRouter.POST("pushData", v2rayNodeApi.PushData)                           // 主动下发数据到节点
	}
	{
		v2rayNodeRouterWithoutRecord.GET("findV2rayNode", v2rayNodeApi.FindV2rayNode)       // 根据ID获取v2rayNode表
		v2rayNodeRouterWithoutRecord.GET("getV2rayNodeList", v2rayNodeApi.GetV2rayNodeList) // 获取v2rayNode表列表
	}
}

func (s *V2rayNodeRouter) InitPublicRouter(Router *gin.RouterGroup) {
	var v2rayNodeApi = v1.ApiGroupApp.V2raynodeApiGroup.V2rayNodeApi
	{
		Router.GET("subscribe", v2rayNodeApi.GenSubscribeUrl)
	}
}
