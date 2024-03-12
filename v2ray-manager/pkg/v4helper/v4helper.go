package v4helper

import (
	"encoding/json"

	v4 "github.com/v2fly/v2ray-core/v5/infra/conf/v4"
)

func ParseToV4(str string) v4.Config {
	jsonConfig := v4.Config{}
	json.Unmarshal([]byte(str), &jsonConfig)
	return jsonConfig
}

func FindVmessTag(str string) string {
	v4Config := ParseToV4(str)
	for _, v := range v4Config.InboundConfigs {
		if v.Protocol == "vmess" {
			return v.Tag
		}
	}
	return ""
}
