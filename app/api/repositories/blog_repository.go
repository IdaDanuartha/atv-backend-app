package repositories

import (
	"github.com/IdaDanuartha/atv-backend-app/app/config"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
)

type BlogRepository interface {
	FindAll(blog models.Blog, search string, currentPage int, pageSize int, exceptID string) ([]models.Blog, int64, int, error)
	Find(ID string) (models.Blog, error)
	Save(blog models.Blog) (models.Blog, error)
	Update(blog models.Blog) (models.Blog, error)
	Delete(blog models.Blog) (models.Blog, error)
}

type blogRepository struct {
	db config.Database
}

// NewBlogRepository : fetching database
func NewBlogRepository(db config.Database) blogRepository {
	return blogRepository{db}
}

// FindAll -> Method for fetching all Blog from database
func (r blogRepository) FindAll(blog models.Blog, search string, currentPage int, pageSize int, exceptID string) ([]models.Blog, int64, int, error) {
	var blogs []models.Blog
	var totalRows int64 = 0

	queryBuilder := r.db.DB.Order("created_at desc").Model(&models.Blog{})

	// Search parameter
	if search != "" {
		querySearch := "%" + search + "%"
		queryBuilder = queryBuilder.Where(
			r.db.DB.Where("blogs.title LIKE ? ", querySearch).
				Or("blogs.description LIKE ? ", querySearch).
				Or("blogs.content LIKE ? ", querySearch))
	}

	if exceptID != "" {
        queryBuilder.Not("id = ?", exceptID)
    }

	if pageSize > 0 {
		// count the total number of rows
		err := queryBuilder.
			Where(blog).
			Count(&totalRows).Error
		if err != nil {
			return nil, 0, 0, err
		}

		// Apply offset and limit to fetch paginated results
		err = queryBuilder.
			Where(blog).
			Offset((currentPage - 1) * pageSize).
			Limit(pageSize).
			Find(&blogs).Error
		return blogs, totalRows, currentPage, err
	} else {
		err := queryBuilder.
			Where(blog).
			Find(&blogs).
			Count(&totalRows).Error
		return blogs, 0, 0, err
	}
}

// Find -> Method for fetching Blog by id
func (r blogRepository) Find(ID string) (models.Blog, error) {
	var blogs models.Blog

	err := r.db.DB.
		Debug().
		Model(&models.Blog{}).
		Where("id = ?", ID).
		Find(&blogs).Error

	return blogs, err
}

// Save -> Method for saving Blog to database
func (r blogRepository) Save(blog models.Blog) (models.Blog, error) {
	err := r.db.DB.Create(&blog).Error
	if err != nil {
		return blog, err
	}

	return blog, nil
}

// Update -> Method for updating Blog
func (r *blogRepository) Update(blog models.Blog) (models.Blog, error) {
	err := r.db.DB.Save(&blog).Error

	if err != nil {
		return blog, err
	}

	return blog, nil
}

// Delete -> Method for deleting Blog
func (r blogRepository) Delete(blog models.Blog) (models.Blog, error) {
	err := r.db.DB.Delete(&blog).Error

	if err != nil {
		return blog, err
	}

	return blog, nil
}
