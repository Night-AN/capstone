package http

import (
	"moon/internal/domain/service"
	"moon/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AssetHandler handles HTTP requests for assets
type AssetHandler struct {
	assetService service.AssetService
}

// NewAssetHandler creates a new asset handler instance
func NewAssetHandler(assetService service.AssetService) *AssetHandler {
	return &AssetHandler{assetService: assetService}
}

// CreateAsset handles POST /assets requests
func (h *AssetHandler) CreateAsset(c *gin.Context) {
	var req usecase.AssetCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.assetService.CreateAsset(context, req)
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

// GetAsset handles GET /assets requests
func (h *AssetHandler) GetAsset(c *gin.Context) {
	assetIDStr := c.Query("asset_id")
	assetID, err := uuid.Parse(assetIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid asset_id: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	req := usecase.AssetGetRequest{AssetID: assetID}
	resp, domainErr := h.assetService.GetAssetByID(context, req)
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

// UpdateAsset handles PUT /assets requests
func (h *AssetHandler) UpdateAsset(c *gin.Context) {
	var req usecase.AssetUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	resp, domainErr := h.assetService.UpdateAsset(context, req)
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

// DeleteAsset handles DELETE /assets requests
func (h *AssetHandler) DeleteAsset(c *gin.Context) {
	assetIDStr := c.Query("asset_id")
	assetID, err := uuid.Parse(assetIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid asset_id: " + err.Error(),
		})
		return
	}
	context := c.Request.Context()
	req := usecase.AssetDeleteRequest{AssetID: assetID}
	resp, domainErr := h.assetService.DeleteAsset(context, req)
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

// ListAssets handles GET /assets/list requests
func (h *AssetHandler) ListAssets(c *gin.Context) {
	context := c.Request.Context()
	req := usecase.AssetListRequest{}
	resp, domainErr := h.assetService.ListAllAssets(context, req)
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

// ListAssetsByOrganization handles GET /assets/organization requests
func (h *AssetHandler) ListAssetsByOrganization(c *gin.Context) {
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

	// 这里需要在AssetService中添加ListAssetsByOrganizationID方法
	// 由于我们还没有实现这个方法，暂时返回空列表
	c.JSON(http.StatusOK, gin.H{
		"code":    "200",
		"message": "success",
		"data":    []interface{}{},
	})
}

// BatchCreateAsset handles POST /assets/batch requests
func (h *AssetHandler) BatchCreateAsset(c *gin.Context) {
	var req struct {
		Assets                 []usecase.AssetCreateRequest `json:"assets"`
		EnableAIClassification bool                         `json:"enable_ai_classification"`
		PromptTemplateID       *string                      `json:"prompt_template_id,omitempty"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: " + err.Error(),
		})
		return
	}

	if len(req.Assets) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "400",
			"message": "invalid request: assets array is empty",
		})
		return
	}

	context := c.Request.Context()
	batchReq := usecase.BatchAssetCreateRequest{
		Assets:                 req.Assets,
		EnableAIClassification: req.EnableAIClassification,
		PromptTemplateID:       req.PromptTemplateID,
	}
	resp, domainErr := h.assetService.BatchCreateAsset(context, batchReq, req.EnableAIClassification)
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
