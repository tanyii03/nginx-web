package config


type Nginx struct {
	ConfigPath      string `mapstructure:"config_path" json:"ConfigPath" yaml:"config_path"`
	CertPath       string `mapstructure:"cert_path" json:"CertPath" yaml:"cert_path"`
}

