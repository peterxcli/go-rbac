package handlers

import (
	"easy-rbac/services"

	"github.com/gin-gonic/gin"
)

type ItemHandler interface {
}

type ItemHandlerImpl struct {
	svc services.ItemService
}

func NewItemHandler(svc services.ItemService) ItemHandler {
	return &ItemHandlerImpl{
		svc: svc,
	}
}

func GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "You have permission to view this!",
	})
}
