package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/datastat"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DataSummarySearch struct{
    datastat.DataSummary
    request.PageInfo
}
