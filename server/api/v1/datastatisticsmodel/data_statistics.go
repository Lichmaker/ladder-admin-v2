package datastatisticsmodel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/datastatisticsmodel"
	datastatisticsmodelReq "github.com/flipped-aurora/gin-vue-admin/server/model/datastatisticsmodel/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/constant"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DataStatisticsApi struct {
}

var dataStatisticsService = service.ServiceGroupApp.DatastatisticsmodelServiceGroup.DataStatisticsService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

// CreateDataStatistics 创建dataStatistics表
// @Tags DataStatistics
// @Summary 创建dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datastatisticsmodel.DataStatistics true "创建dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /dataStatistics/createDataStatistics [post]
func (dataStatisticsApi *DataStatisticsApi) CreateDataStatistics(c *gin.Context) {
	var dataStatistics datastatisticsmodel.DataStatistics
	err := c.ShouldBindJSON(&dataStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dataStatisticsService.CreateDataStatistics(&dataStatistics); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteDataStatistics 删除dataStatistics表
// @Tags DataStatistics
// @Summary 删除dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datastatisticsmodel.DataStatistics true "删除dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /dataStatistics/deleteDataStatistics [delete]
func (dataStatisticsApi *DataStatisticsApi) DeleteDataStatistics(c *gin.Context) {
	id := c.Query("ID")
	if err := dataStatisticsService.DeleteDataStatistics(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteDataStatisticsByIds 批量删除dataStatistics表
// @Tags DataStatistics
// @Summary 批量删除dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /dataStatistics/deleteDataStatisticsByIds [delete]
func (dataStatisticsApi *DataStatisticsApi) DeleteDataStatisticsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := dataStatisticsService.DeleteDataStatisticsByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateDataStatistics 更新dataStatistics表
// @Tags DataStatistics
// @Summary 更新dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body datastatisticsmodel.DataStatistics true "更新dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /dataStatistics/updateDataStatistics [put]
func (dataStatisticsApi *DataStatisticsApi) UpdateDataStatistics(c *gin.Context) {
	var dataStatistics datastatisticsmodel.DataStatistics
	err := c.ShouldBindJSON(&dataStatistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := dataStatisticsService.UpdateDataStatistics(dataStatistics); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindDataStatistics 用id查询dataStatistics表
// @Tags DataStatistics
// @Summary 用id查询dataStatistics表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query datastatisticsmodel.DataStatistics true "用id查询dataStatistics表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /dataStatistics/findDataStatistics [get]
func (dataStatisticsApi *DataStatisticsApi) FindDataStatistics(c *gin.Context) {
	id := c.Query("ID")
	if redataStatistics, err := dataStatisticsService.GetDataStatistics(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"redataStatistics": redataStatistics}, c)
	}
}

// GetDataStatisticsList 分页获取dataStatistics表列表
// @Tags DataStatistics
// @Summary 分页获取dataStatistics表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query datastatisticsmodelReq.DataStatisticsSearch true "分页获取dataStatistics表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /dataStatistics/getDataStatisticsList [get]
func (dataStatisticsApi *DataStatisticsApi) GetDataStatisticsList(c *gin.Context) {
	var pageInfo datastatisticsmodelReq.DataStatisticsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 通过token查询用户信息
	userInfoModel, err := userService.GetUserInfoViaToken(c)
	if err != nil {
		global.GVA_LOG.Error("获取用户信息!", zap.Error(err))
		response.FailWithMessage("读取用户失败", c)
		return
	}

	if userInfoModel.Authority.AuthorityId != constant.AUTHORITY_ID_SUPER_ADMIN {
		// 一般用户访问，仅能查自己的数据
		pageInfo.Email = userInfoModel.Email
	}

	if list, total, err := dataStatisticsService.GetDataStatisticsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
