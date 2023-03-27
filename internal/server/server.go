package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"test_1/internal/service"
)

type Controller struct {
	service *service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{service: s}
}

func RegisterControllers(group *gin.RouterGroup, service *service.Service) error {

	ctl := NewController(service)

	group.GET("/books", ctl.getBooks)
	group.DELETE("/book", ctl.deleteBookByID)

	return nil
}

func (ctl Controller) getBooks(c *gin.Context) {
	books, err := ctl.service.Repo.GetBooks(c)

	if err != nil {
		_ = c.AbortWithError(http.StatusNotFound, err)

		return
	}

	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Header().Add("Accept", "application/json")

	c.JSON(http.StatusOK, books)
}

type deleteBookRequest struct {
	ID string `json:"id"`
}

func (ctl Controller) deleteBookByID(c *gin.Context) {
	var req deleteBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	err := ctl.service.Repo.DeleteBookByID(c, req.ID)

	if err != nil {
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Header().Add("Accept", "application/json")

	c.Status(http.StatusNoContent)
}
