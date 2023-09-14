package server

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func Setup(dbHandler *gorm.DB) {
	fmt.Println("Welcome to the server")

	e := echo.New()

	e.Validator = &CustomValidator{validator: validator.New()}

	//handler
	blogHandler := bindBlog(dbHandler)

	//endpoints
	e.POST("/create", blogHandler.CreateBlog)
	e.PATCH("/update", blogHandler.UpdateBlog)
	e.DELETE("/delete-blog", blogHandler.DeleteBlog)
	e.GET("/blog", blogHandler.GetBlogById)
	e.GET("/blogs", blogHandler.GetAllBlogs)

	e.Start(viper.GetString("server.port"))
}
