package formatters

import (
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

func FormatBlog(blog models.Blog) models.Blog {
	blogFormatter := models.Blog{}

	blogFormatter.ID = blog.ID
	blogFormatter.Title = blog.Title
	blogFormatter.Slug = blog.Slug
	blogFormatter.Description = blog.Description
	blogFormatter.Content = blog.Content
	blogFormatter.ImagePath = blog.ImagePath
	blogFormatter.CreatedAt = blog.CreatedAt
	blogFormatter.UpdatedAt = blog.UpdatedAt
	blogFormatter.DeletedAt = blog.DeletedAt

	return blogFormatter
}

func FormatBlogs(blogs []models.Blog) []models.Blog {
	blogsFormatter := []models.Blog{}

	for _, blog := range blogs {
		blog := FormatBlog(blog)
		blogsFormatter = append(blogsFormatter, blog)
	}

	return blogsFormatter
}
