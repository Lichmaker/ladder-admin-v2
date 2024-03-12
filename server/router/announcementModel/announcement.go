package announcementModel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AnnouncementRouter struct {
}

// InitAnnouncementRouter 初始化 announcement表 路由信息
func (s *AnnouncementRouter) InitAnnouncementRouter(Router *gin.RouterGroup) {
	announcementRouter := Router.Group("announcement").Use(middleware.OperationRecord())
	announcementRouterWithoutRecord := Router.Group("announcement")
	var announcementApi = v1.ApiGroupApp.AnnouncementModelApiGroup.AnnouncementApi
	{
		announcementRouter.POST("createAnnouncement", announcementApi.CreateAnnouncement)   // 新建announcement表
		announcementRouter.DELETE("deleteAnnouncement", announcementApi.DeleteAnnouncement) // 删除announcement表
		announcementRouter.DELETE("deleteAnnouncementByIds", announcementApi.DeleteAnnouncementByIds) // 批量删除announcement表
		announcementRouter.PUT("updateAnnouncement", announcementApi.UpdateAnnouncement)    // 更新announcement表
	}
	{
		announcementRouterWithoutRecord.GET("findAnnouncement", announcementApi.FindAnnouncement)        // 根据ID获取announcement表
		announcementRouterWithoutRecord.GET("getAnnouncementList", announcementApi.GetAnnouncementList)  // 获取announcement表列表
	}
}
