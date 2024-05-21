package domain

import (
	"context"

	"gorm.io/gorm"
)

// IAccountRepository implements all methods of Account to access the database.
type IAccountRepository interface {
	// Use sets the transaction connection.
	// tx is the transaction connection using gorm.DB.
	// Returns an instance of IAccountRepository with the transaction set.
	Use(tx *gorm.DB) IAccountRepository

	// GetByEmail retrieves an account by email.
	// ctx is the context to manage the request's lifecycle.
	// email is the email address to search for.
	// Returns a pointer to the Account object and an error if any issues occur during the retrieval process.
	GetByEmail(ctx context.Context, email string) (*Account, error)

	// Insert adds a new account to the database.
	// ctx is the context to manage the request's lifecycle.
	// acc is a pointer to the Account object to be added.
	// Returns an error if any issues occur during the insertion process.
	Insert(ctx context.Context, acc *Account) error

	// SaveInfo saves the account information to the database.
	// ctx is the context to manage the request's lifecycle.
	// acc is a pointer to the Account object to be saved.
	// Returns an error if any issues occur during the saving process.
	SaveInfo(ctx context.Context, acc *Account) error
}
