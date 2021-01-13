package config

type Server struct {
	JWT     JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis   `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email   `mapstructure:"email" json:"email" yaml:"email"`
	Casbin  Casbin  `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System  `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	// gorm
	Mysql      Mysql      `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	// oss
	Local Local `mapstructure:"local" json:"local" yaml:"local"`
	SSH   SSH   `mapstructure:"ssh" json:"ssh" yaml:"ssh"`
	Nginx Nginx  `mapstructure:"nginx" json:"nginx" yaml:"nginx"`
}
