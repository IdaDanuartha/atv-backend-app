package services

import (
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/repositories"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/app/utils"
)

type BlogService interface {
	FindAll(model models.Blog, search string, currentPage int, pageSize int) ([]models.Blog, int64, int, error)
	Find(input inputs.GetBlogDetailInput) (models.Blog, error)
	SaveImage(ID string, fileLocation string) (models.Blog, error)
	Save(input inputs.BlogInput) (models.Blog, error)
	Update(inputID inputs.GetBlogDetailInput, input inputs.BlogInput) (models.Blog, error)
	Delete(inputID inputs.GetBlogDetailInput) (models.Blog, error)
}

// BlogService BlogService struct
type blogService struct {
	repository repositories.BlogRepository
}

// NewBlogService : returns the BlogService struct instance
func NewBlogService(repository repositories.BlogRepository) blogService {
	return blogService{repository}
}

// FindAll -> calls Blog repo find all method
func (s blogService) FindAll(model models.Blog, search string, currentPage int, pageSize int) ([]models.Blog, int64, int, error) {
	blogs, total, currentPage, err := s.repository.FindAll(model, search, currentPage, pageSize)
	if err != nil {
		return blogs, total, currentPage, err
	}

	return blogs, total, currentPage, nil
}

// Find -> calls Blog repo find method
func (s blogService) Find(input inputs.GetBlogDetailInput) (models.Blog, error) {
	blog, err := s.repository.Find(input.ID)

	if err != nil {
		return blog, err
	}

	return blog, nil
}

func (s blogService) SaveImage(ID string, fileLocation string) (models.Blog, error) {
	blog, err := s.repository.Find(ID)
	if err != nil {
		return blog, err
	}

	blog.ImagePath = &fileLocation

	updatedBlog, err := s.repository.Update(blog)
	if err != nil {
		return updatedBlog, err
	}

	return updatedBlog, nil
}

// Save -> calls Blog repository save method
func (s blogService) Save(input inputs.BlogInput) (models.Blog, error) {
	blog := models.Blog{}

	blog.Title = input.Title
	blog.Slug = utils.CreateSlug(input.Title)
	blog.Description = input.Description
	blog.Content = input.Content

	newBlog, err := s.repository.Save(blog)
	if err != nil {
		return newBlog, err
	}

	return newBlog, nil
}

// Update -> calls Blog repo update method
func (s blogService) Update(inputID inputs.GetBlogDetailInput, input inputs.BlogInput) (models.Blog, error) {
	blog, err := s.repository.Find(inputID.ID)
	if err != nil {
		return blog, err
	}

	blog.Title = input.Title
	blog.Slug = utils.CreateSlug(input.Title)
	blog.Description = input.Description
	blog.Content = input.Content

	updatedBlog, err := s.repository.Update(blog)
	if err != nil {
		return updatedBlog, err
	}

	return updatedBlog, nil
}

// Delete -> calls Blog repo delete method
func (s blogService) Delete(inputID inputs.GetBlogDetailInput) (models.Blog, error) {
	blog, err := s.repository.Find(inputID.ID)
	if err != nil {
		return blog, err
	}

	deletedBlog, err := s.repository.Delete(blog)
	if err != nil {
		return deletedBlog, err
	}

	return deletedBlog, nil
}
