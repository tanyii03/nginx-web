// 自动生成模板NginxDomain
package model

import (
	"nginx-web/global"
)

// 如果含有time.Time 请自行import time包
type NginxDomain struct {
      global.GVA_MODEL
      DomainName  string `json:"DomainName" form:"DomainName" gorm:"uniqueIndex;column:domain_name;comment:"`
      Domain  string `json:"Domain" form:"Domain" gorm:"column:domain;comment:"`
      Port  string `json:"Port" form:"Port" gorm:"column:port;comment:"`
      PoolCluster   NginxPool `json:"PoolCluster" gorm:"foreignKey:ID;references:Pool;comment:PoolCluster"`
      Pool  uint `json:"Pool" form:"Pool" gorm:"column:pool;comment:"`
      NginxCluster   NginxPool `json:"NginxCluster" gorm:"foreignKey:ID;references:Nginx;comment:NginxCluster"`
      Nginx  uint `json:"Nginx" form:"Nginx" gorm:"column:nginx;comment:"`
      UseHttps  bool `json:"UseHttps" form:"UseHttps" gorm:"column:use_https;comment:"`
      HttpsRewrite  string `json:"HttpsRewrite" form:"HttpsRewrite" gorm:"column:https_rewrite;comment:"`
      HttpsPort  string `json:"HttpsPort" form:"HttpsPort" gorm:"column:https_port;comment:"`
      HttpsCert  string `json:"HttpsCert" form:"HttpsCert" gorm:"column:https_cert;comment:"`
}


func (NginxDomain) TableName() string {
  return "nginx_domain"
}
