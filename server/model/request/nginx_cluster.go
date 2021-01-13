package request

import "nginx-web/model"

type NginxClusterSearch struct{
    model.NginxCluster
    PageInfo
}
