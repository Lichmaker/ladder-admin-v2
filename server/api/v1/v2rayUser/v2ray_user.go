package v2rayUser

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/v2rayUser"
	v2rayUserReq "github.com/flipped-aurora/gin-vue-admin/server/model/v2rayUser/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type V2rayUserApi struct {
}

var v2rayUserService = service.ServiceGroupApp.V2rayuserServiceGroup.V2rayUserService

// CreateV2rayUser 创建V2rayUser
// @Tags V2rayUser
// @Summary 创建V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body v2rayUser.V2rayUser true "创建V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v2rayUser/createV2rayUser [post]
func (v2rayUserApi *V2rayUserApi) CreateV2rayUser(c *gin.Context) {
	var v2rayUser v2rayUser.V2rayUser
	_ = c.ShouldBindJSON(&v2rayUser)
	verify := utils.Rules{
		"AdminUserId": {utils.NotEmpty()},
		"Email":       {utils.NotEmpty()},
		"ExpireAt":    {utils.NotEmpty()},
		"ActiveState": {utils.NotEmpty()},
	}
	if err := utils.Verify(v2rayUser, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := v2rayUserService.CreateV2rayUser(v2rayUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteV2rayUser 删除V2rayUser
// @Tags V2rayUser
// @Summary 删除V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body v2rayUser.V2rayUser true "删除V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v2rayUser/deleteV2rayUser [delete]
func (v2rayUserApi *V2rayUserApi) DeleteV2rayUser(c *gin.Context) {
	var v2rayUser v2rayUser.V2rayUser
	_ = c.ShouldBindJSON(&v2rayUser)
	if err := v2rayUserService.DeleteV2rayUser(v2rayUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteV2rayUserByIds 批量删除V2rayUser
// @Tags V2rayUser
// @Summary 批量删除V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /v2rayUser/deleteV2rayUserByIds [delete]
func (v2rayUserApi *V2rayUserApi) DeleteV2rayUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := v2rayUserService.DeleteV2rayUserByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateV2rayUser 更新V2rayUser
// @Tags V2rayUser
// @Summary 更新V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body v2rayUser.V2rayUser true "更新V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v2rayUser/updateV2rayUser [put]
func (v2rayUserApi *V2rayUserApi) UpdateV2rayUser(c *gin.Context) {
	var v2rayUser v2rayUser.V2rayUser
	_ = c.ShouldBindJSON(&v2rayUser)
	verify := utils.Rules{
		"AdminUserId": {utils.NotEmpty()},
		"Email":       {utils.NotEmpty()},
		"ExpireAt":    {utils.NotEmpty()},
		"ActiveState": {utils.NotEmpty()},
	}
	if err := utils.Verify(v2rayUser, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := v2rayUserService.UpdateV2rayUser(v2rayUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindV2rayUser 用id查询V2rayUser
// @Tags V2rayUser
// @Summary 用id查询V2rayUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query v2rayUser.V2rayUser true "用id查询V2rayUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /v2rayUser/findV2rayUser [get]
func (v2rayUserApi *V2rayUserApi) FindV2rayUser(c *gin.Context) {
	var v2rayUser v2rayUser.V2rayUser
	_ = c.ShouldBindQuery(&v2rayUser)
	if rev2rayUser, err := v2rayUserService.GetV2rayUser(uint(*v2rayUser.AdminUserId)); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rev2rayUser": rev2rayUser}, c)
	}
}

// GetV2rayUserList 分页获取V2rayUser列表
// @Tags V2rayUser
// @Summary 分页获取V2rayUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query v2rayUserReq.V2rayUserSearch true "分页获取V2rayUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v2rayUser/getV2rayUserList [get]
func (v2rayUserApi *V2rayUserApi) GetV2rayUserList(c *gin.Context) {
	var pageInfo v2rayUserReq.V2rayUserSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if list, total, err := v2rayUserService.GetV2rayUserInfoList(pageInfo); err != nil {
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
