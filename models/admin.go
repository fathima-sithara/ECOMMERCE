package models

type Admin struct {
	ID       uint   `gorm:"primaryKey;autoincrement" json:"id"`
	Email    string `gorm:"unique;type:varchar(100)" json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=10,max=100"`
	Phone    string `json:"phone" validate:"required,min=10"`
}
