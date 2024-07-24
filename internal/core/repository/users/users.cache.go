package users

import (
	"context"
	"swclabs/swipecore/internal/core/domain/entity"
	"swclabs/swipecore/internal/core/domain/model"
)

type cache struct {
	user IUserRepository
}

var _ IUserRepository = (*cache)(nil)

func useCache(repo IUserRepository) IUserRepository {
	return &cache{user: repo}
}

// GetByEmail implements IUserRepository.
func (c *cache) GetByEmail(ctx context.Context, email string) (*entity.Users, error) {
	return c.user.GetByEmail(ctx, email)
}

// GetByPhone implements IUserRepository.
func (c *cache) GetByPhone(ctx context.Context, nPhone string) (*entity.Users, error) {
	return c.user.GetByPhone(ctx, nPhone)
}

// Info implements IUserRepository.
func (c *cache) Info(ctx context.Context, email string) (*model.Users, error) {
	return c.user.Info(ctx, email)
}

// Insert implements IUserRepository.
func (c *cache) Insert(ctx context.Context, usr entity.Users) error {
	return c.user.Insert(ctx, usr)
}

// OAuth2SaveInfo implements IUserRepository.
func (c *cache) OAuth2SaveInfo(ctx context.Context, user entity.Users) error {
	return c.user.OAuth2SaveInfo(ctx, user)
}

// SaveInfo implements IUserRepository.
func (c *cache) SaveInfo(ctx context.Context, user entity.Users) error {
	return c.user.SaveInfo(ctx, user)
}

// UpdateProperties implements IUserRepository.
func (c *cache) UpdateProperties(ctx context.Context, query string, user entity.Users) error {
	return c.user.UpdateProperties(ctx, query, user)
}

// GetByID implements IUserRepository.
func (c *cache) GetByID(ctx context.Context, id int64) (*entity.Users, error) {
	return c.user.GetByID(ctx, id)
}
