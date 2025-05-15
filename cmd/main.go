package main

import (
	"coupon-system/internal/handlers"
	"coupon-system/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	repository.InitDB()

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.POST("/coupons", handlers.CreateCouponHandler)
	}

	r.Run(":8080")
}
