package rbac

import (
	"be/util/crypt"
	"errors"

	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) GetUserById(id uint) (*User, error) {
	var u User
	err := r.DB.Preload("Roles.Accesses").First(&u, id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) GetUserByEmail(email string) (*User, error) {
	var u User
	err := r.DB.Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) GetUserByUserName(username string) (*User, error) {
	var u User
	err := r.DB.Preload("Roles").
		Preload("Roles.Accesses").
		Where("username = ?", username).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *Repository) CreateUser(req CreateUserRequest) (*User, error) {
	ph, err := crypt.PasswordHasher(req.Password)
	if err != nil {
		return nil, err
	}
	user := &User{
		Name:     req.Name,
		Email:    req.Email,
		UserName: req.UserName,
		Password: ph,
	}
	err = r.DB.Create(user).Error
	if err != nil {
		return nil, err
	}

	role := Role{}
	if err := r.DB.Where("role_name = ?", req.Role).First(&role).Error; err != nil {
		return nil, errors.New("role not found")
	}

	userRole := UserRole{
		UserID: user.ID,
		RoleID: role.RoleID,
	}

	if err := r.DB.Create(&userRole).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (r *Repository) GetUsersByRole(roleName string) ([]User, error) {
	var users []User
	err := r.DB.
		Joins("JOIN user_roles ON user_roles.user_id = users.id").
		Joins("JOIN roles ON roles.role_id = user_roles.role_id").
		Where("roles.role_name = ?", roleName).
		Preload("Roles").
		Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.DB.Preload("Roles").Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *Repository) DeleteUser(id uint) error {
	r.DB.Where("user_id = ?", id).Delete(&UserRole{})
	return r.DB.Delete(&User{}, id).Error
}

func (r *Repository) UpdateUser(id uint, name string, userName string, email string) error {
	return r.DB.Model(&User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"name":     name,
		"username": userName,
		"email":    email,
	}).Error
}

func (r *Repository) UpdateUserRole(userID uint, roleID uint) error {
	r.DB.Where("user_id = ?", userID).Delete(&UserRole{})
	return r.DB.Create(&UserRole{UserID: userID, RoleID: roleID}).Error
}

func (r *Repository) UpdateUserPassword(id uint, currentPassword string, newPassword string) error {
	var u User
	if err := r.DB.First(&u, id).Error; err != nil {
		return err
	}
	ok, _ := crypt.CheckPassword(currentPassword, u.Password)
	if !ok {
		return errors.New("current password is incorrect")
	}
	ph, err := crypt.PasswordHasher(newPassword)
	if err != nil {
		return err
	}
	return r.DB.Model(&u).Update("password", ph).Error
}

func (r *Repository) SetNewPassword(id uint, newPassword string) error {
	ph, err := crypt.PasswordHasher(newPassword)
	if err != nil {
		return err
	}
	return r.DB.Model(&User{}).Where("id = ?", id).Update("password", ph).Error
}

func (r *Repository) SetUserPassword(id uint, password string) error {
	return r.SetNewPassword(id, password)
}

func (r *Repository) UpdateUserProfilePicture(id uint, photo []byte) error {
	return r.DB.Model(&User{}).Where("id = ?", id).Update("profile_picture", photo).Error
}

func (r *Repository) GetRoleByName(name string) (*Role, error) {
	var role Role
	if err := r.DB.Where("role_name = ?", name).First(&role).Error; err != nil {
		return nil, err
	}
	return &role, nil
}
