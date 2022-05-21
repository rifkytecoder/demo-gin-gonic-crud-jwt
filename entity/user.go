package entity

// Represents user table in database
type User struct {
	ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string  `gorm:"type:varchar(255)" json:"name"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string  `gorm:"->;<-;not null" json:"password"`
	Token    string  `gorm:"-" json:"token,omitempty"`
	Books    *[]Book `json:"books,omitempty"`
}
