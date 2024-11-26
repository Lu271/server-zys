package etcd

import (
	hello "github.com/Lu271/rpc-test/hello-server/kitex_gen/hello/helloservice"
	"github.com/Lu271/server-zys/api/middleware"
	"github.com/cloudwego/kitex/client"
	etcd "github.com/kitex-contrib/registry-etcd"
	"sync"
)

var (
	etcdCli hello.Client
	once    sync.Once
)

func GetEtcdClient() hello.Client {
	if etcdCli != nil {
		return etcdCli
	}
	once.Do(func() {
		initEtcd()
	})
	return etcdCli
}

func initEtcd() {
	// 初始化服务发现对象
	r, _ := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})

	// 初始化 RPC 客户端对象
	etcdCli = hello.MustNewClient("hello", client.WithResolver(r),
		client.WithMiddleware(middleware.LogMiddleware))
}
