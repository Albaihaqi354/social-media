package dto

type Response struct {
	Msg     string `json:"msg"`
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Error   string `json:"error,omitempty"`
}
