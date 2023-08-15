package main

import (
	"context"
	"douyinLoginDemo/service"
	"google.golang.org/grpc/grpclog"
	"net"

	//"douyinLoginDemo/service"
	pb "douyinLoginDemo/service/user_login_grpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func initRouter(r *gin.Engine) {
	// public directory is used to serve static resources

	//apiRouter := r.Group("/douyin")

	//创建数据库连接
	db_username := "douyin" //账号
	db_password := "douyin" //密码
	host := "43.143.14.234" //数据库地址，可以是Ip或者域名
	port := 3306            //数据库端口
	Dbname := "mini_douyin" //数据库名
	timeout := "10s"        //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", db_username, db_password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn))
	fmt.Println(db)
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println("数据库连接成功")

	go func() {
		// 创建gRPC服务
		grpcServer := grpc.NewServer()

		// 注册LoginService服务
		//loginSrv := &service.UserLoginService{db: db} // 传入GORM数据库连接
		pb.RegisterLoginServiceServer(grpcServer, &service.UserLoginService{DB: db})
		fmt.Println("grpc server running : 8083 ")

		listen, err := net.Listen("tcp", ":8083")
		if err != nil {
			grpclog.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcServer.Serve(listen); err != nil {

		}
	}()

	//获取请求参数，调用grpc客户端
	r.POST("/douyin/user/login/", func(c *gin.Context) {
		fmt.Println("gin ser running : 8080 ")
		//username := c.Query("username")
		//password := c.Query("password")

		//创建grpc客户端/grpc连接
		addr := ":8083"
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			//log.Fatalf(fmt.Sprintf("grpc connect addr 连接失败"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		}

		defer conn.Close()
		//初始化客户端
		//client := pb.NewLoginServiceClient(conn)
		//result, err := client.Login(context.Background(), &pb.DouyinUserLoginRequest{
		//	Username: username,
		//	Password: password,
		//})
		// 解析请求参数
		var req pb.DouyinUserLoginRequest
		if err := c.Bind(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//defer conn.Close()

		// 调用gRPC服务
		client := pb.NewLoginServiceClient(conn)
		result, err := client.Login(context.Background(), &req)
		if err != nil {
			//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call gRPC service 0814"})
			//c.JSON(http.StatusOK, gin.H{"error": "Failed to call gRPC service"})
			return
		}
		// 处理登录响应
		if result.StatusMsg == "Succeed" {
			c.JSON(http.StatusOK, gin.H{
				"message":    "登录成功",
				"StatusCode": result.StatusCode,
				"StatusMsg":  result.StatusMsg,
				"Token":      "token_test",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "无法访问"})
		}
		fmt.Println(result, err)
		//fmt.Println(result, err)

	})

	// basic apis
	//apiRouter.GET("/feed/", controller.Feed)
	//apiRouter.GET("/user/", controller.UserInfo)
	//apiRouter.POST("/user/register/", controller.Register)
	//apiRouter.POST("/user/login/", service.Login)
	//apiRouter.POST("/publish/action/", service.Publish)

}
