package service

import (
	"context"
	"moon/internal/domain/errors"
	"moon/internal/domain/repository"
	"moon/internal/domain/usecase"
)

type ResourceService interface {
	CreateResource(ctx *context.Context, req usecase.ResourceCreateRequest) (usecase.ResourceCreateResponse, errors.DomainError)
	UpdateResource(ctx *context.Context, req usecase.ResourceUpdateRequest) (usecase.ResourceUpdateResponse, errors.DomainError)
	DeleteResource(ctx *context.Context, req usecase.ResourceDeleteRequest) (usecase.ResourceDeleteResponse, errors.DomainError)
	GetResourceByID(ctx *context.Context, req usecase.ResourceGetRequest) (usecase.ResourceGetResponse, errors.DomainError)
	ListAllResources(ctx *context.Context, req usecase.ResourceListRequest) (usecase.ResourceListResponse, errors.DomainError)
	MoveResource(ctx *context.Context, req usecase.ResourceMoveRequest) (usecase.ResourceMoveResponse, errors.DomainError)
}

func NewResourceService(resource_repo repository.ResourceRepository) ResourceService {
	return &resourceService{resource_repo}
}

type resourceService struct {
	ResourceRepository repository.ResourceRepository
}

func (rs *resourceService) CreateResource(ctx *context.Context, req usecase.ResourceCreateRequest) (usecase.ResourceCreateResponse, errors.DomainError) {
	resource := usecase.ConvertResourceCreateRequestToResourceAggregate(req)
	err := rs.ResourceRepository.SaveResource(ctx, resource)
	if err != nil {
		return usecase.ResourceCreateResponse{}, errors.NewDomainWithError("401", "Create Resource Err", err)
	}
	return usecase.ResourceCreateResponse{
		ResourceID:   resource.ResourceID,
		ResourceName: resource.ResourceName,
		ResourceCode: resource.ResourceCode,
	}, errors.DomainError{}
}

func (rs *resourceService) UpdateResource(ctx *context.Context, req usecase.ResourceUpdateRequest) (usecase.ResourceUpdateResponse, errors.DomainError) {
	// Check if the resource is sensitive
	existingResource, err := rs.ResourceRepository.FindResourceByID(ctx, req.ResourceID)
	if err != nil {
		return usecase.ResourceUpdateResponse{}, errors.NewDomainWithError("401", "Get Resource Err", err)
	}

	// If resource is sensitive, reject the update
	if bool(existingResource.ResourceFlag) {
		return usecase.ResourceUpdateResponse{}, errors.NewDomainError("403", "Sensitive resource cannot be updated")
	}

	resource := usecase.ConvertResourceUpdateRequestToResourceAggregate(req)
	err = rs.ResourceRepository.UpdateResource(ctx, resource)
	if err != nil {
		return usecase.ResourceUpdateResponse{}, errors.NewDomainWithError("401", "Update Resource Err", err)
	}
	return usecase.ResourceUpdateResponse{
		ResourceID:   resource.ResourceID,
		ResourceName: resource.ResourceName,
		ResourceCode: resource.ResourceCode,
	}, errors.DomainError{}
}

func (rs *resourceService) DeleteResource(ctx *context.Context, req usecase.ResourceDeleteRequest) (usecase.ResourceDeleteResponse, errors.DomainError) {
	// Check if the resource is sensitive
	existingResource, err := rs.ResourceRepository.FindResourceByID(ctx, req.ResourceID)
	if err != nil {
		return usecase.ResourceDeleteResponse{Success: false}, errors.NewDomainWithError("401", "Get Resource Err", err)
	}

	// If resource is sensitive, reject the deletion
	if bool(existingResource.ResourceFlag) {
		return usecase.ResourceDeleteResponse{Success: false}, errors.NewDomainError("403", "Sensitive resource cannot be deleted")
	}

	err = rs.ResourceRepository.DeleteResource(ctx, req.ResourceID)
	if err != nil {
		return usecase.ResourceDeleteResponse{Success: false}, errors.NewDomainWithError("401", "Delete Resource Err", err)
	}
	return usecase.ResourceDeleteResponse{Success: true}, errors.DomainError{}
}

func (rs *resourceService) GetResourceByID(ctx *context.Context, req usecase.ResourceGetRequest) (usecase.ResourceGetResponse, errors.DomainError) {
	resource, err := rs.ResourceRepository.FindResourceByID(ctx, req.ResourceID)
	if err != nil {
		return usecase.ResourceGetResponse{}, errors.NewDomainWithError("401", "Get Resource Err", err)
	}
	return usecase.ConvertResourceAggregateToResourceGetResponse(resource), errors.DomainError{}
}

func (rs *resourceService) ListAllResources(ctx *context.Context, req usecase.ResourceListRequest) (usecase.ResourceListResponse, errors.DomainError) {
	resources, err := rs.ResourceRepository.ListAllResources(ctx)
	if err != nil {
		return usecase.ResourceListResponse{}, errors.NewDomainWithError("401", "List Resources Err", err)
	}
	return usecase.ConvertResourceAggregatesToResourceListResponse(resources), errors.DomainError{}
}

func (rs *resourceService) MoveResource(ctx *context.Context, req usecase.ResourceMoveRequest) (usecase.ResourceMoveResponse, errors.DomainError) {
	// Check if the resource is sensitive
	existingResource, err := rs.ResourceRepository.FindResourceByID(ctx, req.ResourceID)
	if err != nil {
		return usecase.ResourceMoveResponse{Success: false}, errors.NewDomainWithError("401", "Get Resource Err", err)
	}

	// If resource is sensitive, reject the move
	if bool(existingResource.ResourceFlag) {
		return usecase.ResourceMoveResponse{Success: false}, errors.NewDomainError("403", "Sensitive resource cannot be moved")
	}

	// Execute the move operation
	err = rs.ResourceRepository.MoveResource(ctx, req.ResourceID, req.NewParentResourceID, req.NewOrganizationID)
	if err != nil {
		return usecase.ResourceMoveResponse{Success: false}, errors.NewDomainWithError("401", "Move Resource Err", err)
	}

	return usecase.ResourceMoveResponse{
		Success:             true,
		ResourceID:          req.ResourceID,
		NewParentResourceID: req.NewParentResourceID,
		NewOrganizationID:   req.NewOrganizationID,
	}, errors.DomainError{}
}
