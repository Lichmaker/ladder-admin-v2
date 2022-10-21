package syscontrol

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysControlRouter struct {
}

func (s *SysControlRouter) InitSysControlRouter(Router *gin.RouterGroup) {
	sysControlRouter := Router.Group("sysControl").Use(middleware.OperationRecord())
	// dataSummaryRouterWithoutRecord := Router.Group("sysControl")
	var sysControlApi = v1.ApiGroupApp.SysControlGroup.SysControlApi
	{
		sysControlRouter.POST("restartV2Ray", sysControlApi.RestartV2ray) // 重启v2ray
		sysControlRouter.POST("playground", sysControlApi.MyPlayground)   // debug 用的playground
	}
}
