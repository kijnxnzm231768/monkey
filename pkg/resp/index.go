package resp

// Response 数据返回结构体
type Response struct {
	Status int         `json:"status"` // 返回状态值
	Msg    string      `json:"msg"`    //返回的提示语
	Data   interface{} `json:"data"`   //返回数据
}

// Success 正确返回
func Success(data interface{}, msg ...string) *Response {
	response := Response{
		Status: 200,
		Data:   data,
		Msg:    "操作成功",
	}
	if len(msg) > 0 {
		response.Msg = msg[0]
	}
	return &response
}

// ErrorResp 错误返回
func ErrorResp(data ...interface{}) *Response {
	response := Response{
		Status: 401,
		Msg:    "操作失败",
		Data:   nil,
	}
	for _, value := range data {
		switch value.(type) {
		case string:
			response.Msg = value.(string)
		case int:
			response.Status = value.(int)
		case interface{}:
			response.Data = value.(interface{})
		}
	}
	return &response
}