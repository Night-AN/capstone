package http

import (
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RoleHandler struct {
	roleService service.RoleService
}

func NewRoleHandler(service service.RoleService) *RoleHandler {
	return &RoleHandler{roleService: service}
}

func (h *RoleHandler) CreateRole(c *gin.Context) {
	var req usecase.RoleCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.roleService.CreateRole(context, req)
	c.JSON(http.StatusCreated, gin.H{
		"code":    "201",
		"message": "success",
		"data":    resp,
	})
}

func (h *RoleHandler) GetRole(c *gin.Context) {
	var req usecase.RoleGetRequest
	roleIDStr := c.Query("role_id")
	roleID, err := uuid.Parse(roleIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid role_id: " + err.Error(),
		})
		return
	}
	req.RoleID = roleID
	context := c.Request.Context()
	resp := h.roleService.GetRole(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *RoleHandler) UpdateRole(c *gin.Context) {
	var req usecase.RoleUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.roleService.UpdateRole(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *RoleHandler) DeleteRole(c *gin.Context) {
	var req usecase.RoleDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.roleService.DeleteRole(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *RoleHandler) AssignPermission(c *gin.Context) {
	var req usecase.RoleAssignPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.roleService.AssignPermission(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *RoleHandler) RemovePermission(c *gin.Context) {
	var req usecase.RoleRemovePermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.roleService.RemovePermission(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *RoleHandler) GetRolePermissions(c *gin.Context) {
	var req usecase.RolePermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.roleService.GetRolePermissions(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *RoleHandler) ListRoles(c *gin.Context) {
	var req usecase.RoleListRequest
	// No need to bind JSON for empty request
	context := c.Request.Context()
	resp := h.roleService.ListRoles(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *RoleHandler) GetRoleUsers(c *gin.Context) {
	var req usecase.RoleUsersRequest
	roleIDStr := c.Query("role_id")
	roleID, err := uuid.Parse(roleIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid role_id: " + err.Error(),
		})
		return
	}
	req.RoleID = roleID
	context := c.Request.Context()
	resp := h.roleService.GetRoleUsers(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}
