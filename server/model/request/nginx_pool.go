package request

import "nginx-web/model"

type NginxPoolSearch struct{
    model.NginxPool
    PageInfo
}