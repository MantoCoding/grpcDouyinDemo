package service

import (
	"context"
	pb "douyinLoginDemo/service/user_login_grpc"
	"fmt"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

func GrpcMiddleware(client pb.LoginServiceClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		//username := c.Query("username")
		//password := c.Query("password")

		//创建grpc客户端/grpc连接
		addr := ":8083"
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			//log.Fatalf(fmt.Sprintf("grpc connect addr 连接失败"))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		}

		//初始化客户端
		//api := pb.NewLoginServiceClient(conn)
		//result, err := api.Login(context.Background(), &pb.DouyinUserLoginRequest{
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
			c.JSON(http.StatusOK, gin.H{"message": "登录成功"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "无法访问"})
		}
		fmt.Println(result, err)
		//fmt.Println(result, err)

	}
}

//func main() {
//	addr := ":5050"
//	// 使用 grpc.Dial 创建一个到指定地址的 gRPC 连接。
//	// 此处使用不安全的证书来实现 SSL/TLS 连接
//	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		log.Fatalf(fmt.Sprintf("grpc connect addr [%s] 连接失败 %s", addr, err))
//	}
//	defer conn.Close()
//
//	//userClient := pb.NewLoginServiceClient(conn)
//}
