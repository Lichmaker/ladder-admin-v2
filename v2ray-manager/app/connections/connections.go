package connections

import (
	"context"
	"fmt"
	"time"
	"v2ray-manager/nodecfg"
	"v2ray-manager/pkg/connect"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var cache map[string]*grpc.ClientConn

func init() {
	cache = make(map[string]*grpc.ClientConn)
}

func GetConnect(nodeID string) (*grpc.ClientConn, error) {
	conn, ok := cache[nodeID]
	if !ok {
		return nil, fmt.Errorf("不存在连接:%s", nodeID)
	}

	return conn, nil
}

func CloseAllConnect() {
	for _, v := range cache {
		go v.Close()
	}
	cache = make(map[string]*grpc.ClientConn)
}

func CloseConnect(nodeID string) {
	conn, ok := cache[nodeID]
	if !ok {
		return
	}
	conn.Close()
	delete(cache, nodeID)
}

func BuildConnect(nodeID string) error {
	nodeInfo := nodecfg.GetByID(nodeID)
	if nodeInfo == nil {
		return fmt.Errorf("不存在的节点:%s", nodeID)
	}

	target := fmt.Sprintf("%s:%d", nodeInfo.Host, nodeInfo.ReverseConfig.RPCPort)
	conn, err := connect.Dial(
		target,
		BuildUnaryClientInterceptor(nodeInfo.ReverseConfig.RPCSecret),
	)
	if err != nil {
		return errors.Wrapf(err, "连接失败[%s]", target)
	}
	cache[nodeID] = conn
	return nil
}

func BuildUnaryClientInterceptor(secret string) grpc.UnaryClientInterceptor {
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
