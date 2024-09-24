package manager

import (
	"context"
	"mime/multipart"
	"swclabs/swix/internal/core/domain/dtos"
	"swclabs/swix/internal/core/domain/model"
)

// IManager : Module for manager with use-cases.
// Actor: Admin & Customer (Users)
type IManager interface {
	// SignUp registers a new user.
	// ctx is the context to manage the request's lifecycle.
	// req contains the sign-up request details.
	// Returns an error if any issues occur during the sign-up process.
	SignUp(ctx context.Context, req dtos.SignUpRequest) error

	// Login authenticates a user and returns a token.
	// ctx is the context to manage the request's lifecycle.
	// req contains the login request details.
	// Returns a token string and an error if any issues occur during the login process.
	Login(ctx context.Context, req dtos.LoginRequest) (string, error)

	// CheckLoginEmail checks if the email is already registered.
	// ctx is the context to manage the request's lifecycle.
	// email is the email address to check.
	// Returns an error if any issues occur during the check process.
	CheckLoginEmail(ctx context.Context, email string) error

	// UserInfo retrieves user information based on email.
	// ctx is the context to manage the request's lifecycle.
	// email is the email address to retrieve user information for.
	// Returns a pointer to the UserInfo object and an error if any issues occur during the retrieval process.
	UserInfo(ctx context.Context, email string) (*model.Users, error)

	// UpdateUserInfo updates the user information.
	// ctx is the context to manage the request's lifecycle.
	// req contains the updated user information details.
	// Returns an error if any issues occur during the update process.
	UpdateUserInfo(ctx context.Context, req dtos.UserUpdate) error

	// UploadAvatar uploads a user's avatar.
	// email is the email address of the user.
	// fileHeader contains the file header of the avatar to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadAvatar(email string, fileHeader *multipart.FileHeader) error

	// OAuth2SaveUser saves user information from an OAuth2 login.
	// ctx is the context to manage the request's lifecycle.
	// req contains the OAuth2 user information details.
	// Returns an error if any issues occur during the save process.
	OAuth2SaveUser(ctx context.Context, req dtos.OAuth2SaveUser) (int64, error)
}
