// 自动生成模板V2rayUser
package v2rayUser

import (
	"time"

	"gorm.io/gorm"
)

// V2rayUser 结构体
type V2rayUser struct {
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间

	AdminUserId       *int       `json:"adminUserId" form:"adminUserId" gorm:"column:admin_user_id;comment:admin_users表的id;size:10;primaryKey;"`
	Email             string     `json:"email" form:"email" gorm:"column:email;comment:;size:64;"`
	BandwidthUsageMax *int       `json:"bandwidthUsageMax" form:"bandwidthUsageMax" gorm:"column:bandwidth_usage_max;comment:每月最大的流量使用，单位MB;"`
	V2rayVmess        string     `json:"v2rayVmess" form:"v2rayVmess" gorm:"column:v2ray_vmess;comment:vemss url;size:191;"`
	Uuid              string     `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;size:255;"`
	ExpireAt          *time.Time `json:"expireAt" form:"expireAt" gorm:"column:expire_at;comment:账户过期时间;"`
	StatUpdatedAt     *time.Time `json:"statUpdatedAt" form:"statUpdatedAt" gorm:"column:stat_updated_at;comment:上一次统计时间;"`
	ActiveState       *int       `json:"activeState" form:"activeState" gorm:"column:active_state;comment:账号激活状态;"`
}

// TableName V2rayUser 表名
func (V2rayUser) TableName() string {
	return "v2ray_user"
}
