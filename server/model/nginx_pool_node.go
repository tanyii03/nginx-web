// 自动生成模板PoolNode
package model

import (
	"nginx-web/global"
)

// 如果含有time.Time 请自行import time包
type PoolNode struct {
      global.GVA_MODEL
      NodeName  string `json:"NodeName" form:"NodeName" gorm:"column:node_name;comment:;type:string;"`
      IP  string `json:"IP" form:"IP" gorm:"column:ip;comment:;type:string;"`
      Port  string `json:"Port" form:"Port" gorm:"column:port;comment:"`
      Weight  string `json:"Weight" form:"Weight" gorm:"column:weight;comment:"`
      MaxFailed  string `json:"MaxFailed" form:"MaxFailed" gorm:"column:max_failed;comment:"`
      TimeOut  string `json:"TimeOut" form:"TimeOut" gorm:"column:time_out;comment:"`
      Status  string `json:"Status" form:"Status" gorm:"column:status;comment:"`
      Cluster   NginxPool `json:"Cluster" gorm:"foreignKey:ID;references:ClusterId;comment:Cluster"`
      ClusterId uint `json:"ClusterId" gorm:"default:1;comment:Cluster ID" sql:"id"`
}


func (PoolNode) TableName() string {
  return "nginx_pool_node"
}
