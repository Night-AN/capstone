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

// MockPermissionResourceRepository is a mock implementation of repository.PermissionResourceRepository
type MockPermissionResourceRepository struct {
	mock.Mock
}

func (m *MockPermissionResourceRepository) Create(ctx context.Context, pr aggregate.PermissionResource) error {
	args := m.Called(ctx, pr)
	return args.Error(0)
}

func (m *MockPermissionResourceRepository) Delete(ctx context.Context, permissionID, resourceID uuid.UUID) error {
	args := m.Called(ctx, permissionID, resourceID)
	return args.Error(0)
}

func (m *MockPermissionResourceRepository) GetByPermissionAndResource(ctx context.Context, permissionID, resourceID uuid.UUID) (aggregate.PermissionResource, error) {
	args := m.Called(ctx, permissionID, resourceID)
	return args.Get(0).(aggregate.PermissionResource), args.Error(1)
}

func (m *MockPermissionResourceRepository) GetPermissionsByResourceID(ctx context.Context, resourceID uuid.UUID) ([]aggregate.Permission, error) {
	args := m.Called(ctx, resourceID)
	return args.Get(0).([]aggregate.Permission), args.Error(1)
}

func (m *MockPermissionResourceRepository) GetResourcesByPermissionID(ctx context.Context, permissionID uuid.UUID) ([]aggregate.Resource, error) {
	args := m.Called(ctx, permissionID)
	return args.Get(0).([]aggregate.Resource), args.Error(1)
}

// MockResourceRepository is a mock implementation of repository.ResourceRepository
type MockResourceRepository struct {
	mock.Mock
}

func (m *MockResourceRepository) Create(ctx context.Context, resource aggregate.Resource) error {
	args := m.Called(ctx, resource)
	return args.Error(0)
}

func (m *MockResourceRepository) GetByID(ctx context.Context, id uuid.UUID) (aggregate.Resource, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(aggregate.Resource), args.Error(1)
}

func (m *MockResourceRepository) Update(ctx context.Context, resource aggregate.Resource) error {
	args := m.Called(ctx, resource)
	return args.Error(0)
}

func (m *MockResourceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockResourceRepository) List(ctx context.Context) ([]aggregate.Resource, error) {
	args := m.Called(ctx)
	return args.Get(0).([]aggregate.Resource), args.Error(1)
}

func TestPermissionResourceService_Create(t *testing.T) {
	// Setup
	mockPermissionResourceRepo := &MockPermissionResourceRepository{}
	mockPermissionRepo := &MockPermissionRepository{}
	mockResourceRepo := &MockResourceRepository{}

	service := NewPermissionResourceService(mockPermissionResourceRepo, mockPermissionRepo, mockResourceRepo)

	ctx := context.Background()
	permissionID := uuid.New()
	resourceID := uuid.New()

	req := usecase.PermissionResourceCreateRequest{
		PermissionID: permissionID,
		ResourceID:   resourceID,
	}

	// Mock responses
	mockPermissionRepo.On("GetByID", ctx, permissionID).Return(aggregate.Permission{ID: permissionID}, nil)
	mockResourceRepo.On("GetByID", ctx, resourceID).Return(aggregate.Resource{ID: resourceID}, nil)
	mockPermissionResourceRepo.On("Create", ctx, mock.Anything).Return(nil)

	// Execute
	resp, err := service.Create(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	mockPermissionRepo.AssertExpectations(t)
	mockResourceRepo.AssertExpectations(t)
	mockPermissionResourceRepo.AssertExpectations(t)
}

func TestPermissionResourceService_Create_PermissionNotFound(t *testing.T) {
	// Setup
	mockPermissionResourceRepo := &MockPermissionResourceRepository{}
	mockPermissionRepo := &MockPermissionRepository{}
	mockResourceRepo := &MockResourceRepository{}

	service := NewPermissionResourceService(mockPermissionResourceRepo, mockPermissionRepo, mockResourceRepo)

	ctx := context.Background()
	permissionID := uuid.New()
	resourceID := uuid.New()

	req := usecase.PermissionResourceCreateRequest{
		PermissionID: permissionID,
		ResourceID:   resourceID,
	}

	// Mock responses
	mockPermissionRepo.On("GetByID", ctx, permissionID).Return(aggregate.Permission{}, errors.New("permission not found"))

	// Execute
	resp, err := service.Create(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.False(t, resp.Success)
	mockPermissionRepo.AssertExpectations(t)
}

func TestPermissionResourceService_Create_ResourceNotFound(t *testing.T) {
	// Setup
	mockPermissionResourceRepo := &MockPermissionResourceRepository{}
	mockPermissionRepo := &MockPermissionRepository{}
	mockResourceRepo := &MockResourceRepository{}

	service := NewPermissionResourceService(mockPermissionResourceRepo, mockPermissionRepo, mockResourceRepo)

	ctx := context.Background()
	permissionID := uuid.New()
	resourceID := uuid.New()

	req := usecase.PermissionResourceCreateRequest{
		PermissionID: permissionID,
		ResourceID:   resourceID,
	}

	// Mock responses
	mockPermissionRepo.On("GetByID", ctx, permissionID).Return(aggregate.Permission{ID: permissionID}, nil)
	mockResourceRepo.On("GetByID", ctx, resourceID).Return(aggregate.Resource{}, errors.New("resource not found"))

	// Execute
	resp, err := service.Create(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.False(t, resp.Success)
	mockPermissionRepo.AssertExpectations(t)
	mockResourceRepo.AssertExpectations(t)
}

func TestPermissionResourceService_Delete(t *testing.T) {
	// Setup
	mockPermissionResourceRepo := &MockPermissionResourceRepository{}
	mockPermissionRepo := &MockPermissionRepository{}
	mockResourceRepo := &MockResourceRepository{}

	service := NewPermissionResourceService(mockPermissionResourceRepo, mockPermissionRepo, mockResourceRepo)

	ctx := context.Background()
	permissionID := uuid.New()
	resourceID := uuid.New()

	req := usecase.PermissionResourceDeleteRequest{
		PermissionID: permissionID,
		ResourceID:   resourceID,
	}

	// Mock responses
	mockPermissionResourceRepo.On("Delete", ctx, permissionID, resourceID).Return(nil)

	// Execute
	resp, err := service.Delete(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	mockPermissionResourceRepo.AssertExpectations(t)
}

func TestPermissionResourceService_GetPermissionsByResourceID(t *testing.T) {
	// Setup
	mockPermissionResourceRepo := &MockPermissionResourceRepository{}
	mockPermissionRepo := &MockPermissionRepository{}
	mockResourceRepo := &MockResourceRepository{}

	service := NewPermissionResourceService(mockPermissionResourceRepo, mockPermissionRepo, mockResourceRepo)

	ctx := context.Background()
	resourceID := uuid.New()

	expectedPermissions := []aggregate.Permission{
		{ID: uuid.New(), Name: "permission1"},
		{ID: uuid.New(), Name: "permission2"},
	}

	// Mock responses
	mockPermissionResourceRepo.On("GetPermissionsByResourceID", ctx, resourceID).Return(expectedPermissions, nil)

	// Execute
	permissions, err := service.GetPermissionsByResourceID(ctx, resourceID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedPermissions, permissions)
	mockPermissionResourceRepo.AssertExpectations(t)
}

func TestPermissionResourceService_GetResourcesByPermissionID(t *testing.T) {
	// Setup
	mockPermissionResourceRepo := &MockPermissionResourceRepository{}
	mockPermissionRepo := &MockPermissionRepository{}
	mockResourceRepo := &MockResourceRepository{}

	service := NewPermissionResourceService(mockPermissionResourceRepo, mockPermissionRepo, mockResourceRepo)

	ctx := context.Background()
	permissionID := uuid.New()

	expectedResources := []aggregate.Resource{
		{ID: uuid.New(), Name: "resource1"},
		{ID: uuid.New(), Name: "resource2"},
	}

	// Mock responses
	mockPermissionResourceRepo.On("GetResourcesByPermissionID", ctx, permissionID).Return(expectedResources, nil)

	// Execute
	resources, err := service.GetResourcesByPermissionID(ctx, permissionID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedResources, resources)
	mockPermissionResourceRepo.AssertExpectations(t)
}
