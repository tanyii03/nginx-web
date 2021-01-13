package request

import "nginx-web/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
