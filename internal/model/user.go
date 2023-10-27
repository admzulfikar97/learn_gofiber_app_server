package model

type User struct {
	ID        uint   `gorm:"primary_key" sql:"AUTO_INCREMENT" json:"id"`
	UserName  string `gorm:"varchar(30);not null" json:"user_name"`
	Gender    string `gorm:"varchar(10);not null" json:"gender"`
	UpdatedAt string `json:"updated_at"`
	CreatedAt string `json:"created_at"`
	Deleted   bool   `json:"deleted"`
}
