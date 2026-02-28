# Moon Backend API Documentation

## Base URL
All API endpoints are prefixed with `/api/v1`.

## Authentication
API authentication is handled through JWT tokens. After logging in, include the token in the `Authorization` header of subsequent requests:

```
Authorization: Bearer <your-token>
```

## API Endpoints

### 1. User Management

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /register | Register a new user | `{"username": "...", "email": "...", "password": "..."}` | N/A |
| POST | /login | Login and get JWT token | `{"email": "...", "password": "..."}` | N/A |
| POST | /users | Create a new user | `{"username": "...", "email": "...", "password": "..."}` | N/A |
| GET | /users/:id | Get user by ID | N/A | N/A |
| PUT | /users | Update user information | `{"user_id": "...", "username": "...", "email": "..."}` | N/A |
| DELETE | /users | Delete a user | `{"user_id": "..."}` | N/A |
| GET | /users/list | List all users | N/A | N/A |
| POST | /users/assign-role | Assign role to user | `{"user_id": "...", "role_id": "..."}` | N/A |
| POST | /users/remove-role | Remove role from user | `{"user_id": "...", "role_id": "..."}` | N/A |
| GET | /users/roles | Get user roles | N/A | `user_id` |

### 2. Organization Management

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /organizations | Create a new organization | `{"organization_name": "...", "organization_code": "...", "organization_description": "..."}` | N/A |
| GET | /organizations | Get organization by ID | N/A | `organization_id` |
| PUT | /organizations | Update organization | `{"organization_id": "...", "organization_name": "...", "organization_code": "...", "organization_description": "..."}` | N/A |
| DELETE | /organizations | Delete organization | `{"organization_id": "..."}` | N/A |
| GET | /organizations/tree | Get organization tree | N/A | `root_organization_code` (default: "ROOT") |
| GET | /organizations/list | List all organizations | N/A | N/A |
| POST | /organizations/move | Move organization | `{"organization_id": "...", "parent_organization_id": "..."}` | N/A |
| POST | /organizations/assign-role | Assign role to organization | `{"organization_id": "...", "role_id": "..."}` | N/A |
| POST | /organizations/remove-role | Remove role from organization | `{"organization_id": "...", "role_id": "..."}` | N/A |
| GET | /organizations/roles | Get organization roles | N/A | `organization_id` |
| GET | /organizations/users | Get organization users | N/A | `organization_id` |

### 3. Role Management

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /roles | Create a new role | `{"role_name": "...", "role_code": "...", "role_description": "..."}` | N/A |
| GET | /roles | Get role by ID | N/A | `role_id` |
| PUT | /roles | Update role | `{"role_id": "...", "role_name": "...", "role_code": "...", "role_description": "..."}` | N/A |
| DELETE | /roles | Delete role | `{"role_id": "..."}` | N/A |
| POST | /roles/assign-permission | Assign permission to role | `{"role_id": "...", "permission_id": "..."}` | N/A |
| POST | /roles/remove-permission | Remove permission from role | `{"role_id": "...", "permission_id": "..."}` | N/A |
| GET | /roles/permissions | Get role permissions | N/A | `role_id` |
| GET | /roles/list | List all roles | N/A | N/A |

### 4. Permission Management

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /permissions | Create a new permission | `{"permission_name": "...", "permission_code": "...", "permission_description": "..."}` | N/A |
| GET | /permissions/:id | Get permission by ID | N/A | N/A |
| PUT | /permissions | Update permission | `{"permission_id": "...", "permission_name": "...", "permission_code": "...", "permission_description": "..."}` | N/A |
| DELETE | /permissions | Delete permission | `{"permission_id": "..."}` | N/A |
| GET | /permissions/list | List all permissions | N/A | N/A |

### 5. Resource Management

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /resources | Create a new resource | `{"resource_name": "...", "resource_code": "...", "resource_description": "..."}` | N/A |
| GET | /resources | Get resource by ID | N/A | `resource_id` |
| PUT | /resources | Update resource | `{"resource_id": "...", "resource_name": "...", "resource_code": "...", "resource_description": "..."}` | N/A |
| DELETE | /resources | Delete resource | `{"resource_id": "..."}` | N/A |
| GET | /resources/list | List all resources | N/A | N/A |
| POST | /resources/move | Move/transfer resource | `{"resource_id": "...", "new_parent_resource_id": "...", "new_organization_id": "..."}` | N/A |

### 6. Asset Management

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /assets | Create a new asset | `{"asset_name": "...", "asset_code": "...", "description": "...", "asset_type": "...", "ip_address": "..."}` | N/A |
| GET | /assets | Get asset by ID | N/A | `asset_id` |
| PUT | /assets | Update asset | `{"asset_id": "...", "asset_name": "...", "asset_code": "...", "description": "..."}` | N/A |
| DELETE | /assets | Delete asset | `{"asset_id": "..."}` | N/A |
| GET | /assets/list | List all assets | N/A | N/A |
| GET | /assets/organization | List assets by organization | N/A | `organization_id` |

### 7. Vulnerability Management

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /vulnerabilities | Create a new vulnerability | `{"cve_id": "...", "title": "...", "description": "...", "severity": "...", "cvss_score": 0.0}` | N/A |
| GET | /vulnerabilities | Get vulnerability by ID | N/A | `vulnerability_id` |
| GET | /vulnerabilities/cve | Get vulnerability by CVE ID | N/A | `cve_id` |
| PUT | /vulnerabilities | Update vulnerability | `{"vulnerability_id": "...", "title": "...", "description": "...", "severity": "...", "cvss_score": 0.0}` | N/A |
| DELETE | /vulnerabilities | Delete vulnerability | `{"vulnerability_id": "..."}` | N/A |
| GET | /vulnerabilities/list | List all vulnerabilities | N/A | N/A |

### 8. Asset-Vulnerability Management

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /asset-vulnerabilities | Create a new asset-vulnerability relationship | `{"asset_id": "...", "vulnerability_id": "...", "detection_date": "...", "status": "..."}` | N/A |
| GET | /asset-vulnerabilities | Get asset-vulnerability by ID | N/A | `asset_vulnerability_id` |
| GET | /asset-vulnerabilities/asset | List asset-vulnerabilities by asset ID | N/A | `asset_id` |
| GET | /asset-vulnerabilities/vulnerability | List asset-vulnerabilities by vulnerability ID | N/A | `vulnerability_id` |
| PUT | /asset-vulnerabilities | Update asset-vulnerability | `{"asset_vulnerability_id": "...", "status": "...", "remediation_plan": "..."}` | N/A |
| DELETE | /asset-vulnerabilities | Delete asset-vulnerability | `{"asset_vulnerability_id": "..."}` | N/A |
| GET | /asset-vulnerabilities/list | List all asset-vulnerabilities | N/A | N/A |

### 9. AI Management

#### 9.1 Model Configuration

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /ai/model-config | Create a new model configuration | `{"provider_name": "...", "model_name": "...", "api_key": "...", "api_endpoint": "...", "max_tokens": 4096, "temperature": 0.7, "timeout_seconds": 30, "is_active": true, "priority": 1}` | N/A |
| GET | /ai/model-config | Get model configuration by ID | N/A | `config_id` |
| GET | /ai/model-config/active | Get active model configuration | N/A | N/A |
| PUT | /ai/model-config | Update model configuration | `{"config_id": "...", "provider_name": "...", "model_name": "...", "api_key": "...", "is_active": true}` | N/A |
| DELETE | /ai/model-config | Delete model configuration | `{"config_id": "..."}` | N/A |
| GET | /ai/model-config/list | List all model configurations | N/A | N/A |

#### 9.2 Prompt Templates

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /ai/prompt-template | Create a new prompt template | `{"template_name": "...", "template_type": "...", "template_content": "...", "variables": {...}, "description": "...", "is_active": true}` | N/A |
| GET | /ai/prompt-template | Get prompt template by ID | N/A | `template_id` |
| GET | /ai/prompt-template/type | Get prompt template by type | N/A | `template_type` |
| PUT | /ai/prompt-template | Update prompt template | `{"template_id": "...", "template_name": "...", "template_content": "...", "is_active": true}` | N/A |
| DELETE | /ai/prompt-template | Delete prompt template | `{"template_id": "..."}` | N/A |
| GET | /ai/prompt-template/list | List all prompt templates | N/A | N/A |

#### 9.3 Asset Classification

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /ai/classify-asset | Classify an asset using AI | `{"asset_id": "..."}` | N/A |
| GET | /ai/classify-asset | Get classification result by asset ID | N/A | `asset_id` |
| PUT | /ai/classify-asset/approve | Approve or adjust classification | `{"classification_id": "...", "manual_category": "...", "is_approved": true}` | N/A |

#### 9.4 Risk Assessment

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /ai/assess-risk | Assess vulnerability risk using AI | `{"vulnerability_id": "...", "asset_id": "..."}` | N/A |
| GET | /ai/assess-risk | Get assessment result by vulnerability ID | N/A | `vulnerability_id` |

#### 9.5 Security Recommendation

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| POST | /ai/generate-recommendation | Generate security recommendation using AI | `{"vulnerability_id": "..."}` | N/A |
| GET | /ai/generate-recommendation | Get recommendation by vulnerability ID | N/A | `vulnerability_id` |
| PUT | /ai/generate-recommendation/feedback | Submit feedback on recommendation | `{"recommendation_id": "...", "is_useful": true, "feedback": "..."}` | N/A |

#### 9.6 API Call Logs

| Method | Endpoint | Description | Request Body | Query Parameters |
|--------|----------|-------------|-------------|------------------|
| GET | /ai/logs | Get AI API call logs | N/A | `limit`, `offset`, `call_type` |
| GET | /ai/logs/:id | Get specific log by ID | N/A | `log_id` |

## Error Responses

All API endpoints return standardized error responses in the following format:

```json
{
  "code": "400",
  "message": "Error message describing the issue"
}
```

Common error codes:

| Code | Description |
|------|-------------|
| 400 | Bad Request - Invalid input parameters |
| 401 | Unauthorized - Missing or invalid authentication |
| 403 | Forbidden - Insufficient permissions |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error - Server-side error |

## Success Responses

Successful responses return data in the following format:

```json
{
  "code": "200",
  "message": "success",
  "data": {...}
}
```

## Testing

To test the API endpoints, you can use tools like Postman or curl. Here's an example using curl to test the login endpoint:

```bash
curl -X POST http://localhost:8080/api/v1/login \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "password": "password123"}'
```

## Rate Limiting

API requests are rate-limited to prevent abuse. Excessive requests will result in a 429 Too Many Requests response.

## Versioning

The API version is specified in the URL path (`/api/v1`). Future versions will be available at `/api/v2`, etc.
