package request

import "nginx-web/model"

type NginxDomainCertSearch struct{
    model.NginxDomainCert
    PageInfo
}