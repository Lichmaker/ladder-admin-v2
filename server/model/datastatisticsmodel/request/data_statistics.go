package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type DataStatisticsSearch struct {
	Email     string `json:"email" form:"email"`
	StartDate string `json:"startDate" form:"startDate"` // example: 2024-01-02
	EndDate   string `json:"endDate" form:"endDate"`     // example: 2024-01-02
	request.PageInfo
}
