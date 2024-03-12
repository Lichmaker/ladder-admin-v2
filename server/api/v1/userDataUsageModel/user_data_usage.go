package userDataUsageModel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userDataUsageModel"
	userDataUsageModelReq "github.com/flipped-aurora/gin-vue-admin/server/model/userDataUsageModel/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserDataUsageApi struct {
}

var userDataUsageService = service.ServiceGroupApp.UserDataUsageModelServiceGroup.UserDataUsageService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

// CreateUserDataUsage 创建userDataUsage表
// @Tags UserDataUsage
// @Summary 创建userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userDataUsageModel.UserDataUsage true "创建userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userDataUsage/createUserDataUsage [post]
func (userDataUsageApi *UserDataUsageApi) CreateUserDataUsage(c *gin.Context) {
	var userDataUsage userDataUsageModel.UserDataUsage
	err := c.ShouldBindJSON(&userDataUsage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 查询uuid是否正确
	_, err = userService.FindUserByUuid(userDataUsage.UserUuid)
	if err != nil {
		response.FailWithMessage("不存在该用户", c)
		return
	}

	if err := userDataUsageService.CreateUserDataUsage(&userDataUsage); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserDataUsage 删除userDataUsage表
// @Tags UserDataUsage
// @Summary 删除userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userDataUsageModel.UserDataUsage true "删除userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userDataUsage/deleteUserDataUsage [delete]
func (userDataUsageApi *UserDataUsageApi) DeleteUserDataUsage(c *gin.Context) {
	id := c.Query("ID")
	if err := userDataUsageService.DeleteUserDataUsage(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserDataUsageByIds 批量删除userDataUsage表
// @Tags UserDataUsage
// @Summary 批量删除userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userDataUsage/deleteUserDataUsageByIds [delete]
func (userDataUsageApi *UserDataUsageApi) DeleteUserDataUsageByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := userDataUsageService.DeleteUserDataUsageByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserDataUsage 更新userDataUsage表
// @Tags UserDataUsage
// @Summary 更新userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userDataUsageModel.UserDataUsage true "更新userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userDataUsage/updateUserDataUsage [put]
func (userDataUsageApi *UserDataUsageApi) UpdateUserDataUsage(c *gin.Context) {
	var userDataUsage userDataUsageModel.UserDataUsage
	err := c.ShouldBindJSON(&userDataUsage)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userDataUsageService.UpdateUserDataUsage(userDataUsage); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserDataUsage 用id查询userDataUsage表
// @Tags UserDataUsage
// @Summary 用id查询userDataUsage表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userDataUsageModel.UserDataUsage true "用id查询userDataUsage表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userDataUsage/findUserDataUsage [get]
func (userDataUsageApi *UserDataUsageApi) FindUserDataUsage(c *gin.Context) {
	id := c.Query("ID")
	if reuserDataUsage, err := userDataUsageService.GetUserDataUsage(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserDataUsage": reuserDataUsage}, c)
	}
}

// GetUserDataUsageList 分页获取userDataUsage表列表
// @Tags UserDataUsage
// @Summary 分页获取userDataUsage表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userDataUsageModelReq.UserDataUsageSearch true "分页获取userDataUsage表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userDataUsage/getUserDataUsageList [get]
func (userDataUsageApi *UserDataUsageApi) GetUserDataUsageList(c *gin.Context) {
	var pageInfo userDataUsageModelReq.UserDataUsageSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userDataUsageService.GetUserDataUsageInfoList(pageInfo); err != nil {
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
