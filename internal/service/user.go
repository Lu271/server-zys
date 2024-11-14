package service

import "server-zys/internal/entity"

func GetUserInfo(id int) *entity.GetUserInfoResp {
	return &entity.GetUserInfoResp{
		UserId:   id,
		UserName: "test-name",
		Gender:   1,
		Age:      12,
	}
}
