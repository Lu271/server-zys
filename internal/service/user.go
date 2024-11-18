package service

import (
	"context"
	"math/rand"
	"server-zys/internal/dao"
	"server-zys/internal/entity"
	"strconv"
)

func GetUserInfo(id int) *entity.GetUserInfoResp {
	return &entity.GetUserInfoResp{
		UserId:   id,
		UserName: "test-name",
		Gender:   1,
		Age:      12,
	}
}

func UserLogin(ctx context.Context, req entity.LoginReq) (resp *entity.LoginResp, err error) {
	// 根据用户名密码查询用户
	user, _ := dao.GetMallUser(ctx, req.Account, req.Password)
	if user.ID == 0 {
		user.Account = req.Account
		user.Password = req.Password
		user.NickName = "匿名用户" + strconv.Itoa(rand.Intn(10000))
		err = dao.CreateMallUser(ctx, user)
	}

	return &entity.LoginResp{
		MallUser: *user,
	}, nil
}
