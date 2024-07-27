// Package dtos (Data transfer object): is a class that encapsulates
// data for transfer between client-server or between services
// in microservices. The purpose of creating DTOs is to reduce
// the amount of unnecessary information that needs to be
// transferred and to enhance security.
package dtos

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
