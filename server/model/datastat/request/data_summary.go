package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DataSummarySearch struct {
	DataSummaryListData
	request.PageInfo
}

type DataSummaryListData struct {
	global.GVA_MODEL
	UplinkByte   *string `json:"uplinkByte" form:"uplinkByte"`
	DownlinkByte *string `json:"downlinkByte" form:"downlinkByte"`
	Date         string  `json:"date" form:"date"`
	Username     string  `json:"username" form:"username"`
}
