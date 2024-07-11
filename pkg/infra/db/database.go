package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"

	"swclabs/swipecore/internal/config"

	"go.uber.org/fx"
)

var (
	pgxConnection *pgxpool.Pool = nil
	lock          *sync.Mutex   = &sync.Mutex{}
	writeLock     *sync.Mutex   = &sync.Mutex{}
)

type Database struct {
	pool      *pgxpool.Pool
	lock      *sync.Mutex
	writeLock *sync.Mutex
}

func GetPool() IDatabase {
	return &Database{
		pool:      pgxConnection,
		lock:      lock,
		writeLock: writeLock,
	}
}

var _ IDatabase = (*Database)(nil)

func New(lc fx.Lifecycle) IDatabase {
	var (
		dsn = fmt.Sprintf(
			"postgres://%s:%s@%s:%s/%s",
			config.DbUser, config.DbPassword, config.DbHost, config.DbPort, config.DbName)
		err error = nil
	)
	if pgxConnection == nil {
		lock.Lock()
		defer lock.Unlock()
		if pgxConnection == nil {
			pgxConnection, err = pgxpool.New(context.Background(), dsn)
		}
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if err != nil {
				return err
			}
			fmt.Printf("[SWIPE]-v%s ===============> connect to PostgreSQL\n", config.Version)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			pgxConnection.Close()
			fmt.Printf("[SWIPE]-v%s ===============> closed PostgreSQL connection\n", config.Version)
			return nil
		},
	})
	return &Database{
		pool:      pgxConnection,
		lock:      lock,
		writeLock: writeLock,
	}
}

// Query implements IDatabase.
func (d *Database) Query(ctx context.Context, sql string, args ...interface{}) (Rows, error) {
	return d.pool.Query(ctx, sql, args...)
}

// SafeWrite implements IDatabase.
func (d *Database) SafeWrite(ctx context.Context, sql string, args ...interface{}) error {
	// lock the connection
	d.writeLock.Lock()
	// after function call return, unlock the write lock
	defer d.writeLock.Unlock()
	_, err := d.pool.Exec(ctx, sql, args...)
	return err
}

// SafeWriteReturn implements IDatabase.
func (d *Database) SafeWriteReturn(ctx context.Context, sql string, args ...interface{}) (int64, error) {
	// lock the connection
	d.writeLock.Lock()
	// after function call return, unlock the write lock
	defer d.writeLock.Unlock()
	var id int64
	if err := d.pool.QueryRow(ctx, sql, args...).Scan(&id); err != nil {
		return -1, err
	}
	return id, nil
}
