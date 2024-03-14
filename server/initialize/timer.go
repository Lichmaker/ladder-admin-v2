package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/task"
	"github.com/flipped-aurora/gin-vue-admin/server/task/blockcheck"
	"github.com/flipped-aurora/gin-vue-admin/server/task/datastat"
	"github.com/flipped-aurora/gin-vue-admin/server/task/usercheck"
	"go.uber.org/zap"

	"github.com/robfig/cron/v3"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

func Timer() {
	go func() {
		var option []cron.Option
		option = append(option, cron.WithSeconds())
		// 清理DB定时任务
		_, err := global.GVA_Timer.AddTaskByFunc("ClearDB", "@daily", func() {
			err := task.ClearTable(global.GVA_DB) // 定时任务方法定在task文件包中
			if err != nil {
				global.GVA_LOG.Error("timer error:", zap.Error(err))
			}
		}, "定时清理数据库【日志，黑名单】内容", option...)
		if err != nil {
			global.GVA_LOG.Error("add timer error:", zap.Error(err))
		}

		// 每20秒检查节点在线状态
		_, err = global.GVA_Timer.AddTaskByFunc("NodeHeartbeat",
			"11,31,51 * * * * *",
			task.NodeHeartbeat,
			"定时检查节点在线", option...)
		if err != nil {
			global.GVA_LOG.Error("add timer error:", zap.Error(err))
		}

		// 每20秒检查节点端口是否被block
		_, err = global.GVA_Timer.AddTaskByFunc("NodeHeartbeat",
			"4,24,44 * * * * *",
			blockcheck.Check,
			"定时检查节点端口", option...)
		if err != nil {
			global.GVA_LOG.Error("add timer error:", zap.Error(err))
		}

		// 每30秒统计流量数据
		_, err = global.GVA_Timer.AddTaskByFunc("NodeDataStat",
			"24,54 * * * * *",
			datastat.Stat,
			"定时查询节点流量数据", option...)
		if err != nil {
			global.GVA_LOG.Error("add timer error:", zap.Error(err))
		}

		// 每分钟检查账号用量和有效期
		_, err = global.GVA_Timer.AddTaskByFunc("NodeDataStat",
			"2 * * * * *",
			usercheck.Check,
			"每分钟检查账号用量和有效期", option...)
		if err != nil {
			global.GVA_LOG.Error("add timer error:", zap.Error(err))
		}

		// 每30分钟，重置所有节点的用户
		_, err = global.GVA_Timer.AddTaskByFunc("NodeDataStat",
			"43 2,32 * * * *",
			usercheck.ReloadUser,
			"定时重置所有节点的用户", option...)
		if err != nil {
			global.GVA_LOG.Error("add timer error:", zap.Error(err))
		}

		// 其他定时任务定在这里 参考上方使用方法

		//_, err := global.GVA_Timer.AddTaskByFunc("定时任务标识", "corn表达式", func() {
		//	具体执行内容...
		//  ......
		//}, option...)
		//if err != nil {
		//	fmt.Println("add timer error:", err)
		//}
	}()
}
