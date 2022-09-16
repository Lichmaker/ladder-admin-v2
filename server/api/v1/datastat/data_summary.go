package datastat

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datastat"
	datastatReq "github.com/flipped-aurora/gin-vue-admin/server/model/datastat/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DataSummaryApi struct {
}

var dataSummaryService = service.ServiceGroupApp.DatastatServiceGroup.DataSummaryService

// CreateDataSummary 创建DataSummary
// @Tags DataSummary
// @Summary 创建DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datastat.DataSummary true "创建DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataSummary/createDataSummary [post]
func (dataSummaryApi *DataSummaryApi) CreateDataSummary(c *gin.Context) {
	var dataSummary datastat.DataSummary
	_ = c.ShouldBindJSON(&dataSummary)
	if err := dataSummaryService.CreateDataSummary(dataSummary); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDataSummary 删除DataSummary
// @Tags DataSummary
// @Summary 删除DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datastat.DataSummary true "删除DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataSummary/deleteDataSummary [delete]
func (dataSummaryApi *DataSummaryApi) DeleteDataSummary(c *gin.Context) {
	var dataSummary datastat.DataSummary
	_ = c.ShouldBindJSON(&dataSummary)
	if err := dataSummaryService.DeleteDataSummary(dataSummary); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDataSummaryByIds 批量删除DataSummary
// @Tags DataSummary
// @Summary 批量删除DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dataSummary/deleteDataSummaryByIds [delete]
func (dataSummaryApi *DataSummaryApi) DeleteDataSummaryByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := dataSummaryService.DeleteDataSummaryByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDataSummary 更新DataSummary
// @Tags DataSummary
// @Summary 更新DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datastat.DataSummary true "更新DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataSummary/updateDataSummary [put]
func (dataSummaryApi *DataSummaryApi) UpdateDataSummary(c *gin.Context) {
	var dataSummary datastat.DataSummary
	_ = c.ShouldBindJSON(&dataSummary)
	if err := dataSummaryService.UpdateDataSummary(dataSummary); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDataSummary 用id查询DataSummary
// @Tags DataSummary
// @Summary 用id查询DataSummary
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query datastat.DataSummary true "用id查询DataSummary"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataSummary/findDataSummary [get]
func (dataSummaryApi *DataSummaryApi) FindDataSummary(c *gin.Context) {
	var dataSummary datastat.DataSummary
	_ = c.ShouldBindQuery(&dataSummary)
	if redataSummary, err := dataSummaryService.GetDataSummary(dataSummary.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redataSummary": redataSummary}, c)
	}
}

// GetDataSummaryList 分页获取DataSummary列表
// @Tags DataSummary
// @Summary 分页获取DataSummary列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query datastatReq.DataSummarySearch true "分页获取DataSummary列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataSummary/getDataSummaryList [get]
func (dataSummaryApi *DataSummaryApi) GetDataSummaryList(c *gin.Context) {
	var pageInfo datastatReq.DataSummarySearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := dataSummaryService.GetDataSummaryInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		type TestDataSummary struct {
			global.GVA_MODEL
			UplinkByte   *int       `json:"uplinkByte" form:"uplinkByte" gorm:"column:uplink_byte;comment:;size:20;"`
			DownlinkByte *int       `json:"downlinkByte" form:"downlinkByte" gorm:"column:downlink_byte;comment:;size:20;"`
			Date         *time.Time `json:"date" form:"date" gorm:"column:date;comment:;"`
			Username     string     `json:"username" form:"username" gorm:"column:username;comment:;size:64;"`
			JinDuTiao    string     `json:"jinDuTiao"`
		}
		var testList []TestDataSummary
		for _, v := range list {
			testList = append(testList, TestDataSummary{
				GVA_MODEL:    v.GVA_MODEL,
				UplinkByte:   v.UplinkByte,
				DownlinkByte: v.DownlinkByte,
				Date:         v.Date,
				Username:     v.Username,
				JinDuTiao:    "55",
			})
		}

		response.OkWithDetailed(response.PageResult{
			List:     testList,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
