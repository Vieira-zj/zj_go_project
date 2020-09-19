package service

import (
	"fmt"

	"src/tools.app/apps/webmvc/model"
)

// UserService user实体类的dao服务
type UserService struct {
}

// LoginWithOpenID user登录服务
func (service *UserService) LoginWithOpenID(ID int64) (*model.User, error) {
	// mock
	if ID == 111 {
		return &model.User{
			ID:       111,
			NickName: "test_user_01",
			Role:     model.USER,
		}, nil
	}
	return nil, fmt.Errorf("user id=%d not exist", ID)
}
