package rbaccontroller

type RBACService interface{}

func New(rbacService RBACService) *Controller {
	return &Controller{RBACService: rbacService}
}

type Controller struct {
	RBACService RBACService
}
