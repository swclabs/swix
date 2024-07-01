package domain

// request -> [name]
// response -> [name]Schema

type HealthCheck struct {
	Status string `json:"status"`
}

type Error struct {
	Msg string `json:"msg"`
}

type OK struct {
	Msg string `json:"msg"`
}

type Slices[T any] struct {
	Body []T `json:"body"`
}
