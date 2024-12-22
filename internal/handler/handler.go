package handler

import (
	"net/http"

	"github.com/Crampustallin/redisProj/internal/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db DataBase
}

func NewServer(db DataBase) *Handler {
	return &Handler{
		db: db,
	}
}

func (h *Handler) GetUser(c *gin.Context) {
	var u *model.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	u, err := h.db.GetUser(u.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": u})
}

func (h *Handler) SetUser(c *gin.Context) {
	var u *model.User
	if err := c.BindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		c.Abort()
		return
	}

	if err := h.db.SetUser(u); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": "saved"})
}

type DataBase interface {
	SetUser(u *model.User) (err error)
	GetUser(name string) (u *model.User, err error)
}
