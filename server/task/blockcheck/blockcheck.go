package blockcheck

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/pkg/rpcconnect"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/vpsproxy"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

var v2rayNodeService = service.ServiceGroupApp.V2raynodeServiceGroup.V2rayNodeService
var failedTimeBuffer map[string]int

const BLOCK_STANDER_TIMES = 5

func init() {
	failedTimeBuffer = make(map[string]int)
}

func Check() {
	allNode, err := v2rayNodeService.GetAllNode()
	if err != nil {
		global.GVA_LOG.Warn("数据查询失败", zap.Error(err))
		return
	}

	for _, v := range allNode {
		ctx := context.Background()
		if !v.CheckOnline() {
			continue
		}

		vmessTargetAddr := fmt.Sprintf("%s:%d", v.Host, v.VmessPort)
		if err := dialAndCheckPort(vmessTargetAddr, time.Second*3); err != nil {
			global.GVA_LOG.Debug("dial失败", zap.Error(err))
			failedTimeBuffer[v.UniqueId]++
		} else {
			failedTimeBuffer[v.UniqueId] = 0
		}

		if failedTimeBuffer[v.UniqueId] > BLOCK_STANDER_TIMES {
			// 触发更新端口并重启
			newPort := v.VmessPort + 1
			secret, _ := v.GetSecret()
			vpsproxyAddr := fmt.Sprintf("%s:%d", v.Host, v.VPSProxyPort)

			conn, err := rpcconnect.Dial(vpsproxyAddr, buildUnaryClientInterceptor(secret))
			if err != nil {
				global.GVA_LOG.Warn("vpsproxy 连接失败", zap.String("nodeID", v.UniqueId), zap.Error(err))
				continue
			}

			client := vpsproxy.NewVPSProxyServiceClient(conn)
			if _, err := client.UpdateHttpProxy(ctx, &vpsproxy.UpdateHttpProxyReq{
				Port: uint32(newPort),
			}); err != nil {
				global.GVA_LOG.Warn("远程更新端口失败", zap.String("nodeId", v.UniqueId), zap.Error(err))
				continue
			}

			if _, err := client.Shutdown(ctx, &emptypb.Empty{}); err != nil {
				global.GVA_LOG.Warn("远程关闭程序失败", zap.String("nodeId", v.UniqueId), zap.Error(err))
				continue
			}

			if err := v2rayNodeService.UpdateVmessPort(v.UniqueId, newPort); err != nil {
				global.GVA_LOG.Warn("更新数据失败", zap.String("nodeId", v.UniqueId), zap.Error(err))
			}

			// 下发给manager更新
			v.VmessPort = newPort
			if _, err := global.V2RAY_MANAGER_CLIENT.AddNode(ctx, v.ToManagerNode()); err != nil {
				global.GVA_LOG.Warn("下发数据到manager失败", zap.String("nodeId", v.UniqueId), zap.Error(err))
			}
		}
	}
}

func dialAndCheckPort(addr string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil

}

func buildUnaryClientInterceptor(secret string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 获取要发送给服务端的`metadata`
		md, ok := metadata.FromOutgoingContext(ctx)
		if ok && len(md.Get("secret")) == 0 {
			ctx = metadata.AppendToOutgoingContext(ctx, "secret", secret)
		} else {
			newMd := make(metadata.MD)
			newMd.Set("secret", secret)
			ctx = metadata.NewOutgoingContext(ctx, newMd)
		}

		// 设置5秒的超时
		ctx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		// Invoking the remote method
		err := invoker(ctx, method, req, reply, cc, opts...)
		return err
	}
}
