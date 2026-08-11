package rbac

//RBAC tables

type User struct {
	ID             uint   `gorm:"primaryKey"`
	Name           string
	Email          string
	UserName 	   string `gorm:"column:username"`
	Password       string
	ProfilePicture []byte `gorm:"type:bytea"`
	Roles          []Role `gorm:"many2many:user_roles;foreignKey:id;References:role_id;user_roles;joinForeignKey:user_id;joinReferences:role_id"`
}

type UserRole struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}

type Role struct {
	RoleID   uint 	  `gorm:"primaryKey;column:role_id"`
	RoleName string	  `gorm:"column:role_name"`
	Accesses []Access `gorm:"many2many:role_access;foreignKey:role_id;References:access_id;role_access;joinForeignKey:role_id;joinReferences:access_id"`
}

type RoleAccess struct {
	UserID uint `gorm:"primaryKey"`
	RoleID uint `gorm:"primaryKey"`
}

type Access struct {
	AccessID   uint `gorm:"primaryKey"`
	AccessName string
}

// RBAC request
type CreateUserRequest struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	UserName        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required,min=1"`
	PasswordConfirm string `json:"passwordconfirm" binding:"required,eqfield=Password"`
	Role            string `json:"role" binding:"required"`
}

type LoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}