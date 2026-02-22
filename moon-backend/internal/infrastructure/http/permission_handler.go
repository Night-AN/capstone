package http

import (
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PermissionHandler struct {
	permissionService service.PermissionService
}

func NewPermissionHandler(service service.PermissionService) *PermissionHandler {
	return &PermissionHandler{permissionService: service}
}

func (h *PermissionHandler) CreatePermission(c *gin.Context) {
	var req usecase.PermissionCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.permissionService.CreatePermission(context, req)
	c.JSON(http.StatusCreated, gin.H{
		"code":    "201",
		"message": "success",
		"data":    resp,
	})
}

func (h *PermissionHandler) GetPermission(c *gin.Context) {
	// 从URL路径参数中获取permissionId
	permissionIdStr := c.Param("id")
	permissionId, err := uuid.Parse(permissionIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid permission id: " + err.Error(),
		})
		return
	}

	var req usecase.PermissionGetRequest
	req.PermissionID = permissionId

	context := c.Request.Context()
	resp := h.permissionService.GetPermission(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *PermissionHandler) UpdatePermission(c *gin.Context) {
	var req usecase.PermissionUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.permissionService.UpdatePermission(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *PermissionHandler) DeletePermission(c *gin.Context) {
	var req usecase.PermissionDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp := h.permissionService.DeletePermission(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}

func (h *PermissionHandler) ListPermissions(c *gin.Context) {
	var req usecase.PermissionListRequest

	// 从查询参数中获取limit和offset
	req.Limit, _ = strconv.Atoi(c.DefaultQuery("limit", "100"))
	req.Offset, _ = strconv.Atoi(c.DefaultQuery("offset", "0"))

	context := c.Request.Context()
	resp := h.permissionService.ListPermissions(context, req)
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    resp,
	})
}
