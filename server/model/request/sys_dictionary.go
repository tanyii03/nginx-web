package request

import "nginx-web/model"

type SysDictionarySearch struct{
    model.SysDictionary
    PageInfo
}