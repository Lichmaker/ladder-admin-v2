package config

type V2RayManager struct {
	Host           string `mapstructure:"host" json:"host" yaml:"host"`
	Port           uint   `mapstructure:"port" json:"port" yaml:"port"`
	DefaultAlterId uint32 `mapstructure:"defaultAlterId" json:"defaultAlterId" yaml:"defaultAlterId"`
	DefaultLevel   uint32 `mapstructure:"defaultLevel" json:"defaultLevel" yaml:"defaultLevel"`
	SecretKey      string `mapstructure:"secretKey" json:"secretKey" yaml:"secretKey"`
}
