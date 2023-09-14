package server

import (
	"rest/pkg/api/blog"

	"gorm.io/gorm"
)

func bindBlog(dbHandler *gorm.DB) blog.Handler {
	repository := blog.NewRepository(dbHandler)
	service := blog.NewService(repository)
	return blog.NewHandler(service)
}
