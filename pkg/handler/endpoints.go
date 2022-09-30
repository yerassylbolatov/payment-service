package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) startingPage(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!!! I'm alive!")
}

func (h *Handler) newTransaction(c *gin.Context) {
	
}

func (h *Handler) cancelByID(c *gin.Context) {

}

func (h *Handler) getByID(c *gin.Context) {

}

func (h *Handler) getByEmail(c *gin.Context) {

}
