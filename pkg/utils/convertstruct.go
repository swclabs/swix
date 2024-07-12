// Package utils provides utils functionality
package utils

// NxN2Nx1 converts NxN map to Nx1 map
func NxN2Nx1(data map[string][]string) map[string]interface{} {
	resp := make(map[string]interface{}, len(data))
	for k, v := range data {
		resp[k] = v[0]
	}
	return resp
}
