package controller

import (
	"net/http"

	"github.com/Albaihaqi354/FinalPhase3.git/internal/dto"
	"github.com/Albaihaqi354/FinalPhase3.git/internal/service"
	"github.com/gin-gonic/gin"
)

type ProfileController struct {
	profileService *service.ProfileService
}

func NewProfileController(profileService *service.ProfileService) *ProfileController {
	return &ProfileController{
		profileService: profileService,
	}
}

// GetProfile godoc
// @Summary      Get user profile
// @Description  Get the profile of the currently logged-in user
// @Tags         profile
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} dto.Response{data=[]dto.ProfileResponse}
// @Failure      401 {object} dto.Response
// @Failure      500 {object} dto.Response
// @Router       /profile [get]
func (ctrl *ProfileController) GetProfile(c *gin.Context) {
	accountId := c.GetInt("user_id")
	if accountId == 0 {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized",
			Success: false,
			Error:   "User Id Not Found",
		})
		return
	}

	profile, err := ctrl.profileService.GetProfile(c.Request.Context(), accountId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Msg:     "Internal Server Error",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Get Profile Succes",
		Success: true,
		Data:    []any{profile},
	})
}

// UpdateProfile godoc
// @Summary      Update user profile
// @Description  Update name, bio, and avatar of the currently logged-in user
// @Tags         profile
// @Accept       multipart/form-data
// @Produce      json
// @Param        name    formData  string  false  "New name"
// @Param        bio     formData  string  false  "New bio"
// @Param        avatar  formData  file    false  "New avatar image (max 2MB, jpg/png/gif)"
// @Security     BearerAuth
// @Success      200 {object} dto.Response{data=[]dto.ProfileResponse}
// @Failure      400 {object} dto.Response
// @Failure      401 {object} dto.Response
// @Router       /profile [patch]
func (ctrl *ProfileController) UpdateProfile(c *gin.Context) {
	accountId := c.GetInt("user_id")
	if accountId == 0 {
		c.JSON(http.StatusUnauthorized, dto.Response{
			Msg:     "Unauthorized",
			Success: false,
			Error:   "User Id not Found",
		})
		return
	}

	var req dto.UpdateProfileRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	file, _ := c.FormFile("avatar")

	profile, err := ctrl.profileService.UpdateProfile(c.Request.Context(), accountId, req, file)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Msg:     "Bad Request",
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Msg:     "Profile Update Success",
		Success: true,
		Data:    []any{profile},
	})
}
