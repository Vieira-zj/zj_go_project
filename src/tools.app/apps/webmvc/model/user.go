package model

// 用户性别和角色
const (
	FEMALE = 2
	MALE   = 1

	UNKNOW = 0
	ADMIN  = 1
	USER   = 2
)

// User 实体类
type User struct {
	ID       int64  `xorm:"pk autoincr BIGINT(20)" form:"id" json:"id"`
	NickName string `xorm:"VARCHAR(40)" form:"nickName" json:"nickName"`
	Role     int    `xorm:"int(11)" form:"role" json:"role"`
	// Openid   string `xorm:"VARCHAR(40)" form:"openid" json:"openid"`
	// Mobile   string `xorm:"VARCHAR(15)" form:"mobile" json:"mobile"`
	// Passwd   string `xorm:"VARCHAR(40)" form:"passwd" json:"-"`
	// Enable   int    `xorm:"int(11)" form:"enable" json:"enable"`
	// Gender   int    `xorm:"int(11)" form:"gender" json:"gender"`
}
