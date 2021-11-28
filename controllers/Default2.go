package controllers

import "github.com/gin-gonic/gin"

type Default2 struct{}

func (c *Default2) Index(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "this is Default2 Struct Index",
	})
}
