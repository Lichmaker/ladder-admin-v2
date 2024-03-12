package announcementModel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/announcementModel"
    announcementModelReq "github.com/flipped-aurora/gin-vue-admin/server/model/announcementModel/request"
)

type AnnouncementService struct {
}

// CreateAnnouncement 创建announcement表记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService) CreateAnnouncement(announcement *announcementModel.Announcement) (err error) {
	err = global.GVA_DB.Create(announcement).Error
	return err
}

// DeleteAnnouncement 删除announcement表记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService)DeleteAnnouncement(id string) (err error) {
	err = global.GVA_DB.Delete(&announcementModel.Announcement{},"id = ?",id).Error
	return err
}

// DeleteAnnouncementByIds 批量删除announcement表记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService)DeleteAnnouncementByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]announcementModel.Announcement{},"id in ?",ids).Error
	return err
}

// UpdateAnnouncement 更新announcement表记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService)UpdateAnnouncement(announcement announcementModel.Announcement) (err error) {
	err = global.GVA_DB.Save(&announcement).Error
	return err
}

// GetAnnouncement 根据id获取announcement表记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService)GetAnnouncement(id string) (announcement announcementModel.Announcement, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&announcement).Error
	return
}

// GetAnnouncementInfoList 分页获取announcement表记录
// Author [piexlmax](https://github.com/piexlmax)
func (announcementService *AnnouncementService)GetAnnouncementInfoList(info announcementModelReq.AnnouncementSearch) (list []announcementModel.Announcement, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&announcementModel.Announcement{})
    var announcements []announcementModel.Announcement
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	db = db.Order("id DESC")
	
	err = db.Find(&announcements).Error
	return  announcements, total, err
}
