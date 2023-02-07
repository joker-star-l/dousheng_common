package common

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenUser struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
