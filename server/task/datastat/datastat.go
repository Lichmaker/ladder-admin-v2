package datastat

import (
	"context"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/manager"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/v2fly/v2ray-core/v5/app/stats/command"
	"go.uber.org/zap"
)

var v2rayNodeService = service.ServiceGroupApp.V2raynodeServiceGroup.V2rayNodeService
var userService = service.ServiceGroupApp.SystemServiceGroup.UserService
var dataStatisticsService = service.ServiceGroupApp.DatastatisticsmodelServiceGroup.DataStatisticsService
var userDataUsageService = service.ServiceGroupApp.UserDataUsageModelServiceGroup

func Stat() {
	current := time.Now()

	nodeQuery, err := v2rayNodeService.GetAllNode()
	if err != nil {
		global.GVA_LOG.Error("查询节点失败", zap.Error(err))
		return
	}

	for _, nodeItem := range nodeQuery {
		if !nodeItem.CheckOnline() {
			// 不在线直接忽略
			global.GVA_LOG.Debug("节点离线，不进行流量统计", zap.String("nodeID", nodeItem.UniqueId))
			continue
		}

		ctx := context.Background()
		statQuery, err := global.V2RAY_MANAGER_CLIENT.QueryStats(ctx, &manager.QueryStatsRequest{
			NodeID: nodeItem.UniqueId,
			Condition: &command.QueryStatsRequest{
				Reset_: true,
			},
		})
		if err != nil {
			global.GVA_LOG.Error("查询流量统计失败", zap.String("nodeID", nodeItem.UniqueId), zap.Error(err))
			continue
		}

		groupData := make(map[string]int64)
		for _, statItem := range statQuery.Stat {
			email, value, _ := parseUserData(statItem)
			if len(email) == 0 || value == 0 {
				continue
			}
			groupData[email] += value
		}
		global.GVA_LOG.Debug("节点统计数据", zap.String("nodeID", nodeItem.UniqueId), zap.Any("data", groupData))

		for emailItem, valueItem := range groupData {
			if err := dataStatisticsService.SaveTargetDateData(emailItem, nodeItem.UniqueId, current, int(valueItem)); err != nil {
				global.GVA_LOG.Error("统计数据写入失败！", zap.Error(err))
				continue
			}

			user, err := userService.GetUserInfoByEmail(emailItem)
			if err != nil {
				global.GVA_LOG.Warn("不存在的用户或者数据库异常", zap.String("email", emailItem), zap.Error(err))
				continue
			}

			if err := userDataUsageService.SaveTargetDateData(user.UUID.String(), current, int(valueItem)); err != nil {
				global.GVA_LOG.Error("用量数据写入失败！", zap.Error(err))
			}
		}
	}
}

func parseUserData(statItem *command.Stat) (email string, value int64, uplink bool) {
	parseName := strings.Split(statItem.Name, ">>>")
	if len(parseName) != 4 {
		return
	}
	if parseName[0] != "user" {
		return
	}
	email = parseName[1]
	uplink = parseName[3] == "uplink"
	value = statItem.Value
	return
}
