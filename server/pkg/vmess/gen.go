package vmess

import (
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/model/v2raynode"
)

type VMESS struct {
	V    string `json:"v"`
	Ps   string `json:"ps"`
	Add  string `json:"add"`
	Port string `json:"port"`
	ID   string `json:"id"`
	Aid  uint32 `json:"aid"`
	Scy  string `json:"scy"`
	Net  string `json:"net"`
	Type string `json:"type"`
	Host string `json:"host"`
	Path string `json:"path"`
	TLS  string `json:"tls"`
	Sni  string `json:"sni"`
	Alpn string `json:"alpn"`
	Fp   string `json:"fp"`
}

func (v *VMESS) Default() {
	v.V = "2"
	v.Scy = "auto"
	v.Net = "ws"
	v.Type = "none"
	v.Host = "www.bilibili.com" // TODO 根据时间变更伪装host
	v.TLS = "tls"
	v.Alpn = "h2"
	v.Fp = "chrome"
}

func (v *VMESS) Build(UUID string, alterId uint32, nodeInfo v2raynode.V2rayNode) string {
	v.Port = fmt.Sprint(nodeInfo.VmessPort)
	v.Add = nodeInfo.Host
	v.ID = UUID
	v.Aid = alterId
	v.Path = nodeInfo.GetWSPath()
	v.Sni = nodeInfo.Host
	v.Ps = fmt.Sprintf("%s[由Ladder-Admin-V2导入]", nodeInfo.Name)
	jsonBytes, _ := json.Marshal(v)
	base64Str := base64.StdEncoding.EncodeToString(jsonBytes)
	return fmt.Sprintf("vmess://%s", base64Str)
}
