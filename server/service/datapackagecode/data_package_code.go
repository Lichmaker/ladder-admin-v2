package datapackagecode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datapackagecode"
	datapackagecodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/datapackagecode/request"
	"gorm.io/gorm"
)

type DataPackageCodeService struct {
}

// CreateDataPackageCode 创建dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *DataPackageCodeService) CreateDataPackageCode(dataPackageCode *datapackagecode.DataPackageCode) (err error) {
	err = global.GVA_DB.Create(dataPackageCode).Error
	return err
}

// DeleteDataPackageCode 删除dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *DataPackageCodeService) DeleteDataPackageCode(id string) (err error) {
	err = global.GVA_DB.Delete(&datapackagecode.DataPackageCode{}, "id = ?", id).Error
	return err
}

// DeleteDataPackageCodeByIds 批量删除dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *DataPackageCodeService) DeleteDataPackageCodeByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]datapackagecode.DataPackageCode{}, "id in ?", ids).Error
	return err
}

// UpdateDataPackageCode 更新dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *DataPackageCodeService) UpdateDataPackageCode(dataPackageCode datapackagecode.DataPackageCode) (err error) {
	err = global.GVA_DB.Save(&dataPackageCode).Error
	return err
}

func (dataPackageCodeService *DataPackageCodeService) UpdateDataPackageCodeWithDB(db *gorm.DB, dataPackageCode datapackagecode.DataPackageCode) (err error) {
	err = db.Save(&dataPackageCode).Error
	return err
}

// GetDataPackageCode 根据id获取dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *DataPackageCodeService) GetDataPackageCode(id string) (dataPackageCode datapackagecode.DataPackageCode, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dataPackageCode).Error
	return
}

// GetDataPackageCodeInfoList 分页获取dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *DataPackageCodeService) GetDataPackageCodeInfoList(info datapackagecodeReq.DataPackageCodeSearch, orderBy string) (list []datapackagecode.DataPackageCode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&datapackagecode.DataPackageCode{})
	var dataPackageCodes []datapackagecode.DataPackageCode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UniqueCode != "" {
		db = db.Where("unique_code LIKE ?", "%"+info.UniqueCode+"%")
	}
	if info.Enable != nil {
		db = db.Where("enable = ?", info.Enable)
	}
	if info.SearchUnused != nil {
		if *info.SearchUnused {
			db = db.Where("used_at IS NULL")
		} else {
			db = db.Where("used_at IS NOT NULL")
		}
	}
	if info.SearchUnconsumed != nil {
		if *info.SearchUnconsumed {
			db = db.Where("consumed_at IS NULL")
		} else {
			db = db.Where("consumed_at IS NOT NULL")
		}
	}
	if info.StartUsedAt != nil && info.EndUsedAt != nil {
		db = db.Where("used_at BETWEEN ? AND ? ", info.StartUsedAt, info.EndUsedAt)
	}
	if len(info.UsedBy) > 0 {
		db = db.Where("used_by = ?", info.UsedBy)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	if orderBy != "" {
		db = db.Order(orderBy)
	}

	err = db.Find(&dataPackageCodes).Error
	return dataPackageCodes, total, err
}

func (dataPackageCodeService *DataPackageCodeService) BatchCreateDataPackageCode(db *gorm.DB, dataPackageCode []datapackagecode.DataPackageCode) (err error) {
	err = db.CreateInBatches(dataPackageCode, 20).Error
	return err
}

func (dataPackageCodeService *DataPackageCodeService) GetByCode(code string) (data datapackagecode.DataPackageCode, err error) {
	err = global.GVA_DB.Model(datapackagecode.DataPackageCode{}).Where("`unique_code` = ?", code).First(&data).Error
	return
}
