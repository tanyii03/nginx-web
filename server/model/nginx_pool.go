// 自动生成模板NginxPool
package model

import "nginx-web/global"

// 如果含有time.Time 请自行import time包
type NginxPool struct {
      global.GVA_MODEL
      PoolName  string  `json:"PoolName" form:"PoolName" gorm:"uniqueIndex;column:pool_name;comment:pool名称;type:string"`
      Policy  string    `json:"Policy" form:"Policy" gorm:"column:policy;type:string"`
      Keepalive  int `json:"Keepalive" form:"Keepalive" gorm:"column:keepalive;type:int;default:20"`
      CheckType  string `json:"CheckType" form:"CheckType" gorm:"column:check_type;type:string;default:tcp"`
      CheckInterval  int `json:"CheckInterval" form:"CheckInterval" gorm:"column:CheckInterval;type:int;default:3000"`
      CheckTimeout  int `json:"CheckTimeout" form:"CheckTimeout" gorm:"column:check_timeout;type:int;default:3000"`
      Authority   SysAuthority `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户部门"`
      AuthorityId string       `json:"authorityId" gorm:"default:888;comment:用户部门ID"`
}


func (NginxPool) TableName() string {
      return "nginx_pool"
}
