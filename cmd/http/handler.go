package http

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/bariasabda/kumparan/domain/entity"
	"github.com/bariasabda/kumparan/domain/service"
	"github.com/gin-gonic/gin"
)

type httpHandler struct {
	svc service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) *httpHandler {
	return &httpHandler{
		svc: service,
	}
}

func (h *httpHandler) GetArticle(c *gin.Context) {
	search := strings.Join(c.Request.URL.Query()["search"], "")
	author := strings.Join(c.Request.URL.Query()["author"], "")
	article, err := h.svc.GetArticle(search, author)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, article)
}

type RequestCreateArticle struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (h *httpHandler) CreateArticle(c *gin.Context) {
	var req RequestCreateArticle
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&req)
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	if req.Author == "" || req.Title == "" || req.Body == "" {
		c.JSON(http.StatusBadRequest, "All input is required")
		return
	}
	article, err := h.svc.CreateArticle(entity.Article{
		Author: req.Author,
		Title:  req.Title,
		Body:   req.Body,
	})
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success create",
		"data":    article,
	})
}
