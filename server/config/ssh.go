package config

type SSH struct {
	Private_key      string `mapstructure:"private_key" json:"privateKey" yaml:"private_key"`
	Public_key       string `mapstructure:"public_key" json:"publicKey" yaml:"public_key"`
}
