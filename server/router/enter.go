package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/article"
	"github.com/flipped-aurora/gin-vue-admin/server/router/datastat"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/syscontrol"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
	"github.com/flipped-aurora/gin-vue-admin/server/router/v2rayUser"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Article    article.RouterGroup
	Datastat   datastat.RouterGroup
	SysControl syscontrol.RouterGroup
	V2rayuser  v2rayUser.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
