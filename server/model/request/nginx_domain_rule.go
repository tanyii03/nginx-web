package request

import "nginx-web/model"

type DomainRuleSearch struct{
    model.DomainRule
    PageInfo
}