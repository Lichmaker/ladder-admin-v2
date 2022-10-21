package v2rayUser

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/v2rayUser"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    v2rayUserReq "github.com/flipped-aurora/gin-vue-admin/server/model/v2rayUser/request"
)

type V2rayUserService struct {
}

// CreateV2rayUser 创建V2rayUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayUserService *V2rayUserService) CreateV2rayUser(v2rayUser v2rayUser.V2rayUser) (err error) {
	err = global.GVA_DB.Create(&v2rayUser).Error
	return err
}

// DeleteV2rayUser 删除V2rayUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayUserService *V2rayUserService)DeleteV2rayUser(v2rayUser v2rayUser.V2rayUser) (err error) {
	err = global.GVA_DB.Delete(&v2rayUser).Error
	return err
}

// DeleteV2rayUserByIds 批量删除V2rayUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayUserService *V2rayUserService)DeleteV2rayUserByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]v2rayUser.V2rayUser{},"id in ?",ids.Ids).Error
	return err
}

// UpdateV2rayUser 更新V2rayUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayUserService *V2rayUserService)UpdateV2rayUser(v2rayUser v2rayUser.V2rayUser) (err error) {
	err = global.GVA_DB.Save(&v2rayUser).Error
	return err
}

// GetV2rayUser 根据id获取V2rayUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayUserService *V2rayUserService)GetV2rayUser(id uint) (v2rayUser v2rayUser.V2rayUser, err error) {
	err = global.GVA_DB.Where("admin_user_id = ?", id).First(&v2rayUser).Error
	return
}

// GetV2rayUserInfoList 分页获取V2rayUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayUserService *V2rayUserService)GetV2rayUserInfoList(info v2rayUserReq.V2rayUserSearch) (list []v2rayUser.V2rayUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&v2rayUser.V2rayUser{})
    var v2rayUsers []v2rayUser.V2rayUser
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.AdminUserId != nil {
        db = db.Where("admin_user_id = ?",info.AdminUserId)
    }
    if info.Email != "" {
        db = db.Where("email = ?",info.Email)
    }
    if info.ActiveState != nil {
        db = db.Where("active_state = ?",info.ActiveState)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&v2rayUsers).Error
	return  v2rayUsers, total, err
}
