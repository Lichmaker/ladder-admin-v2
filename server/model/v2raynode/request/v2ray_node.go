package request

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type V2rayNodeSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type GenVemssUrlRequest struct {
	UUID string `form:"UUID"`
}

type V2rayNodePushDataReq struct {
	All    bool   `json:"all" form:"all"`
	NodeID string `json:"nodeID" form:"nodeID"`
}
