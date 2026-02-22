package http

import (
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{userService: service}
}

func (h *UserHandler) Login(c *gin.Context) {
	var req usecase.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.userService.Login(&context, req.Email, req.Password)
	if domainErr.Code != "" {
		c.JSON(http.StatusUnauthorized, gin.H{
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

func (h *UserHandler) Register(c *gin.Context) {
	var req usecase.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.userService.Register(&context, req)
	if domainErr.Code != "" {
		status := http.StatusInternalServerError
		if domainErr.Code == "409" {
			status = http.StatusConflict
		}
		c.JSON(status, gin.H{
			"code":    domainErr.Code,
			"message": domainErr.Message,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"code":    "201",
		"message": "success",
		"data":    resp,
	})
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var req usecase.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.userService.CreateUser(&context, req)
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

func (h *UserHandler) UpdateUser(c *gin.Context) {
	var req usecase.UserUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.userService.UpdateUser(&context, req)
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

func (h *UserHandler) DeleteUser(c *gin.Context) {
	var req usecase.UserDeleteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.userService.DeleteUser(&context, req)
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

func (h *UserHandler) ListUsers(c *gin.Context) {
	context := c.Request.Context()
	req := usecase.UserListRequest{}
	resp, domainErr := h.userService.ListUsers(&context, req)
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

func (h *UserHandler) GetUser(c *gin.Context) {
	userID := c.Param("id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: user ID is required",
		})
		return
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: invalid user ID",
		})
		return
	}

	context := c.Request.Context()
	req := usecase.UserGetRequest{UserID: userUUID}
	resp, domainErr := h.userService.GetUserByID(&context, req)
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

func (h *UserHandler) AssignRoleToUser(c *gin.Context) {
	var req struct {
		UserID uuid.UUID `json:"user_id" binding:"required"`
		RoleID uuid.UUID `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}

	context := c.Request.Context()
	success, domainErr := h.userService.AssignRoleToUser(&context, req.UserID, req.RoleID)
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
		"data":    gin.H{"success": success},
	})
}

func (h *UserHandler) RemoveRoleFromUser(c *gin.Context) {
	var req struct {
		UserID uuid.UUID `json:"user_id" binding:"required"`
		RoleID uuid.UUID `json:"role_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}

	context := c.Request.Context()
	success, domainErr := h.userService.RemoveRoleFromUser(&context, req.UserID, req.RoleID)
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
		"data":    gin.H{"success": success},
	})
}

func (h *UserHandler) GetUserRoles(c *gin.Context) {
	userIDStr := c.Query("user_id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: user ID is required",
		})
		return
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: invalid user ID",
		})
		return
	}

	context := c.Request.Context()
	roles, domainErr := h.userService.GetUserRoles(&context, userID)
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
		"data":    roles,
	})
}
