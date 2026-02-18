package usecase

import (
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type ResourceGetRequest struct {
	ResourceID uuid.UUID `json:"resource_id"`
}

type ResourceGetResponse struct {
	ResourceID          uuid.UUID `json:"resource_id"`
	ResourceName        string    `json:"resource_name"`
	ResourceCode        string    `json:"resource_code"`
	ResourceDescription string    `json:"resource_description"`
	ResourceFlag        string    `json:"resource_flag"`
	ResourceType        string    `json:"resource_type"`
	ResourcePath        string    `json:"resource_path"`
	RequestMethod       string    `json:"request_method"`
}

func ConvertResourceAggregateToResourceGetResponse(resource aggregate.Resource) ResourceGetResponse {
	response := ResourceGetResponse{
		ResourceID:   resource.ResourceID,
		ResourceName: resource.ResourceName,
		ResourceCode: resource.ResourceCode,
		ResourceType: resource.ResourceType,
		ResourceFlag: "false",
	}

	if resource.ResourceDescription != nil {
		response.ResourceDescription = *resource.ResourceDescription
	}

	if bool(resource.ResourceFlag) {
		response.ResourceFlag = "true"
	}

	if resource.ResourcePath != nil {
		response.ResourcePath = *resource.ResourcePath
	}

	if resource.RequestMethod != nil {
		response.RequestMethod = *resource.RequestMethod
	}

	return response
}
