package controller

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func newResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func successResponse(data interface{}) *Response {
	return newResponse(0, "success", data)
}

func failResponse(code int, msg string) *Response {
	return newResponse(code, msg, nil)
}
