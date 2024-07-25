package inputs

type GetBlogDetailInput struct {
	ID string `uri:"id" binding:"required"`
}

type BlogInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Content     string `json:"content" binding:"required"`
}
