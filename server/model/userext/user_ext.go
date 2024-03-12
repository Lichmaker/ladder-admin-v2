// 自动生成模板UserExt
package userext

import (
	"time"

	"gorm.io/gorm"
)

// userExt表 结构体  UserExt
type UserExt struct {
	CreatedAt    time.Time      // 创建时间
	UpdatedAt    time.Time      // 更新时间
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`                                                                    // 删除时间
	UserUUID     string         `json:"user_uuid" form:"user_uuid" gorm:"column:user_uuid;comment:;size:255;primary_key;"` //uuid字段
	ValidStart   *time.Time     `json:"validStart" form:"validStart" gorm:"column:valid_start;type:DATE;comment:;"`        //validStart字段
	ValidEnd     *time.Time     `json:"validEnd" form:"validEnd" gorm:"column:valid_end;type:DATE;comment:;"`              //validEnd字段
	StandardData int            `json:"standardData" form:"standardData" gorm:"column:standard_data;type:INT NOT NULL;comment:基础用量(MB);"`
	Enable       uint           `json:"enable" gorm:"column:enable;type:TINYINT UNSIGNED NOT NULL;default:0;"` // 1启用，0禁用
}

// TableName userExt表 UserExt自定义表名 la_user_ext
func (UserExt) TableName() string {
	return "la_user_ext"
}
