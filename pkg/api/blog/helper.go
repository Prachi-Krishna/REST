package blog

type CreateBlogRequest struct {
	BlogTitle    string
	BlogCategory string
	Content      string
}

type UpdateBlogRequest struct {
	BlogId    uint64
	BlogTitle string
}

type UpdateConfigRequest struct {
	Port    string
	Service string
}

type DeleteBlogRequest struct {
	BlogId    uint64
	BlogTitle string
}

func CreateErrorResponse(err error) map[string]string {
	return map[string]string{
		"Error": err.Error(),
	}
}
