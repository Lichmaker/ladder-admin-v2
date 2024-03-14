package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/announcementModel"
	"github.com/flipped-aurora/gin-vue-admin/server/router/datastatisticsmodel"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/invitecode"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/userDataUsageModel"
	"github.com/flipped-aurora/gin-vue-admin/server/router/userext"
	"github.com/flipped-aurora/gin-vue-admin/server/router/v2raynode"
)

type RouterGroup struct {
	System              system.RouterGroup
	Example             example.RouterGroup
	Datastatisticsmodel datastatisticsmodel.RouterGroup
	Userext             userext.RouterGroup
	UserDataUsageModel  userDataUsageModel.RouterGroup
	AnnouncementModel   announcementModel.RouterGroup
	V2raynode           v2raynode.RouterGroup
	Invitecode          invitecode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
