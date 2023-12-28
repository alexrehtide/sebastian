package model

type AccountRole struct {
	ID        uint
	AccountID uint
	Role      Role
}

type ReadAccountRoleOptions struct {
	ID        uint
	AccountID uint
	Role      Role
}

type CreateAccountRoleOptions struct {
	AccountID uint `validate:"required"`
	Role      Role `validate:"required"`
}

var RolePermission = map[Role][]Permission{
	User: {
		AuthAuthorize,
		AuthLogout,
		AuthRefresh,
	},
	Admin: {
		AccountCreate,
		AccountDelete,
		AccountRead,
		AccountUpdate,

		RBACAddAccountRole,
		RBACReadAccountRoles,
		RBACRemoveAccountRole,
	},
}

var Roles = []Role{
	User,
	Admin,
}

type Role string

const (
	User  Role = "user"
	Admin Role = "admin"
)

var Permissions = []Permission{
	AccountCreate,
	AccountDelete,
	AccountRead,
	AccountUpdate,
	AuthAuthorize,
	AuthLogout,
	AuthRefresh,
	RBACAddAccountRole,
	RBACReadAccountRoles,
	RBACRemoveAccountRole,
}

type Permission int

const (
	AccountCreate Permission = iota
	AccountDelete
	AccountRead
	AccountUpdate
	AuthAuthorize
	AuthLogout
	AuthRefresh
	RBACAddAccountRole
	RBACReadAccountRoles
	RBACRemoveAccountRole
)
