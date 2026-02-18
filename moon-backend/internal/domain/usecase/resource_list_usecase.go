package usecase

import (
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type ResourceListRequest struct {
}

type ResourceListResponse struct {
	Resources []ResourceListItem `json:"resources"`
}

type ResourceListItem struct {
	ResourceID          uuid.UUID `json:"resource_id"`
	ResourceName        string    `json:"resource_name"`
	ResourceCode        string    `json:"resource_code"`
	ResourceDescription *string   `json:"resource_description"`
	ResourceFlag        bool      `json:"resource_flag"`
	ResourceType        string    `json:"resource_type"`
	ResourcePath        *string   `json:"resource_path"`
	RequestMethod       *string   `json:"request_method"`
	CreatedAt           string    `json:"created_at"`
	UpdatedAt           string    `json:"updated_at"`
}

func ConvertResourceAggregatesToResourceListResponse(resources []aggregate.Resource) ResourceListResponse {
	response := ResourceListResponse{
		Resources: make([]ResourceListItem, len(resources)),
	}

	for i, resource := range resources {
		response.Resources[i] = ResourceListItem{
			ResourceID:          resource.ResourceID,
			ResourceName:        resource.ResourceName,
			ResourceCode:        resource.ResourceCode,
			ResourceDescription: resource.ResourceDescription,
			ResourceFlag:        bool(resource.ResourceFlag),
			ResourceType:        resource.ResourceType,
			ResourcePath:        resource.ResourcePath,
			RequestMethod:       resource.RequestMethod,
			CreatedAt:           resource.CreatedAt.Format("2006-01-02T15:04:05Z"),
			UpdatedAt:           resource.UpdatedAt.Format("2006-01-02T15:04:05Z"),
		}
	}

	return response
}
