package v2raynode

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/v2raynode"
	v2raynodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/v2raynode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/pkg/errors"
)

type V2rayNodeService struct {
}

// CreateV2rayNode 创建v2rayNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayNodeService *V2rayNodeService) CreateV2rayNode(v2rayNode *v2raynode.V2rayNode) (err error) {
	v2rayNode.LastContactAt = nil

	// 保存密码
	if len(global.GVA_CONFIG.V2RayManager.SecretKey) > 0 {
		secretRaw := v2rayNode.Secret
		v2rayNode.SecretIv = utils.GenAESRandomIV()
		v2rayNode.Secret, err = utils.AESEncrypt(global.GVA_CONFIG.V2RayManager.SecretKey, v2rayNode.SecretIv, secretRaw)
		if err != nil {
			return errors.Wrap(err, "密码加密失败")
		}
	}

	err = global.GVA_DB.Create(v2rayNode).Error
	return err
}

// DeleteV2rayNode 删除v2rayNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayNodeService *V2rayNodeService) DeleteV2rayNode(id string) (err error) {
	err = global.GVA_DB.Delete(&v2raynode.V2rayNode{}, "id = ?", id).Error
	return err
}

// DeleteV2rayNodeByIds 批量删除v2rayNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayNodeService *V2rayNodeService) DeleteV2rayNodeByIds(ids []string) (err error) {
	err = global.GVA_DB.Delete(&[]v2raynode.V2rayNode{}, "id in ?", ids).Error
	return err
}

// UpdateV2rayNode 更新v2rayNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayNodeService *V2rayNodeService) UpdateV2rayNode(v2rayNode *v2raynode.V2rayNode) (err error) {
	db := global.GVA_DB.Model(v2raynode.V2rayNode{}).Omit("last_contact_at")

	// 保存密码
	if len(v2rayNode.Secret) > 0 && len(global.GVA_CONFIG.V2RayManager.SecretKey) > 0 {
		secretRaw := v2rayNode.Secret
		v2rayNode.SecretIv = utils.GenAESRandomIV()
		v2rayNode.Secret, err = utils.AESEncrypt(global.GVA_CONFIG.V2RayManager.SecretKey, v2rayNode.SecretIv, secretRaw)
		if err != nil {
			return errors.Wrap(err, "密码加密失败")
		}
	} else {
		db = db.Omit("secret", "secret_iv")
	}

	err = db.Where("id = ?", v2rayNode.ID).Save(&v2rayNode).Error
	return err
}

// GetV2rayNode 根据id获取v2rayNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayNodeService *V2rayNodeService) GetV2rayNode(id string) (v2rayNode v2raynode.V2rayNode, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&v2rayNode).Error
	return
}

func (v2rayNodeService *V2rayNodeService) GetByUniqueID(id string) (v2rayNode v2raynode.V2rayNode, err error) {
	err = global.GVA_DB.Where("unique_id = ?", id).First(&v2rayNode).Error
	return
}

// GetV2rayNodeInfoList 分页获取v2rayNode表记录
// Author [piexlmax](https://github.com/piexlmax)
func (v2rayNodeService *V2rayNodeService) GetV2rayNodeInfoList(info v2raynodeReq.V2rayNodeSearch) (list []v2raynode.V2rayNode, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&v2raynode.V2rayNode{})
	var v2rayNodes []v2raynode.V2rayNode
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

	err = db.Find(&v2rayNodes).Error
	for k, v := range v2rayNodes {
		v2rayNodes[k].Online = v.CheckOnline()
	}
	return v2rayNodes, total, err
}

func (v2rayNodeService *V2rayNodeService) GetAllNode() (list []v2raynode.V2rayNode, err error) {
	db := global.GVA_DB.Model(&v2raynode.V2rayNode{})
	var v2rayNodes []v2raynode.V2rayNode
	err = db.Find(&v2rayNodes).Error
	return v2rayNodes, err
}

func (v2rayNodeService *V2rayNodeService) UpdateLastContactAt(uniqueID string, t time.Time) error {
	db := global.GVA_DB.Model(&v2raynode.V2rayNode{})
	return db.Where("unique_id", uniqueID).UpdateColumn("last_contact_at", t).Error
}

func (v2rayNodeService *V2rayNodeService) UpdateVmessPort(uniqueID string, port uint) error {
	db := global.GVA_DB.Model(&v2raynode.V2rayNode{})
	return db.Where("unique_id", uniqueID).UpdateColumn("vmess_port", port).Error
}
