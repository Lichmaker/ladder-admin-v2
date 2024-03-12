package userDataUsageModel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserDataUsageRouter struct {
}

// InitUserDataUsageRouter 初始化 userDataUsage表 路由信息
func (s *UserDataUsageRouter) InitUserDataUsageRouter(Router *gin.RouterGroup) {
	userDataUsageRouter := Router.Group("userDataUsage").Use(middleware.OperationRecord())
	userDataUsageRouterWithoutRecord := Router.Group("userDataUsage")
	var userDataUsageApi = v1.ApiGroupApp.UserDataUsageModelApiGroup.UserDataUsageApi
	{
		userDataUsageRouter.POST("createUserDataUsage", userDataUsageApi.CreateUserDataUsage)   // 新建userDataUsage表
		userDataUsageRouter.DELETE("deleteUserDataUsage", userDataUsageApi.DeleteUserDataUsage) // 删除userDataUsage表
		userDataUsageRouter.DELETE("deleteUserDataUsageByIds", userDataUsageApi.DeleteUserDataUsageByIds) // 批量删除userDataUsage表
		userDataUsageRouter.PUT("updateUserDataUsage", userDataUsageApi.UpdateUserDataUsage)    // 更新userDataUsage表
	}
	{
		userDataUsageRouterWithoutRecord.GET("findUserDataUsage", userDataUsageApi.FindUserDataUsage)        // 根据ID获取userDataUsage表
		userDataUsageRouterWithoutRecord.GET("getUserDataUsageList", userDataUsageApi.GetUserDataUsageList)  // 获取userDataUsage表列表
	}
}
