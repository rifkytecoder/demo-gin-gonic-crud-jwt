package dto

// Used when client POST from /login url
type LoginDTO struct {
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password" form:"password" binding:"required"`
}
