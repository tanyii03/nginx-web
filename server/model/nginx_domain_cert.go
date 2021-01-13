// 自动生成模板NginxDomainCert
package model

import (
	"nginx-web/global"
)

// 如果含有time.Time 请自行import time包
type NginxDomainCert struct {
      global.GVA_MODEL
      CertName  string `json:"CertName" form:"CertName" gorm:"uniqueIndex;column:CertName;comment:"`
      Issued  string `json:"Issued" form:"Issued" gorm:"column:issued;comment:"`
      Dns  string `json:"Dns" form:"Dns" gorm:"column:dns;comment:"`
      Validity  string `json:"Validity" form:"Validity" gorm:"column:validity;comment:"`
      Deadline  string `json:"Deadline" form:"Deadline" gorm:"column:deadline;comment:"`
}

type NginxDomainCertFile struct {
      global.GVA_MODEL
      CertName  string `json:"CertName" form:"CertName"`
      Pem  string `json:"Pem" form:"Pem"`
      Key  string `json:"Key" form:"Key"`
}


func (NginxDomainCert) TableName() string {
  return "nginx_domain_cert"
}
