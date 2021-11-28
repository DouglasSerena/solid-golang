package server

import (
	"github.com/gin-gonic/gin"
)

type IServer interface {
	SetupRouter() *gin.Engine
}

type Server struct {
	IServer
	Route *gin.Engine
}

func NewServer() *Server {
	route := gin.Default()

	route.SetTrustedProxies([]string{"192.168.1.2"})

	route.GET("/ping", func(context *gin.Context) {
		context.String(200, "pong")
	})

	return &Server{
		Route: route,
	}
}
