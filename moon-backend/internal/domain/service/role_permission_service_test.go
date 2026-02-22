package service

import (
	"context"
	"errors"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/usecase"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockRolePermissionRepository is a mock implementation of repository.RolePermissionRepository
type MockRolePermissionRepository struct {
	mock.Mock
}

func (m *MockRolePermissionRepository) Create(ctx context.Context, rp aggregate.RolePermission) error {
	args := m.Called(ctx, rp)
	return args.Error(0)
}

func (m *MockRolePermissionRepository) Delete(ctx context.Context, roleID, permissionID uuid.UUID) error {
	args := m.Called(ctx, roleID, permissionID)
	return args.Error(0)
}

func (m *MockRolePermissionRepository) GetByRoleAndPermission(ctx context.Context, roleID, permissionID uuid.UUID) (aggregate.RolePermission, error) {
	args := m.Called(ctx, roleID, permissionID)
	return args.Get(0).(aggregate.RolePermission), args.Error(1)
}

func (m *MockRolePermissionRepository) GetPermissionsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Permission, error) {
	args := m.Called(ctx, roleID)
	return args.Get(0).([]aggregate.Permission), args.Error(1)
}

func (m *MockRolePermissionRepository) GetRolesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Role, error) {
	args := m.Called(ctx, permissionID)
	return args.Get(0).([]aggregate.Role), args.Error(1)
}

// MockRoleRepository is a mock implementation of repository.RoleRepository
type MockRoleRepository struct {
	mock.Mock
}

func (m *MockRoleRepository) Create(ctx context.Context, role aggregate.Role) error {
	args := m.Called(ctx, role)
	return args.Error(0)
}

func (m *MockRoleRepository) GetByID(ctx context.Context, id uuid.UUID) (aggregate.Role, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(aggregate.Role), args.Error(1)
}

func (m *MockRoleRepository) Update(ctx context.Context, role aggregate.Role) error {
	args := m.Called(ctx, role)
	return args.Error(0)
}

func (m *MockRoleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockRoleRepository) List(ctx context.Context) ([]aggregate.Role, error) {
	args := m.Called(ctx)
	return args.Get(0).([]aggregate.Role), args.Error(1)
}

// MockPermissionRepository is a mock implementation of repository.PermissionRepository
type MockPermissionRepository struct {
	mock.Mock
}

func (m *MockPermissionRepository) Create(ctx context.Context, permission aggregate.Permission) error {
	args := m.Called(ctx, permission)
	return args.Error(0)
}

func (m *MockPermissionRepository) GetByID(ctx context.Context, id uuid.UUID) (aggregate.Permission, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(aggregate.Permission), args.Error(1)
}

func (m *MockPermissionRepository) Update(ctx context.Context, permission aggregate.Permission) error {
	args := m.Called(ctx, permission)
	return args.Error(0)
}

func (m *MockPermissionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockPermissionRepository) List(ctx context.Context) ([]aggregate.Permission, error) {
	args := m.Called(ctx)
	return args.Get(0).([]aggregate.Permission), args.Error(1)
}

func TestRolePermissionService_Create(t *testing.T) {
	// Setup
	mockRolePermissionRepo := &MockRolePermissionRepository{}
	mockRoleRepo := &MockRoleRepository{}
	mockPermissionRepo := &MockPermissionRepository{}

	service := NewRolePermissionService(mockRolePermissionRepo, mockRoleRepo, mockPermissionRepo)

	ctx := context.Background()
	roleID := uuid.New()
	permissionID := uuid.New()

	req := usecase.RolePermissionCreateRequest{
		RoleID:       roleID,
		PermissionID: permissionID,
	}

	// Mock responses
	mockRoleRepo.On("GetByID", ctx, roleID).Return(aggregate.Role{ID: roleID}, nil)
	mockPermissionRepo.On("GetByID", ctx, permissionID).Return(aggregate.Permission{ID: permissionID}, nil)
	mockRolePermissionRepo.On("Create", ctx, mock.Anything).Return(nil)

	// Execute
	resp, err := service.Create(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	mockRoleRepo.AssertExpectations(t)
	mockPermissionRepo.AssertExpectations(t)
	mockRolePermissionRepo.AssertExpectations(t)
}

func TestRolePermissionService_Create_RoleNotFound(t *testing.T) {
	// Setup
	mockRolePermissionRepo := &MockRolePermissionRepository{}
	mockRoleRepo := &MockRoleRepository{}
	mockPermissionRepo := &MockPermissionRepository{}

	service := NewRolePermissionService(mockRolePermissionRepo, mockRoleRepo, mockPermissionRepo)

	ctx := context.Background()
	roleID := uuid.New()
	permissionID := uuid.New()

	req := usecase.RolePermissionCreateRequest{
		RoleID:       roleID,
		PermissionID: permissionID,
	}

	// Mock responses
	mockRoleRepo.On("GetByID", ctx, roleID).Return(aggregate.Role{}, errors.New("role not found"))

	// Execute
	resp, err := service.Create(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.False(t, resp.Success)
	mockRoleRepo.AssertExpectations(t)
}

func TestRolePermissionService_Create_PermissionNotFound(t *testing.T) {
	// Setup
	mockRolePermissionRepo := &MockRolePermissionRepository{}
	mockRoleRepo := &MockRoleRepository{}
	mockPermissionRepo := &MockPermissionRepository{}

	service := NewRolePermissionService(mockRolePermissionRepo, mockRoleRepo, mockPermissionRepo)

	ctx := context.Background()
	roleID := uuid.New()
	permissionID := uuid.New()

	req := usecase.RolePermissionCreateRequest{
		RoleID:       roleID,
		PermissionID: permissionID,
	}

	// Mock responses
	mockRoleRepo.On("GetByID", ctx, roleID).Return(aggregate.Role{ID: roleID}, nil)
	mockPermissionRepo.On("GetByID", ctx, permissionID).Return(aggregate.Permission{}, errors.New("permission not found"))

	// Execute
	resp, err := service.Create(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.False(t, resp.Success)
	mockRoleRepo.AssertExpectations(t)
	mockPermissionRepo.AssertExpectations(t)
}

func TestRolePermissionService_Delete(t *testing.T) {
	// Setup
	mockRolePermissionRepo := &MockRolePermissionRepository{}
	mockRoleRepo := &MockRoleRepository{}
	mockPermissionRepo := &MockPermissionRepository{}

	service := NewRolePermissionService(mockRolePermissionRepo, mockRoleRepo, mockPermissionRepo)

	ctx := context.Background()
	roleID := uuid.New()
	permissionID := uuid.New()

	req := usecase.RolePermissionDeleteRequest{
		RoleID:       roleID,
		PermissionID: permissionID,
	}

	// Mock responses
	mockRolePermissionRepo.On("Delete", ctx, roleID, permissionID).Return(nil)

	// Execute
	resp, err := service.Delete(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	mockRolePermissionRepo.AssertExpectations(t)
}

func TestRolePermissionService_GetPermissionsByRoleID(t *testing.T) {
	// Setup
	mockRolePermissionRepo := &MockRolePermissionRepository{}
	mockRoleRepo := &MockRoleRepository{}
	mockPermissionRepo := &MockPermissionRepository{}

	service := NewRolePermissionService(mockRolePermissionRepo, mockRoleRepo, mockPermissionRepo)

	ctx := context.Background()
	roleID := uuid.New()

	expectedPermissions := []aggregate.Permission{
		{ID: uuid.New(), Name: "permission1"},
		{ID: uuid.New(), Name: "permission2"},
	}

	// Mock responses
	mockRolePermissionRepo.On("GetPermissionsByRoleID", ctx, roleID).Return(expectedPermissions, nil)

	// Execute
	permissions, err := service.GetPermissionsByRoleID(ctx, roleID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedPermissions, permissions)
	mockRolePermissionRepo.AssertExpectations(t)
}

func TestRolePermissionService_GetRolesByPermissionID(t *testing.T) {
	// Setup
	mockRolePermissionRepo := &MockRolePermissionRepository{}
	mockRoleRepo := &MockRoleRepository{}
	mockPermissionRepo := &MockPermissionRepository{}

	service := NewRolePermissionService(mockRolePermissionRepo, mockRoleRepo, mockPermissionRepo)

	ctx := context.Background()
	permissionID := uuid.New()

	expectedRoles := []aggregate.Role{
		{ID: uuid.New(), Name: "role1"},
		{ID: uuid.New(), Name: "role2"},
	}

	// Mock responses
	mockRolePermissionRepo.On("GetRolesByPermissionID", ctx, permissionID).Return(expectedRoles, nil)

	// Execute
	roles, err := service.GetRolesByPermissionID(ctx, permissionID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedRoles, roles)
	mockRolePermissionRepo.AssertExpectations(t)
}
