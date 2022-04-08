package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nick/api-database-jwt/model"
	"github.com/nick/api-database-jwt/service"
	"gorm.io/gorm"
)

type Bookx struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var db *gorm.DB

type BookController struct {
}

func (e BookController) ListBooksHandler(c *gin.Context) {
	db = service.GetDatabaseConnection()

	var books []model.Book

	if result := db.Find(&books); result.Error != nil {
		return
	}

	c.JSON(http.StatusOK, &books)
}
func listBooksHandler(c *gin.Context) {
	db = service.GetDatabaseConnection()
	var books []model.Book

	if result := db.Find(&books); result.Error != nil {
		return
	}

	c.JSON(http.StatusOK, &books)
}
func (e BookController) CreateBookHandler(c *gin.Context) {
	var book model.Book
	db = service.GetDatabaseConnection()
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if result := db.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &book)

}

func (e BookController) DeleteBookHandler(c *gin.Context) {
	db = service.GetDatabaseConnection()
	id := c.Param("id")

	if result := db.Where("id = ?", id).Delete(&model.Book{}); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
