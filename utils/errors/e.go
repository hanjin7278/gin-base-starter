package errors

const (
	SUCCESS = 200
	ERROR   = 500
)

var errormsg map[int]string

/**
初始化错误码
*/
func init() {
	errormsg = make(map[int]string)

	errormsg[200] = "success"
	errormsg[500] = "errors"
}

/**
获取错误描述信息
*/
func GetError(code int) string {
	for k, v := range errormsg {
		if k == code {
			return v
		}
	}
	return errormsg[500]
}
