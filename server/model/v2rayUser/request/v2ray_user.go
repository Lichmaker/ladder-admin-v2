package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/v2rayUser"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type V2rayUserSearch struct{
    v2rayUser.V2rayUser
    request.PageInfo
}
