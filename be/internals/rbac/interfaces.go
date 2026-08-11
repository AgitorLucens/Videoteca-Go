package rbac

type UserRepository interface {
	GetUsersByRole(roleName string) ([]User, error)
	CreateUser(req CreateUserRequest) (*User, error)
	SetUserPassword(id uint, password string) error
	GetUserById(id uint) (*User, error)
}
