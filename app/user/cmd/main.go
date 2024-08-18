package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"grpc_todolist/app/user/internal/repository/db/dao"
	"grpc_todolist/app/user/internal/service"
	"grpc_todolist/config"
	pb "grpc_todolist/idl/pb/user"
	"grpc_todolist/pkg/discovery"
	"net"
)

func main() {
	config.InitConfig()
	dao.InitDB()

	// etcd 地址
	etcdAddress := []string{config.Conf.Etcd.Address}
	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	// 127.0.0.1:10002
	grpcAddress := config.Conf.Services["user"].Addr[0]

	defer etcdRegister.Stop()
	userNode := discovery.Server{
		Name: config.Conf.Domain["user"].Name,
		Addr: grpcAddress,
	}
	// 正常的 grpc 服务启动
	server := grpc.NewServer()
	defer server.Stop()

	// 绑定 service
	pb.RegisterUserServiceServer(server, service.GetUserSrv())
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}

	if _, err := etcdRegister.Register(userNode, 10); err != nil {
		panic(fmt.Sprintf("start server failed, err: %v", err))
	}

	logrus.Info("server started listen on ", grpcAddress)
	// 简单而言，就是个死循环，Accept() 之后，处理请求
	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
