package serializer

import (
	"GoCloud/pkg/dao"
	"GoCloud/pkg/rbac"
)

// User 用户序列化器
type User struct {
	ID             string `json:"id"`
	Email          string `json:"email"`
	UserName       string `json:"nickname"`
	Avatar         string `json:"avatar,omitempty"`
	PreferredTheme string `json:"preferred_theme,omitempty"`
	Anonymous      bool   `json:"anonymous,omitempty"`
	Role           string `json:"role,omitempty"`
}

// BuildUser 序列化用户
func BuildUser(user dao.User) (User, error) {
	u := User{
		ID:       user.UUID,
		Email:    user.Email,
		UserName: user.NickName,
		Avatar:   user.Avatar,
	}
	role, err := rbac.GetRolesForUser(user.UUID)
	if err != nil {
		return User{}, err
	}
	u.Role = role
	return u, nil
}
