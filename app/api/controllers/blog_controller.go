package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/IdaDanuartha/atv-backend-app/app/api/formatters"
	"github.com/IdaDanuartha/atv-backend-app/app/api/inputs"
	"github.com/IdaDanuartha/atv-backend-app/app/api/services"
	"github.com/IdaDanuartha/atv-backend-app/app/models"
	"github.com/IdaDanuartha/atv-backend-app/app/utils"
	"github.com/gin-gonic/gin"
)

// BlogController -> BlogController
type BlogController struct {
	service services.BlogService
}

// NewBlogController : NewBlogController
func NewBlogController(service services.BlogService) *BlogController {
	return &BlogController{service}
}

// GetBlogs : GetBlogs controller
func (h *BlogController) GetBlogs(ctx *gin.Context) {
	var blogs models.Blog

	search := ctx.Query("search")
	currentPage, err := strconv.Atoi(ctx.Query("current_page"))
	if err != nil {
		currentPage = 1
	}

	pageSize, err := strconv.Atoi(ctx.Query("page_size"))
	if err != nil {
		pageSize = 0
	}

	getBlogs, total, _, err := h.service.FindAll(blogs, search, currentPage, pageSize)

	if err != nil {
		response := utils.APIResponse("Failed to find blog", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	if pageSize > 0 {
		response := utils.APIResponseWithPagination("Blogs result set", http.StatusOK, "success", total, currentPage, pageSize, formatters.FormatBlogs(getBlogs))
		ctx.JSON(http.StatusOK, response)
	} else {
		response := utils.APIResponse("Blogs result set", http.StatusOK, "success", formatters.FormatBlogs(getBlogs))
		ctx.JSON(http.StatusOK, response)
	}

}

// GetBlog : get blog by id
func (h *BlogController) GetBlog(c *gin.Context) {
	var input inputs.GetBlogDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of blog", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	blog, err := h.service.Find(input)
	if err != nil {
		response := utils.APIResponse("Failed to get detail of blog", http.StatusBadRequest, "error", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Blog detail", http.StatusOK, "success", formatters.FormatBlog(blog))
	c.JSON(http.StatusOK, response)

}

func (h *BlogController) UploadImage(ctx *gin.Context) {
	var inputID inputs.GetBlogDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to get blog id", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := ctx.FormFile("image")
	if err != nil {
		data := gin.H{"message": err.Error()}
		response := utils.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	blog, err := h.service.Find(inputID)
	if err != nil {
		data := gin.H{"message": err.Error()}
		response := utils.APIResponse("Failed to get blog id", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	ID := blog.ID

	// Check if ImagePath is not nil before proceeding
	if blog.ImagePath != nil {
		// Check if the old avatar image exists for the user
		_, err := os.Stat(*blog.ImagePath)
		if err == nil {
			// If the old avatar image exists, delete it
			err := os.Remove(*blog.ImagePath)
			if err != nil {
				data := gin.H{
					"is_uploaded": false,
					"message":     err.Error(),
				}
				response := utils.APIResponse("Failed to delete old avatar image", http.StatusBadRequest, "error", data)
				ctx.JSON(http.StatusBadRequest, response)
				return
			}
		} else if !os.IsNotExist(err) {
			// Handle other possible errors from os.Stat
			data := gin.H{
				"is_uploaded": false,
				"message":     err.Error(),
			}
			response := utils.APIResponse("Error checking old avatar image", http.StatusInternalServerError, "error", data)
			ctx.JSON(http.StatusInternalServerError, response)
			return
		}
	}

	path := fmt.Sprintf("uploads/blogs/%s-%s", ID, file.Filename)

	err = ctx.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := utils.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.service.SaveImage(ID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := utils.APIResponse("Failed to upload image", http.StatusBadRequest, "error", data)

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := utils.APIResponse("Image successfuly uploaded", http.StatusOK, "success", data)

	ctx.JSON(http.StatusOK, response)
}

// AddBlog : AddBlog controller
func (h *BlogController) AddBlog(ctx *gin.Context) {
	var input inputs.BlogInput
	customizer := g.Validator(inputs.BlogInput{})

	// Check if request body is empty or has no content type
	if ctx.Request.Body == nil || ctx.Request.ContentLength == 0 || ctx.GetHeader("Content-Type") == "" {
		errorMessage := gin.H{"errors": "No fields sent"}
		response := utils.APIResponse("No fields sent", http.StatusBadRequest, "error", errorMessage)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to store blog", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBlog, err := h.service.Save(input)
	if err != nil {
		response := utils.APIResponse("Failed to store blog", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to store blog", http.StatusCreated, "success", formatters.FormatBlog(newBlog))
	ctx.JSON(http.StatusCreated, response)
}

// UpdateBlog : get update by id
func (h *BlogController) UpdateBlog(ctx *gin.Context) {
	var inputID inputs.GetBlogDetailInput
	customizer := g.Validator(inputs.BlogInput{})

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to update blog", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData inputs.BlogInput

	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		// errors := utils.FormatValidationError(err)
		errorMessage := gin.H{"errors": customizer.DecryptErrors(err)}

		response := utils.APIResponse("Failed to update blog", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	updatedBlog, err := h.service.Update(inputID, inputData)
	if err != nil {
		response := utils.APIResponse("Failed to update blog", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to update blog", http.StatusOK, "success", formatters.FormatBlog(updatedBlog))
	ctx.JSON(http.StatusOK, response)
}

// DeleteBlog : Deletes blog
func (h *BlogController) DeleteBlog(ctx *gin.Context) {
	var inputID inputs.GetBlogDetailInput

	err := ctx.ShouldBindUri(&inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete blog", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	deletedBlog, err := h.service.Delete(inputID)
	if err != nil {
		response := utils.APIResponse("Failed to delete blog", http.StatusBadRequest, "error", err.Error())
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.APIResponse("Success to delete blog", http.StatusOK, "success", formatters.FormatBlog(deletedBlog))
	ctx.JSON(http.StatusOK, response)
}
