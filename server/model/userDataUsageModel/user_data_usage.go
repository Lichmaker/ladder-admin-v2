// 自动生成模板UserDataUsage
package userDataUsageModel

import (
	"time"
)

// userDataUsage表 结构体  UserDataUsage
type UserDataUsage struct {
	ID           uint       `gorm:"primaryKey;autoIncrement;" json:"ID"` // 主键ID
	CreatedAt    time.Time  // 创建时间
	UpdatedAt    time.Time  // 更新时间
	UserUuid     string     `json:"userUuid" form:"userUuid" gorm:"column:user_uuid;comment:;size:255;index;"`                //
	StartDate    *time.Time `json:"startDate" form:"startDate" gorm:"column:start_date;comment:;"`                            //
	EndDate      *time.Time `json:"endDate" form:"endDate" gorm:"column:end_date;comment:;"`                                  //
	StandardData int        `json:"standardData" form:"standardData" gorm:"column:standard_data;comment:;type:INT NOT NULL;"` // 单位是MB
	Usage        int        `json:"usage" form:"usage" gorm:"column:usage;comment:;type:BIGINT NOT NULL;"`                    // 单位是Byte
}

// TableName userDataUsage表 UserDataUsage自定义表名 la_user_data_usage
func (UserDataUsage) TableName() string {
	return "la_user_data_usage"
}
