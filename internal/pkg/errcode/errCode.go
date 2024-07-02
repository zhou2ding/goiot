package errcode

import "github.com/pkg/errors"

// Response 错误时返回自定义结构
type Response struct {
	Code      ErrCode `json:"code"`      // 错误码
	Msg       string  `json:"msg"`       // 错误信息
	RequestId string  `json:"requestId"` // 请求ID
}

func (e *Response) Error() string {
	return e.Code.String()
}

type ErrCode int

//go:generate stringer -type ErrCode -linecomment

const (
	// ServerError 内部错误
	ServerError      ErrCode = iota + 10001 // 服务内部错误
	ParamError                              // 参数信息有误
	FileError                               // 上传的文件有误
	TooMandyRequests                        // 请求频率超出限制
	APIKeyAuthFail                          // API Key校验失败
)

// NewCustomError 新建自定义error实例化
func NewCustomError(code ErrCode) error {
	// 初次调用得用Wrap方法，进行实例化
	return errors.Wrap(&Response{
		Code: code,
		Msg:  code.String(),
	}, "")
}
