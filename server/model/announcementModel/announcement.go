// 自动生成模板Announcement
package announcementModel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// announcement表 结构体  Announcement
type Announcement struct {
	global.GVA_MODEL
	Title   string `json:"title" form:"title" gorm:"column:title;comment:;type:VARCHAR(255);"` //title字段
	Content string `json:"content" form:"content" gorm:"column:content;type:TEXT;comment:;"`   //content字段
}

// TableName announcement表 Announcement自定义表名 announcement
func (Announcement) TableName() string {
	return "la_announcement"
}
