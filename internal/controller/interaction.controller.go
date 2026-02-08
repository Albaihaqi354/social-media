package controller

import (
	"net/http"
	"strconv"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/service"
	"github.com/gin-gonic/gin"
)

type InteractionController struct {
	interactionService *service.InteractionService
}

func NewInteractionController(interactionService *service.InteractionService) *InteractionController {
	return &InteractionController{
		interactionService: interactionService,
	}
}

// LikePost godoc
// @Summary      Like a post
// @Description  Like a post by its ID
// @Tags         interactions
// @Security     BearerAuth
// @Param        id  path  int  true  "Post ID to like"
// @Success      200 {object} dto.Response
// @Router       /posts/{id}/like [post]
func (ctrl *InteractionController) LikePost(c *gin.Context) {
	userId := c.GetInt("user_id")
	postId, _ := strconv.Atoi(c.Param("id"))

	err := ctrl.interactionService.LikePost(c.Request.Context(), postId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Error Liking Post",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Successfully Liked Post",
		Success: true,
	})
}

// UnlikePost godoc
// @Summary      Unlike a post
// @Description  Unlike a post by its ID
// @Tags         interactions
// @Security     BearerAuth
// @Param        id  path  int  true  "Post ID to unlike"
// @Success      200 {object} dto.Response
// @Router       /posts/{id}/like [delete]
func (ctrl *InteractionController) UnlikePost(c *gin.Context) {
	userId := c.GetInt("user_id")
	postId, _ := strconv.Atoi(c.Param("id"))

	err := ctrl.interactionService.UnlikePost(c.Request.Context(), postId, userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Error Unliking Post",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Successfully Unliked Post",
		Success: true,
	})
}

// CreateComment godoc
// @Summary      Add a comment to a post
// @Description  Add a comment to a post by its ID
// @Tags         interactions
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  int               true  "Post ID"
// @Param        req      body  dto.CommentRequest  true  "Comment content"
// @Success      201 {object} dto.Response
// @Router       /posts/{id}/comments [post]
func (ctrl *InteractionController) CreateComment(c *gin.Context) {
	userId := c.GetInt("user_id")
	postId, _ := strconv.Atoi(c.Param("id"))

	var req dto.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	id, err := ctrl.interactionService.CreateComment(c.Request.Context(), postId, userId, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Error Creating Comment",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Msg:     "Comment Created Successfully",
		Success: true,
		Data:    map[string]any{"id": id},
	})
}

// GetComments godoc
// @Summary      Get comments for a post
// @Description  Get all comments for a post by its ID
// @Tags         interactions
// @Produce      json
// @Security     BearerAuth
// @Param        id  path  int  true  "Post ID"
// @Success      200 {object} dto.Response{data=[]dto.CommentResponse}
// @Router       /posts/{id}/comments [get]
func (ctrl *InteractionController) GetComments(c *gin.Context) {
	postId, _ := strconv.Atoi(c.Param("id"))

	comments, err := ctrl.interactionService.GetCommentsByPostId(c.Request.Context(), postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Error Fetching Comments",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Comments Fetched Successfully",
		Success: true,
		Data:    comments,
	})
}
