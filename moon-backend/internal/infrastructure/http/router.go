package http

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter(
	userHandler *UserHandler,
	organizationHandler *OrganizationHandler,
	roleHandler *RoleHandler,
	permissionHandler *PermissionHandler,
	resourceHandler *ResourceHandler,
	assetHandler *AssetHandler,
	vulnerabilityHandler *VulnerabilityHandler,
	assetVulnerabilityHandler *AssetVulnerabilityHandler,
	aiHandler *AIHandler,
) *gin.Engine {
	r := gin.Default()
	r.Use(CORSMiddleware())

	v1 := r.Group("/api/v1")
	{
		// User routes
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)
		v1.POST("/users", userHandler.CreateUser)
		v1.GET("/users/:id", userHandler.GetUser)
		v1.PUT("/users", userHandler.UpdateUser)
		v1.DELETE("/users", userHandler.DeleteUser)
		v1.GET("/users/list", userHandler.ListUsers)
		v1.POST("/users/assign-role", userHandler.AssignRoleToUser)
		v1.POST("/users/remove-role", userHandler.RemoveRoleFromUser)
		v1.GET("/users/roles", userHandler.GetUserRoles)

		// Organization routes
		v1.POST("/organizations", organizationHandler.CreateOrganization)
		v1.GET("/organizations", organizationHandler.GetOrganization)
		v1.PUT("/organizations", organizationHandler.UpdateOrganization)
		v1.DELETE("/organizations", organizationHandler.DeleteOrganization)
		v1.GET("/organizations/tree", organizationHandler.GetOrganizationTree)
		v1.GET("/organizations/list", organizationHandler.ListOrganizations)
		v1.POST("/organizations/move", organizationHandler.MoveOrganization)
		v1.POST("/organizations/assign-role", organizationHandler.AssignRoleToOrganization)
		v1.POST("/organizations/remove-role", organizationHandler.RemoveRoleFromOrganization)
		v1.GET("/organizations/roles", organizationHandler.GetOrganizationRoles)
		v1.GET("/organizations/users", organizationHandler.GetOrganizationUsers)

		// 角色管理
		v1.POST("/roles", roleHandler.CreateRole)
		v1.GET("/roles", roleHandler.GetRole)
		v1.PUT("/roles", roleHandler.UpdateRole)
		v1.DELETE("/roles", roleHandler.DeleteRole)
		v1.POST("/roles/assign-permission", roleHandler.AssignPermission)
		v1.POST("/roles/remove-permission", roleHandler.RemovePermission)
		v1.GET("/roles/permissions", roleHandler.GetRolePermissions)
		v1.GET("/roles/list", roleHandler.ListRoles)
		v1.GET("/roles/users", roleHandler.GetRoleUsers)

		// Permission routes
		v1.POST("/permissions", permissionHandler.CreatePermission)
		v1.GET("/permissions/:id", permissionHandler.GetPermission)
		v1.PUT("/permissions", permissionHandler.UpdatePermission)
		v1.DELETE("/permissions", permissionHandler.DeletePermission)
		v1.GET("/permissions/list", permissionHandler.ListPermissions)

		// Resource routes
		v1.POST("/resources", resourceHandler.CreateResource)
		v1.GET("/resources", resourceHandler.GetResource)
		v1.PUT("/resources", resourceHandler.UpdateResource)
		v1.DELETE("/resources", resourceHandler.DeleteResource)
		v1.GET("/resources/list", resourceHandler.ListResources)
		v1.POST("/resources/move", resourceHandler.MoveResource)

		// Asset routes
		v1.POST("/assets", assetHandler.CreateAsset)
		v1.GET("/assets", assetHandler.GetAsset)
		v1.PUT("/assets", assetHandler.UpdateAsset)
		v1.DELETE("/assets", assetHandler.DeleteAsset)
		v1.GET("/assets/list", assetHandler.ListAssets)
		v1.GET("/assets/organization", assetHandler.ListAssetsByOrganization)
		v1.POST("/assets/batch", assetHandler.BatchCreateAsset)

		// Vulnerability routes
		v1.POST("/vulnerabilities", vulnerabilityHandler.CreateVulnerability)
		v1.GET("/vulnerabilities", vulnerabilityHandler.GetVulnerability)
		v1.GET("/vulnerabilities/cve", vulnerabilityHandler.GetVulnerabilityByCVEID)
		v1.PUT("/vulnerabilities", vulnerabilityHandler.UpdateVulnerability)
		v1.DELETE("/vulnerabilities", vulnerabilityHandler.DeleteVulnerability)
		v1.GET("/vulnerabilities/list", vulnerabilityHandler.ListVulnerabilities)

		// Asset-Vulnerability routes
		v1.POST("/asset-vulnerabilities", assetVulnerabilityHandler.CreateAssetVulnerability)
		v1.GET("/asset-vulnerabilities", assetVulnerabilityHandler.GetAssetVulnerability)
		v1.GET("/asset-vulnerabilities/asset", assetVulnerabilityHandler.ListAssetVulnerabilitiesByAssetID)
		v1.GET("/asset-vulnerabilities/vulnerability", assetVulnerabilityHandler.ListAssetVulnerabilitiesByVulnerabilityID)
		v1.PUT("/asset-vulnerabilities", assetVulnerabilityHandler.UpdateAssetVulnerability)
		v1.DELETE("/asset-vulnerabilities", assetVulnerabilityHandler.DeleteAssetVulnerability)
		v1.GET("/asset-vulnerabilities/list", assetVulnerabilityHandler.ListAssetVulnerabilities)

		// AI routes
		v1.POST("/ai/model-config", aiHandler.CreateModelConfig)
		v1.GET("/ai/model-config", aiHandler.GetModelConfig)
		v1.GET("/ai/model-config/active", aiHandler.GetActiveModelConfig)
		v1.PUT("/ai/model-config", aiHandler.UpdateModelConfig)
		v1.DELETE("/ai/model-config", aiHandler.DeleteModelConfig)
		v1.GET("/ai/model-config/list", aiHandler.ListModelConfigs)

		v1.POST("/ai/prompt-template", aiHandler.CreatePromptTemplate)
		v1.GET("/ai/prompt-template", aiHandler.GetPromptTemplate)
		v1.GET("/ai/prompt-template/type", aiHandler.GetPromptTemplateByType)
		v1.PUT("/ai/prompt-template", aiHandler.UpdatePromptTemplate)
		v1.DELETE("/ai/prompt-template", aiHandler.DeletePromptTemplate)
		v1.GET("/ai/prompt-template/list", aiHandler.ListPromptTemplates)

		v1.POST("/ai/classify-asset", aiHandler.ClassifyAsset)
		v1.GET("/ai/classify-asset", aiHandler.GetClassificationByAssetID)
		v1.PUT("/ai/classify-asset/approve", aiHandler.ApproveClassification)

		v1.POST("/ai/assess-risk", aiHandler.AssessRisk)
		v1.GET("/ai/assess-risk", aiHandler.GetAssessmentByVulnerabilityID)

		v1.POST("/ai/generate-recommendation", aiHandler.GenerateRecommendation)
		v1.GET("/ai/generate-recommendation", aiHandler.GetRecommendationByVulnerabilityID)
		v1.PUT("/ai/generate-recommendation/feedback", aiHandler.SubmitFeedback)

		v1.GET("/ai/logs", aiHandler.ListAPICallLogs)
		v1.GET("/ai/logs/:id", aiHandler.GetAPICallLog)
		v1.POST("/ai/chat", aiHandler.Chat)
		v1.GET("/ai/chat", aiHandler.GetConversation)
		v1.GET("/ai/chat/list", aiHandler.ListConversations)
		v1.DELETE("/ai/chat", aiHandler.DeleteConversation)
	}

	return r
}
