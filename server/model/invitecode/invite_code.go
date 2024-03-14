// 自动生成模板InviteCode
package invitecode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// inviteCode表 结构体  InviteCode
type InviteCode struct {
	global.GVA_MODEL
	Code        string `json:"code" form:"code" gorm:"column:code;comment:邀请码;size:32;unique;"`                           //邀请码
	CreatorUuid string `json:"creatorUuid" form:"creatorUuid" gorm:"column:creator_uuid;comment:创建人uuid;size:64;"` //创建人uuid
	NewUuid     string `json:"newUuid" form:"newUuid" gorm:"column:new_uuid;comment:使用者uuid;size:64;"`             //使用者uuid
	FreeData    *int   `json:"freeData" form:"freeData" gorm:"column:free_data;comment:;type:INT;"`                //freeData字段
	FreeDays    *int   `json:"freeDays" form:"freeDays" gorm:"column:free_days;comment:;type:INT"`                 //freeDays字段
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName inviteCode表 InviteCode自定义表名 la_invite_code
func (InviteCode) TableName() string {
	return "la_invite_code"
}
