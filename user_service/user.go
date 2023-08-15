package user_service

import (
	"context"
	"fmt"
	"github.com/MantoCoding/grpcDouyinDemo/user_service/dao"
	"github.com/MantoCoding/grpcDouyinDemo/user_service/pojo"
	pb "github.com/MantoCoding/grpcDouyinDemo/user_service/user_login_grpc"
	"github.com/MantoCoding/grpcDouyinDemo/utils"
	"gorm.io/gorm"
)

type UserLoginService struct {
	pb.UnimplementedLoginServiceServer
	DB *gorm.DB
}

func NewUserLoginService() *UserLoginService {
	db, err := dao.GetDB()
	if err != nil {
		panic("NewUserLoginService失败")
	}
	return &UserLoginService{
		DB: db,
	}
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
		resp.StatusCode = 400
		resp.StatusMsg = "Bad request"
		return
	}

	//用户登录验证逻辑
	user := &pojo.User{}

	db := u.DB.WithContext(ctx)
	db = db.Table("user")
	db = db.Where("name = ?", username).Find(user)
	if db.Error != nil {
		resp.StatusCode = 500
		resp.StatusMsg = "Internal server error"
		return
	}

	// 密码错误，登陆失败
	if !utils.ComparePasswords(user.Password, password) {
		fmt.Println(user.Password, password)
		resp.StatusCode = 403
		resp.StatusMsg = "Permission denied"
		return
	}

	fmt.Println(user)

	// 请求成功
	resp.StatusCode = 200
	resp.Token = token
	resp.StatusMsg = "Succeed"
	resp.UserId = user.Id
	return
}
