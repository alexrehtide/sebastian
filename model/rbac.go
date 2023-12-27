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
	AccountID uint
	Role      Role
}

var RolePermission = map[Role][]Permission{
	Guest: {
		AuthAuthenticate,
	},
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

		AuthAuthorize,
		AuthLogout,
		AuthRefresh,

		RBACAddAccountRole,
		RBACReadAccountRoles,
		RBACRemoveAccountRole,
	},
}

type Role string

const (
	Guest Role = "guest"
	User  Role = "user"
	Admin Role = "admin"
)

type Permission int

const (
	AccountCreate Permission = iota
	AccountDelete
	AccountRead
	AccountUpdate
	AuthAuthenticate
	AuthAuthorize
	AuthLogout
	AuthRefresh
	RBACAddAccountRole
	RBACReadAccountRoles
	RBACRemoveAccountRole
)
