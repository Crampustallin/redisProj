package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewServer(h Handler) *http.Server {
	r := gin.Default()

	g := r.Group("/user")
	{
		g.GET("/:userName", h.GetUserHandler)
		g.POST("/new", h.SetUserHandler)
	}
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: r,
	}
}

type Handler interface {
	GetUserHandler(c *gin.Context)
	SetUserHandler(c *gin.Context)
}
