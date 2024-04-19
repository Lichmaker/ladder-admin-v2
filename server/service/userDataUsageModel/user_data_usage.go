package userDataUsageModel

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userDataUsageModel"
	userDataUsageModelReq "github.com/flipped-aurora/gin-vue-admin/server/model/userDataUsageModel/request"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

type UserDataUsageService struct {
}

// CreateUserDataUsage 创建userDataUsage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataUsageService *UserDataUsageService) CreateUserDataUsage(userDataUsage *userDataUsageModel.UserDataUsage) (err error) {
	err = global.GVA_DB.Create(userDataUsage).Error
	return err
}

func (userDataUsageService *UserDataUsageService) CreateUserDataUsageWithDB(db *gorm.DB, userDataUsage *userDataUsageModel.UserDataUsage) (err error) {
	err = db.Create(userDataUsage).Error
	return err
}

// DeleteUserDataUsage 删除userDataUsage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataUsageService *UserDataUsageService) DeleteUserDataUsage(id string) (err error) {
	err = global.GVA_DB.Delete(&userDataUsageModel.UserDataUsage{}, "id = ?", id).Error
	return err
}

// DeleteUserDataUsageByIds 批量删除userDataUsage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataUsageService *UserDataUsageService) DeleteUserDataUsageByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]userDataUsageModel.UserDataUsage{}, "id in ?", ids).Error
	return err
}

// UpdateUserDataUsage 更新userDataUsage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataUsageService *UserDataUsageService) UpdateUserDataUsage(userDataUsage userDataUsageModel.UserDataUsage) (err error) {
	err = global.GVA_DB.Save(&userDataUsage).Error
	return err
}

// GetUserDataUsage 根据id获取userDataUsage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataUsageService *UserDataUsageService) GetUserDataUsage(id string) (userDataUsage userDataUsageModel.UserDataUsage, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userDataUsage).Error
	return
}

// GetUserDataUsageInfoList 分页获取userDataUsage表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userDataUsageService *UserDataUsageService) GetUserDataUsageInfoList(info userDataUsageModelReq.UserDataUsageSearch) (list []userDataUsageModel.UserDataUsage, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&userDataUsageModel.UserDataUsage{})
	var userDataUsages []userDataUsageModel.UserDataUsage
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if len(info.UUID) > 0 {
		db = db.Where("`user_uuid` = ?", info.UUID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	db = db.Order("id DESC")

	err = db.Find(&userDataUsages).Error
	return userDataUsages, total, err
}

func (userDataUsageService *UserDataUsageService) GetUserUsingData(uuid string, date time.Time) (result userDataUsageModel.UserDataUsage, err error) {
	if len(uuid) == 0 {
		return
	}
	current := carbon.CreateFromStdTime(date).Layout("2006-01-02")

	var query []userDataUsageModel.UserDataUsage
	db := global.GVA_DB.Model(result)
	db = db.Where("`user_uuid` = ? AND start_date <= ? AND end_date >= ?", uuid, current, current).Order("id ASC")
	if err = db.Find(&query).Error; err != nil {
		return
	}
	// 遍历，找出有剩余流量的
	for _, v := range query {
		if v.StandardData*1024*1024-v.Usage > 0 {
			result = v
			return
		}
	}

	return result, err
}

func (userDataUsageService *UserDataUsageService) SaveTargetDateData(uuid string, date time.Time, value int) error {
	currentModel, err := userDataUsageService.GetUserUsingData(uuid, date)
	if err != nil {
		return err
	}

	if currentModel.ID > 0 {
		return global.GVA_DB.Model(currentModel).Update("usage", gorm.Expr("`usage` + ?", value)).Error
	}
	return nil
}
