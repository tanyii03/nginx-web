// 自动生成模板NginxCluster
package model

import (
	"nginx-web/global"
)

// 如果含有time.Time 请自行import time包
type NginxCluster struct {
      global.GVA_MODEL
      ClusterName string `json:"ClusterName" form:"ClusterName" gorm:"uniqueIndex;column:cluster_name;comment:;type:string"`
	  Vip         string `json:"Vip" form:"Vip" gorm:"column:Vip;type:string"`
      Department  string `json:"Department" form:"Department" gorm:"column:department;type:string"`
	  Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户部门"`
	  AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户部门ID"`
}


func (NginxCluster) TableName() string {
  return "nginx_cluster"
}


