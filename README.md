# Gin 框架 route 文件抽离

在使用 [gin](https://gin-gonic.com/zh-cn/docs/) 框架的时候，该框架为我们的项目提供了很方便的路由处理功能。但是如下示例代码，把路由处理和逻辑处理写在一块，在开发项目的时候造成了很大的不便，所以本文是针对于 gin 框架路由抽离的文章。


```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
```

## 项目的目录展示

```text

```

## 方法一

该方法在 controller 中定义了方法，然后再路由的重写入控制器的方法。在实际开发过程中，不推荐使用此方法，因为项目比较庞大的时候，如果在同一个文件夹下，如果无法规范化命名，则会产生命名冲突。

**main.go**

```go
package main

import (
	"gin_router/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/r1",controllers.Index)
	r.Run()
}
```

**Default1.go**

```go
package controllers

import (
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is Default1 Index",
	})
}
```

## 方法二

该方法的实现和方法三是一致的，先初始化结构体，然后传入结构体的方法。

**main.go**

```go
package main

import (
	"gin_router/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/r2", new(controllers.Default2).Index)
	r.Run()
}

```

**Default2.go**

```go
package controllers

import "github.com/gin-gonic/gin"

type Default2 struct {}

func (c *Default2) Index(g *gin.Context)  {
	g.JSON(200, gin.H{
		"message": "this is Default2 Struct Index",
	})
}

```

## 方法三

改方法编写参照了，beego 对于 [route](https://github.com/beego/beego/blob/v2.0.0/server/web/router.go) (204行)的处理，用 `&` 获取结构体的指针，传入结构体的方法。

**main.go**

```go
package main

import (
	"gin_router/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/r3", (&controllers.Default3{}).Index)
	r.Run()
}

```

**Default3.go**

```go
package controllers

import "github.com/gin-gonic/gin"

type Default3 struct {}

func (c *Default3) Index(g *gin.Context)  {
	g.JSON(200, gin.H{
		"message": "this is Default3 Struct Index function",
	})
}
```