// 自动生成模板DataStatistics
package datastatisticsmodel

import "time"

// dataStatistics表 结构体  DataStatistics
type DataStatistics struct {
	ID            uint      `gorm:"primaryKey;autoIncrement;" json:"ID"` // 主键ID
	CreatedAt     time.Time // 创建时间
	UpdatedAt     time.Time // 更新时间
	Email         string    `json:"email" form:"email" gorm:"column:email;comment:;type:varchar(191) NOT NULL;"`
	NodeID        string    `json:"nodeID" form:"nodeID" gorm:"column:node_id;comment:;type:varchar(64) NOT NULL;"`
	Date          time.Time `json:"date" form:"date" gorm:"column:date;type:DATE not null;comment:;"`      //date字段
	Value         int       `json:"value" form:"value" gorm:"column:value;type:BIGINT UNSIGNED NOT NULL;"` //value字段
	ValueHumanize string    `json:"valueHumanize,omitempty" gorm:"-:all"`
}

// TableName dataStatistics表 DataStatistics自定义表名 data_statistics
func (DataStatistics) TableName() string {
	return "la_data_statistics"
}
