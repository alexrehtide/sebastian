package accountcontroller_test

import (
	"context"
	"testing"

	accountcontroller "github.com/alexrehtide/sebastian/internal/account-controller"
	testutils "github.com/alexrehtide/sebastian/internal/controller-testutils"
	"github.com/alexrehtide/sebastian/model"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
)

func TestController(t *testing.T) {
	app := initApp()

	tcases := []testutils.TestCase{
		{
			Req: testutils.Req{Method: "POST", Route: "/create", Body: `{"email":"test@test.test","password":"test1234"}`},
			Res: testutils.Res{Code: 200, Body: "1"},
		},
	}

	for _, tcase := range tcases {
		testutils.TestJsonT(t, app, tcase)
	}
}

func initApp() testutils.Testable {
	app := fiber.New()

	accountService := new(MockedAccountService)
	accountService.On("Create", "test@test.test", "test1234").Return(1, nil)
	ctrl := accountcontroller.New(accountService)

	app.Post("/create", ctrl.Create)
	return app
}

type MockedAccountService struct {
	mock.Mock
}

func (s *MockedAccountService) Count(ctx context.Context, ops model.ReadAccountOptions) (int, error) {
	args := s.Called(ops.ID, ops.Email)
	return args.Int(0), args.Error(1)
}

func (s *MockedAccountService) Create(ctx context.Context, ops model.CreateAccountOptions) (uint, error) {
	args := s.Called(ops.Email, ops.Password)
	return uint(args.Int(0)), args.Error(1)
}

func (s *MockedAccountService) Delete(ctx context.Context, id uint) error {
	args := s.Called(id)
	return args.Error(0)
}

func (s *MockedAccountService) Read(ctx context.Context, ops model.ReadAccountOptions, pgOps model.PaginationOptions) ([]model.Account, error) {
	args := s.Called(ops.ID, ops.Email, pgOps.Limit, pgOps.Offset)
	return args.Get(0).([]model.Account), args.Error(1)
}

func (s *MockedAccountService) ReadByID(ctx context.Context, id uint) (model.Account, error) {
	args := s.Called(id)
	return args.Get(0).(model.Account), args.Error(1)
}

func (s *MockedAccountService) Update(ctx context.Context, id uint, ops model.UpdateAccountOptions) error {
	args := s.Called(id, ops.Email)
	return args.Error(0)
}
