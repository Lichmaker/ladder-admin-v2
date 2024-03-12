package task

import (
	"context"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/v2raynode"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"go.uber.org/zap"
)

func NodeHeartbeat() {
	if global.V2RAY_MANAGER_CLIENT == nil {
		return
	}

	v2rayNodeService := &service.ServiceGroupApp.V2raynodeServiceGroup.V2rayNodeService
	allNode, err := v2rayNodeService.GetAllNode()
	if err != nil {
		global.GVA_LOG.Error("获取所有节点失败", zap.Error(err))
		return
	}

	for _, v := range allNode {
		go singleNodeHeartbeat(v)
	}
}

func singleNodeHeartbeat(model v2raynode.V2rayNode) {
	v2rayNodeService := &service.ServiceGroupApp.V2raynodeServiceGroup.V2rayNodeService

	ctx := context.Background()
	sysStat, err := model.GetSysStat(ctx)
	if err != nil {
		global.GVA_LOG.Warn("节点连接失败", zap.String("id", model.UniqueId), zap.Error(err))
		return
	}
	if sysStat.Alloc <= 0 {
		global.GVA_LOG.Warn("节点数据异常", zap.String("id", model.UniqueId), zap.Any("stat", sysStat))
		return
	}
	if err := v2rayNodeService.UpdateLastContactAt(model.UniqueId, time.Now()); err != nil {
		global.GVA_LOG.Error("数据更新失败", zap.String("id", model.UniqueId), zap.Error(err))
	}
}
