package rbac

import (
	"github.com/pkg/errors"
)

// AddRoles 添加角色
// cRole 子角色,fRole 父角色
func AddRoles(cRole string, fRoles []string) error {

	for _, fRole := range fRoles {
		if ok, _ := RBAC().HasRole(fRole); !ok {
			return errors.New(fRole + " 不存在")
		}

	}
	if result, err := sCasbinInstance.enforcer.AddRolesForUser(cRole, fRoles); err != nil || !result {
		err = errors.Wrap(err, "添加角色失败")
	}
	return nil
}
