package res

import "net/http"

type ResponseJson struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}
type H map[string]interface{}

func InputParamBad(message string) ResponseJson {
	return ResponseJson{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func Error(code int, message string) ResponseJson {
	return ResponseJson{
		Code:    code,
		Message: message,
	}
}

func OKNoMessage() ResponseJson {
	return ResponseJson{
		Code:    1,
		Data:    []interface{}{},
		Message: "ok",
	}
}

func Data(data interface{}, msg ...string) ResponseJson {
	rmsg := "ok"
	if len(msg) != 0 {
		rmsg = msg[0]
	}
	return ResponseJson{
		Code:    1,
		Data:    data,
		Message: rmsg,
	}
}
func NoLogin() ResponseJson {
	return ResponseJson{
		Code:    4059,
		Data:    "",
		Message: "该接口需要登录才可调用",
	}
}
