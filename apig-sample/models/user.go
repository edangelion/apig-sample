package models

import "time"

type User struct {
    ID        uint       `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
    Name      string     `json:"name"`
    Emails    []Email    `json:"emails"`
    CreatedAt *time.Time `json:"created_at"`
    UpdatedAt *time.Time `json:"updated_at"`
}
