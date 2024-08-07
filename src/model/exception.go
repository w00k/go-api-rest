package model

// Exception message for some errors
type Exception struct {
	Status  int    `json:"status"`
	Code    string `json:"code"`
	Message string `json:"error"`
}
