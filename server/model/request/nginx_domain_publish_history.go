package request

import "nginx-web/model"

type PublishHistorySearch struct{
    model.PublishHistory
    PageInfo
}