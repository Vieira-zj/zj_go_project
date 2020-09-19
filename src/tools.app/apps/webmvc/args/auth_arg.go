package args

import "src/tools.app/apps/webmvc/model"

// AuthArg 登录接口请求数据
type AuthArg struct {
	PageArg
	model.User
	Code string `json:"code" form:"code"`
	// Kword  string `json:"kword" form:"kword"`
	// Passwd string `json:"passwd" form:"passwd"`
}
