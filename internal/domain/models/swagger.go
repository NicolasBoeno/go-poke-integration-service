package models

type Response struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Operation completed successfully"`
	Data    interface{} `json:"data"`
}
