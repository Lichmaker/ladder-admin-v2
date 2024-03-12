package v4helper

import "encoding/json"

type SimpleV2rayV4Config struct {
	Inbounds  []Inbounds   `json:"inbounds"`
	Outbounds []Outbounds  `json:"outbounds"`
	Log       Log          `json:"log"`
	Policy    PolicyConfig `json:"policy"`
	Stats     Stats        `json:"stats"`
	Routing   Routing      `json:"routing"`
	API       API          `json:"api"`
}
type WsSettings struct {
	Path                 string            `json:"path"`
	Headers              map[string]string `json:"headers,omitempty"`
	AcceptProxyProtocol  bool              `json:"acceptProxyProtocol,omitempty"`
	MaxEarlyData         int32             `json:"maxEarlyData,omitempty"`
	UseBrowserForwarding bool              `json:"useBrowserForwarding,omitempty"`
	EarlyDataHeaderName  string            `json:"earlyDataHeaderName,omitempty"`
}
type StreamSettings struct {
	Network    string     `json:"network"`
	WsSettings WsSettings `json:"wsSettings"`
}
type Inbounds struct {
	Tag            string           `json:"tag"`
	Port           int              `json:"port"`
	Listen         string           `json:"listen"`
	Protocol       string           `json:"protocol"`
	Settings       *json.RawMessage `json:"settings"`
	StreamSettings StreamSettings   `json:"streamSettings,omitempty"`
}
type Outbounds struct {
	Protocol string           `json:"protocol"`
	Settings *json.RawMessage `json:"settings"`
}
type Log struct {
	Access   string `json:"access"`
	Error    string `json:"error"`
	Loglevel string `json:"loglevel"`
}
type SystemPolicy struct {
	StatsInboundUplink    bool `json:"statsInboundUplink"`
	StatsInboundDownlink  bool `json:"statsInboundDownlink"`
	StatsOutboundUplink   bool `json:"statsOutboundUplink"`
	StatsOutboundDownlink bool `json:"statsOutboundDownlink"`
	OverrideAccessLogDest bool `json:"overrideAccessLogDest"`
}

type PolicyConfig struct {
	Levels map[uint32]*Policy `json:"levels"`
	System *SystemPolicy      `json:"system"`
}
type Policy struct {
	Handshake         *uint32 `json:"handshake"`
	ConnectionIdle    *uint32 `json:"connIdle"`
	UplinkOnly        *uint32 `json:"uplinkOnly"`
	DownlinkOnly      *uint32 `json:"downlinkOnly"`
	StatsUserUplink   bool    `json:"statsUserUplink"`
	StatsUserDownlink bool    `json:"statsUserDownlink"`
	BufferSize        *int32  `json:"bufferSize"`
}

type Stats struct {
}
type Rules struct {
	Type        string   `json:"type"`
	InboundTag  []string `json:"inboundTag"`
	OutboundTag string   `json:"outboundTag"`
}
type Routing struct {
	Rules []Rules `json:"rules"`
}
type API struct {
	Tag      string   `json:"tag"`
	Services []string `json:"services"`
}
