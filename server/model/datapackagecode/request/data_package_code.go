package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DataPackageCodeSearch struct {
	StartCreatedAt   *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt     *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	UniqueCode       string     `json:"uniqueCode" form:"uniqueCode" `
	Enable           *bool      `json:"enable" form:"enable"`
	SearchUnused     *bool      `json:"searchUnused" form:"searchUnused"`
	SearchUnconsumed *bool      `json:"searchUnconsumed" form:"searchUnconsumed"`
	UsedBy           string     `json:"usedBy" form:"usedBy"`
	StartUsedAt      *time.Time `json:"startUsedAt" form:"startUsedAt"`
	EndUsedAt        *time.Time `json:"endUsedAt" form:"endUsedAt"`
	request.PageInfo
}

type BatchCreateDataPackageCode struct {
	CodePrefix   string `json:"codePrefix"`
	StandardData uint   `json:"standardData"` //基础流量(MB)
	Days         uint   `json:"days"`         //有效天数
	Remark       string `json:"remark"`       //有效天数
	Total        uint   `json:"total"`
}

type ExchangeDataPackageCode struct {
	CaptchaId string `json:"captchaId"` // 验证码ID
	Captcha   string `json:"captcha"`
	Code      string `json:"code"`
}

type ConsumePackageReq struct {
	PackageID int `json:"packageID"`
}
