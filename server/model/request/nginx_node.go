package request

import "nginx-web/model"

type NginxNodeSearch struct{
    model.NginxNode
    PageInfo
}