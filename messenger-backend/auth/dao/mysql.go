package dao

import (
	"context"
	"messenger-backend/share/model"

	"gorm.io/gorm"
)

type MySQL struct {
	mysql *gorm.DB
}

func NewMySQL(mysql *gorm.DB) *MySQL {
	return &MySQL{
		mysql: mysql,
	}
}

func (m *MySQL) GetUser(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	if err := m.mysql.WithContext(ctx).Model(&model.User{Email: email}).Take(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (m *MySQL) CreateUser(ctx context.Context, email string, password, salt []byte) (*model.User, error) {
	user := &model.User{
		Email:    email,
		Password: password,
		Salt:     salt,
	}
	result := m.mysql.WithContext(ctx).Create(&user)
	if err := result.Error; err != nil {
		return nil, err
	}
	return user, nil
}
