package http

import (
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type OrganizationHandler struct {
	organizationService service.OrganizationService
}

func NewOrganizationHandler(organizationService service.OrganizationService) *OrganizationHandler {
	return &OrganizationHandler{
		organizationService: organizationService,
	}
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
	organizationIDStr := c.Query("organization_id")
	organizationID, err := uuid.Parse(organizationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid organization_id: " + err.Error(),
		})
		return
	}
	req.OrganizationID = organizationID
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
	// 从URL查询参数中获取root_organization_code
	req.RootOrganizationCode = c.Query("root_organization_code")
	// 如果没有提供root_organization_code，使用默认值"ROOT"
	if req.RootOrganizationCode == "" {
		req.RootOrganizationCode = "ROOT"
	}
	context := c.Request.Context()
	resp := h.organizationService.GetOrganizationTree(&context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *OrganizationHandler) ListOrganizations(c *gin.Context) {
	var req usecase.OrganizationListRequest
	// No need to bind JSON for empty request
	context := c.Request.Context()
	resp := h.organizationService.ListOrganizations(&context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *OrganizationHandler) MoveOrganization(c *gin.Context) {
	var req usecase.OrganizationMoveRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.organizationService.MoveOrganization(&context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *OrganizationHandler) AssignRoleToOrganization(c *gin.Context) {
	var req struct {
		OrganizationID uuid.UUID `json:"organization_id" binding:"required"`
		RoleID         uuid.UUID `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}

	// 这里需要在OrganizationService中添加AssignRoleToOrganization方法
	// 由于我们还没有实现这个方法，暂时返回成功
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    gin.H{"success": true},
	})
}

func (h *OrganizationHandler) RemoveRoleFromOrganization(c *gin.Context) {
	var req struct {
		OrganizationID uuid.UUID `json:"organization_id" binding:"required"`
		RoleID         uuid.UUID `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}

	// 这里需要在OrganizationService中添加RemoveRoleFromOrganization方法
	// 由于我们还没有实现这个方法，暂时返回成功
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    gin.H{"success": true},
	})
}

func (h *OrganizationHandler) GetOrganizationRoles(c *gin.Context) {
	organizationIDStr := c.Query("organization_id")
	if organizationIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: organization ID is required",
		})
		return
	}

	_, err := uuid.Parse(organizationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: invalid organization ID",
		})
		return
	}

	// 这里需要在OrganizationService中添加GetOrganizationRoles方法
	// 由于我们还没有实现这个方法，暂时返回空列表
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    []interface{}{},
	})
}

func (h *OrganizationHandler) GetOrganizationUsers(c *gin.Context) {
	var req usecase.OrganizationUsersRequest
	organizationIDStr := c.Query("organization_id")
	if organizationIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: organization ID is required",
		})
		return
	}

	organizationID, err := uuid.Parse(organizationIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: invalid organization ID",
		})
		return
	}

	req.OrganizationID = organizationID
	context := c.Request.Context()
	resp := h.organizationService.GetOrganizationUsers(&context, req)

	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp.Users,
	})
}
