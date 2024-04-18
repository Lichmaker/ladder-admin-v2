// 自动生成模板LaDataPackageCode
package datapackagecode

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// dataPackageCode表 结构体  LaDataPackageCode
type DataPackageCode struct {
	global.GVA_MODEL
	UniqueCode   string     `json:"uniqueCode" form:"uniqueCode" gorm:"column:unique_code;comment:;size:64;" binding:"required"`       //唯一编号
	StandardData *int       `json:"standardData" form:"standardData" gorm:"column:standard_data;comment:;size:10;" binding:"required"` //基础流量(MB)
	Days         *int       `json:"days" form:"days" gorm:"column:days;comment:;size:10;" binding:"required"`                          //有效天数
	Enable       *bool      `json:"enable" form:"enable" gorm:"column:enable;comment:;" binding:"required"`                            //是否启用
	UsedAt       *time.Time `json:"usedAt" form:"usedAt" gorm:"column:used_at;comment:;"`                                              //被使用时间
	Remark       string     `json:"remark" form:"remark" gorm:"column:remark;comment:;size:255;"`                                      //备注
	UsedBy       string     `json:"usedBy" form:"usedBy" gorm:"column:used_by;comment:;size:255;"`                                     // 被使用人uuid
}

// TableName dataPackageCode表 LaDataPackageCode自定义表名 la_data_package_code
func (DataPackageCode) TableName() string {
	return "la_data_package_code"
}
