package blog

import (
	"net/http"
	"strconv"

	goerrors "github.com/go-errors/errors"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	CreateBlog(c echo.Context) error
	UpdateBlog(c echo.Context) error
	GetBlogById(c echo.Context) error
	DeleteBlog(c echo.Context) error
	GetAllBlogs(c echo.Context) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) CreateBlog(c echo.Context) error {
	req := CreateBlogRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(goerrors.New(err)))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(goerrors.New(err)))
	}

	err := h.service.CreateBlog(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(err))
	}

	return c.NoContent(http.StatusOK)
}

func (h *handler) UpdateBlog(c echo.Context) error {
	req := UpdateBlogRequest{}
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(goerrors.New(err)))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(goerrors.New(err)))
	}

	err := h.service.UpdateBlog(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(err))
	}

	return c.NoContent(http.StatusOK)
}

func (h *handler) DeleteBlog(c echo.Context) error {
	request := &DeleteBlogRequest{}
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(goerrors.New(err)))
	}

	if err := c.Validate(request); err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(goerrors.New(err)))
	}

	err := h.service.DeleteBlog(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(err))
	}

	return c.NoContent(http.StatusOK)
}

func (h *handler) GetAllBlogs(c echo.Context) error {

	blogs, err := h.service.GetAllBlogs()
	if err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(err))
	}

	return c.JSON(http.StatusOK, blogs)
}

func (h *handler) GetBlogById(c echo.Context) error {

	blog_id, err := strconv.ParseUint(c.QueryParam("id"), 10, 64)
	resp, err := h.service.GetBlogById(blog_id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, CreateErrorResponse(err))
	}

	return c.JSON(http.StatusOK, resp)
}
