#!/bin/bash

# 父子组织API的curl命令示例

# 1. 获取组织树 API
echo "=== 获取组织树 API ==="
curl -X POST "http://localhost:8080/api/v1/organizations/tree" \
  -H "Content-Type: application/json" \
  -d '{
    "root_organization_code": "ROOT"
  }'

# 2. 移动组织 API
echo "\n=== 移动组织 API ==="
curl -X POST "http://localhost:8080/api/v1/organizations/move" \
  -H "Content-Type: application/json" \
  -d '{
    "organization_id": "YOUR_ORGANIZATION_ID",
    "new_parent_id": "YOUR_NEW_PARENT_ID"
  }'

# 3. 创建组织 API
echo "\n=== 创建组织 API ==="
curl -X POST "http://localhost:8080/api/v1/organizations" \
  -H "Content-Type: application/json" \
  -d '{
    "organization_name": "新组织",
    "organization_code": "ROOT::NewOrg",
    "organization_description": "新组织描述",
    "organization_flag": "新组织标志",
    "sensitive_flag": false,
    "parent_id": "YOUR_PARENT_ORGANIZATION_ID"
  }'

# 4. 更新组织 API
echo "\n=== 更新组织 API ==="
curl -X PUT "http://localhost:8080/api/v1/organizations" \
  -H "Content-Type: application/json" \
  -d '{
    "organization_id": "YOUR_ORGANIZATION_ID",
    "organization_name": "更新后的组织",
    "organization_code": "ROOT::UpdatedOrg",
    "organization_description": "更新后的组织描述",
    "organization_flag": "更新后的组织标志",
    "sensitive_flag": false,
    "parent_id": "YOUR_PARENT_ORGANIZATION_ID"
  }'
