package service

import (
	"context"
	"github.com/Lu271/rpc-test/hello-server/kitex_gen/hello"
	"github.com/Lu271/rpc-test/hello-server/kitex_gen/hello/helloservice"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"math/rand"
	"server-zys/internal/dao"
	"server-zys/internal/entity"
	"strconv"
	"time"
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
	dao.Lock(ctx, req.Account, 5*time.Second)
	defer dao.Unlock(ctx, req.Account)

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

func SayHello(ctx context.Context) (*hello.HelloResponse, error) {
	r, _ := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	cli, err := helloservice.NewClient("hello", client.WithResolver(r))
	if err != nil {
		return nil, err
	}

	req := &hello.HelloRequest{
		Message: "client request",
	}
	return cli.SayHello(ctx, req)
}
