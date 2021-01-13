// 自动生成模板DomainRuleIns
package model

import (
	"nginx-web/global"
)

// 如果含有time.Time 请自行import time包
type DomainRuleIns struct {
      global.GVA_MODEL
      InsType   string `json:"InsType" form:"InsType" gorm:"column:ins_type;comment:"`
      InsValue  string `json:"InsValue" form:"InsValue" gorm:"column:ins_value;comment:"`
	  Rule      DomainRule  `json:"Rule" gorm:"foreignKey:ID;references:RuleId;comment:Domain"`
	  RuleId    uint  `json:"RuleId" gorm:"default:1;comment:Rule ID"`
}


func (DomainRuleIns) TableName() string {
  return "nginx_domain_rule_ins"
}
