// 自动生成模板DataSummary
package datastat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// DataSummary 结构体
type DataSummary struct {
      global.GVA_MODEL
      UplinkByte  *int `json:"uplinkByte" form:"uplinkByte" gorm:"column:uplink_byte;comment:;size:20;"`
      DownlinkByte  *int `json:"downlinkByte" form:"downlinkByte" gorm:"column:downlink_byte;comment:;size:20;"`
      Date  *time.Time `json:"date" form:"date" gorm:"column:date;comment:;"`
      Username  string `json:"username" form:"username" gorm:"column:username;comment:;size:64;"`
}


// TableName DataSummary 表名
func (DataSummary) TableName() string {
  return "data_summary"
}

