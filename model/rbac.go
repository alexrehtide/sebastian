package model

type AccountRole struct {
	ID        uint `db:"id"`
	AccountID uint `db:"account_id"`
	Role      Role `db:"role"`
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

		TOTPGenerate,
		TOTPValidate,
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
	TOTPGenerate,
	TOTPValidate,
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
	TOTPGenerate
	TOTPValidate
)
