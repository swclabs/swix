// Package init start all deps packages
package init

import (
	_ "swclabs/swix/docs" // start swagger

	_ "github.com/golang-migrate/migrate/v4/database/postgres" // prepare migration
	_ "github.com/golang-migrate/migrate/v4/source/file"       // prepare migration
)
