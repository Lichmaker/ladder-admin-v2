package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/article"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/datastat"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/syscontrol"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/v2rayUser"
)

type ApiGroup struct {
	SystemApiGroup    system.ApiGroup
	ExampleApiGroup   example.ApiGroup
	ArticleApiGroup   article.ApiGroup
	DatastatApiGroup  datastat.ApiGroup
	SysControlGroup   syscontrol.ApiGroup
	V2rayuserApiGroup v2rayUser.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
