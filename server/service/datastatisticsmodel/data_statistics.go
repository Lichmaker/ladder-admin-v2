package datastatisticsmodel

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datastatisticsmodel"
	datastatisticsmodelReq "github.com/flipped-aurora/gin-vue-admin/server/model/datastatisticsmodel/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type DataStatisticsService struct {
}

// CreateDataStatistics 创建dataStatistics表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataStatisticsService *DataStatisticsService) CreateDataStatistics(dataStatistics *datastatisticsmodel.DataStatistics) (err error) {
	err = global.GVA_DB.Create(dataStatistics).Error
	return err
}

// DeleteDataStatistics 删除dataStatistics表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataStatisticsService *DataStatisticsService) DeleteDataStatistics(id string) (err error) {
	err = global.GVA_DB.Delete(&datastatisticsmodel.DataStatistics{}, "id = ?", id).Error
	return err
}

// DeleteDataStatisticsByIds 批量删除dataStatistics表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataStatisticsService *DataStatisticsService) DeleteDataStatisticsByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]datastatisticsmodel.DataStatistics{}, "id in ?", ids).Error
	return err
}

// UpdateDataStatistics 更新dataStatistics表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataStatisticsService *DataStatisticsService) UpdateDataStatistics(dataStatistics datastatisticsmodel.DataStatistics) (err error) {
	err = global.GVA_DB.Save(&dataStatistics).Error
	return err
}

// GetDataStatistics 根据id获取dataStatistics表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataStatisticsService *DataStatisticsService) GetDataStatistics(id string) (dataStatistics datastatisticsmodel.DataStatistics, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dataStatistics).Error
	return
}

// GetDataStatisticsInfoList 分页获取dataStatistics表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataStatisticsService *DataStatisticsService) GetDataStatisticsInfoList(info datastatisticsmodelReq.DataStatisticsSearch) (list []datastatisticsmodel.DataStatistics, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&datastatisticsmodel.DataStatistics{})
	var dataStatisticss []datastatisticsmodel.DataStatistics
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartDate != "" && info.EndDate != "" {
		db = db.Where("`date` BETWEEN ? AND ?", info.StartDate, info.EndDate)
	}
	if info.Email != "" {
		db = db.Where("`email` = ?", info.Email)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	db = db.Order("id DESC")

	err = db.Find(&dataStatisticss).Error
	dataStatisticsService.ByteValueHumanize(dataStatisticss)
	return dataStatisticss, total, err
}

func (dataStatisticsService *DataStatisticsService) ByteValueHumanize(data []datastatisticsmodel.DataStatistics) {
	for k, v := range data {
		data[k].ValueHumanize = utils.ParseByteSize(v.Value)
	}
}

func (dataStatisticsService *DataStatisticsService) SaveTargetDateData(email, nodeID string, date time.Time, value int) error {
	var currentModel datastatisticsmodel.DataStatistics
	if err := global.GVA_DB.Where("email = ? AND node_id = ? AND date = ?", email, nodeID, date.Format("2006-01-02")).First(&currentModel).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if currentModel.ID > 0 {
		return global.GVA_DB.Model(currentModel).Update("value", gorm.Expr("`value` + ?", value)).Error
	} else {
		currentModel.Email = email
		currentModel.NodeID = nodeID
		currentModel.Date = date
		currentModel.Value = value
		return global.GVA_DB.Model(currentModel).Create(&currentModel).Error
	}
}
