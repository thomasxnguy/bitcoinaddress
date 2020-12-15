package address

type healthCheckResponse struct {
	Status string `json:"status"`
}

func newHealthCheckResponse() *healthCheckResponse {
	resp := &healthCheckResponse{
		Status: "Alive",
	}
	return resp
}
