package main

import (
	"gin_router/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 方法一
	r.GET("/r1", controllers.Index)

	// 方法二
	r.GET("/r2", new(controllers.Default2).Index)

	// 方法三
	r.GET("/r3", (&controllers.Default3{}).Index)

	r.Run()
}
