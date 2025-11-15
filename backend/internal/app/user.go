package app

import "time"

type User struct {
	ID        int        `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time  `gorm:"type:timestamp" json:"updated_at"`
	DeletedAt *time.Time `gorm:"type:timestamp" json:"deleted_at,omitempty"`
	Name      string     `gorm:"type:varchar(100)" json:"name"`
	Position  string     `gorm:"type:varchar(100)" json:"position"`
	Salary    float64    `gorm:"type:numeric" json:"salary"`
}

type CreateUserDTO struct {
	Name     string  `json:"name"`
	Position string  `json:"position"`
	Salary   float64 `json:"salary"`
}
