package utils

func NxN2Nx1(data map[string][]string) map[string]string {
	resp := make(map[string]string, len(data))
	for k, v := range data {
		resp[k] = v[0]
	}
	return resp
}
