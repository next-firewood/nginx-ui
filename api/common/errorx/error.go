package errorx

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/go-sql-driver/mysql"
	"rui/common"
)

const (
	Success       = 0
	DefaultCode   = iota + 1000
	Param         = 1002
	TokenExpired  = 1003
	TokenRefresh  = 1004
	NonPermission = 1005
	// Business 业务错误
	Business = 1006

	SqlDupKey    = 1007
	UserNotExist = 1008

	PasswordErr = 1009
)

var ErrMap = map[int]error{
	DefaultCode:   fmt.Errorf("发生未知错误"),
	Param:         fmt.Errorf("参数错误"),
	TokenExpired:  fmt.Errorf("信息过期, 请重新登录"),
	TokenRefresh:  fmt.Errorf("刷新token"),
	NonPermission: fmt.Errorf("没有接口权限"),
	PasswordErr:   fmt.Errorf("密码错误"),
	SqlDupKey:     fmt.Errorf("字段值重复"),
	UserNotExist:  fmt.Errorf("用户不存在"),
}

type CodeError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CodeErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (e *CodeError) Error() string {
	return e.Msg
}

func Is(err, target error) bool {
	e, ok := err.(*CodeError)
	if !ok {
		return false
	}

	et, ok := target.(*CodeError)
	if !ok {
		return false
	}

	if et.Code == e.Code && et.Msg == e.Msg {
		return true
	}

	return false
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}

func NewCodeError(code int) error {
	return &CodeError{Code: code, Msg: ErrMap[code].Error()}
}

func NewValidateError(err validator.ValidationErrorsTranslations) error {
	return &CodeError{Code: Param, Msg: fmt.Sprintf("%v", err)}
}

func NewParamError(err error) error {
	return &CodeError{Code: Param, Msg: fmt.Sprintf("%v", err)}
}
func SqlDupErr(err error) (isDup bool, res error) {
	e, ok := err.(*mysql.MySQLError)
	if ok {
		if e.Number == common.SqlDupNum {
			return true, NewCodeError(SqlDupKey)
		}
	}

	return false, err
}

func BusinessErr(msg string) (err error) {
	return &CodeError{Code: Business, Msg: msg}
}
