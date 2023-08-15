package user_service

import (
	"context"
	"fmt"
	"github.com/MantoCoding/grpcDouyinDemo/user_service/pojo"
	pb "github.com/MantoCoding/grpcDouyinDemo/user_service/user_login_grpc"
	"github.com/MantoCoding/grpcDouyinDemo/utils"
	"gorm.io/gorm"
)

type UserLoginService struct {
	pb.UnimplementedLoginServiceServer
	DB *gorm.DB
}

func (u *UserLoginService) Login(ctx context.Context, req *pb.DouyinUserLoginRequest) (resp *pb.DouyinUserLoginResponse, err error) {
	fmt.Println("微服务调用成功，开始查询")
	fmt.Printf("username : %v ", req.Username)
	fmt.Printf("password : %v ", req.Password)
	username := req.Username
	password := req.Password
	token, err := utils.GenerateJWT(username)
	resp = new(pb.DouyinUserLoginResponse)

	// 参数为空，请求失败
	if len(username) == 0 || len(password) == 0 {

	}

	//用户登录验证逻辑
	user := &pojo.User{}

	db := u.DB.WithContext(ctx)
	db = db.Table("user")
	db = db.Where("name = ?", username).Find(user)
	if db.Error != nil {
		fmt.Println(db.Error)
		resp.StatusCode = 403
		resp.StatusMsg = "Permission denied"
		return
	}

	// 密码错误，登陆失败
	if !utils.ComparePasswords(user.Password, password) {

	}

	fmt.Println(user)

	// 请求成功
	resp.StatusCode = 200
	resp.Token = token
	resp.StatusMsg = "Succeed"
	resp.UserId = 1
	return
}

//func main() {
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
//	//reflection.Register(s)
//	if err := s.Serve(listen); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//
//}
