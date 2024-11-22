package dao

import (
	"context"
	"github.com/Lu271/server-zys/internal/entity"
)

const (
	TABLE_MALL_USER = "mall_user"
)

// CreateMallUser 新增用户
func CreateMallUser(ctx context.Context, dto *entity.MallUser) error {
	db := GetDbInstance("")
	return db.Table(TABLE_MALL_USER).WithContext(ctx).Create(dto).Error
}

// GetMallUser 根据账号密码查询用户
func GetMallUser(ctx context.Context, account, pwd string) (*entity.MallUser, error) {
	mallUser := &entity.MallUser{}

	db := GetDbInstance("")
	err := db.Table(TABLE_MALL_USER).WithContext(ctx).
		Where("account = ?", account).
		Where("password = ?", pwd).
		Find(mallUser).Error
	return mallUser, err
}
