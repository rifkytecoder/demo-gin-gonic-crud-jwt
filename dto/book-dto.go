package dto

// Used client when PUT updating a book
type BookUpdateDTO struct {
	ID          uint   `json:"id" binding:"required"`
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint   `json:"user_id,omitempty" form:"user_id,omitempty"`
}

// Used client when create a new book
type BookCreateDTO struct {
	Title       string `json:"title" form:"title" binding:"required"`
	Description string `json:"description" form:"description" binding:"required"`
	UserID      uint   `json:"user_id,omitempty" form:"user_id,omitempty"`
}
