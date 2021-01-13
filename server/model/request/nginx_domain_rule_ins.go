package request

import "nginx-web/model"

type DomainRuleInsSearch struct{
    model.DomainRuleIns
    PageInfo
}