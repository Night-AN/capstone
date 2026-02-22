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

// MockOrganizationRoleRepository is a mock implementation of repository.OrganizationRoleRepository
type MockOrganizationRoleRepository struct {
	mock.Mock
}

func (m *MockOrganizationRoleRepository) Create(ctx context.Context, or aggregate.OrganizationRole) error {
	args := m.Called(ctx, or)
	return args.Error(0)
}

func (m *MockOrganizationRoleRepository) Delete(ctx context.Context, organizationID, roleID uuid.UUID) error {
	args := m.Called(ctx, organizationID, roleID)
	return args.Error(0)
}

func (m *MockOrganizationRoleRepository) GetByOrganizationAndRole(ctx context.Context, organizationID, roleID uuid.UUID) (aggregate.OrganizationRole, error) {
	args := m.Called(ctx, organizationID, roleID)
	return args.Get(0).(aggregate.OrganizationRole), args.Error(1)
}

func (m *MockOrganizationRoleRepository) GetRolesByOrganizationID(ctx context.Context, organizationID uuid.UUID) ([]aggregate.Role, error) {
	args := m.Called(ctx, organizationID)
	return args.Get(0).([]aggregate.Role), args.Error(1)
}

func (m *MockOrganizationRoleRepository) GetOrganizationsByRoleID(ctx context.Context, roleID uuid.UUID) ([]aggregate.Organization, error) {
	args := m.Called(ctx, roleID)
	return args.Get(0).([]aggregate.Organization), args.Error(1)
}

// MockOrganizationRepository is a mock implementation of repository.OrganizationRepository
type MockOrganizationRepository struct {
	mock.Mock
}

func (m *MockOrganizationRepository) Create(ctx context.Context, organization aggregate.Organization) error {
	args := m.Called(ctx, organization)
	return args.Error(0)
}

func (m *MockOrganizationRepository) GetByID(ctx context.Context, id uuid.UUID) (aggregate.Organization, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(aggregate.Organization), args.Error(1)
}

func (m *MockOrganizationRepository) Update(ctx context.Context, organization aggregate.Organization) error {
	args := m.Called(ctx, organization)
	return args.Error(0)
}

func (m *MockOrganizationRepository) Delete(ctx context.Context, id uuid.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockOrganizationRepository) List(ctx context.Context) ([]aggregate.Organization, error) {
	args := m.Called(ctx)
	return args.Get(0).([]aggregate.Organization), args.Error(1)
}

func TestOrganizationRoleService_Create(t *testing.T) {
	// Setup
	mockOrganizationRoleRepo := &MockOrganizationRoleRepository{}
	mockOrganizationRepo := &MockOrganizationRepository{}
	mockRoleRepo := &MockRoleRepository{}

	service := NewOrganizationRoleService(mockOrganizationRoleRepo, mockOrganizationRepo, mockRoleRepo)

	ctx := context.Background()
	organizationID := uuid.New()
	roleID := uuid.New()

	req := usecase.OrganizationRoleCreateRequest{
		OrganizationID: organizationID,
		RoleID:         roleID,
	}

	// Mock responses
	mockOrganizationRepo.On("GetByID", ctx, organizationID).Return(aggregate.Organization{ID: organizationID}, nil)
	mockRoleRepo.On("GetByID", ctx, roleID).Return(aggregate.Role{ID: roleID}, nil)
	mockOrganizationRoleRepo.On("Create", ctx, mock.Anything).Return(nil)

	// Execute
	resp, err := service.Create(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	mockOrganizationRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
	mockOrganizationRoleRepo.AssertExpectations(t)
}

func TestOrganizationRoleService_Create_OrganizationNotFound(t *testing.T) {
	// Setup
	mockOrganizationRoleRepo := &MockOrganizationRoleRepository{}
	mockOrganizationRepo := &MockOrganizationRepository{}
	mockRoleRepo := &MockRoleRepository{}

	service := NewOrganizationRoleService(mockOrganizationRoleRepo, mockOrganizationRepo, mockRoleRepo)

	ctx := context.Background()
	organizationID := uuid.New()
	roleID := uuid.New()

	req := usecase.OrganizationRoleCreateRequest{
		OrganizationID: organizationID,
		RoleID:         roleID,
	}

	// Mock responses
	mockOrganizationRepo.On("GetByID", ctx, organizationID).Return(aggregate.Organization{}, errors.New("organization not found"))

	// Execute
	resp, err := service.Create(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.False(t, resp.Success)
	mockOrganizationRepo.AssertExpectations(t)
}

func TestOrganizationRoleService_Create_RoleNotFound(t *testing.T) {
	// Setup
	mockOrganizationRoleRepo := &MockOrganizationRoleRepository{}
	mockOrganizationRepo := &MockOrganizationRepository{}
	mockRoleRepo := &MockRoleRepository{}

	service := NewOrganizationRoleService(mockOrganizationRoleRepo, mockOrganizationRepo, mockRoleRepo)

	ctx := context.Background()
	organizationID := uuid.New()
	roleID := uuid.New()

	req := usecase.OrganizationRoleCreateRequest{
		OrganizationID: organizationID,
		RoleID:         roleID,
	}

	// Mock responses
	mockOrganizationRepo.On("GetByID", ctx, organizationID).Return(aggregate.Organization{ID: organizationID}, nil)
	mockRoleRepo.On("GetByID", ctx, roleID).Return(aggregate.Role{}, errors.New("role not found"))

	// Execute
	resp, err := service.Create(ctx, req)

	// Assert
	assert.Error(t, err)
	assert.False(t, resp.Success)
	mockOrganizationRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
}

func TestOrganizationRoleService_Delete(t *testing.T) {
	// Setup
	mockOrganizationRoleRepo := &MockOrganizationRoleRepository{}
	mockOrganizationRepo := &MockOrganizationRepository{}
	mockRoleRepo := &MockRoleRepository{}

	service := NewOrganizationRoleService(mockOrganizationRoleRepo, mockOrganizationRepo, mockRoleRepo)

	ctx := context.Background()
	organizationID := uuid.New()
	roleID := uuid.New()

	req := usecase.OrganizationRoleDeleteRequest{
		OrganizationID: organizationID,
		RoleID:         roleID,
	}

	// Mock responses
	mockOrganizationRoleRepo.On("Delete", ctx, organizationID, roleID).Return(nil)

	// Execute
	resp, err := service.Delete(ctx, req)

	// Assert
	assert.NoError(t, err)
	assert.True(t, resp.Success)
	mockOrganizationRoleRepo.AssertExpectations(t)
}

func TestOrganizationRoleService_GetRolesByOrganizationID(t *testing.T) {
	// Setup
	mockOrganizationRoleRepo := &MockOrganizationRoleRepository{}
	mockOrganizationRepo := &MockOrganizationRepository{}
	mockRoleRepo := &MockRoleRepository{}

	service := NewOrganizationRoleService(mockOrganizationRoleRepo, mockOrganizationRepo, mockRoleRepo)

	ctx := context.Background()
	organizationID := uuid.New()

	expectedRoles := []aggregate.Role{
		{ID: uuid.New(), Name: "role1"},
		{ID: uuid.New(), Name: "role2"},
	}

	// Mock responses
	mockOrganizationRoleRepo.On("GetRolesByOrganizationID", ctx, organizationID).Return(expectedRoles, nil)

	// Execute
	roles, err := service.GetRolesByOrganizationID(ctx, organizationID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedRoles, roles)
	mockOrganizationRoleRepo.AssertExpectations(t)
}

func TestOrganizationRoleService_GetOrganizationsByRoleID(t *testing.T) {
	// Setup
	mockOrganizationRoleRepo := &MockOrganizationRoleRepository{}
	mockOrganizationRepo := &MockOrganizationRepository{}
	mockRoleRepo := &MockRoleRepository{}

	service := NewOrganizationRoleService(mockOrganizationRoleRepo, mockOrganizationRepo, mockRoleRepo)

	ctx := context.Background()
	roleID := uuid.New()

	expectedOrganizations := []aggregate.Organization{
		{ID: uuid.New(), Name: "org1"},
		{ID: uuid.New(), Name: "org2"},
	}

	// Mock responses
	mockOrganizationRoleRepo.On("GetOrganizationsByRoleID", ctx, roleID).Return(expectedOrganizations, nil)

	// Execute
	organizations, err := service.GetOrganizationsByRoleID(ctx, roleID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expectedOrganizations, organizations)
	mockOrganizationRoleRepo.AssertExpectations(t)
}
