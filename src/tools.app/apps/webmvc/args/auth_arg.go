package args

import "tools.app/apps/webmvc/model"

// AuthArg 用户登录接口请求参数
type AuthArg struct {
	PageArg
	model.User
	Code string `json:"code" form:"code"`
	// Kword  string `json:"kword" form:"kword"`
	// Passwd string `json:"passwd" form:"passwd"`
}
