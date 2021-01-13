package response

import "nginx-web/model"

type ExaFileResponse struct {
	File model.ExaFileUploadAndDownload `json:"file"`
}
