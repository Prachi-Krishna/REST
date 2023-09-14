package blog

import (
	"rest/pkg/models"
	"time"
)

type Service interface {
	CreateBlog(req *CreateBlogRequest) error
	UpdateBlog(req *UpdateBlogRequest) error
	DeleteBlog(req *DeleteBlogRequest) error
	GetBlogById(id uint64) (*models.Blog, error)
	GetAllBlogs() ([]models.Blog, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) CreateBlog(req *CreateBlogRequest) error {
	blog := &models.Blog{
		Category:  req.BlogCategory,
		Title:     req.BlogTitle,
		Content:   req.Content,
		CreatedAt: time.Now(),
	}

	err := s.repository.CreateBlog(blog)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) UpdateBlog(req *UpdateBlogRequest) error {
	err := s.repository.UpdateBlog(req.BlogId, req.BlogTitle)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteBlog(req *DeleteBlogRequest) error {
	err := s.repository.DeleteBlog(req.BlogId, req.BlogTitle)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetBlogById(id uint64) (*models.Blog, error) {
	blog, err := s.repository.GetBlogById(id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (s *service) GetAllBlogs() ([]models.Blog, error) {
	return s.repository.GetAllBlogs()
}
