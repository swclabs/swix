package accounts

import (
	"context"
	"swclabs/swix/internal/core/domain/entity"
)

// IAccounts implements all methods of Account to access the database.
type IAccounts interface {
	// GetByEmail retrieves an account by email.
	// ctx is the context to manage the request's lifecycle.
	// email is the email address to search for.
	// Returns a pointer to the Account object and an error if any issues occur during the retrieval process.
	GetByEmail(ctx context.Context, email string) (*entity.Account, error)

	// Insert adds a new account to the database.
	// ctx is the context to manage the request's lifecycle.
	// acc is a pointer to the Account object to be added.
	// Returns an error if any issues occur during the insertion process.
	Insert(ctx context.Context, acc entity.Account) (int64, error)

	// SaveInfo saves the account information to the database.
	// ctx is the context to manage the request's lifecycle.
	// acc is a pointer to the Account object to be saved.
	// Returns an error if any issues occur during the saving process.
	SaveInfo(ctx context.Context, acc entity.Account) error
}
