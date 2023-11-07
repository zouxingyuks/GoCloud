package rbac

import (
	"GoCloud/pkg/dao"
	"GoCloud/pkg/log"
	"GoCloud/pkg/rbac"
	"GoCloud/service/serializer"
)

func AddRole(cRole string, fRoles []string) serializer.Response {
	//todo 会出现添加角色到一半的情况，是不是需要考虑
	entry := log.NewEntry("service.rbac.addRole")
	err := rbac.AddRoles(cRole, fRoles)
	if err != nil {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("添加失败"), serializer.WithErr(err))
	}
	return serializer.NewResponse(entry, 200, serializer.WithMsg("添加成功"))

}

// GetRoles 查看全部角色
func GetRoles() serializer.Response {
	entry := log.NewEntry("service.rbac.readRoles")
	roles := rbac.GetAllRoles()
	return serializer.NewResponse(entry, 200, serializer.WithData(roles))
}

// GetPermissions 查看角色权限
func GetPermissions(role string) serializer.Response {
	entry := log.NewEntry("service.rbac.getPermissions")
	permissions := rbac.GetPermissions(role)
	if len(permissions) == 0 {
		serializer.NewResponse(entry, 400, serializer.WithMsg("角色不存在"))
	}
	return serializer.NewResponse(entry, 200, serializer.WithData(permissions))
}

// GetRolesForUser 查看
func GetRolesForUser(uuid string) serializer.Response {
	entry := log.NewEntry("service.rbac.getRolesForUser")
	// 判断用户是否存在
	if _, err := dao.GetUser(dao.WithUUID(uuid)); err != nil {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("用户不存在"), serializer.WithErr(err))
	}
	// 获取用户角色
	roles, err := rbac.GetRolesForUser(uuid)
	if err != nil {

		return serializer.NewResponse(entry, 400, serializer.WithMsg("获取失败"), serializer.WithErr(err))
	}
	return serializer.NewResponse(entry, 200, serializer.WithData(roles))
}

// AssignRolesToUser 给用户分配角色的接口。
func AssignRolesToUser(UUID, NewRole string) serializer.Response {
	entry := log.NewEntry("service.rbac.assignRolesToUser")
	if _, err := dao.GetUser(dao.WithUUID(UUID)); err != nil {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("用户不存在"), serializer.WithErr(err))
	}
	if ok := rbac.HasRole(NewRole); !ok {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("角色不存在"))
	}
	rbac.ChangeRoleForUser(UUID, NewRole)
	return serializer.NewResponse(entry, 200, serializer.WithMsg("修改成功"))
}

// DeleteRole 删除角色
func DeleteRole(role string) serializer.Response {
	entry := log.NewEntry("service.rbac.deleteRole")
	//1. 检查此角色是否能被删除
	//1.1 检查是否是基础角色
	if rbac.IsBasicRole(role) {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("基础角色不能删除"))
	}
	//1.2 检查是否是管理员角色
	if role == rbac.Admin {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("管理员角色不能删除"))
	}
	//1.3 检查是否存在
	if ok := rbac.HasRole(role); !ok {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("角色不存在"))
	}
	//2. 检查是否有使用或者是依赖于此角色的角色
	users, err := rbac.GetUsersForRole(role)
	if err != nil || len(users) > 1 {
		return serializer.NewResponse(entry, 400, serializer.WithMsg("删除失败,仍有用户使用此角色"), serializer.WithErr(err))
	}
	//3. 删除角色
	err = rbac.DeleteRole(role)
	if err != nil {
		return serializer.NewResponse(entry, 500, serializer.WithMsg("删除失败"), serializer.WithErr(err))
	}
	return serializer.NewResponse(entry, 200)
}
