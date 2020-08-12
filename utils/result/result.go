package result

import (
	"github.com/hanjin7278/gin-base-starter/utils/errors"
	"github.com/hanjin7278/gin-base-starter/utils/lib"
)

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
	Time string      `json:"time"`
}

/**
返回正确信息
*/
func Ok(data interface{}) R {
	r := R{errors.SUCCESS, "success", data, lib.GetTimeStr()}
	return r
}

/**
返回错误信息
*/
func Error(code int, msg string) R {
	r := R{Code: errors.ERROR, Msg: "error", Time: lib.GetTimeStr()}
	return r
}

/**
返回自定义错误信息
*/
func Result(code int, msg string, data interface{}) R {
	r := R{Code: errors.ERROR, Msg: "error", Time: lib.GetTimeStr(), Data: data}
	return r
}
