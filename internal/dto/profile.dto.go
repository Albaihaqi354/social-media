package dto

type ProfileResponse struct {
	Id        int    `json:"id"`
	Email     string `json:"email"`
	Name      string `json:"name"`
	AvatarUrl string `json:"avatar_url"`
	Bio       string `json:"bio"`
}

type UpdateProfileRequest struct {
	Name string `form:"name"`
	Bio  string `form:"bio"`
}
