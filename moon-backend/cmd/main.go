package main

import (
	"fmt"
	"moon/internal/domain/service"
	"moon/internal/infrastructure/http"
	"moon/internal/infrastructure/persistence/postgres"

	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(driver.Open("host=localhost user=capstone password=capstone dbname=capstone port=5432 sslmode=disable"))
	if err != nil {
		fmt.Printf(err.Error())
	}

	// Initialize repositories
	userRepository := postgres.NewUserRepository(db)
	organizationRepository := postgres.NewOrganizationRepository(db)
	roleRepository := postgres.NewRoleRepository(db)
	permissionRepository := postgres.NewPermissionRepository(db)
	resourceRepository := postgres.NewResourceRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepository)
	organizationService := service.NewOrganizationService(organizationRepository)
	roleService := service.NewRoleService(roleRepository)
	permissionService := service.NewPermissionService(permissionRepository)
	resourceService := service.NewResourceService(resourceRepository)

	// Initialize handlers
	userHandler := http.NewUserHandler(userService)
	organizationHandler := http.NewOrganizationHandler(organizationService)
	roleHandler := http.NewRoleHandler(roleService)
	permissionHandler := http.NewPermissionHandler(permissionService)
	resourceHandler := http.NewResourceHandler(resourceService)

	// Setup router
	r := http.SetupRouter(
		userHandler,
		organizationHandler,
		roleHandler,
		permissionHandler,
		resourceHandler,
	)

	r.Run(":8080")
}
