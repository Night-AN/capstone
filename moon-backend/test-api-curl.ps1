# Test script for API endpoints using curl and PowerShell

# Base URL for the API
$baseUrl = "http://localhost:8080/api/v1"

# Function to test an API endpoint using curl
function Test-Endpoint {
    param(
        [string]$method,
        [string]$endpoint,
        [string]$body = "",
        [string]$description
    )

    Write-Host "`n=== Testing $description ==="
    Write-Host "Method: $method"
    Write-Host "Endpoint: $endpoint"
    
    if ($body) {
        Write-Host "Body: $body"
        curl -X $method -H "Content-Type: application/json" -d $body "$baseUrl$endpoint"
    } else {
        curl -X $method -H "Content-Type: application/json" "$baseUrl$endpoint"
    }
}

# Test user endpoints
Write-Host "`n================ Testing User Endpoints ================"

# Test register endpoint
$registerBody = '{"Username": "testuser", "Email": "test@example.com", "Password": "password123", "OrganizationID": "00000000-0000-0000-0000-000000000000"}'
Test-Endpoint -method "POST" -endpoint "/register" -body $registerBody -description "Register User"

# Test login endpoint
$loginBody = '{"Email": "test@example.com", "Password": "password123"}'
Test-Endpoint -method "POST" -endpoint "/login" -body $loginBody -description "Login User"

# Test organization endpoints
Write-Host "`n================ Testing Organization Endpoints ================"

# Test create organization endpoint
$createOrgBody = '{"OrganizationName": "Test Organization", "Description": "A test organization", "OrganizationCode": "TEST", "OrganizationType": "Company", "ParentOrganizationID": null}'
Test-Endpoint -method "POST" -endpoint "/organizations" -body $createOrgBody -description "Create Organization"

# Test get organization endpoint
$getOrgBody = '{"OrganizationID": "00000000-0000-0000-0000-000000000000"}'
Test-Endpoint -method "GET" -endpoint "/organizations" -body $getOrgBody -description "Get Organization"

# Test get organization tree endpoint
Test-Endpoint -method "GET" -endpoint "/organizations/tree" -description "Get Organization Tree"

# Test role endpoints
Write-Host "`n================ Testing Role Endpoints ================"

# Test create role endpoint
$createRoleBody = '{"RoleName": "Test Role", "Description": "A test role", "RoleCode": "TEST_ROLE", "RoleFlag": "User", "SensitiveFlag": false}'
Test-Endpoint -method "POST" -endpoint "/roles" -body $createRoleBody -description "Create Role"

# Test get role endpoint
$getRoleBody = '{"RoleID": "00000000-0000-0000-0000-000000000000"}'
Test-Endpoint -method "GET" -endpoint "/roles" -body $getRoleBody -description "Get Role"

# Test permission endpoints
Write-Host "`n================ Testing Permission Endpoints ================"

# Test create permission endpoint
$createPermBody = '{"PermissionName": "Test Permission", "Description": "A test permission", "PermissionCode": "TEST_PERMISSION", "SensitiveFlag": false}'
Test-Endpoint -method "POST" -endpoint "/permissions" -body $createPermBody -description "Create Permission"

# Test get permission endpoint
$getPermBody = '{"PermissionID": "00000000-0000-0000-0000-000000000000"}'
Test-Endpoint -method "GET" -endpoint "/permissions" -body $getPermBody -description "Get Permission"

# Test list permissions endpoint
$listPermBody = '{"Limit": 10, "Offset": 0}'
Test-Endpoint -method "GET" -endpoint "/permissions/list" -body $listPermBody -description "List Permissions"

Write-Host "`n================ Test Completed ================"
