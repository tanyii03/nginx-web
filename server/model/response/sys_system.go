package response

import "nginx-web/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
