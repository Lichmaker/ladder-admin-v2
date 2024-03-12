package v2raynode

import (
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/v2raynode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/v2raynode/request"
	v2raynodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/v2raynode/request"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/vmess"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/manager"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type V2rayNodeApi struct {
}

var v2rayNodeService = service.ServiceGroupApp.V2raynodeServiceGroup.V2rayNodeService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

const ALL_NODE_LIMIT_KEY = "la_all_node_limit"
const ALL_NODE_LIMIT_KEY_TTL = time.Second * 5

// CreateV2rayNode 创建v2rayNode表
// @Tags V2rayNode
// @Summary 创建v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body v2raynode.V2rayNode true "创建v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v2rayNode/createV2rayNode [post]
func (v2rayNodeApi *V2rayNodeApi) CreateV2rayNode(c *gin.Context) {
	var v2rayNode v2raynode.V2rayNode
	err := c.ShouldBindJSON(&v2rayNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := v2rayNodeService.CreateV2rayNode(&v2rayNode); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	}

	// 写入到manager
	if _, err := global.V2RAY_MANAGER_CLIENT.AddNode(c.Request.Context(), v2rayNode.ToManagerNode()); err != nil {
		global.GVA_LOG.Error("写入新节点到manager失败", zap.Error(err))
	}
	// 把当前所有用户查询出来然后写入节点
	allUser := userService.GetAllManagerUserInfo()
	if _, err := global.V2RAY_MANAGER_CLIENT.SetUser(c.Request.Context(), &manager.SetUserRequest{NodeID: v2rayNode.UniqueId, Users: allUser}); err != nil {
		global.GVA_LOG.Error("写入用户到新节点失败", zap.Error(err))
	}

	response.OkWithMessage("创建成功", c)
}

// DeleteV2rayNode 删除v2rayNode表
// @Tags V2rayNode
// @Summary 删除v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body v2raynode.V2rayNode true "删除v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /v2rayNode/deleteV2rayNode [delete]
func (v2rayNodeApi *V2rayNodeApi) DeleteV2rayNode(c *gin.Context) {
	id := c.Query("ID")

	query, err := v2rayNodeService.GetV2rayNode(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.FailWithMessage("不存在的数据", c)
		} else {
			global.GVA_LOG.Error("查询失败!", zap.Error(err))
			response.FailWithMessage("删除失败", c)
		}
	}

	if err := v2rayNodeService.DeleteV2rayNode(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	}

	// 从 manager 删除
	if _, err := global.V2RAY_MANAGER_CLIENT.DeleteNode(c.Request.Context(), &manager.DeleteNodeRequest{ID: query.UniqueId}); err != nil {
		global.GVA_LOG.Error("写入新节点到manager失败", zap.Error(err))
	}

	response.OkWithMessage("删除成功", c)
}

// DeleteV2rayNodeByIds 批量删除v2rayNode表
// @Tags V2rayNode
// @Summary 批量删除v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /v2rayNode/deleteV2rayNodeByIds [delete]
func (v2rayNodeApi *V2rayNodeApi) DeleteV2rayNodeByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := v2rayNodeService.DeleteV2rayNodeByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateV2rayNode 更新v2rayNode表
// @Tags V2rayNode
// @Summary 更新v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body v2raynode.V2rayNode true "更新v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /v2rayNode/updateV2rayNode [put]
func (v2rayNodeApi *V2rayNodeApi) UpdateV2rayNode(c *gin.Context) {
	var v2rayNode v2raynode.V2rayNode
	err := c.ShouldBindJSON(&v2rayNode)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := v2rayNodeService.UpdateV2rayNode(&v2rayNode); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	// 写入manager
	nodeInfo, _ := v2rayNodeService.GetByUniqueID(v2rayNode.UniqueId)
	if _, err := global.V2RAY_MANAGER_CLIENT.AddNode(c.Request.Context(), nodeInfo.ToManagerNode()); err != nil {
		global.GVA_LOG.Error("更新节点到manager失败", zap.Error(err))
	}

	response.OkWithMessage("更新成功", c)
}

// FindV2rayNode 用id查询v2rayNode表
// @Tags V2rayNode
// @Summary 用id查询v2rayNode表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query v2raynode.V2rayNode true "用id查询v2rayNode表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /v2rayNode/findV2rayNode [get]
func (v2rayNodeApi *V2rayNodeApi) FindV2rayNode(c *gin.Context) {
	id := c.Query("ID")
	if v2rayNode, err := v2rayNodeService.GetV2rayNode(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"v2rayNode": v2rayNode}, c)
	}
}

// GetV2rayNodeList 分页获取v2rayNode表列表
// @Tags V2rayNode
// @Summary 分页获取v2rayNode表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query v2raynodeReq.V2rayNodeSearch true "分页获取v2rayNode表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v2rayNode/getV2rayNodeList [get]
func (v2rayNodeApi *V2rayNodeApi) GetV2rayNodeList(c *gin.Context) {
	var pageInfo v2raynodeReq.V2rayNodeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := v2rayNodeService.GetV2rayNodeInfoList(pageInfo); err != nil {
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

func (v2rayNodeApi *V2rayNodeApi) GenSubscribeUrl(c *gin.Context) {
	var req v2raynodeReq.GenVemssUrlRequest
	err := c.ShouldBindQuery(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 通过uuid查询用户是否可用
	userInfo, err := userService.GetUserInfo(uuid.FromStringOrNil(req.UUID))
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if userInfo.UserExt.Enable != 1 {
		response.FailWithMessage("账号已过期或未激活", c)
		return
	}

	result := make([]string, 0)
	nodeQuery, _ := v2rayNodeService.GetAllNode()
	for _, v := range nodeQuery {
		vmessTmp := vmess.VMESS{}
		vmessTmp.Default()
		tmp := vmessTmp.Build(userInfo.UUID.String(), global.GVA_CONFIG.V2RayManager.DefaultAlterId, v)
		result = append(result, tmp)
	}
	resultStr := strings.Join(result, "\n")
	base64Str := base64.StdEncoding.EncodeToString([]byte(resultStr))
	c.String(http.StatusOK, base64Str)
}

func (v2rayNodeApi *V2rayNodeApi) PushData(c *gin.Context) {
	var req request.V2rayNodePushDataReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var nodeInfoArr []v2raynode.V2rayNode
	if req.All {
		val, err := global.GVA_REDIS.SetNX(c.Request.Context(), ALL_NODE_LIMIT_KEY, 1, ALL_NODE_LIMIT_KEY_TTL).Result()
		if err != nil {
			global.GVA_LOG.Error("redis服务异常", zap.Error(err))
			response.FailWithMessage("REDIS服务异常", c)
			return
		}
		if !val {
			response.FailWithMessage("五秒内请勿重复操作", c)
			return
		}
		nodeInfoArr, err = v2rayNodeService.GetAllNode()
		if err != nil {
			global.GVA_LOG.Error("查询节点失败", zap.Error(err))
			response.FailWithMessage("数据库异常", c)
			return
		}
	} else {
		query, err := v2rayNodeService.GetByUniqueID(req.NodeID)
		if err != nil {
			global.GVA_LOG.Error("查询节点失败", zap.Error(err))
			response.FailWithMessage("数据库异常", c)
			return
		}
		nodeInfoArr = append(nodeInfoArr, query)
	}

	for _, nodeItem := range nodeInfoArr {
		// 写入到manager
		if _, err := global.V2RAY_MANAGER_CLIENT.AddNode(c.Request.Context(), nodeItem.ToManagerNode()); err != nil {
			global.GVA_LOG.Warn("主动push数据到manager失败", zap.Error(err))
		}
		// 把当前所有用户查询出来然后写入节点
		allUser := userService.GetAllManagerUserInfo()
		if _, err := global.V2RAY_MANAGER_CLIENT.SetUser(c.Request.Context(), &manager.SetUserRequest{NodeID: nodeItem.UniqueId, Users: allUser}); err != nil {
			global.GVA_LOG.Warn("主动push用户到manager", zap.Error(err))
		}
	}

	response.OkWithMessage("创建成功", c)
}
