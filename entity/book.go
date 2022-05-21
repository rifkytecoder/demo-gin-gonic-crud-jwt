package entity

// Represents book table in database
type Book struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:varchar(255)" json:"description"`
	UserID      uint   `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
