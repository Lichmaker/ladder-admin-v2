// 自动生成模板V2rayNode
package v2raynode

import (
	"context"
	"encoding/json"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/protobuf/manager"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/golang-module/carbon/v2"
	command "github.com/v2fly/v2ray-core/v5/app/stats/command"
	"go.uber.org/zap"

	v4 "github.com/v2fly/v2ray-core/v5/infra/conf/v4"
)

// v2rayNode表 结构体  V2rayNode
type V2rayNode struct {
	ID            uint       `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt     time.Time  // 创建时间
	UpdatedAt     time.Time  // 更新时间
	UniqueId      string     `json:"uniqueId" form:"uniqueId" gorm:"column:unique_id;comment:;size:64;unique;"` //uniqueId字段
	ConfigRaw     string     `json:"configRaw" form:"configRaw" gorm:"column:config_raw;type:TEXT;comment:;"`   //configRaw字段
	Name          string     `json:"name" form:"name" gorm:"column:name;comment:;size:64;"`                     //name字段
	Host          string     `json:"host" form:"host" gorm:"column:host;comment:;size:128;"`
	Secret        string     `json:"secret" form:"secret" gorm:"column:secret;comment:;size:64;"`                //secret字段
	SecretIv      string     `json:"secretIv" form:"secretIv" gorm:"column:secret_iv;comment:;size:64;"`         //secretIv字段
	VmessPort     uint       `json:"vmessPort" form:"vmessPort" gorm:"column:vmess_port;type:INT(10);comment:;"` //vmessPort字段
	RpcPort       uint       `json:"rpcPort" form:"rpcPort" gorm:"column:rpc_port;type:INT(10);comment:;"`       //rpcPort字段
	VPSProxyPort  uint       `json:"vpsproxyPort" form:"vpsproxyPort" gorm:"column:vpsproxy_port;type:INT(10);comment:;notnull;default:0"`
	LastContactAt *time.Time `json:"lastContactAt" form:"lastContactAt" gorm:"column:last_contact_at;comment:;"`
	Online        bool       `json:"online" form:"online" gorm:"-"` // 当 LastContactAt 在1分钟内，视为在线
}

// TableName v2rayNode表 V2rayNode自定义表名 la_v2ray_node
func (V2rayNode) TableName() string {
	return "la_v2ray_node"
}

func (v V2rayNode) GetSecret() (string, error) {
	if len(v.SecretIv) == 0 {
		return v.Secret, nil
	}
	d, err := utils.AESDecrypt(global.GVA_CONFIG.V2RayManager.SecretKey, v.SecretIv, v.Secret)
	global.GVA_LOG.Debug("secret", zap.String("d", d), zap.Any("raw", v))
	return d, err
}

func (v V2rayNode) GetSysStat(ctx context.Context) (*command.SysStatsResponse, error) {
	return global.V2RAY_MANAGER_CLIENT.SysStat(ctx, &manager.SysStatRequest{NodeID: v.UniqueId})
}

func (v V2rayNode) CheckOnline() bool {
	if v.LastContactAt == nil {
		return false
	}
	return carbon.Now().SubMinute().ToStdTime().Unix()-v.LastContactAt.Unix() < 60
}

func (v V2rayNode) ToManagerNode() *manager.AddNodeRequest {
	secret, err := v.GetSecret()
	if err != nil {
		global.GVA_LOG.Error("get secret failed", zap.Error(err))
	}

	newNode := &manager.AddNodeRequest{}
	newNode.NodeInfo = &manager.SingleNodeInfo{}
	newNode.NodeInfo.ID = v.UniqueId
	newNode.NodeInfo.Name = v.Name
	newNode.NodeInfo.Host = v.Host
	newNode.NodeInfo.NodeConfigV4 = v.ConfigRaw
	newNode.NodeInfo.ReverseConfig = &manager.NodeReverseConfig{
		VMessPort: int64(v.VmessPort),
		RPCPort:   int64(v.RpcPort),
		RPCSecret: secret,
	}
	return newNode
}

func (v V2rayNode) ParseRawConfigToV4() v4.Config {
	jsonConfig := v4.Config{}
	json.Unmarshal([]byte(v.ConfigRaw), &jsonConfig)
	return jsonConfig
}

func (v V2rayNode) GetWSPath() string {
	v4Config := v.ParseRawConfigToV4()
	for _, v := range v4Config.InboundConfigs {
		if v.StreamSetting != nil && v.StreamSetting.WSSettings != nil {
			return v.StreamSetting.WSSettings.Path
		}
	}
	return ""
}
