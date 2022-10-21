package v2rayUser

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type V2rayUserRouter struct {
}

// InitV2rayUserRouter 初始化 V2rayUser 路由信息
func (s *V2rayUserRouter) InitV2rayUserRouter(Router *gin.RouterGroup) {
	v2rayUserRouter := Router.Group("v2rayUser").Use(middleware.OperationRecord())
	v2rayUserRouterWithoutRecord := Router.Group("v2rayUser")
	var v2rayUserApi = v1.ApiGroupApp.V2rayuserApiGroup.V2rayUserApi
	{
		v2rayUserRouter.POST("createV2rayUser", v2rayUserApi.CreateV2rayUser)   // 新建V2rayUser
		v2rayUserRouter.DELETE("deleteV2rayUser", v2rayUserApi.DeleteV2rayUser) // 删除V2rayUser
		v2rayUserRouter.DELETE("deleteV2rayUserByIds", v2rayUserApi.DeleteV2rayUserByIds) // 批量删除V2rayUser
		v2rayUserRouter.PUT("updateV2rayUser", v2rayUserApi.UpdateV2rayUser)    // 更新V2rayUser
	}
	{
		v2rayUserRouterWithoutRecord.GET("findV2rayUser", v2rayUserApi.FindV2rayUser)        // 根据ID获取V2rayUser
		v2rayUserRouterWithoutRecord.GET("getV2rayUserList", v2rayUserApi.GetV2rayUserList)  // 获取V2rayUser列表
	}
}
