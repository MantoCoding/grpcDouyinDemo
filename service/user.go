package service

import (
	"context"
	pb "douyinLoginDemo/service/user_login_grpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"gorm.io/gorm"
	"log"
	"net"
)

type UserLoginService struct {
	pb.UnimplementedLoginServiceServer
	DB *gorm.DB
}

func (UserLoginService) Login(ctx context.Context, req *pb.DouyinUserLoginRequest) (resp *pb.DouyinUserLoginResponse, err error) {
	fmt.Println("微服务调用成功，开始查询")
	fmt.Printf("username : %v ", req.Username)
	fmt.Printf("password : %v ", req.Password)
	username := req.Username
	password := req.Password
	token := username + password
	resp = new(pb.DouyinUserLoginResponse)
	resp.StatusCode = 0
	resp.StatusMsg = "Succeed"
	resp.UserId = 1
	resp.Token = token
	//用户登录验证逻辑
	//数据库查询

	return
}

func main() {
	// 监听端口
	listen, err := net.Listen("tcp", ":8083")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	// 创建一个gRPC服务器实例。
	s := grpc.NewServer()
	server := UserLoginService{}
	// 将server结构体注册为gRPC服务。
	pb.RegisterLoginServiceServer(s, &server)
	fmt.Println("grpc server running :8083")
	// 开始处理客户端请求。
	//reflection.Register(s)
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}

//func LoginServiceLis() {
//	defer func() {
//		err := recover()
//		fmt.Println(err)
//	}()
//	// 监听端口
//	listen, err := net.Listen("tcp", ":8083")
//	if err != nil {
//		grpclog.Fatalf("Failed to listen: %v", err)
//	}
//
//	// 创建一个gRPC服务器实例。
//	s := grpc.NewServer()
//	server := UserLoginService{}
//	// 将server结构体注册为gRPC服务。
//	pb.RegisterLoginServiceServer(s, &server)
//	fmt.Println("grpc server running :8083")
//	// 开始处理客户端请求。
//	err = s.Serve(listen)
//}
