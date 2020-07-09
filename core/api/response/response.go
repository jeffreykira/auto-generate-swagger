package response

//HTTPError common error response
type HTTPError struct {
	Code    int    `json:"code" example:"401"`
	Message string `json:"message" example:"Unauthorized"`
}

//SigninSuccess ...
type SigninSuccess struct {
	Token string `json:"token"`
}

//SystemStatus ...
type SystemStatus struct {
	Alive bool `json:"isAlive"`
}
