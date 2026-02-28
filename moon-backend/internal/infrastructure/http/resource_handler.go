package http

import (
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ResourceHandler struct {
	resourceService service.ResourceService
}

func NewResourceHandler(service service.ResourceService) *ResourceHandler {
	return &ResourceHandler{resourceService: service}
}

func (h *ResourceHandler) CreateResource(c *gin.Context) {
	var req usecase.ResourceCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.resourceService.CreateResource(&context, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    domainErr.Code,
			"message": domainErr.Message + domainErr.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    "201",
		"message": "success",
		"data":    resp,
	})
}

func (h *ResourceHandler) GetResource(c *gin.Context) {
	resourceIDStr := c.Query("resource_id")
	resourceID, err := uuid.Parse(resourceIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid resource_id: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	req := usecase.ResourceGetRequest{ResourceID: resourceID}
	resp, domainErr := h.resourceService.GetResourceByID(&context, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    domainErr.Code,
			"message": domainErr.Message + domainErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *ResourceHandler) UpdateResource(c *gin.Context) {
	var req usecase.ResourceUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.resourceService.UpdateResource(&context, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    domainErr.Code,
			"message": domainErr.Message + domainErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *ResourceHandler) DeleteResource(c *gin.Context) {
	resourceIDStr := c.Query("resource_id")
	resourceID, err := uuid.Parse(resourceIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid resource_id: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	req := usecase.ResourceDeleteRequest{ResourceID: resourceID}
	resp, domainErr := h.resourceService.DeleteResource(&context, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    domainErr.Code,
			"message": domainErr.Message + domainErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *ResourceHandler) ListResources(c *gin.Context) {
	context := c.Request.Context()
	req := usecase.ResourceListRequest{}
	resp, domainErr := h.resourceService.ListAllResources(&context, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    domainErr.Code,
			"message": domainErr.Message + domainErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *ResourceHandler) MoveResource(c *gin.Context) {
	var req usecase.ResourceMoveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.resourceService.MoveResource(&context, req)
	if domainErr.Code != "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    domainErr.Code,
			"message": domainErr.Message + domainErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}
