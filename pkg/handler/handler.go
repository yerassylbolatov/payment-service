package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yerassylbolatov/payment-service/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("", h.startingPage)
	router.POST("/transaction/create/:id/:email/:amount/:currency", h.newTransaction)
	router.POST("/transaction/cancel/:id", h.cancelByID)
	router.GET("/by-id/:id", h.getByID)
	router.GET("/by-email/:email/", h.getByEmail)
	return router
}
