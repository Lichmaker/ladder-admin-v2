package userext

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type UserExtRouter struct {
}

// InitUserExtRouter 初始化 userExt表 路由信息
func (s *UserExtRouter) InitUserExtRouter(Router *gin.RouterGroup) {
	userExtRouter := Router.Group("userExt").Use(middleware.OperationRecord())
	userExtRouterWithoutRecord := Router.Group("userExt")
	var userExtApi = v1.ApiGroupApp.UserextApiGroup.UserExtApi
	{
		userExtRouter.POST("createUserExt", userExtApi.CreateUserExt)             // 新建userExt表
		userExtRouter.DELETE("deleteUserExt", userExtApi.DeleteUserExt)           // 删除userExt表
		userExtRouter.DELETE("deleteUserExtByIds", userExtApi.DeleteUserExtByIds) // 批量删除userExt表
		userExtRouter.PUT("updateUserExt", userExtApi.UpdateUserExt)              // 更新userExt表
	}
	{
		userExtRouterWithoutRecord.GET("findUserExt", userExtApi.FindUserExt)       // 根据ID获取userExt表
		userExtRouterWithoutRecord.GET("getUserExtList", userExtApi.GetUserExtList) // 获取userExt表列表
	}

	// 结合 user_ext 的用户增删改查
	userExtParentGroup := Router.Group("userExt")
	userExtV2Router := userExtParentGroup.Group("v2").Use(middleware.OperationRecord())
	userExtV2RouterWithoutRecord := userExtParentGroup.Group("v2")
	{
		userExtV2Router.POST("create", userExtApi.CreateUserExtV2)
		userExtV2Router.PUT("update", userExtApi.UpdateUserExtV2)
	}
	{
		userExtV2RouterWithoutRecord.GET("getList", userExtApi.GetUserList)
		userExtV2RouterWithoutRecord.GET("dashboard", userExtApi.Dashboard)
	}
}
