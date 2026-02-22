package main

import (
	"fmt"
	"time"

	"moon/internal/domain/aggregate"
	"moon/internal/domain/service"
	"moon/internal/infrastructure/http"
	"moon/internal/infrastructure/persistence/postgres"

	"github.com/google/uuid"

	driver "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initSampleData(db *gorm.DB) {
	// Check if there are any assets already
	var assetCount int64
	db.Model(&aggregate.Asset{}).Count(&assetCount)
	if assetCount > 0 {
		fmt.Println("Sample data already exists, skipping initialization")
		return
	}

	fmt.Println("Initializing sample data...")

	// Create sample assets
	assets := []aggregate.Asset{
		{
			AssetID:      uuid.New(),
			AssetName:    "Web Server 1",
			AssetCode:    "WEB-001",
			Description:  "Main web server for the company website",
			AssetType:    "server",
			Manufacturer: "Dell",
			Model:        "PowerEdge R740",
			SerialNumber: "S123456789",
			IPAddress:    "192.168.1.100",
			MACAddress:   "00:11:22:33:44:55",
			Location:     "Data Center",
			Department:   "IT Department",
			Owner:        "John Doe",
			Status:       "active",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			AssetID:      uuid.New(),
			AssetName:    "Database Server",
			AssetCode:    "DB-001",
			Description:  "Main database server",
			AssetType:    "server",
			Manufacturer: "HP",
			Model:        "ProLiant DL380",
			SerialNumber: "S987654321",
			IPAddress:    "192.168.1.101",
			MACAddress:   "00:11:22:33:44:56",
			Location:     "Data Center",
			Department:   "IT Department",
			Owner:        "Jane Smith",
			Status:       "active",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			AssetID:      uuid.New(),
			AssetName:    "Workstation 1",
			AssetCode:    "WS-001",
			Description:  "Developer workstation",
			AssetType:    "workstation",
			Manufacturer: "Lenovo",
			Model:        "ThinkCentre M70t",
			SerialNumber: "S135792468",
			IPAddress:    "192.168.1.200",
			MACAddress:   "00:11:22:33:44:57",
			Location:     "Office",
			Department:   "Development",
			Owner:        "Bob Johnson",
			Status:       "active",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			AssetID:      uuid.New(),
			AssetName:    "Network Switch",
			AssetCode:    "NW-001",
			Description:  "Core network switch",
			AssetType:    "network_device",
			Manufacturer: "Cisco",
			Model:        "Catalyst 9300",
			SerialNumber: "S246813579",
			IPAddress:    "192.168.1.1",
			MACAddress:   "00:11:22:33:44:58",
			Location:     "Data Center",
			Department:   "IT Department",
			Owner:        "Alice Brown",
			Status:       "active",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
		{
			AssetID:      uuid.New(),
			AssetName:    "Firewall",
			AssetCode:    "FW-001",
			Description:  "Corporate firewall",
			AssetType:    "network_device",
			Manufacturer: "Palo Alto",
			Model:        "PA-220",
			SerialNumber: "S975310864",
			IPAddress:    "192.168.1.2",
			MACAddress:   "00:11:22:33:44:59",
			Location:     "Data Center",
			Department:   "IT Department",
			Owner:        "Charlie Davis",
			Status:       "active",
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		},
	}

	for _, asset := range assets {
		db.Create(&asset)
	}

	// Create sample vulnerabilities
	vulnerabilities := []aggregate.Vulnerability{
		{
			VulnerabilityID:  uuid.New(),
			CVEID:            "CVE-2023-21706",
			NISTCVEID:        "CVE-2023-21706",
			Title:            "Windows Kerberos Elevation of Privilege Vulnerability",
			Description:      "An elevation of privilege vulnerability exists in Windows Kerberos.",
			Severity:         "Critical",
			CVSSScore:        9.8,
			AffectedSoftware: "Windows Server 2019, Windows Server 2022",
			AffectedVersion:  "10.0.17763, 10.0.20348",
			Mitigation:       "Apply security update KB5022282",
			Reference:        "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-21706",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			VulnerabilityID:  uuid.New(),
			CVEID:            "CVE-2023-22515",
			NISTCVEID:        "CVE-2023-22515",
			Title:            "Atlassian Confluence Server and Data Center Remote Code Execution Vulnerability",
			Description:      "A remote code execution vulnerability exists in Atlassian Confluence Server and Data Center.",
			Severity:         "Critical",
			CVSSScore:        10.0,
			AffectedSoftware: "Atlassian Confluence Server, Atlassian Confluence Data Center",
			AffectedVersion:  "7.18.0 - 7.19.16, 7.20.0 - 7.20.12, 7.21.0 - 7.21.8, 7.22.0 - 7.22.3",
			Mitigation:       "Update to Confluence 7.19.17, 7.20.13, 7.21.9, or 7.22.4",
			Reference:        "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-22515",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			VulnerabilityID:  uuid.New(),
			CVEID:            "CVE-2023-1389",
			NISTCVEID:        "CVE-2023-1389",
			Title:            "Microsoft Exchange Server Remote Code Execution Vulnerability",
			Description:      "A remote code execution vulnerability exists in Microsoft Exchange Server.",
			Severity:         "Critical",
			CVSSScore:        9.8,
			AffectedSoftware: "Microsoft Exchange Server 2013, Microsoft Exchange Server 2016, Microsoft Exchange Server 2019",
			AffectedVersion:  "15.0.1497.0 - 15.0.1497.32, 15.1.2308.0 - 15.1.2308.24, 15.2.986.0 - 15.2.986.22",
			Mitigation:       "Apply security update KB5024204",
			Reference:        "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-1389",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			VulnerabilityID:  uuid.New(),
			CVEID:            "CVE-2023-20198",
			NISTCVEID:        "CVE-2023-20198",
			Title:            "Cisco IOS XE Software Web UI Privilege Escalation Vulnerability",
			Description:      "A privilege escalation vulnerability exists in the web UI feature of Cisco IOS XE Software.",
			Severity:         "High",
			CVSSScore:        8.8,
			AffectedSoftware: "Cisco IOS XE Software",
			AffectedVersion:  "17.6.1 - 17.6.3, 17.7.1 - 17.7.2, 17.8.1",
			Mitigation:       "Update to Cisco IOS XE Software 17.6.4, 17.7.3, or 17.8.2",
			Reference:        "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-20198",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
		{
			VulnerabilityID:  uuid.New(),
			CVEID:            "CVE-2023-3519",
			NISTCVEID:        "CVE-2023-3519",
			Title:            "Microsoft Windows Secure Boot Security Feature Bypass Vulnerability",
			Description:      "A security feature bypass vulnerability exists in Microsoft Windows Secure Boot.",
			Severity:         "High",
			CVSSScore:        8.2,
			AffectedSoftware: "Microsoft Windows 10, Microsoft Windows 11, Microsoft Windows Server 2019, Microsoft Windows Server 2022",
			AffectedVersion:  "10.0.19044, 10.0.19045, 10.0.20348",
			Mitigation:       "Apply security update KB5028185",
			Reference:        "https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-3519",
			CreatedAt:        time.Now(),
			UpdatedAt:        time.Now(),
		},
	}

	for _, vulnerability := range vulnerabilities {
		db.Create(&vulnerability)
	}

	// Get the created assets and vulnerabilities
	var createdAssets []aggregate.Asset
	db.Find(&createdAssets)

	var createdVulnerabilities []aggregate.Vulnerability
	db.Find(&createdVulnerabilities)

	// Create sample asset-vulnerability relationships
	currentDate := time.Now().Format("2006-01-02")
	assetVulnerabilities := []aggregate.AssetVulnerability{
		{
			AssetVulnerabilityID: uuid.New(),
			AssetID:              createdAssets[0].AssetID,
			VulnerabilityID:      createdVulnerabilities[0].VulnerabilityID,
			DetectionDate:        currentDate,
			Status:               "open",
			RemediationPlan:      "Apply security update KB5022282",
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			AssetVulnerabilityID: uuid.New(),
			AssetID:              createdAssets[1].AssetID,
			VulnerabilityID:      createdVulnerabilities[0].VulnerabilityID,
			DetectionDate:        currentDate,
			Status:               "open",
			RemediationPlan:      "Apply security update KB5022282",
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			AssetVulnerabilityID: uuid.New(),
			AssetID:              createdAssets[0].AssetID,
			VulnerabilityID:      createdVulnerabilities[2].VulnerabilityID,
			DetectionDate:        currentDate,
			Status:               "open",
			RemediationPlan:      "Apply security update KB5024204",
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			AssetVulnerabilityID: uuid.New(),
			AssetID:              createdAssets[3].AssetID,
			VulnerabilityID:      createdVulnerabilities[3].VulnerabilityID,
			DetectionDate:        currentDate,
			Status:               "open",
			RemediationPlan:      "Update to Cisco IOS XE Software 17.6.4",
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
		{
			AssetVulnerabilityID: uuid.New(),
			AssetID:              createdAssets[2].AssetID,
			VulnerabilityID:      createdVulnerabilities[4].VulnerabilityID,
			DetectionDate:        currentDate,
			Status:               "open",
			RemediationPlan:      "Apply security update KB5028185",
			CreatedAt:            time.Now(),
			UpdatedAt:            time.Now(),
		},
	}

	for _, assetVulnerability := range assetVulnerabilities {
		db.Create(&assetVulnerability)
	}

	fmt.Println("Sample data initialization completed")
}

func main() {
	db, err := gorm.Open(driver.Open("host=localhost user=capstone password=capstone dbname=capstone port=5432 sslmode=disable"))
	if err != nil {
		fmt.Printf(err.Error())
	}

	// Auto migrate database tables
	fmt.Println("Running database migrations...")
	err = db.AutoMigrate(
		&aggregate.User{},
		&aggregate.Organization{},
		&aggregate.Role{},
		&aggregate.Permission{},
		&aggregate.UserOrganization{},
	)
	if err != nil {
		fmt.Printf("Database migration error: %v\n", err)
	}
	fmt.Println("Database migration completed")

	// Skip database initialization script due to permission issues
	// The database tables should already be created
	fmt.Println("Skipping database initialization script")

	// Initialize sample data
	initSampleData(db)

	// Initialize repositories
	userRepository := postgres.NewUserRepository(db)
	organizationRepository := postgres.NewOrganizationRepository(db)
	roleRepository := postgres.NewRoleRepository(db)
	permissionRepository := postgres.NewPermissionRepository(db)
	resourceRepository := postgres.NewResourceRepository(db)
	assetRepository := postgres.NewAssetRepository(db)
	vulnerabilityRepository := postgres.NewVulnerabilityRepository(db)
	assetVulnerabilityRepository := postgres.NewAssetVulnerabilityRepository(db)

	// Initialize services
	userService := service.NewUserService(userRepository)
	organizationService := service.NewOrganizationService(organizationRepository, userRepository)
	roleService := service.NewRoleService(roleRepository)
	permissionService := service.NewPermissionService(permissionRepository)
	resourceService := service.NewResourceService(resourceRepository)
	assetService := service.NewAssetService(assetRepository)
	vulnerabilityService := service.NewVulnerabilityService(vulnerabilityRepository)
	assetVulnerabilityService := service.NewAssetVulnerabilityService(assetVulnerabilityRepository)

	// Initialize handlers
	userHandler := http.NewUserHandler(userService)
	organizationHandler := http.NewOrganizationHandler(organizationService)
	roleHandler := http.NewRoleHandler(roleService)
	permissionHandler := http.NewPermissionHandler(permissionService)
	resourceHandler := http.NewResourceHandler(resourceService)
	assetHandler := http.NewAssetHandler(assetService)
	vulnerabilityHandler := http.NewVulnerabilityHandler(vulnerabilityService)
	assetVulnerabilityHandler := http.NewAssetVulnerabilityHandler(assetVulnerabilityService)

	// Setup router
	r := http.SetupRouter(
		userHandler,
		organizationHandler,
		roleHandler,
		permissionHandler,
		resourceHandler,
		assetHandler,
		vulnerabilityHandler,
		assetVulnerabilityHandler,
	)

	r.Run(":8080")
}
