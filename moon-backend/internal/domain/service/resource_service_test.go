package service

import (
	"context"
	"moon/internal/domain/aggregate"
	"moon/internal/domain/usecase"
	"testing"

	"github.com/google/uuid"
)

// MockResourceRepository is a mock implementation of repository.ResourceRepository
type MockResourceRepository struct {
	saveResourceErr     error
	updateResourceErr   error
	deleteResourceErr   error
	findResourceByIDErr error
	listAllResourcesErr error
	mockResource        aggregate.Resource
	mockResources       []aggregate.Resource
}

func (m *MockResourceRepository) SaveResource(ctx *context.Context, resource aggregate.Resource) error {
	m.mockResource = resource
	return m.saveResourceErr
}

func (m *MockResourceRepository) UpdateResource(ctx *context.Context, resource aggregate.Resource) error {
	m.mockResource = resource
	return m.updateResourceErr
}

func (m *MockResourceRepository) DeleteResource(ctx *context.Context, resource_id uuid.UUID) error {
	return m.deleteResourceErr
}

func (m *MockResourceRepository) FindResourceByID(ctx *context.Context, resource_id uuid.UUID) (aggregate.Resource, error) {
	if m.findResourceByIDErr != nil {
		return aggregate.Resource{}, m.findResourceByIDErr
	}
	m.mockResource.ResourceID = resource_id
	return m.mockResource, nil
}

func (m *MockResourceRepository) ListAllResources(ctx *context.Context) ([]aggregate.Resource, error) {
	return m.mockResources, m.listAllResourcesErr
}

func TestResourceService_CreateResource(t *testing.T) {
	// Setup
	mockRepo := &MockResourceRepository{}
	service := NewResourceService(mockRepo)
	ctx := context.Background()

	// Create test request
	req := usecase.ResourceCreateRequest{
		ResourceName: "Test Resource",
		ResourceCode: "test-resource",
		ResourceType: "api",
	}

	// Execute
	result, err := service.CreateResource(&ctx, req)

	// Assert
	if err.Code != "" {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.ResourceName != req.ResourceName {
		t.Errorf("Expected resource name %s, got %s", req.ResourceName, result.ResourceName)
	}

	if result.ResourceCode != req.ResourceCode {
		t.Errorf("Expected resource code %s, got %s", req.ResourceCode, result.ResourceCode)
	}

	if result.ResourceID == uuid.Nil {
		t.Errorf("Expected resource ID to be set, got nil")
	}
}

func TestResourceService_GetResourceByID(t *testing.T) {
	// Setup
	resourceID := uuid.New()
	mockResource := aggregate.Resource{
		ResourceID:   resourceID,
		ResourceName: "Test Resource",
		ResourceCode: "test-resource",
		ResourceType: "api",
	}

	mockRepo := &MockResourceRepository{
		mockResource: mockResource,
	}

	service := NewResourceService(mockRepo)
	ctx := context.Background()

	// Create test request
	req := usecase.ResourceGetRequest{
		ResourceID: resourceID,
	}

	// Execute
	result, err := service.GetResourceByID(&ctx, req)

	// Assert
	if err.Code != "" {
		t.Errorf("Expected no error, got %v", err)
	}

	if result.ResourceName != mockResource.ResourceName {
		t.Errorf("Expected resource name %s, got %s", mockResource.ResourceName, result.ResourceName)
	}

	if result.ResourceCode != mockResource.ResourceCode {
		t.Errorf("Expected resource code %s, got %s", mockResource.ResourceCode, result.ResourceCode)
	}

	if result.ResourceID != resourceID {
		t.Errorf("Expected resource ID %v, got %v", resourceID, result.ResourceID)
	}
}

func TestResourceService_ListAllResources(t *testing.T) {
	// Setup
	mockResources := []aggregate.Resource{
		{
			ResourceID:   uuid.New(),
			ResourceName: "Resource 1",
			ResourceCode: "resource-1",
			ResourceType: "api",
		},
		{
			ResourceID:   uuid.New(),
			ResourceName: "Resource 2",
			ResourceCode: "resource-2",
			ResourceType: "menu",
		},
	}

	mockRepo := &MockResourceRepository{
		mockResources: mockResources,
	}

	service := NewResourceService(mockRepo)
	ctx := context.Background()

	// Create test request
	req := usecase.ResourceListRequest{}

	// Execute
	result, err := service.ListAllResources(&ctx, req)

	// Assert
	if err.Code != "" {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(result.Resources) != len(mockResources) {
		t.Errorf("Expected %d resources, got %d", len(mockResources), len(result.Resources))
	}

	for i, expected := range mockResources {
		if i >= len(result.Resources) {
			t.Errorf("Expected resource at index %d, but got only %d resources", i, len(result.Resources))
			break
		}

		actual := result.Resources[i]
		if actual.ResourceName != expected.ResourceName {
			t.Errorf("Expected resource name %s at index %d, got %s", expected.ResourceName, i, actual.ResourceName)
		}

		if actual.ResourceCode != expected.ResourceCode {
			t.Errorf("Expected resource code %s at index %d, got %s", expected.ResourceCode, i, actual.ResourceCode)
		}
	}
}

func TestResourceService_DeleteResource(t *testing.T) {
	// Setup
	resourceID := uuid.New()
	mockRepo := &MockResourceRepository{}
	service := NewResourceService(mockRepo)
	ctx := context.Background()

	// Create test request
	req := usecase.ResourceDeleteRequest{
		ResourceID: resourceID,
	}

	// Execute
	result, err := service.DeleteResource(&ctx, req)

	// Assert
	if err.Code != "" {
		t.Errorf("Expected no error, got %v", err)
	}

	if !result.Success {
		t.Errorf("Expected success to be true, got false")
	}
}
