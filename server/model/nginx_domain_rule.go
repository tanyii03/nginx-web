// 自动生成模板DomainRule
package model

import (
	"nginx-web/global"
)

// 如果含有time.Time 请自行import time包
type DomainRule struct {
      global.GVA_MODEL
      MatchType  string `json:"MatchType" form:"MatchType" gorm:"column:match_type;comment:"`
      Path  string   `json:"Path" form:"Path" gorm:"column:path;comment:"`
      HttpsRewrite  string `json:"HttpsRewrite" form:"HttpsRewrite" gorm:"column:https_rewrite;comment:"`
      Domain  NginxDomain  `json:"Domain" gorm:"foreignKey:ID;references:DomainId;comment:Domain"`
      DomainId uint  `json:"DomainId" gorm:"default:1;comment:Domain ID"`
}


func (DomainRule) TableName() string {
  return "nginx_domain_rule"
}
