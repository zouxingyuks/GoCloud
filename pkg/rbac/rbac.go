package rbac

import (
	"sync"
)

// 第一，在实现接口的时候需要使用结构体进行封装，这样可以避免在实现接口的时候，需要实现接口中的所有方法，而是只需要实现自己需要的方法即可。但是这带来的一个问题就是可能出现外部调用的时候暴露了结构体的一些信息，这样就破坏了封装性。

type enforcer interface {
	// 修改
	addRoles(cRole string, fRoles []string) error
	changeRoleForUser(uuid string, role string) (bool, error)

	// 查询
	enforce(role, obj, act string) (bool, error, []string)
	getAllRoles() []string
	getRolesForUser(uuid string) (string, error)
	getUsersForRole(role string) ([]string, error)
	hasRole(role string) bool
	getRoleBasicPermissions(role string) []string

	// 删除
	deleteRole(role string) error
	clearPolicy()
}

var instance = new(struct {
	enforcer
	sync.Once
})

func RBAC() enforcer {
	instance.Do(
		func() {
			instance.enforcer = newCasbin()
		})
	return instance
}

// AddRoles 添加角色
// cRole 子角色,fRole 父角色
func AddRoles(cRole string, fRoles []string) error {
	return RBAC().addRoles(cRole, fRoles)
}

// Enforce 检测角色是否有权限
func Enforce(role, obj, act string) (bool, error, []string) {
	return RBAC().enforce(role, obj, act)
}

// HasRole 判断角色是否存在
func HasRole(role string) bool {
	return RBAC().hasRole(role)
}

// GetAllRoles 获取全部角色
func GetAllRoles() []string {
	return RBAC().getAllRoles()
}

// GetRolesForUser 获取用户的角色
func GetRolesForUser(user string) (string, error) {
	return RBAC().getRolesForUser(user)
}

// GetUsersForRole 获取角色的用户
func GetUsersForRole(role string) ([]string, error) {
	return RBAC().getUsersForRole(role)
}

// GetRoleBasicPermissions 获取角色的基本权限
func GetRoleBasicPermissions(role string) []string {
	return RBAC().getRoleBasicPermissions(role)
}

// ChangeRoleForUser 修改用户的角色
func ChangeRoleForUser(user, role string) (bool, error) {
	return RBAC().changeRoleForUser(user, role)
}

// DeleteRole 删除角色
func DeleteRole(role string) error {
	return RBAC().deleteRole(role)
}
