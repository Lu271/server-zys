package entity

import "time"

type MallUser struct {
	ID        int       `gorm:"primaryKey;column:id"`
	NickName  string    `gorm:"column:nick_name"`
	Account   string    `gorm:"index:idx_account;column:account"`
	Password  string    `gorm:"column:password"`
	Icon      string    `gorm:"column:icon"`
	Gender    int       `gorm:"column:gender"`
	Status    int       `gorm:"column:status;default:1"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
