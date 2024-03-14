package invitecode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/invitecode"
	invitecodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/invitecode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type InviteCodeApi struct {
}

var inviteCodeService = service.ServiceGroupApp.InvitecodeServiceGroup.InviteCodeService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

// CreateInviteCode 创建inviteCode表
// @Tags InviteCode
// @Summary 创建inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body invitecode.InviteCode true "创建inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /inviteCode/createInviteCode [post]
func (inviteCodeApi *InviteCodeApi) CreateInviteCode(c *gin.Context) {
	var inviteCode invitecode.InviteCode
	err := c.ShouldBindJSON(&inviteCode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	inviteCode.CreatedBy = utils.GetUserID(c)
	if len(inviteCode.Code) < 8 {
		inviteCode.Code = utils.RandomString(8)
	}
	if len(inviteCode.CreatorUuid) == 0 {
		inviteCode.CreatorUuid = utils.GetUserUuid(c).String()
	} else {
		user, err := userService.GetUserInfoViaToken(c)
		if err != nil {
			response.FailWithMessage(err.Error(), c)
			return
		}
		inviteCode.CreatorUuid = user.UUID.String()
	}

	if err := inviteCodeService.CreateInviteCode(&inviteCode); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteInviteCode 删除inviteCode表
// @Tags InviteCode
// @Summary 删除inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body invitecode.InviteCode true "删除inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /inviteCode/deleteInviteCode [delete]
func (inviteCodeApi *InviteCodeApi) DeleteInviteCode(c *gin.Context) {
	id := c.Query("ID")
	userID := utils.GetUserID(c)
	if err := inviteCodeService.DeleteInviteCode(id, userID); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteInviteCodeByIds 批量删除inviteCode表
// @Tags InviteCode
// @Summary 批量删除inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /inviteCode/deleteInviteCodeByIds [delete]
func (inviteCodeApi *InviteCodeApi) DeleteInviteCodeByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	userID := utils.GetUserID(c)
	if err := inviteCodeService.DeleteInviteCodeByIds(ids, userID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateInviteCode 更新inviteCode表
// @Tags InviteCode
// @Summary 更新inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body invitecode.InviteCode true "更新inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /inviteCode/updateInviteCode [put]
func (inviteCodeApi *InviteCodeApi) UpdateInviteCode(c *gin.Context) {
	var inviteCode invitecode.InviteCode
	err := c.ShouldBindJSON(&inviteCode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	inviteCode.UpdatedBy = utils.GetUserID(c)

	if err := inviteCodeService.UpdateInviteCode(inviteCode); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindInviteCode 用id查询inviteCode表
// @Tags InviteCode
// @Summary 用id查询inviteCode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query invitecode.InviteCode true "用id查询inviteCode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /inviteCode/findInviteCode [get]
func (inviteCodeApi *InviteCodeApi) FindInviteCode(c *gin.Context) {
	id := c.Query("ID")
	if reinviteCode, err := inviteCodeService.GetInviteCode(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reinviteCode": reinviteCode}, c)
	}
}

// GetInviteCodeList 分页获取inviteCode表列表
// @Tags InviteCode
// @Summary 分页获取inviteCode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query invitecodeReq.InviteCodeSearch true "分页获取inviteCode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /inviteCode/getInviteCodeList [get]
func (inviteCodeApi *InviteCodeApi) GetInviteCodeList(c *gin.Context) {
	var pageInfo invitecodeReq.InviteCodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := inviteCodeService.GetInviteCodeInfoList(pageInfo); err != nil {
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

func (inviteCodeApi *InviteCodeApi) GetMyInviteCodeList(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	reqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var pageInfo invitecodeReq.InviteCodeSearch
	err = c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	pageInfo.CreatorUUID = reqUser.UUID.String()
	if list, total, err := inviteCodeService.GetInviteCodeInfoList(pageInfo); err != nil {
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

func (inviteCodeApi *InviteCodeApi) DeleteMyInviteCodeByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	uuid := utils.GetUserUuid(c)
	reqUser, err := userService.GetUserInfo(uuid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if pass, err := inviteCodeService.VerifyMyInviteCodeIds(reqUser.UUID.String(), ids); err != nil || !pass {
		response.FailWithMessage("数据异常，请刷新页面后重试 "+err.Error(), c)
		return
	}

	if err := inviteCodeService.DeleteInviteCodeByIds(ids, reqUser.ID); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}
