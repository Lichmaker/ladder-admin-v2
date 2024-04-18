// 自动生成模板LaDataPackageCode
package datapackagecode

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// dataPackageCode表 结构体  LaDataPackageCode
type DataPackageCode struct {
	global.GVA_MODEL
	UniqueCode   string     `json:"uniqueCode" form:"uniqueCode" gorm:"column:unique_code;comment:;size:64;unique;" binding:"required"` //唯一编号
	StandardData *int       `json:"standardData" form:"standardData" gorm:"column:standard_data;comment:;size:10;" binding:"required"`  //基础流量(MB)
	Days         *int       `json:"days" form:"days" gorm:"column:days;comment:;size:10;" binding:"required"`                           //有效天数
	Enable       bool       `json:"enable" form:"enable" gorm:"column:enable;comment:;notnull;default:0" binding:"required"`            //是否启用（禁用后不能兑换也不能使用）
	UsedAt       *time.Time `json:"usedAt" form:"usedAt" gorm:"column:used_at;comment:;"`                                               //被兑换时间
	Remark       string     `json:"remark" form:"remark" gorm:"column:remark;comment:;size:255;"`                                       //备注
	UsedBy       string     `json:"usedBy" form:"usedBy" gorm:"column:used_by;comment:;size:255;index;"`                                // 被兑换人uuid
	ConsumedAt   *time.Time `json:"consumedAt" form:"consumedAt" gorm:"column:consumed_at;comment:;"`                                   // 使用时间（兑换之后可以不立即使用）
}

// TableName dataPackageCode表 LaDataPackageCode自定义表名 la_data_package_code
func (DataPackageCode) TableName() string {
	return "la_data_package_code"
}
