package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/article"
	"github.com/flipped-aurora/gin-vue-admin/server/service/datastat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/syscontrol"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service/v2rayUser"
)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	ArticleServiceGroup    article.ServiceGroup
	DatastatServiceGroup   datastat.ServiceGroup
	SysControlServiceGroup syscontrol.ServiceGroup
	V2rayuserServiceGroup  v2rayUser.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
