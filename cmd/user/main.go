package main

import (
	"log"
	"net"

	user "github.com/808-not-found/tik_duck/kitex_gen/user/userservice"
	"github.com/808-not-found/tik_duck/pkg/consts"
	server "github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.EtcdAddress})
	if err != nil {
		panic(err)
	}
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:10001")
	svr := user.NewServer(new(UserServiceImpl), server.WithServiceAddr(addr), server.WithRegistry(r))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
