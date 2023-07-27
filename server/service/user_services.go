package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/wqh66886/vue-gin-admin/server/server/model"
)

type UserService struct {
	UserRepository model.UserRepository
}

type USConfig struct {
	UserRepository model.UserRepository
}

func NewUserService(c *USConfig) model.UserService {
	return &UserService{
		UserRepository: c.UserRepository,
	}
}

// 绑定 UserService 并且实现 model.UserService 接口的 Get 方法
func (s *UserService) Get(ctx context.Context, uid uuid.UUID) (*model.User, error) {
	u, err := s.UserRepository.FindByID(ctx, uid)
	return u, err
}
