package request

import "nginx-web/model"

type SysDictionaryDetailSearch struct{
    model.SysDictionaryDetail
    PageInfo
}