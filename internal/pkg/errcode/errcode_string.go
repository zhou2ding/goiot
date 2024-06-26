// Code generated by "stringer -type ErrCode -linecomment"; DO NOT EDIT.

package errcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ServerError-10001]
	_ = x[ParamError-10002]
	_ = x[FileError-10003]
	_ = x[TooMandyRequests-10004]
	_ = x[APIKeyAuthFail-10005]
}

const _ErrCode_name = "服务内部错误参数信息有误上传的文件有误请求频率超出限制API Key校验失败"

var _ErrCode_index = [...]uint8{0, 18, 36, 57, 81, 100}

func (i ErrCode) String() string {
	i -= 10001
	if i < 0 || i >= ErrCode(len(_ErrCode_index)-1) {
		return "ErrCode(" + strconv.FormatInt(int64(i+10001), 10) + ")"
	}
	return _ErrCode_name[_ErrCode_index[i]:_ErrCode_index[i+1]]
}
