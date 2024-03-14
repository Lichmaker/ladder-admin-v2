package invitecode

import (
	"context"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/invitecode"
	invitecodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/invitecode/request"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type InviteCodeService struct {
}

// CreateInviteCode 创建inviteCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (inviteCodeService *InviteCodeService) CreateInviteCode(inviteCode *invitecode.InviteCode) (err error) {
	err = global.GVA_DB.Create(inviteCode).Error
	return err
}

// DeleteInviteCode 删除inviteCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (inviteCodeService *InviteCodeService) DeleteInviteCode(id string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&invitecode.InviteCode{}).Where("id = ?", id).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&invitecode.InviteCode{}, "id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteInviteCodeByIds 批量删除inviteCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (inviteCodeService *InviteCodeService) DeleteInviteCodeByIds(ids []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&invitecode.InviteCode{}).Where("id in ?", ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids).Delete(&invitecode.InviteCode{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateInviteCode 更新inviteCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (inviteCodeService *InviteCodeService) UpdateInviteCode(inviteCode invitecode.InviteCode) (err error) {
	err = global.GVA_DB.Save(&inviteCode).Error
	return err
}

// GetInviteCode 根据id获取inviteCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (inviteCodeService *InviteCodeService) GetInviteCode(id string) (inviteCode invitecode.InviteCode, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&inviteCode).Error
	return
}

// GetInviteCodeInfoList 分页获取inviteCode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (inviteCodeService *InviteCodeService) GetInviteCodeInfoList(info invitecodeReq.InviteCodeSearch) (list []invitecode.InviteCode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&invitecode.InviteCode{})
	var inviteCodes []invitecode.InviteCode
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if len(info.CreatorUUID) > 0 {
		db = db.Where("creator_uuid = ?", info.CreatorUUID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&inviteCodes).Error
	return inviteCodes, total, err
}

func (inviteCodeService *InviteCodeService) VerifyMyInviteCodeIds(uuid string, ids []string) (bool, error) {
	db := global.GVA_DB.Model(&invitecode.InviteCode{})
	var count int64
	if err := db.Where("creator_uuid = ? AND id in ?", uuid, ids).Count(&count).Error; err != nil {
		return false, err
	}
	if count != int64(len(ids)) {
		return false, nil
	}
	return true, nil
}

func (inviteCodeService *InviteCodeService) inviteCodeLockKey(code string) string {
	return fmt.Sprintf("lock_%s", code)
}

func (inviteCodeService *InviteCodeService) LockInviteCode(code string) (bool, error) {
	result, err := global.GVA_REDIS.SetNX(context.Background(), inviteCodeService.inviteCodeLockKey(code), 1, time.Second*10).Result()
	if err != nil {
		return false, errors.Wrap(err, "redis异常")
	}
	return result, nil
}

func (inviteCodeService *InviteCodeService) UnlockInviteCode(code string) error {
	_, err := global.GVA_REDIS.Del(context.Background(), inviteCodeService.inviteCodeLockKey(code)).Result()
	return err
}

func (inviteCodeService *InviteCodeService) GetByCode(code string) (result invitecode.InviteCode, err error) {
	db := global.GVA_DB.Model(&invitecode.InviteCode{})
	err = db.Where("`code` = ?", code).First(&result).Error
	return
}
