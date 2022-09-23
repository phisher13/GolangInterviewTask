package routes

import (
	"github.com/gin-gonic/gin"
	handler "urls/pkg/handlers"
)

func Router(g *gin.RouterGroup) {
	g.GET("/api/v1/:token/", handler.GetHandler())
	g.POST("/api/v1/", handler.PostHandler())
}