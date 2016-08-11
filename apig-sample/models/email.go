package models

type Email struct {
    ID      uint   `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
    UserID  uint   `json:"user_id"`
    Address string `json:"address"`
    User    *User  `json:"user`
}
