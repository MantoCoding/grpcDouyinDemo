package main

import (
	"douyinLoginDemo/service"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	go service.LoginServiceLis()

	//访问地址，处理我们的请求 Request Response

	initRouter(r)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
