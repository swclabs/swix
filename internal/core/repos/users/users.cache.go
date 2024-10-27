package users

import (
	"context"
	"swclabs/swipex/internal/core/domain/entity"
	"swclabs/swipex/internal/core/domain/model"
	"swclabs/swipex/pkg/infra/cache"
)

type _cache struct {
	cache cache.ICache
	user  IUsers
}

var _ IUsers = (*_cache)(nil)

func useCache(cache cache.ICache, repo IUsers) IUsers {
	return &_cache{
		user:  repo,
		cache: cache,
	}
}

// GetByEmail implements IUserRepository.
func (c *_cache) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	return c.user.GetByEmail(ctx, email)
}

// GetByPhone implements IUserRepository.
func (c *_cache) GetByPhone(ctx context.Context, nPhone string) (*entity.User, error) {
	return c.user.GetByPhone(ctx, nPhone)
}

// Info implements IUserRepository.
func (c *_cache) Info(ctx context.Context, email string) (*model.Users, error) {
	return c.user.Info(ctx, email)
}

// Insert implements IUserRepository.
func (c *_cache) Insert(ctx context.Context, usr entity.User) (int64, error) {
	return c.user.Insert(ctx, usr)
}

// OAuth2SaveInfo implements IUserRepository.
func (c *_cache) OAuth2SaveInfo(ctx context.Context, user entity.User) error {
	return c.user.OAuth2SaveInfo(ctx, user)
}

// SaveInfo implements IUserRepository.
func (c *_cache) Save(ctx context.Context, user entity.User) error {
	return c.user.Save(ctx, user)
}

// GetByID implements IUserRepository.
func (c *_cache) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	return c.user.GetByID(ctx, id)
}
