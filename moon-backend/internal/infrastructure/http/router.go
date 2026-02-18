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

		// Organization routes
		v1.POST("/organizations", organizationHandler.CreateOrganization)
		v1.GET("/organizations", organizationHandler.GetOrganization)
		v1.PUT("/organizations", organizationHandler.UpdateOrganization)
		v1.DELETE("/organizations", organizationHandler.DeleteOrganization)
		v1.GET("/organizations/tree", organizationHandler.GetOrganizationTree)

		// Role routes
		v1.POST("/roles", roleHandler.CreateRole)
		v1.GET("/roles", roleHandler.GetRole)
		v1.PUT("/roles", roleHandler.UpdateRole)
		v1.DELETE("/roles", roleHandler.DeleteRole)
		v1.POST("/roles/assign-permission", roleHandler.AssignPermission)
		v1.POST("/roles/remove-permission", roleHandler.RemovePermission)
		v1.GET("/roles/permissions", roleHandler.GetRolePermissions)

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
	}

	return r
}
