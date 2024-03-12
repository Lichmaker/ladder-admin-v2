package userext

import (
	"encoding/base64"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userext"
	userExtReq "github.com/flipped-aurora/gin-vue-admin/server/model/userext/request"
	userExtRes "github.com/flipped-aurora/gin-vue-admin/server/model/userext/response"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/vmess"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/manager"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-module/carbon/v2"
	qrcode "github.com/skip2/go-qrcode"
	"go.uber.org/zap"
)

type UserExtApi struct {
}

var userExtService = service.ServiceGroupApp.UserextServiceGroup.UserExtService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var userDataUsageService = service.ServiceGroupApp.UserDataUsageModelServiceGroup
var v2rayNodeService = service.ServiceGroupApp.V2raynodeServiceGroup.V2rayNodeService

// CreateUserExt 创建userExt表
// @Tags UserExt
// @Summary 创建userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userext.UserExt true "创建userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userExt/createUserExt [post]
func (userExtApi *UserExtApi) CreateUserExt(c *gin.Context) {
	var userExt userext.UserExt
	err := c.ShouldBindJSON(&userExt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userExtService.CreateUserExt(&userExt); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteUserExt 删除userExt表
// @Tags UserExt
// @Summary 删除userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userext.UserExt true "删除userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /userExt/deleteUserExt [delete]
func (userExtApi *UserExtApi) DeleteUserExt(c *gin.Context) {
	id := c.Query("ID")
	if err := userExtService.DeleteUserExt(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUserExtByIds 批量删除userExt表
// @Tags UserExt
// @Summary 批量删除userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /userExt/deleteUserExtByIds [delete]
func (userExtApi *UserExtApi) DeleteUserExtByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := userExtService.DeleteUserExtByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUserExt 更新userExt表
// @Tags UserExt
// @Summary 更新userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userext.UserExt true "更新userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /userExt/updateUserExt [put]
func (userExtApi *UserExtApi) UpdateUserExt(c *gin.Context) {
	var userExt userext.UserExt
	err := c.ShouldBindJSON(&userExt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := userExtService.UpdateUserExt(userExt); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUserExt 用id查询userExt表
// @Tags UserExt
// @Summary 用id查询userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userext.UserExt true "用id查询userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /userExt/findUserExt [get]
func (userExtApi *UserExtApi) FindUserExt(c *gin.Context) {
	id := c.Query("ID")
	if reuserExt, err := userExtService.GetUserExt(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reuserExt": reuserExt}, c)
	}
}

// GetUserExtList 分页获取userExt表列表
// @Tags UserExt
// @Summary 分页获取userExt表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query userextReq.UserExtSearch true "分页获取userExt表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /userExt/getUserExtList [get]
func (userExtApi *UserExtApi) GetUserExtList(c *gin.Context) {
	var pageInfo userExtReq.UserExtSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := userExtService.GetUserExtInfoList(pageInfo); err != nil {
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

// CreateUserExt 创建userExt表
// @Tags UserExt
// @Summary 创建userExt表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body userext.UserExt true "创建userExt表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /userExt/createUserExt [post]
func (userExtApi *UserExtApi) CreateUserExtV2(c *gin.Context) {
	var userExt userExtReq.UserExtInfoV2
	err := c.ShouldBindJSON(&userExt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(userExt, utils.RegisterVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var authorities []system.SysAuthority
	for _, v := range userExt.AuthorityIds {
		authorities = append(authorities, system.SysAuthority{
			AuthorityId: v,
		})
	}
	// 创建时，默认都不激活
	userExt.Enable = 0

	global.GVA_LOG.Info("新用户注册", zap.String("body", utils.JsonMarshal(userExt, true)))

	validStartTime := carbon.ParseByLayout(userExt.ValidStart, "2006-01-02").ToStdTime()
	validEndTime := carbon.ParseByLayout(userExt.ValidEnd, "2006-01-02").ToStdTime()
	user := &system.SysUser{
		Username:    userExt.Username,
		NickName:    userExt.NickName,
		Password:    userExt.Password,
		HeaderImg:   "",
		AuthorityId: userExt.AuthorityId,
		Authorities: authorities,
		Enable:      1, // 账号默认启用
		Phone:       userExt.Phone,
		Email:       userExt.Email,
		UserExt: userext.UserExt{
			ValidStart:   &validStartTime,
			ValidEnd:     &validEndTime,
			StandardData: userExt.StandardDataMB,
			Enable:       uint(userExt.Enable),
		},
	}
	userReturn, err := userService.Register(*user)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		// response.FailWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册失败", c)
		response.FailWithMessage(err.Error(), c)
		return
	}

	if userExt.Enable == 1 {
		// 把用户加入到所有节点中
		nodeQuery, err := v2rayNodeService.GetAllNode()
		if err != nil {
			global.GVA_LOG.Error("查询节点失败!", zap.Error(err))
		} else {
			for _, v := range nodeQuery {
				if _, err := global.V2RAY_MANAGER_CLIENT.DeleteUser(c.Request.Context(), &manager.DeleteUserRequest{
					NodeID: v.UniqueId,
					Email:  userReturn.Email,
				}); err != nil {
					global.GVA_LOG.Error("从节点删除用户失败!", zap.String("nodeId", v.UniqueId), zap.Error(err))
				}
				if _, err := global.V2RAY_MANAGER_CLIENT.AddUser(c.Request.Context(), &manager.AddUserRequest{
					NodeID: v.UniqueId,
					User:   userReturn.BuildManagerUserInfo(),
				}); err != nil {
					global.GVA_LOG.Error("写入用户到节点失败!", zap.String("nodeId", v.UniqueId), zap.Error(err))
				}
			}
		}
	}

	global.GVA_LOG.Info("注册成功", zap.String("body", utils.JsonMarshal(userReturn, true)))
	response.OkWithMessage("注册成功", c)
	// response.OkWithDetailed(systemRes.SysUserResponse{User: userReturn}, "注册成功", c)
}

func (userExtApi *UserExtApi) GetUserList(c *gin.Context) {
	var pageInfo userExtReq.UserListSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := userService.GetUserInfoListV2(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (userExtApi *UserExtApi) UpdateUserExtV2(c *gin.Context) {
	var user userExtReq.UpdateUserInfo
	err := c.ShouldBindJSON(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(user, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	global.GVA_LOG.Info(utils.JsonMarshal(user, true))

	err = userExtService.UpdateUserInfo(user)
	if err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}

	userInfo, err := userService.FindUserById(user.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
		return
	}

	// 把用户加入到所有节点中
	go func() {
		nodeQuery, err := v2rayNodeService.GetAllNode()
		if err != nil {
			global.GVA_LOG.Error("查询节点失败!", zap.Error(err))
		} else {
			if user.Enable != 1 {
				// 从所有节点中删除用户
				for _, v := range nodeQuery {
					if _, err := global.V2RAY_MANAGER_CLIENT.DeleteUser(c.Request.Context(), &manager.DeleteUserRequest{
						NodeID: v.UniqueId,
						Email:  user.Email,
					}); err != nil {
						global.GVA_LOG.Error("从节点删除用户失败!", zap.String("nodeId", v.UniqueId), zap.Error(err))
					}
				}
			} else {
				// 从所有节点中添加用户
				for _, v := range nodeQuery {
					if _, err := global.V2RAY_MANAGER_CLIENT.AddUser(c.Request.Context(), &manager.AddUserRequest{
						NodeID: v.UniqueId,
						User:   userInfo.BuildManagerUserInfo(),
					}); err != nil {
						global.GVA_LOG.Error("节点新增用户失败!", zap.String("nodeId", v.UniqueId), zap.Error(err))
					}
				}
			}
		}
	}()

	response.OkWithMessage("设置成功", c)
}

func (userExtApi *UserExtApi) Dashboard(c *gin.Context) {
	var resp userExtRes.UserExtDashboardResponse

	// 通过token查询用户信息
	userInfoModel, err := userService.GetUserInfoViaToken(c)
	if err != nil {
		global.GVA_LOG.Error("获取用户信息!", zap.Error(err))
		response.FailWithDetailed(resp, "读取用户失败", c)
		return
	}

	// 查询当前用量信息
	currentUserUsage, err := userDataUsageService.GetUserUsingData(userInfoModel.UUID.String(), time.Now())
	if err != nil {
		global.GVA_LOG.Error("查询用量信息失败", zap.Error(err))
	}

	vmessResult := make([]string, 0)
	nodeQuery, _ := v2rayNodeService.GetAllNode()
	for _, v := range nodeQuery {
		vmessTmp := vmess.VMESS{}
		vmessTmp.Default()
		tmp := vmessTmp.Build(userInfoModel.UUID.String(), global.GVA_CONFIG.V2RayManager.DefaultAlterId, v)
		vmessResult = append(vmessResult, tmp)
	}
	resultStr := strings.Join(vmessResult, "\n")
	genQRCodeBytes, _ := qrcode.Encode(resultStr, qrcode.Medium, 256)
	genQRCodeString := base64.StdEncoding.EncodeToString(genQRCodeBytes)

	resp.Base.QRCode = "data:image/png;base64," + genQRCodeString
	resp.Base.SubUrl = fmt.Sprintf("https://%s/subscribe?UUID=%s", c.Request.Host, userInfoModel.UUID.String())
	resp.Base.ValidStart = userInfoModel.UserExt.ValidStart.Format("2006-01-02")
	resp.Base.ValidEnd = userInfoModel.UserExt.ValidEnd.Format("2006-01-02")
	resp.Base.StandardData = utils.ParseByteSize(userInfoModel.UserExt.StandardData * (1024 * 1024))
	resp.Base.Enable = userInfoModel.UserExt.Enable == 1

	if currentUserUsage.ID > 0 {
		totalByte := currentUserUsage.StandardData * (1024 * 1024)
		remainByte := totalByte - currentUserUsage.Usage

		resp.Using.Start = currentUserUsage.StartDate.Format("2006-01-02")
		resp.Using.End = currentUserUsage.EndDate.Format("2006-01-02")
		resp.Using.StandardData = utils.ParseByteSize(totalByte)
		resp.Using.Usage = utils.ParseByteSize(currentUserUsage.Usage)
		resp.Using.Remain = utils.ParseByteSize(remainByte)
		resp.Using.Percentage = int(math.Round(float64(currentUserUsage.Usage) / float64(totalByte) * 100))
		if resp.Using.Percentage >= 100 {
			resp.Using.Percentage = 100
		}
	} else {
		resp.Using.Start = "未启用"
		resp.Using.End = "未启用"
		resp.Using.StandardData = "未启用"
		resp.Using.Usage = "未启用"
		resp.Using.Remain = "未启用"
		resp.Using.Percentage = 0
	}

	response.OkWithDetailed(resp, "获取成功", c)

}
