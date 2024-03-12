package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type UserExtSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type UserExtInfoV2 struct {
	Username       string `json:"userName" example:"用户名"`
	Password       string `json:"passWord" example:"密码"`
	NickName       string `json:"nickName" example:"昵称"`
	Enable         int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	Phone          string `json:"phone" example:"电话号码"`
	Email          string `json:"email" example:"电子邮箱"`
	ValidStart     string `json:"validStart" form:"validStart"` // 开始日期 2006-01-02
	ValidEnd       string `json:"validEnd" form:"validEnd"`     // 结束日期 2006-01-02
	StandardDataMB int    `json:"standardDataMB" form:"standardDataMB"`
	AuthorityId    uint   `json:"authorityId" swaggertype:"string" example:"int 角色id"`
	AuthorityIds   []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
}

type UpdateUserInfo struct {
	ID             int    `json:"ID"` // 主键ID
	NickName       string `json:"nickName" example:"昵称"`
	Enable         int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	Phone          string `json:"phone" example:"电话号码"`
	Email          string `json:"email" example:"电子邮箱"`
	ValidStart     string `json:"validStart" form:"validStart"` // 开始日期 2006-01-02
	ValidEnd       string `json:"validEnd" form:"validEnd"`     // 结束日期 2006-01-02
	StandardDataMB int    `json:"standardDataMB" form:"standardDataMB"`
}

type UserListSearch struct {
	UUID   string `json:"uuid" form:"uuid"`
	Enable *int   `json:"enable,omitempty" form:"enable"` // user_ext.enable
	request.PageInfo
}
