package adapters

import (
	"fmt"
	"hexagonal_go/core"
	"net/http"

	"github.com/gin-gonic/gin"
)

// primary adapter
type HttpOrderHandler struct {
	service core.OrderService
}

func NewHttpOrderHandler(service core.OrderService) *HttpOrderHandler {
	return &HttpOrderHandler{
		service: service,
	}
}

func (h *HttpOrderHandler) CreateOrder(c *gin.Context) {
	var order core.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if err := h.service.CreateOrder(order); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}
