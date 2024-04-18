package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type LaDataPackageCodeSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
                      UniqueCode  string `json:"uniqueCode" form:"uniqueCode" `
                      Enable  *bool `json:"enable" form:"enable" `
                StartUsedAt  *time.Time  `json:"startUsedAt" form:"startUsedAt"`
                EndUsedAt  *time.Time  `json:"endUsedAt" form:"endUsedAt"`
    request.PageInfo
}
