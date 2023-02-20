package response

import "github.com/defeng-hub/ByOfficeAutomatic/server/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
