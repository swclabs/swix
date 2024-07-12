package domain

// request -> [name]
// response -> [name]Schema

// HealthCheck schema for response
type HealthCheck struct {
	Status string `json:"status"`
}

// Error schema for response
type Error struct {
	Msg string `json:"msg"`
}

// OK schema for response
type OK struct {
	Msg string `json:"msg"`
}

// Slices schema for response
type Slices[T any] struct {
	Body []T `json:"body"`
}
