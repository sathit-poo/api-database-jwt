package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nick/api-database-jwt/controller"
	"github.com/nick/api-database-jwt/model"
	"github.com/nick/api-database-jwt/service"
	"github.com/unrolled/secure"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {

	service.CreateConnect()
	migrate()

	r := gin.New()
	//r.Use(LoadTls())

	login := controller.LoginController{}
	r.POST("/login", login.LoginHandler)

	book := controller.BookController{}

	protected := r.Group("/", authorizationMiddleware)
	protected.GET("/books", book.ListBooksHandler)
	r.POST("/books", book.CreateBookHandler)
	r.DELETE("/books/:id", book.DeleteBookHandler)

	//r.Run()
	r.Run(":3000")
	//r.RunTLS(":8080", "cer/example.com.pem", "cer/example.com-key.pem")

}
func LoadTls() gin.HandlerFunc {
	return func(c *gin.Context) {
		middleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     "localhost:8000",
		})
		err := middleware.Process(c.Writer, c.Request)
		if err != nil {
			//If an error occurs, do not continue.
			fmt.Println(err)
			return
		}
		//Continue processing
		c.Next()
	}
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
