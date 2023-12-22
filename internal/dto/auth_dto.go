package dto

type RegisterUserRequest struct {
	Name     string `json:"name" binding:""`
	Username string `json:"username" binding:""`
	Email    string `json:"email" binding:""`
	Password string `json:"password" binding:""`
}

type RegisterUserResponse struct{}
