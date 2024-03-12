package initialize

import (
	"context"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/rpcconnect"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/manager"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

var v2rayNodeService = service.ServiceGroupApp.V2raynodeServiceGroup.V2rayNodeService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService

func V2rayManager() {
	managerConfig := global.GVA_CONFIG.V2RayManager
	conn, err := rpcconnect.Dial(fmt.Sprintf("%s:%d", managerConfig.Host, managerConfig.Port), nil)
	if err != nil {
		global.GVA_LOG.Panic("v2ray-manager连接失败，请检查是否已启动", zap.Error(err))
	}
	global.V2RAY_MANAGER_CONN = conn
	global.V2RAY_MANAGER_CLIENT = manager.NewServeServiceClient(conn)

	LoadAllV2rayNode()
}

func LoadAllV2rayNode() {
	ctx := context.Background()

	// 从db中读取所有节点信息，然后下发到manager
	nodeQuery, err := v2rayNodeService.GetAllNode()
	if err != nil {
		global.GVA_LOG.Panic("读取所有节点信息失败", zap.Error(err))
	}

	setNodeReq := manager.SetNodeRequest{}
	for _, v := range nodeQuery {
		secretStr, err := v.GetSecret()
		if err != nil {
			global.GVA_LOG.Error("读取节点secret失败", zap.Error(err))
		}

		tmp := manager.SingleNodeInfo{}
		tmp.ID = v.UniqueId
		tmp.Name = v.Name
		tmp.Host = v.Host
		tmp.NodeConfigV4 = v.ConfigRaw
		tmp.ReverseConfig = &manager.NodeReverseConfig{
			VMessPort: int64(v.VmessPort),
			RPCPort:   int64(v.RpcPort),
			RPCSecret: secretStr,
		}
		setNodeReq.NodeInfoArr = append(setNodeReq.NodeInfoArr, &tmp)
	}

	if _, err := global.V2RAY_MANAGER_CLIENT.SetNode(ctx, &setNodeReq); err != nil {
		global.GVA_LOG.Error("往manager下发node失败", zap.Error(err))
	}

	// 查询所有用户，写入每个节点中
	userInfoArr := userService.GetAllManagerUserInfo()
	for _, v := range nodeQuery {
		setUserReq := manager.SetUserRequest{}
		setUserReq.NodeID = v.UniqueId
		setUserReq.Users = userInfoArr
		if _, err := global.V2RAY_MANAGER_CLIENT.SetUser(ctx, &setUserReq); err != nil {
			global.GVA_LOG.Error("往manager下发用户失败", zap.Error(err), zap.String("nodeId", v.UniqueId))
		}
	}

}
