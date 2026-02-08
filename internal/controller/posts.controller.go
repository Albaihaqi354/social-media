package controller

import (
	"net/http"
	"strconv"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/model"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/service"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService *service.PostService
}

func NewPostController(postService *service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

// CreatePost godoc
// @Summary      Create a new post
// @Description  Create a post with text and/or an image
// @Tags         posts
// @Accept       multipart/form-data
// @Produce      json
// @Param        content  formData  string  false  "Post text content"
// @Param        image    formData  file    false  "Post image (max 5MB, jpg/png/gif)"
// @Security     BearerAuth
// @Success      201 {object} dto.Response
// @Failure      400 {object} dto.Response
// @Failure      401 {object} dto.Response
// @Router       /posts [post]
func (ctrl *PostController) CreatePost(c *gin.Context) {
	userId := c.GetInt("user_id")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized",
			Success: false,
			Error:   "User ID not found",
		})
		return
	}

	var req dto.CreatePostRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	file, _ := c.FormFile("image")

	post := model.Post{
		UserId:  userId,
		Content: req.Content,
	}

	id, err := ctrl.postService.CreatePost(c.Request.Context(), post, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Error Creating Post",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Msg:     "Post Created Successfully",
		Success: true,
		Data:    map[string]any{"id": id},
	})
}

// GetFeed godoc
// @Summary      Get post feed
// @Description  Get latest posts from followed users and yourself
// @Tags         posts
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} dto.Response{data=[]dto.PostResponse}
// @Failure      401 {object} dto.Response
// @Router       /feed [get]
func (ctrl *PostController) GetFeed(c *gin.Context) {
	userId := c.GetInt("user_id")
	if userId == 0 {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized",
			Success: false,
			Error:   "User ID not found",
		})
		return
	}

	posts, err := ctrl.postService.GetFeed(c.Request.Context(), userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Error Fetching Feed",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Feed Fetched Successfully",
		Success: true,
		Data:    posts,
	})
}

// GetPostById godoc
// @Summary      Get post by ID
// @Description  Get a single post details by ID
// @Tags         posts
// @Produce      json
// @Param        id   path      int  true  "Post ID"
// @Security     BearerAuth
// @Success      200 {object} dto.Response{data=dto.PostResponse}
// @Router       /posts/{id} [get]
func (ctrl *PostController) GetPostById(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	post, err := ctrl.postService.GetPostById(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, dto.Response{
			Msg:     "Post Not Found",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Post Fetched Successfully",
		Success: true,
		Data:    post,
	})
}
