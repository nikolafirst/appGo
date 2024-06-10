package usergrpc_test

import (
	"appGo/internal/database"
	"appGo/internal/user/usergrpc"
	"appGo/pkg/pb"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHandler_CreateUser(t *testing.T) {
	mockRepository := &mockUsersRepository{}
	timeout := time.Second
	handler := usergrpc.New(mockRepository, timeout)

	ctx := context.Background()
	request := &pb.CreateUserRequest{
		Id:       "123",
		Username: "testuser",
		Password: "password",
	}

	mockRepository.On("Create", ctx, database.CreateUserReq{
		ID:       uuid.MustParse(request.Id),
		Username: request.Username,
		Password: request.Password,
	}).Return(nil)

	response, err := handler.CreateUser(ctx, request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestHandler_GetUser(t *testing.T) {
	mockRepository := &mockUsersRepository{}
	timeout := time.Second
	handler := usergrpc.New(mockRepository, timeout)

	ctx := context.Background()
	request := &pb.GetUserRequest{
		Id: "123",
	}

	userID := uuid.MustParse(request.Id)
	mockRepository.On("FindByID", ctx, userID).Return(database.User{
		ID:        userID,
		Username:  "testuser",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil)

	response, err := handler.GetUser(ctx, request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestHandler_UpdateUser(t *testing.T) {
	mockRepository := &mockUsersRepository{}
	timeout := time.Second
	handler := usergrpc.New(mockRepository, timeout)

	ctx := context.Background()
	request := &pb.UpdateUserRequest{
		Id:       "123",
		Username: "testuser",
		Password: "password",
	}

	mockRepository.On("Update", ctx, database.UpdateUserReq{
		ID:       uuid.MustParse(request.Id),
		Username: request.Username,
		Password: request.Password,
	}).Return(nil)

	response, err := handler.UpdateUser(ctx, request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestHandler_DeleteUser(t *testing.T) {
	mockRepository := &mockUsersRepository{}
	timeout := time.Second
	handler := usergrpc.New(mockRepository, timeout)

	ctx := context.Background()
	request := &pb.DeleteUserRequest{
		Id: "123",
	}

	userID := uuid.MustParse(request.Id)
	mockRepository.On("DeleteByUserID", ctx, userID).Return(nil)

	response, err := handler.DeleteUser(ctx, request)

	assert.NoError(t, err)
	assert.NotNil(t, response)
}

func TestHandler_ListUsers(t *testing.T) {
	mockRepository := &mockUsersRepository{}
	timeout := time.Second
	handler := usergrpc.New(mockRepository, timeout)

	ctx := context.Background()

	users := []database.User{
		{
			ID:        uuid.New(),
			Username:  "user1",
			Password:  "password1",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        uuid.New(),
			Username:  "user2",
			Password:  "password2",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepository.On("FindAll", ctx).Return(users, nil)

	response, err := handler.ListUsers(ctx, &pb.Empty{})

	assert.NoError(t, err)
	assert.NotNil(t, response)
	assert.Equal(t, len(users), len(response.Users))
}

type mockUsersRepository struct {
	mock.Mock
}

func (m *mockUsersRepository) Create(ctx context.Context, req database.CreateUserReq) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *mockUsersRepository) FindByID(ctx context.Context, id uuid.UUID) (database.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *mockUsersRepository) Update(ctx context.Context, req database.UpdateUserReq) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *mockUsersRepository) DeleteByUserID(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *mockUsersRepository) FindAll(ctx context.Context) ([]database.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]database.User), args.Error(1)
}
