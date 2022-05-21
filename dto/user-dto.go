package dto

// Used client when PUT update profile
type UserUpdateDTO struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

// Used client when register a new user
// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required,email"`
// 	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required,min:6"`
// }
