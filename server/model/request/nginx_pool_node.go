package request

import "nginx-web/model"

type PoolNodeSearch struct{
    model.PoolNode
    PageInfo
}