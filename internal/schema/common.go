package schema

type HealthCheckResponse struct {
	Status string `json:"status"`
}

type Error struct {
	Msg string `json:"msg"`
}

type OK struct {
	Msg string `json:"msg"`
}
