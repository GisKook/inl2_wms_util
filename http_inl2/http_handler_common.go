package http_inl2

import (
	"encoding/json"
)

const (
	UNKNOW_ERROR string = "未知错误"

	HTTP_RESPONSE_RESULT_SUCCESS       uint8 = 0
	HTTP_RESPONSE_RESULT_PARAMTER_ERR  uint8 = 255
	HTTP_RESPONSE_RESULT_TIMEOUT       uint8 = 254
	HTTP_RESPONSE_RESULT_SERVER_FAILED uint8 = 253
)

var HTTP_RESULT map[uint8]string = map[uint8]string{
	HTTP_RESPONSE_RESULT_SUCCESS:       "成功",
	HTTP_RESPONSE_RESULT_PARAMTER_ERR:  "参数错误",
	HTTP_RESPONSE_RESULT_TIMEOUT:       "超时",
	HTTP_RESPONSE_RESULT_SERVER_FAILED: "失败,服务器内部错误"}

type GeneralResponse struct {
	Code uint8  `json:"code"`
	Desc string `json:"desc"`
}

func EncodingGeneralResponse(code uint8, errmsg string) string {
	if errmsg == "" {
		errmsg = HTTP_RESULT[code]
		errmsg = UNKNOW_ERROR
	}
	general_response := &GeneralResponse{
		Code: code,
		Desc: errmsg,
	}

	response, _ := json.Marshal(general_response)

	return string(response)
}
