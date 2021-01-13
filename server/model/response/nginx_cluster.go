package response

import "nginx-web/model"

type NginxClusterResponse struct {
	NginxCluster model.NginxCluster `json:"nginxCluster"`
}

type NginxClusterListResponse struct {
	NginxClusters []model.NginxCluster `json:"nginxClusters"`
}


type NginxPoolResponse struct {
	NginxPool model.NginxPool `json:"nginxPool"`
}

type NginxPoolListResponse struct {
	NginxPools []model.NginxPool `json:"nginxPools"`
}



type NginxDomainResponse struct {
	NginxDomain model.NginxDomain `json:"nginxDomain"`
}

type NginxDomainListResponse struct {
	NginxDomains []model.NginxDomain `json:"nginxDomains"`
}