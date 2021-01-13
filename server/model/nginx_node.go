// 自动生成模板NginxNode
package model

import (
	"nginx-web/global"
)

// 如果含有time.Time 请自行import time包
type NginxNode struct {
      global.GVA_MODEL
      NodeName  string `json:"NodeName" form:"NodeName" gorm:"column:node_name;comment:;type:string;"`
	  IP        string `json:"IP" form:"IP" gorm:"column:ip;comment:;type:string;"`
	  Port      string `json:"Port" form:"Port" gorm:"column:port;comment:;type:string;"`
      Status    string `json:"Status" form:"Status" gorm:"column:status;comment:;type:string;"`
      UserName  string `json:"UserName" form:"UserName" gorm:"column:user_name;comment:;type:string;"`
      Passwd    string `json:"Passwd" form:"Passwd" gorm:"-"`
	  Cluster   NginxCluster `json:"Cluster" gorm:"foreignKey:ID;references:ClusterId;comment:Cluster"`
	  ClusterId uint       `json:"ClusterId" gorm:"default:1;comment:Cluster ID"`
}


func (NginxNode) TableName() string {
  return "nginx_node"
}
