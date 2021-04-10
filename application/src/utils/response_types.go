package utils

type HttpResponse struct {
	Message    string `json:"msg"`
	StatusCode int    `json:"status"`
}
