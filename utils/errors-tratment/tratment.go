package errorstratment

type ResError struct {
	Code    *int64  `json:"code"`
	Message *string `json:"message"`
	Error   *string `json:"error"`
}

func KeyError(err, message string, code int64) (res *ResError) {
	res = &ResError{}
	res.Code = &code
	res.Message = &message
	
	res.Error = &err
	return
}
