package usecase

import (
	"moon/internal/domain/aggregate"

	"github.com/google/uuid"
)

type ResourceUpdateRequest struct {
	ResourceID          uuid.UUID `json:"resource_id"`
	ResourceName        string    `json:"resource_name"`
	ResourceCode        string    `json:"resource_code"`
	ResourceDescription string    `json:"resource_description"`
	ResourceFlag        string    `json:"resource_flag"`
	ResourceType        string    `json:"resource_type"`
	ResourcePath        string    `json:"resource_path"`
	RequestMethod       string    `json:"request_method"`
}

type ResourceUpdateResponse struct {
	ResourceID   uuid.UUID `json:"resource_id"`
	ResourceName string    `json:"resource_name"`
	ResourceCode string    `json:"resource_code"`
}

func ConvertResourceUpdateRequestToResourceAggregate(req ResourceUpdateRequest) aggregate.Resource {
	return aggregate.Resource{
		ResourceID:          req.ResourceID,
		ResourceName:        req.ResourceName,
		ResourceCode:        req.ResourceCode,
		ResourceDescription: &req.ResourceDescription,
		ResourceFlag:        aggregate.BoolString(req.ResourceFlag == "true"),
		ResourceType:        req.ResourceType,
		ResourcePath:        &req.ResourcePath,
		RequestMethod:       &req.RequestMethod,
	}
}
