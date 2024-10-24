package models

type User struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"-"`
	Role      string `json:"role"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

const (
	AdminRole   = "admin"
	ManagerRole = "manager"
	UserRole    = "user"
)
