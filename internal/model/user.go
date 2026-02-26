package model

import "time"

type User struct {
	ID          int       `json:"id"`
	FullName    string    `json:"full_name"`
	Phone       string    `json:"phone"`
	NationalID  string    `json:"national_id"`
	IDImagePath string    `json:"id_image_path"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
}
