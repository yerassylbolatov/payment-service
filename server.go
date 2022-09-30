package payment_service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler *gin.Engine) error {
	s.httpServer = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return s.httpServer.ListenAndServe()
}
