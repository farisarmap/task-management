package login_test

import (
	"context"
	"task-management-be/src/domain"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindOne(ctx context.Context, colName string, value string) (domain.User, error) {
	args := m.Called()
	return args.Get(0).(domain.User), args.Error(1)
}

func (m *MockRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	args := m.Called()
	return args.Get(0).(domain.User), nil
}

func (m *MockRepository) Update(ctx context.Context, id string, user domain.User) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRepository) FindById(ctx context.Context, id string) (domain.User, error) {
	args := m.Called()
	return args.Get(0).(domain.User), nil
}

func (m *MockRepository) Delete(ctx context.Context, id string) error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockRepository) List(ctx context.Context) ([]domain.User, error) {
	args := m.Called()
	return args.Get(0).([]domain.User), nil
}

func (m *MockRepository) Exist(ctx context.Context, colName string, value string) error {
	args := m.Called()
	return args.Error(0)
}
