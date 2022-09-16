package errorstratment

type ResOk struct {
	Code    *int64  `json:"code"`
	Message *string `json:"message"`
}

func KeyOk(message string, code int64) (res *ResOk) {
	res = &ResOk{}
	res.Code = &code
	res.Message = &message
	
	return
}