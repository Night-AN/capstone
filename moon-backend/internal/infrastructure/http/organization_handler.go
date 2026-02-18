package http

import (
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrganizationHandler struct {
	organizationService service.OrganizationService
}

func NewOrganizationHandler(service service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{organizationService: service}
}

func (h *OrganizationHandler) CreateOrganization(c *gin.Context) {
	var req usecase.OrganizationCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.organizationService.CreateOrganization(&context, req)
	c.JSON(http.StatusCreated, gin.H{
		"code":    "201",
		"message": "success",
		"data":    resp,
	})
}

func (h *OrganizationHandler) GetOrganization(c *gin.Context) {
	var req usecase.OrganizationGetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.organizationService.GetOrganization(&context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *OrganizationHandler) UpdateOrganization(c *gin.Context) {
	var req usecase.OrganizationUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.organizationService.UpdateOrganization(&context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *OrganizationHandler) DeleteOrganization(c *gin.Context) {
	var req usecase.OrganizationDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.organizationService.DeleteOrganization(&context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *OrganizationHandler) GetOrganizationTree(c *gin.Context) {
	var req usecase.OrganizationTreeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.organizationService.GetOrganizationTree(&context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}
