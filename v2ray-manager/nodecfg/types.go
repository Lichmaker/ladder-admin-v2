package nodecfg

import (
	"v2ray-manager/protobuf/manager"
)

type NodeInfo struct {
	Name   string `json:"name"`
	ID     string `json:"id"`
	Host   string `json:"host"`
	Config string `json:"config"` // 原始配置json，v4版本
	// Users         []serve.UserInfo  `json:"users"`  // 需加载的用户列表
	ReverseConfig *manager.NodeReverseConfig `json:"reverseConfig"`
}
