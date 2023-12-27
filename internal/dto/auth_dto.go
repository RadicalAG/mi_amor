package dto

type RegisterUserRequest struct {
	Name     string `json:"name" binding:""`
	Email    string `json:"email" binding:""`
	Password string `json:"password" binding:""`
}

type RegisterUserResponse struct{}

type LoginUserRequest struct {
	Email    string `json:"email" binding:""`
	Password string `json:"password" binding:""`
}

type LoginUserResponse struct {
	Token string `json:"token" binding:""`
}
