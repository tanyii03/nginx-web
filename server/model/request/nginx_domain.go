package request

import "nginx-web/model"

type NginxDomainSearch struct{
    model.NginxDomain
    PageInfo
}