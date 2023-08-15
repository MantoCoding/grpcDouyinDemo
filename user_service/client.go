package user_service

import (
	"context"
	"fmt"
	api "github.com/MantoCoding/grpcDouyinDemo/api/client"
	pb "github.com/MantoCoding/grpcDouyinDemo/user_service/user_login_grpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 实现了 client 端调用 server 端的方法

func UserLoginAction() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 客户端
		client := api.C.GetUserLoginClient()

		// 解析请求参数
		var req pb.DouyinUserLoginRequest
		req.Username = c.Query("username")
		req.Password = c.Query("password")

		// 调用gRPC服务
		result, err := client.Login(context.Background(), &req)
		if err != nil {
			//c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to call gRPC user_service 0814"})
			//c.JSON(http.StatusOK, gin.H{"error": "Failed to call gRPC user_service"})
			return
		}
		// 处理登录响应
		if result.StatusMsg == "Succeed" {
			c.JSON(http.StatusOK, gin.H{
				"message":    "登录成功",
				"StatusCode": result.StatusCode,
				"StatusMsg":  result.StatusMsg,
				"Token":      result.Token,
			})
		} else {
			c.JSON(int(result.StatusCode), gin.H{"StatusMsg": result.StatusMsg})
		}
		fmt.Println(result, err)
	}
}
