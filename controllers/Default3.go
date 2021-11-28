package controllers

import "github.com/gin-gonic/gin"

type Default3 struct{}

func (c *Default3) Index(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "this is Default3 Struct Index function",
	})
}
