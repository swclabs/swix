package config

import (
	"fmt"
	"swclabs/swix/pkg/lib/logger"
)

// Version of the application
const Version = "0.0.1"

// banner is the banner of the package.
const banner = ` 
SWIX - %s                          
`

const version = "v0.0.1"

// Info prints the version information.
func Info() {
	fmt.Printf(banner, logger.Red.Add(version))
}
