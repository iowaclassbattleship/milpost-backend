package model

import "time"

type Post struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Grade     string    `gorm:"type:varchar(64)" json:"grade"`
	Name      string    `gorm:"type:varchar(128);not null" json:"name"`
	Company   string    `gorm:"type:varchar(32)" json:"company"`
	Section   string    `gorm:"type:varchar(32)" json:"section"`
	ItemType  uint8     `gorm:"not null" json:"itemType"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
