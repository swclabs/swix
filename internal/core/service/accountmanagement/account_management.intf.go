package accountmanagement

import (
	"context"
	"mime/multipart"
	"swclabs/swipecore/internal/core/domain"
)

type IAccountManagementTask interface {
	DelaySignUp(req *domain.SignUpReq) error
	DelayUpdateUserInfo(req *domain.UserUpdate) error
	DelayOAuth2SaveUser(req *domain.OAuth2SaveUser) error
}

// IAccountManagement : Module for Account Management with use-cases.
// Actor: Admin & Customer (User)
type IAccountManagement interface {
	// SignUp registers a new user.
	// ctx is the context to manage the request's lifecycle.
	// req contains the sign-up request details.
	// Returns an error if any issues occur during the sign-up process.
	SignUp(ctx context.Context, req *domain.SignUpReq) error

	// Login authenticates a user and returns a token.
	// ctx is the context to manage the request's lifecycle.
	// req contains the login request details.
	// Returns a token string and an error if any issues occur during the login process.
	Login(ctx context.Context, req *domain.LoginReq) (string, error)

	// CheckLoginEmail checks if the email is already registered.
	// ctx is the context to manage the request's lifecycle.
	// email is the email address to check.
	// Returns an error if any issues occur during the check process.
	CheckLoginEmail(ctx context.Context, email string) error

	// UserInfo retrieves user information based on email.
	// ctx is the context to manage the request's lifecycle.
	// email is the email address to retrieve user information for.
	// Returns a pointer to the UserInfo object and an error if any issues occur during the retrieval process.
	UserInfo(ctx context.Context, email string) (*domain.UserInfo, error)

	// UpdateUserInfo updates the user information.
	// ctx is the context to manage the request's lifecycle.
	// req contains the updated user information details.
	// Returns an error if any issues occur during the update process.
	UpdateUserInfo(ctx context.Context, req *domain.UserUpdate) error

	// UploadAvatar uploads a user's avatar.
	// email is the email address of the user.
	// fileHeader contains the file header of the avatar to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadAvatar(email string, fileHeader *multipart.FileHeader) error

	// OAuth2SaveUser saves user information from an OAuth2 login.
	// ctx is the context to manage the request's lifecycle.
	// req contains the OAuth2 user information details.
	// Returns an error if any issues occur during the save process.
	OAuth2SaveUser(ctx context.Context, req *domain.OAuth2SaveUser) error

	// UploadAddress uploads a user's address.
	// ctx is the context to manage the request's lifecycle.
	// data contains the address details to be uploaded.
	// Returns an error if any issues occur during the upload process.
	UploadAddress(ctx context.Context, data *domain.Addresses) error
}
