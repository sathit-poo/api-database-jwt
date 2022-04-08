package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nick/api-database-jwt/controller"
	"github.com/nick/api-database-jwt/model"
	"github.com/nick/api-database-jwt/service"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	service.CreateConnect()
	migrate()

	r := gin.New()
	login := controller.LoginController{}
	r.POST("/login", login.LoginHandler)

	book := controller.BookController{}

	protected := r.Group("/", authorizationMiddleware)
	protected.GET("/books", book.ListBooksHandler)
	r.POST("/books", book.CreateBookHandler)
	r.DELETE("/books/:id", book.DeleteBookHandler)

	r.Run()

}

func migrate() {
	db = service.GetDatabaseConnection()

	db.AutoMigrate(&model.Book{})
	db.AutoMigrate(&model.Employee{})
}

func authorizationMiddleware(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
func validateToken(token string) error {
	if token != "ACCESS_TOKEN" {
		return fmt.Errorf("token provided was invalid")
	}

	return nil
}

var emp = model.Employee{
	Id:        11,
	FirstName: "John",
}
var location = model.Location{
	Lat:  1000.11,
	Long: 13000.22,
}
