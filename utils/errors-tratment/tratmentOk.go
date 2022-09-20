package errorstratment

type ResOk struct {
	Code    *int64  `json:"code" example:"200"`
	Message *string `json:"message" example:"Project deleted successfully"`
}

func KeyOk(message string, code int64) (res *ResOk) {
	res = &ResOk{}
	res.Code = &code
	res.Message = &message
	
	return
}