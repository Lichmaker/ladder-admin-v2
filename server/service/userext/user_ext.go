package userext

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userext"
	userextReq "github.com/flipped-aurora/gin-vue-admin/server/model/userext/request"
	"github.com/pkg/errors"
)

type UserExtService struct {
}

// CreateUserExt 创建userExt表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userExtService *UserExtService) CreateUserExt(userExt *userext.UserExt) (err error) {
	err = global.GVA_DB.Create(userExt).Error
	return err
}

// DeleteUserExt 删除userExt表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userExtService *UserExtService) DeleteUserExt(id string) (err error) {
	err = global.GVA_DB.Delete(&userext.UserExt{}, "id = ?", id).Error
	return err
}

// DeleteUserExtByIds 批量删除userExt表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userExtService *UserExtService) DeleteUserExtByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]userext.UserExt{}, "id in ?", ids).Error
	return err
}

// UpdateUserExt 更新userExt表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userExtService *UserExtService) UpdateUserExt(userExt userext.UserExt) (err error) {
	err = global.GVA_DB.Save(&userExt).Error
	return err
}

// GetUserExt 根据id获取userExt表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userExtService *UserExtService) GetUserExt(id string) (userExt userext.UserExt, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&userExt).Error
	return
}

// GetUserExtInfoList 分页获取userExt表记录
// Author [piexlmax](https://github.com/piexlmax)
func (userExtService *UserExtService) GetUserExtInfoList(info userextReq.UserExtSearch) (list []userext.UserExt, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&userext.UserExt{})
	var userExts []userext.UserExt
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userExts).Error
	return userExts, total, err
}

func (userExtService *UserExtService) UpdateUserInfo(req userextReq.UpdateUserInfo) error {
	var user system.SysUser
	if err := global.GVA_DB.Model(user).First(&user, req.ID).Error; err != nil {
		return errors.Wrap(err, "查询用户失败")
	}

	err := global.GVA_DB.Model(&system.SysUser{}).
		Select("updated_at", "nick_name", "header_img", "phone", "email", "sideMode", "enable").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"nick_name":  req.NickName,
			"phone":      req.Phone,
			"email":      req.Email,
		}).Error
	if err != nil {
		return errors.Wrap(err, "保存用户信息失败")
	}

	if err := global.GVA_DB.Model(userext.UserExt{}).Select("valid_start", "valid_end", "standard_data", "enable").
		Where("`user_uuid` = ?", user.UUID).Updates(map[string]interface{}{
		"valid_start":   req.ValidStart,
		"valid_end":     req.ValidEnd,
		"standard_data": req.StandardDataMB,
		"enable":        req.Enable,
	}).Error; err != nil {
		return errors.Wrap(err, "保存用户扩展信息失败")
	}
	return nil
}

func (userExtService *UserExtService) UpdateEnable(uuid string, enableValue int) error {
	db := global.GVA_DB.Model(&userext.UserExt{})
	return db.Select("enable").Where("user_uuid = ?", uuid).Update("enable", enableValue).Error
}
