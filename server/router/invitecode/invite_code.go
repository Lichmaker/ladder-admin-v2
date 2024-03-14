package invitecode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type InviteCodeRouter struct {
}

// InitInviteCodeRouter 初始化 inviteCode表 路由信息
func (s *InviteCodeRouter) InitInviteCodeRouter(Router *gin.RouterGroup) {
	inviteCodeRouter := Router.Group("inviteCode").Use(middleware.OperationRecord())
	inviteCodeRouterWithoutRecord := Router.Group("inviteCode")
	var inviteCodeApi = v1.ApiGroupApp.InvitecodeApiGroup.InviteCodeApi
	{
		inviteCodeRouter.POST("createInviteCode", inviteCodeApi.CreateInviteCode)   // 新建inviteCode表
		inviteCodeRouter.DELETE("deleteInviteCode", inviteCodeApi.DeleteInviteCode) // 删除inviteCode表
		inviteCodeRouter.DELETE("deleteInviteCodeByIds", inviteCodeApi.DeleteInviteCodeByIds) // 批量删除inviteCode表
		inviteCodeRouter.DELETE("deleteMyInviteCodeByIds", inviteCodeApi.DeleteMyInviteCodeByIds) // 批量删除inviteCode表
		inviteCodeRouter.PUT("updateInviteCode", inviteCodeApi.UpdateInviteCode)    // 更新inviteCode表
	}
	{
		inviteCodeRouterWithoutRecord.GET("findInviteCode", inviteCodeApi.FindInviteCode)        // 根据ID获取inviteCode表
		inviteCodeRouterWithoutRecord.GET("getInviteCodeList", inviteCodeApi.GetInviteCodeList)  // 获取inviteCode表列表
		inviteCodeRouterWithoutRecord.GET("getMyInviteCodeList", inviteCodeApi.GetMyInviteCodeList)  // 获取inviteCode表列表
	}
}
