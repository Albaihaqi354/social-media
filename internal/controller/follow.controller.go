package controller

import (
	"net/http"
	"strconv"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/service"
	"github.com/gin-gonic/gin"
)

type FollowController struct {
	followService *service.FollowService
}

func NewFollowController(followService *service.FollowService) *FollowController {
	return &FollowController{
		followService: followService,
	}
}

// FollowUser godoc
// @Summary      Follow a user
// @Description  Follow another user by their ID
// @Tags         follows
// @Security     BearerAuth
// @Param        following_id  path  int  true  "User ID to follow"
// @Success      200 {object} dto.Response
// @Failure      400 {object} dto.Response
// @Router       /follows/{following_id} [post]
func (ctrl *FollowController) FollowUser(c *gin.Context) {
	followerId := c.GetInt("user_id")
	followingId, _ := strconv.Atoi(c.Param("following_id"))

	if followerId == followingId {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Cannot follow yourself",
			Success: false,
		})
		return
	}

	err := ctrl.followService.FollowUser(c.Request.Context(), followerId, followingId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Error Following User",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Successfully Followed User",
		Success: true,
	})
}

// UnfollowUser godoc
// @Summary      Unfollow a user
// @Description  Unfollow another user by their ID
// @Tags         follows
// @Security     BearerAuth
// @Param        following_id  path  int  true  "User ID to unfollow"
// @Success      200 {object} dto.Response
// @Router       /follows/{following_id} [delete]
func (ctrl *FollowController) UnfollowUser(c *gin.Context) {
	followerId := c.GetInt("user_id")
	followingId, _ := strconv.Atoi(c.Param("following_id"))

	err := ctrl.followService.UnfollowUser(c.Request.Context(), followerId, followingId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Error Unfollowing User",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Successfully Unfollowed User",
		Success: true,
	})
}
