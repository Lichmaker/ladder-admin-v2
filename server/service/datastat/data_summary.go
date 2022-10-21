package datastat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datastat"
	datastatReq "github.com/flipped-aurora/gin-vue-admin/server/model/datastat/request"
)

type DataSummaryService struct {
}

// CreateDataSummary 创建DataSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataSummaryService *DataSummaryService) CreateDataSummary(dataSummary datastat.DataSummary) (err error) {
	err = global.GVA_DB.Create(&dataSummary).Error
	return err
}

// DeleteDataSummary 删除DataSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataSummaryService *DataSummaryService) DeleteDataSummary(dataSummary datastat.DataSummary) (err error) {
	err = global.GVA_DB.Delete(&dataSummary).Error
	return err
}

// DeleteDataSummaryByIds 批量删除DataSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataSummaryService *DataSummaryService) DeleteDataSummaryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]datastat.DataSummary{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateDataSummary 更新DataSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataSummaryService *DataSummaryService) UpdateDataSummary(dataSummary datastat.DataSummary) (err error) {
	err = global.GVA_DB.Save(&dataSummary).Error
	return err
}

// GetDataSummary 根据id获取DataSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataSummaryService *DataSummaryService) GetDataSummary(id uint) (dataSummary datastat.DataSummary, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dataSummary).Error
	return
}

// GetDataSummaryInfoList 分页获取DataSummary记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataSummaryService *DataSummaryService) GetDataSummaryInfoList(info datastatReq.DataSummarySearch) (list []datastat.DataSummary, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&datastat.DataSummary{})
	var dataSummarys []datastat.DataSummary
	// 如果有条件搜索 下方会自动创建搜索语句
	if len(info.Date) > 0 {
		db = db.Where("date = ?", info.Date)
	}
	if info.Username != "" {
		db = db.Where("username = ?", info.Username)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&dataSummarys).Error
	return dataSummarys, total, err
}
