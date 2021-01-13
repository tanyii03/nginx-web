// 自动生成模板PublishHistory
package model

import (
	"nginx-web/global"
)

// 如果含有time.Time 请自行import time包
type PublishHistory struct {
      global.GVA_MODEL
      Uuid     string `json:"Uuid" form:"Uuid" gorm:"column:uuid;comment:"`
      Type     string `json:"Type" form:"Type" gorm:"column:type;comment:"`
      Comment  string `json:"Comment" form:"Comment" gorm:"column:comment;comment:"`
      User     string `json:"User" form:"User" gorm:"column:user;comment:"`
      Operate  string `json:"Operate" form:"Operate" gorm:"column:operate;comment:"`
      Status   string `json:"Status" form:"Status" gorm:"column:status;comment:"`
	  Domain  NginxDomain  `json:"Domain" gorm:"foreignKey:ID;references:DomainId;comment:Domain"`
      Config   string  `json:"Config" form:"Config"`
	  DomainId uint  `json:"DomainId" gorm:"default:0;comment:Domain ID"`
	  ClusterId uint  `json:"ClusterId" gorm:"default:0;comment:Cluster ID"`
}


func (PublishHistory) TableName() string {
  return "nginx_domain_publish_history"
}
