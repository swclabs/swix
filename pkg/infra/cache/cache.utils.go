package cache

import (
	"context"
	"encoding/json"
)

// Get retrieves a value from the cache.
func Get[T any](ctx context.Context, cache ICache, key string) (*T, error) {
	raw, err := cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var result T
	if err = json.Unmarshal([]byte(raw), &result); err != nil {
		return nil, err
	}
	return &result, nil
}

// GetSlice retrieves a slice from the cache.
func GetSlice[T any](ctx context.Context, cache ICache, key string) ([]T, error) {
	raw, err := cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	var result []T
	if err = json.Unmarshal([]byte(raw), &result); err != nil {
		return nil, err
	}
	return result, nil
}

// Set stores a value in the cache.
func Set[T any](ctx context.Context, cache ICache, key string, val T) error {
	raw, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return cache.Set(ctx, key, string(raw))
}

// Delete removes a value from the cache.
func Delete(ctx context.Context, cache ICache, key string) error {
	return cache.Del(ctx, key)
}
