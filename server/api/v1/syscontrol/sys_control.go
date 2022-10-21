package syscontrol

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
)

type SysControlApi struct {
}

var sysControlService = service.ServiceGroupApp.SysControlServiceGroup.SysControlService

func (sysControlApi *SysControlApi) RestartV2ray(c *gin.Context) {
	if false {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (sysControlApi *SysControlApi) MyPlayground(c *gin.Context) {
	if false {
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("hey", c)
	}
}
