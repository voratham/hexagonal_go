package main

import (
	"fmt"
	"hexagonal_go/adapters"
	"hexagonal_go/core"
	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	db, err := gorm.Open(sqlite.Open("orders.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&core.Order{})

	orderRepo := adapters.NewGormOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.POST("/order", orderHandler.CreateOrder)

	srv := &http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("listen: %s\n", err)
	}

}
