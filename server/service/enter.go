package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/announcementModel"
	"github.com/flipped-aurora/gin-vue-admin/server/service/datastatisticsmodel"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/invitecode"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/userDataUsageModel"
	"github.com/flipped-aurora/gin-vue-admin/server/service/userext"
	"github.com/flipped-aurora/gin-vue-admin/server/service/v2raynode"
)

type ServiceGroup struct {
	SystemServiceGroup              system.ServiceGroup
	ExampleServiceGroup             example.ServiceGroup
	DatastatisticsmodelServiceGroup datastatisticsmodel.ServiceGroup
	UserextServiceGroup             userext.ServiceGroup
	UserDataUsageModelServiceGroup  userDataUsageModel.ServiceGroup
	AnnouncementModelServiceGroup   announcementModel.ServiceGroup
	V2raynodeServiceGroup           v2raynode.ServiceGroup
	InvitecodeServiceGroup          invitecode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
