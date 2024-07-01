package users

import (
	"context"
	"swclabs/swipecore/internal/core/domain"
)

type cache struct {
	user IUserRepository
}

var _ IUserRepository = (*cache)(nil)

func useCache(repo IUserRepository) IUserRepository {
	return &cache{user: repo}
}

// GetByEmail implements IUserRepository.
func (c *cache) GetByEmail(ctx context.Context, email string) (*domain.Users, error) {
	return c.user.GetByEmail(ctx, email)
}

// GetByPhone implements IUserRepository.
func (c *cache) GetByPhone(ctx context.Context, nPhone string) (*domain.Users, error) {
	return c.user.GetByPhone(ctx, nPhone)
}

// Info implements IUserRepository.
func (c *cache) Info(ctx context.Context, email string) (*domain.UserSchema, error) {
	return c.user.Info(ctx, email)
}

// Insert implements IUserRepository.
func (c *cache) Insert(ctx context.Context, usr domain.Users) error {
	return c.user.Insert(ctx, usr)
}

// OAuth2SaveInfo implements IUserRepository.
func (c *cache) OAuth2SaveInfo(ctx context.Context, user domain.Users) error {
	return c.user.OAuth2SaveInfo(ctx, user)
}

// SaveInfo implements IUserRepository.
func (c *cache) SaveInfo(ctx context.Context, user domain.Users) error {
	return c.user.SaveInfo(ctx, user)
}

// TransactionSaveOAuth2 implements IUserRepository.
func (c *cache) TransactionSaveOAuth2(ctx context.Context, data domain.Users) error {
	return c.user.TransactionSaveOAuth2(ctx, data)
}

// TransactionSignUp implements IUserRepository.
func (c *cache) TransactionSignUp(ctx context.Context, user domain.Users, password string) error {
	return c.user.TransactionSignUp(ctx, user, password)
}

// UpdateProperties implements IUserRepository.
func (c *cache) UpdateProperties(ctx context.Context, query string, user domain.Users) error {
	return c.user.UpdateProperties(ctx, query, user)
}
