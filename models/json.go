package models

type JSONResponse struct {
	HttpRes int                    `json:"http_res,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Status  string                 `json:"status,omitempty"`
}

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
