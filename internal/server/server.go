package server

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"test_1/internal/service"
)

type Controller struct {
	service *service.Service
}

func NewController(s *service.Service) *Controller {
	return &Controller{service: s}
}

func RegisterControllers(group *gin.Engine, service *service.Service) error {

	ctl := NewController(service)

	group.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group.GET("/books", ctl.getBooks)
	group.DELETE("/book", ctl.deleteBookByID)

	return nil
}

// getBooks получение списка книг
// @Description получение списка книг.
// @Accept json
// @Produce json
// @Success 200 {object} []domain.Book
// @Failure 400
// @Failure 500
// @Router /v1/books [GET].
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

// deleteBookByID удаление книги по ID
// @Description удаление книги по ID
// @Accept json
// @Produce json
// @Success 200
// @Param search body deleteBookRequest true "Search request"
// @Failure 400
// @Failure 500
// @Router /v1/book [DELETE].
func (ctl Controller) deleteBookByID(c *gin.Context) {
	var req deleteBookRequest

	if err := c.ShouldBind(&req); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	err := ctl.service.Repo.DeleteBookByID(c, req.ID)

	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)

		return
	}

	c.Writer.Header().Add("Content-Type", "application/json")
	c.Writer.Header().Add("Accept", "application/json")

	c.Status(http.StatusNoContent)
}
