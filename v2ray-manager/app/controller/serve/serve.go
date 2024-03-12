package serve

import (
	"context"
	"encoding/json"
	"v2ray-manager/app/connections"
	"v2ray-manager/nodecfg"
	"v2ray-manager/pkg/logger"
	"v2ray-manager/pkg/v4helper"
	"v2ray-manager/protobuf/manager"

	core "github.com/v2fly/v2ray-core/v5"
	proxymanCommand "github.com/v2fly/v2ray-core/v5/app/proxyman/command"
	statsCommand "github.com/v2fly/v2ray-core/v5/app/stats/command"
	"github.com/v2fly/v2ray-core/v5/common/protocol"
	"github.com/v2fly/v2ray-core/v5/common/serial"
	v4 "github.com/v2fly/v2ray-core/v5/infra/conf/v4"
	"github.com/v2fly/v2ray-core/v5/proxy/vmess"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ServeServer struct {
	manager.UnimplementedServeServiceServer
}

var _ manager.ServeServiceServer = (*ServeServer)(nil)

func (s *ServeServer) SetNode(ctx context.Context, req *manager.SetNodeRequest) (*emptypb.Empty, error) {
	connections.CloseAllConnect()

	nodeInfoArr := make([]*nodecfg.NodeInfo, 0)
	for _, v := range req.NodeInfoArr {
		var v4Config = v4.Config{}
		json.Unmarshal([]byte(v.NodeConfigV4), &v4Config)

		nodeInfoArr = append(nodeInfoArr, &nodecfg.NodeInfo{
			ID:            v.ID,
			Name:          v.Name,
			Host:          v.Host,
			Config:        v.NodeConfigV4,
			ReverseConfig: v.ReverseConfig,
		})
	}
	if err := nodecfg.SetNode(nodeInfoArr); err != nil {
		logger.Logger.Error("SetNode写入配置失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "写入配置失败:%s", err)
	}

	// 建立连接
	go func() {
		for _, v := range req.NodeInfoArr {
			if err := connections.BuildConnect(v.ID); err != nil {
				logger.Logger.Error("节点建立连接失败", zap.String("nodeID", v.ID), zap.Error(err))
			}
		}
	}()

	return &emptypb.Empty{}, nil
}

func (s *ServeServer) AddNode(ctx context.Context, req *manager.AddNodeRequest) (*emptypb.Empty, error) {
	lg, _ := logger.StartTrace(ctx, "AddNode")
	defer lg.TraceSpanEnd()

	nodeInfo := &nodecfg.NodeInfo{}
	nodeInfo.ID = req.NodeInfo.ID
	nodeInfo.Name = req.NodeInfo.Name
	nodeInfo.Config = req.NodeInfo.NodeConfigV4
	nodeInfo.Host = req.NodeInfo.Host
	nodeInfo.ReverseConfig = req.NodeInfo.ReverseConfig
	if err := nodecfg.AddNode(nodeInfo); err != nil {
		lg.Error("添加节点写入失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "写入失败:%s", err)
	}
	connections.CloseConnect(req.NodeInfo.ID)
	if err := connections.BuildConnect(req.NodeInfo.ID); err != nil {
		logger.Logger.Error("节点建立连接失败", zap.String("nodeID", req.NodeInfo.ID), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "连接节点失败:%s", err)
	}
	return &emptypb.Empty{}, nil
}

func (s *ServeServer) DeleteNode(ctx context.Context, req *manager.DeleteNodeRequest) (*emptypb.Empty, error) {
	if err := nodecfg.DeleteNode(req.ID); err != nil {
		logger.Logger.Error("节点删除失败", zap.String("nodeID", req.ID), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "节点删除失败:%s", err)
	}
	connections.CloseConnect(req.ID)
	return &emptypb.Empty{}, nil
}

func (s *ServeServer) SetUser(ctx context.Context, req *manager.SetUserRequest) (*emptypb.Empty, error) {
	lg, _ := logger.StartTrace(ctx, "SetUser")
	defer lg.TraceSpanEnd()

	nodeInfo := nodecfg.GetByID(req.NodeID)
	if nodeInfo == nil {
		return nil, status.Errorf(codes.NotFound, "不存在的节点")
	}

	if err := nodecfg.SetUsers(req.NodeID, req.Users); err != nil {
		rpcError := status.Errorf(codes.Internal, "写入失败:%s", err)
		lg.Error("写入用户失败", zap.Error(err))
		return nil, rpcError
	}

	vmessTag := v4helper.FindVmessTag(nodeInfo.Config)
	if len(vmessTag) == 0 {
		return nil, status.Errorf(codes.NotFound, "找不到vmess配置tag")
	}

	// 写入用户数据到v2ray
	users := nodecfg.GetUsers(req.NodeID)
	if len(users) == 0 {
		rpcError := status.Errorf(codes.Internal, "数据异常")
		return nil, rpcError
	}

	rpcConn, err := connections.GetConnect(req.NodeID)
	if err != nil {
		lg.Error("节点连接获取失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "连接异常")
	}
	client := proxymanCommand.NewHandlerServiceClient(rpcConn)

	// 删除inbound
	if _, err := client.RemoveInbound(ctx, &proxymanCommand.RemoveInboundRequest{Tag: vmessTag}); err != nil {
		lg.Warn("执行inbound删除失败", zap.Error(err))
		// return nil, status.Errorf(codes.Internal, "inbound重置异常")
	}

	// 重新写入inbound
	var v5InboundConfig *core.InboundHandlerConfig
	inboundArr := v4helper.ParseToV4(nodeInfo.Config).InboundConfigs
	for _, v := range inboundArr {
		if v.Tag == vmessTag {
			v5InboundConfig, err = v.Build()
			if err != nil {
				lg.Error("inbound转换为v5结构体失败", zap.Error(err))
				return nil, status.Errorf(codes.Internal, "inbound解析异常")
			}
		}
	}
	if v5InboundConfig == nil {
		lg.Error("inbound配置获取失败", zap.String("tag", vmessTag))
		return nil, status.Errorf(codes.Internal, "inbound解析异常")
	}
	if _, err := client.AddInbound(ctx, &proxymanCommand.AddInboundRequest{Inbound: v5InboundConfig}); err != nil {
		lg.Error("inbound配置设置失败", zap.Error(err))
		// return nil, status.Errorf(codes.Internal, "inbound写入异常")
	}

	// 写入用户
	for _, v := range users {
		if _, err := client.AlterInbound(context.Background(), &proxymanCommand.AlterInboundRequest{
			Tag:       vmessTag,
			Operation: serial.ToTypedMessage(&proxymanCommand.RemoveUserOperation{Email: v.Email}),
		}); err != nil {
			lg.Warn("inbound删除用户失败", zap.Error(err))
		}

		_, err := client.AlterInbound(ctx, &proxymanCommand.AlterInboundRequest{
			Tag: vmessTag,
			Operation: serial.ToTypedMessage(
				&proxymanCommand.AddUserOperation{
					User: &protocol.User{
						Email: v.Email,
						Account: serial.ToTypedMessage(&vmess.Account{
							Id:      v.ID,
							AlterId: v.AlterID,
						}),
						Level: v.Level,
					},
				}),
		})
		if err != nil {
			lg.Warn("往v2ray写入用户失败", zap.Error(err))
		}
	}

	return &emptypb.Empty{}, nil
}

func (s *ServeServer) AddUser(ctx context.Context, req *manager.AddUserRequest) (*emptypb.Empty, error) {
	lg, _ := logger.StartTrace(ctx, "AddUser")
	defer lg.TraceSpanEnd()
	if err := nodecfg.AddUser(req.NodeID, req.User); err != nil {
		rpcError := status.Errorf(codes.Internal, "写入失败:%s", err)
		lg.Error("写入用户失败", zap.Error(err))
		return nil, rpcError
	}

	nodeInfo := nodecfg.GetByID(req.NodeID)
	if nodeInfo == nil {
		return nil, status.Errorf(codes.NotFound, "不存在的节点")
	}

	vmessTag := v4helper.FindVmessTag(nodeInfo.Config)
	if len(vmessTag) == 0 {
		return nil, status.Errorf(codes.NotFound, "找不到vmess配置tag")
	}

	// 写入数据到v2ray
	rpcConn, err := connections.GetConnect(req.NodeID)
	if err != nil {
		lg.Error("节点连接获取失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "连接异常")
	}
	client := proxymanCommand.NewHandlerServiceClient(rpcConn)
	if _, err := client.AlterInbound(context.Background(), &proxymanCommand.AlterInboundRequest{
		Tag:       vmessTag,
		Operation: serial.ToTypedMessage(&proxymanCommand.RemoveUserOperation{Email: req.User.Email}),
	}); err != nil {
		lg.Warn("inbound删除用户失败", zap.Error(err))
	}
	_, err = client.AlterInbound(ctx, &proxymanCommand.AlterInboundRequest{
		Tag: vmessTag,
		Operation: serial.ToTypedMessage(
			&proxymanCommand.AddUserOperation{
				User: &protocol.User{
					Email: req.User.Email,
					Account: serial.ToTypedMessage(&vmess.Account{
						Id:      req.User.ID,
						AlterId: req.User.AlterID,
					}),
					Level: req.User.Level,
				},
			}),
	})
	if err != nil {
		lg.Warn("inbound写入用户失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "inbound写入异常")
	}

	return &emptypb.Empty{}, nil
}

func (s *ServeServer) DeleteUser(ctx context.Context, req *manager.DeleteUserRequest) (*emptypb.Empty, error) {
	lg, _ := logger.StartTrace(ctx, "DeleteUser")
	defer lg.TraceSpanEnd()

	if err := nodecfg.DeleteUser(req.NodeID, req.Email); err != nil {
		logger.Logger.Error("用户删除失败", zap.String("nodeID", req.NodeID), zap.String("email", req.Email), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "用户删除失败:%s", err)
	}

	nodeInfo := nodecfg.GetByID(req.NodeID)
	if nodeInfo == nil {
		return nil, status.Errorf(codes.NotFound, "不存在的节点")
	}

	vmessTag := v4helper.FindVmessTag(nodeInfo.Config)
	if len(vmessTag) == 0 {
		return nil, status.Errorf(codes.NotFound, "找不到vmess配置tag")
	}

	// 写入数据到v2ray
	rpcConn, err := connections.GetConnect(req.NodeID)
	if err != nil {
		lg.Error("节点连接获取失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "连接异常")
	}
	client := proxymanCommand.NewHandlerServiceClient(rpcConn)

	if _, err := client.AlterInbound(context.Background(), &proxymanCommand.AlterInboundRequest{
		Tag:       vmessTag,
		Operation: serial.ToTypedMessage(&proxymanCommand.RemoveUserOperation{Email: req.Email}),
	}); err != nil {
		lg.Error("inbound删除用户失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "inbound删除用户异常")
	}

	return &emptypb.Empty{}, nil
}

func (s *ServeServer) SysStat(ctx context.Context, req *manager.SysStatRequest) (*statsCommand.SysStatsResponse, error) {
	lg, _ := logger.StartTrace(ctx, "SysStat")
	defer lg.TraceSpanEnd()

	conn, err := connections.GetConnect(req.NodeID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	client := statsCommand.NewStatsServiceClient(conn)

	resp, err := client.GetSysStats(ctx, &statsCommand.SysStatsRequest{})
	if err != nil {
		lg.Warn("系统状况查询失败", zap.String("nodeID", req.NodeID), zap.Error(err))
		return nil, status.Errorf(codes.Internal, "查询失败")
	}
	return resp, nil
}

func (s *ServeServer) QueryStats(ctx context.Context, req *manager.QueryStatsRequest) (*statsCommand.QueryStatsResponse, error) {
	lg, _ := logger.StartTrace(ctx, "QueryStats")
	defer lg.TraceSpanEnd()

	if req.Condition == nil {
		req.Condition = &statsCommand.QueryStatsRequest{} // 需要给一个默认值
	}

	conn, err := connections.GetConnect(req.NodeID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	client := statsCommand.NewStatsServiceClient(conn)

	resp, err := client.QueryStats(ctx, req.Condition)
	if err != nil {
		lg.Warn("流量统计查询失败", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "查询失败")
	}
	return resp, nil
}
