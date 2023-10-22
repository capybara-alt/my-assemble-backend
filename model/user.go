package model

import (
	"time"
)

type User struct {
	ID        string `gorm:"primaryKey;index:user_idx"`
	Name      string
	Icon      string
	CreatedAt time.Time
}
