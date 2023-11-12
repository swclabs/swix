package sentry

import (
	"example/komposervice/internal/config"
	"fmt"

	sentrygin "github.com/getsentry/sentry-go/gin"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

func Init() {
	if config.StageStatus != "dev" {
		if err := sentry.Init(sentry.ClientOptions{
			Dsn:           config.SentryDsn,
			EnableTracing: true,
			// Set TracesSampleRate to 1.0 to capture 100%
			// of transactions for performance monitoring.
			// We recommend adjusting this value in production,
			TracesSampleRate: 1.0,
			AttachStacktrace: true,
		}); err != nil {
			fmt.Printf("Sentry initialization failed: %v", err)
		}
	}
}

func CaptureMessage(ctx *gin.Context, message string) {
	if config.StageStatus != "dev" {
		if hub := sentrygin.GetHubFromContext(ctx); hub != nil {
			hub.WithScope(func(scope *sentry.Scope) {
				// scope.SetExtra("unwantedQuery", "someQueryDataMaybe")
				hub.CaptureMessage(message)
			})
		}
	}
}
