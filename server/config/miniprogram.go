package config

type Miniprogram struct {
	AppID     string `mapstructure:"app-id" json:"app-id" yaml:"app-id"`         // 小程序AppID
	AppSecret string `mapstructure:"app-secret" json:"app-secret" yaml:"app-secret"` // 小程序AppSecret
}