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
	id := c.Param("id")
	message := "hi " + id
	c.String(http.StatusOK, message)
}

func (h *Handler) getByEmail(c *gin.Context) {

}
