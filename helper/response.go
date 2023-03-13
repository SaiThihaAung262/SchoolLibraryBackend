package helper

type Response struct {
	ErrorCode    uint64      `json:"err_code"`
	ErrorMessage string      `json:"err_msg"`
	Data         interface{} `json:"data"`
}

type ResponseErr struct {
	ErrorCode    uint64 `json:"err_code"`
	ErrorMessage string `json:"err_msg"`
}

type ResponseUserData struct {
	Type     uint64      `json:"type" form:"type"`
	UserData interface{} `json:"user_data" form:"user_data"`
}

type EmptyObj struct{}

func ResponseData(err_code uint64, err string, data interface{}) Response {
	response := Response{
		ErrorCode:    err_code,
		ErrorMessage: err,
		Data:         data,
	}
	return response
}

func ResponseErrorData(err_code uint64, err string) ResponseErr {
	response := ResponseErr{
		ErrorCode:    err_code,
		ErrorMessage: err,
	}
	return response
}
