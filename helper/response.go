package helper

import "MyGO.com/m/model"

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

type ResponseSummaryData struct {
	BorrowCount uint64     `json:"borrow_count"`
	BookDetail  model.Book `json:"book_data"`
	TotalBook   uint64     `json:"total_book" form:"total_book"`
}

type ResponseSummaryDataList struct {
	List  []ResponseSummaryData `json:"list" form:"list"`
	Total uint64                `json:"total" form:"total"`
}

type ResponBookDetailByUUID struct {
	BookDetail  model.Book `json:"book_data"`
	BorrowCount uint64     `json:"borrow_count"`
}
