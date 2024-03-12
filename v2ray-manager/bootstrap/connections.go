package bootstrap

import (
	"v2ray-manager/app/connections"
	"v2ray-manager/nodecfg"
)

func SetupConnections() {
	for _, v := range nodecfg.GetAllNode() {
		connections.BuildConnect(v.ID)
	}
}
