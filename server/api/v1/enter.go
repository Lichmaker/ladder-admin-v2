package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/announcementModel"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/datastatisticsmodel"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/invitecode"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/userDataUsageModel"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/userext"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/v2raynode"
)

type ApiGroup struct {
	SystemApiGroup              system.ApiGroup
	ExampleApiGroup             example.ApiGroup
	DatastatisticsmodelApiGroup datastatisticsmodel.ApiGroup
	UserextApiGroup             userext.ApiGroup
	UserDataUsageModelApiGroup  userDataUsageModel.ApiGroup
	AnnouncementModelApiGroup   announcementModel.ApiGroup
	V2raynodeApiGroup           v2raynode.ApiGroup
	InvitecodeApiGroup          invitecode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
