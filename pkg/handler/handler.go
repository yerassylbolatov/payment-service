package handler

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("", h.startingPage)
	router.POST("/list/create/:id/:email/:amount/:currency", h.newTransaction)
	router.POST("/list/cancel/:id", h.cancelByID)
	router.GET("/list/:id", h.getByID)
	router.GET("/list/:email/", h.getByEmail)
	return router
}
