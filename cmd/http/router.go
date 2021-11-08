package http

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
	router  *gin.Engine
	Handler httpHandler
}

func NewRouter(handler httpHandler) *Router {
	return &Router{
		router:  gin.Default(),
		Handler: handler,
	}
}

func (r *Router) routes() {
	r.router.GET("/articles", r.Handler.GetArticle)
	r.router.POST("/articles", r.Handler.CreateArticle)

}
