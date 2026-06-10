package model

import "time"

type User struct {
	ID            int64      `gorm:"primaryKey;autoIncrement"`
	PublicKeyHash string     `gorm:"column:public_key_hash;type:char(64);not null;uniqueIndex"`
	CreatedAt     time.Time  `gorm:"column:created_at;autoCreateTime"`
	LastLoginAt   *time.Time `gorm:"column:last_login_at"`
}
