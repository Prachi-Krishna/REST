package blog

import (
	"fmt"
	"rest/pkg/models"

	goerrors "github.com/go-errors/errors"
	"gorm.io/gorm"
)

type Repository interface {
	CreateBlog(blog *models.Blog) error
	UpdateBlog(id uint64, title string) error
	DeleteBlog(id uint64, title string) error
	GetBlogById(id uint64) (blog *models.Blog, err error)
	GetAllBlogs() (blogs []models.Blog, err error)
	GetConfigByService(service string) (config *models.Config, err error)
	UpdateConfig(port, db string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetConfigByService(service string) (config *models.Config, err error) {
	err = r.db.Find(&config, "service=?", service).Error
	if err != nil {
		return nil, goerrors.New(err)
	}

	return
}

func (r *repository) UpdateConfig(port, service string) error {
	err := r.db.Model(&models.Config{}).Where("service", service).Update("port", port).Error
	if err != nil {
		return goerrors.New(err)
	}

	return nil
}

func (r *repository) CreateBlog(blog *models.Blog) error {
	err := r.db.Create(blog).Error
	if err != nil {
		return goerrors.New(err)
	}
	fmt.Println("Blog Added!!")

	return nil
}

func (r *repository) UpdateBlog(id uint64, title string) error {
	err := r.db.Model(&models.Blog{}).Where("id", id).Update("title", title).Error
	if err != nil {
		return goerrors.New(err)
	}

	return nil
}

func (r *repository) DeleteBlog(id uint64, title string) error {
	err := r.db.Model(&models.Blog{}).Where("id = ? and title = ?", id, title).Delete(&models.Blog{}).Error
	if err != nil {
		return goerrors.New(err)
	}

	return nil
}

func (r *repository) GetBlogById(id uint64) (blog *models.Blog, err error) {
	err = r.db.Find(&blog, "id=?", id).Error
	if err != nil {
		return nil, goerrors.New(err)
	}

	return
}

func (r *repository) GetAllBlogs() (blogs []models.Blog, err error) {
	qry := `select * from blogs`
	err = r.db.Raw(qry).Scan(&blogs).Error
	if err != nil {
		return nil, goerrors.New(err)
	}

	return blogs, nil
}
