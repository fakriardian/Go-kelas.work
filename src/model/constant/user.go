package constant

type ResigesterUserRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
