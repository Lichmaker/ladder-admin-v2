package usercheck

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/userext/request"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/manager"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

var userDataUsageService = service.ServiceGroupApp.UserDataUsageModelServiceGroup.UserDataUsageService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var userExtService = service.ServiceGroupApp.UserextServiceGroup.UserExtService
var v2rayNodeService = service.ServiceGroupApp.V2raynodeServiceGroup.V2rayNodeService

func Check() {
	ctx := context.Background()
	currentTime := time.Now()

	userReq := request.UserListSearch{}
	userReq.Page = 1
	userReq.PageSize = -1
	userQuery, _, err := userService.GetUserInfoListV2(userReq)
	if err != nil {
		global.GVA_LOG.Error("查询用户失败", zap.Error(err))
		return
	}
	userMap := make(map[string]system.SysUser)
	for _, v := range userQuery {
		userMap[v.UUID.String()] = v
	}

	// 两个map的key是uuid
	var needEnableUser map[string]system.SysUser = make(map[string]system.SysUser)
	var needDisableUser map[string]system.SysUser = make(map[string]system.SysUser)

	// 查用户看是否在有效期内
	for k, v := range userMap {

		if v.UserExt.ValidStart.After(currentTime) || v.UserExt.ValidEnd.Before(currentTime) {
			// 不可用
			if v.UserExt.Enable == 1 {
				needDisableUser[v.UUID.String()] = v
				delete(userMap, k)
				global.GVA_LOG.Debug("账号过期，禁用", zap.String("uuid", v.UUID.String()))
			}
		}
	}

	// 查用户是否有可用的流量包
	for k, v := range userMap {
		queryData, err := userDataUsageService.GetUserUsingData(v.UUID.String(), currentTime)
		if err != nil {
			global.GVA_LOG.Error("查询失败", zap.Error(err))
			continue
		}
		if queryData.ID > 0 && v.UserExt.Enable != 1 {
			needEnableUser[v.UUID.String()] = v
			delete(userMap, k)
			global.GVA_LOG.Debug("账号由流量包，启用", zap.String("uuid", v.UUID.String()))

		} else if queryData.ID == 0 && v.UserExt.Enable == 1 {
			needDisableUser[v.UUID.String()] = v
			delete(userMap, k)
			global.GVA_LOG.Debug("账号无流量包，禁用", zap.String("uuid", v.UUID.String()))
		}
	}

	// 发生变更的用户，马上下发到所有节点
	allNode, err := v2rayNodeService.GetAllNode()
	if err != nil {
		global.GVA_LOG.Error("读取节点失败", zap.Error(err))
		return
	}

	for _, v := range needEnableUser {
		if err := userExtService.UpdateEnable(v.UUID.String(), 1); err != nil {
			global.GVA_LOG.Error("账户激活失败", zap.Error(err))
			continue
		}
		global.GVA_LOG.Debug("账户执行激活", zap.String("uuid", v.UUID.String()))
		for _, nodeInfo := range allNode {
			if !nodeInfo.CheckOnline() {
				continue
			}
			if _, err := global.V2RAY_MANAGER_CLIENT.AddUser(ctx, &manager.AddUserRequest{
				NodeID: nodeInfo.UniqueId,
				User:   v.BuildManagerUserInfo(),
			}); err != nil {
				global.GVA_LOG.Warn("写入用户到节点失败", zap.Error(err))
			}
		}
	}
	for _, v := range needDisableUser {
		if err := userExtService.UpdateEnable(v.UUID.String(), 0); err != nil {
			global.GVA_LOG.Error("账户冻结失败", zap.Error(err))
			continue
		}
		global.GVA_LOG.Debug("账户执行冻结", zap.String("uuid", v.UUID.String()))
		for _, nodeInfo := range allNode {
			if !nodeInfo.CheckOnline() {
				continue
			}
			if _, err := global.V2RAY_MANAGER_CLIENT.DeleteUser(ctx, &manager.DeleteUserRequest{
				NodeID: nodeInfo.UniqueId,
				Email:  v.Email,
			}); err != nil {
				global.GVA_LOG.Warn("从节点删除用户失败", zap.Error(err))
			}
		}
	}
}

func ReloadUser() {
	ctx := context.Background()
	allNode, err := v2rayNodeService.GetAllNode()
	if err != nil {
		global.GVA_LOG.Error("读取节点失败", zap.Error(err))
		return
	}

	// 查询所有用户，写入每个节点中
	userInfoArr := userService.GetAllManagerUserInfo()
	for _, v := range allNode {
		if !v.CheckOnline() {
			continue
		}
		setUserReq := manager.SetUserRequest{}
		setUserReq.NodeID = v.UniqueId
		setUserReq.Users = userInfoArr
		if _, err := global.V2RAY_MANAGER_CLIENT.SetUser(ctx, &setUserReq); err != nil {
			global.GVA_LOG.Error("往manager下发用户失败", zap.Error(err), zap.String("nodeId", v.UniqueId))
		}
	}
}
