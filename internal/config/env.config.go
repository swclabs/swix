// Package config contains all environment variables
package config

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func init() {
	numberOfWorker := os.Getenv("NUMBER_OF_WORKER")
	Num, err := strconv.Atoi(numberOfWorker)
	if err != nil {
		NumberOfWorker = Num // default
	}
}

var (
	// StageStatus Dev or Prod stage status
	StageStatus = os.Getenv("STAGE_STATUS")

	// Host server host
	Host = os.Getenv("HOST")

	// Port server port
	Port = os.Getenv("PORT")
)

// authentication environment variables
var (
	JwtSecret = os.Getenv("JWT_SECRET_KEY")
	JwtCost   = os.Getenv("JWT_COST")
	Secret    = os.Getenv("SECRET")
)

// database environment variables
var (
	DbHost     = os.Getenv("DB_HOST")
	DbPort     = os.Getenv("DB_PORT")
	DbUser     = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbName     = os.Getenv("DB_NAME")
	DbSSLMode  = os.Getenv("DB_SSL_MODE")
)

// cache environment variables
var (
	RedisHost     = os.Getenv("REDIS_HOST")
	RedisPort     = os.Getenv("REDIS_PORT")
	RedisPassword = os.Getenv("REDIS_PASSWORD")
)

// SentryDsn environment variables
var (
	SentryDsn = os.Getenv("SENTRY_DSN")
)

// email
var (
	Email            = os.Getenv("EMAIL")
	EmailAppPassword = os.Getenv("EMAIL_APP_PASSWORD")
)

// Auth0
var (
	FeHomepage        = os.Getenv("FE_HOMEPAGE")
	Auth0ClientID     = os.Getenv("AUTH0_CLIENT_ID")
	Auth0ClientSecret = os.Getenv("AUTH0_CLIENT_SECRET")
	Auth0CallbackURL  = os.Getenv("AUTH0_CALLBACK_URL")
)

// CloudinaryURL Cloudinary
var (
	CloudinaryURL = os.Getenv("CLOUDINARY_URL")
)

var DeliveryTokenAPI = os.Getenv("DELIVERY_TOKEN_API")

// NumberOfWorker Number of worker
var NumberOfWorker = 10
