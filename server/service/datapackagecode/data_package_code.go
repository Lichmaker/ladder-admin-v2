package datapackagecode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datapackagecode"
	datapackagecodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/datapackagecode/request"
)

type LaDataPackageCodeService struct {
}

// CreateLaDataPackageCode 创建dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *LaDataPackageCodeService) CreateLaDataPackageCode(dataPackageCode *datapackagecode.LaDataPackageCode) (err error) {
	err = global.GVA_DB.Create(dataPackageCode).Error
	return err
}

// DeleteLaDataPackageCode 删除dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *LaDataPackageCodeService) DeleteLaDataPackageCode(id string) (err error) {
	err = global.GVA_DB.Delete(&datapackagecode.LaDataPackageCode{}, "id = ?", id).Error
	return err
}

// DeleteLaDataPackageCodeByIds 批量删除dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *LaDataPackageCodeService) DeleteLaDataPackageCodeByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]datapackagecode.LaDataPackageCode{}, "id in ?", ids).Error
	return err
}

// UpdateLaDataPackageCode 更新dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *LaDataPackageCodeService) UpdateLaDataPackageCode(dataPackageCode datapackagecode.LaDataPackageCode) (err error) {
	err = global.GVA_DB.Save(&dataPackageCode).Error
	return err
}

// GetLaDataPackageCode 根据id获取dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *LaDataPackageCodeService) GetLaDataPackageCode(id string) (dataPackageCode datapackagecode.LaDataPackageCode, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&dataPackageCode).Error
	return
}

// GetLaDataPackageCodeInfoList 分页获取dataPackageCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (dataPackageCodeService *LaDataPackageCodeService) GetLaDataPackageCodeInfoList(info datapackagecodeReq.LaDataPackageCodeSearch) (list []datapackagecode.LaDataPackageCode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&datapackagecode.LaDataPackageCode{})
	var dataPackageCodes []datapackagecode.LaDataPackageCode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UniqueCode != "" {
		db = db.Where("unique_code = ?", info.UniqueCode)
	}
	if info.Enable != nil {
		db = db.Where("enable = ?", info.Enable)
	}
	if info.StartUsedAt != nil && info.EndUsedAt != nil {
		db = db.Where("used_at BETWEEN ? AND ? ", info.StartUsedAt, info.EndUsedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&dataPackageCodes).Error
	return dataPackageCodes, total, err
}
