package bootstrap

import (
	"v2ray-manager/pkg/config"
	"v2ray-manager/pkg/logger"
)

// SetupLogger 初始化 Logger
func SetupLogger() {
	logger.InitLogger(
		config.GetString("log.filename"),
		config.GetInt("log.max_size"),
		config.GetInt("log.max_backup"),
		config.GetInt("log.max_age"),
		config.GetBool("log.compress"),
		config.GetString("log.encoder"),
		config.GetString("log.level"),
	)
}
