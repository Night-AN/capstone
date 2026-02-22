# 网络资产管理系统设计

## 1. 数据库模型设计

### 1.1 资产表 (systems.asset)

| 字段名 | 数据类型 | 约束 | 描述 |
| :--- | :--- | :--- | :--- |
| `asset_id` | `UUID` | `PRIMARY KEY` | 资产唯一标识符 |
| `asset_name` | `VARCHAR(255)` | `NOT NULL` | 资产名称 |
| `asset_code` | `VARCHAR(255)` | `UNIQUE NOT NULL` | 资产编码 |
| `asset_type` | `VARCHAR(50)` | `NOT NULL` | 资产类型（如：server, workstation, network_device, storage等） |
| `manufacturer` | `VARCHAR(255)` | | 制造商 |
| `model` | `VARCHAR(255)` | | 型号 |
| `serial_number` | `VARCHAR(255)` | | 序列号 |
| `ip_address` | `VARCHAR(50)` | | IP地址 |
| `mac_address` | `VARCHAR(50)` | | MAC地址 |
| `os_name` | `VARCHAR(255)` | | 操作系统名称 |
| `os_version` | `VARCHAR(100)` | | 操作系统版本 |
| `location` | `VARCHAR(255)` | | 物理位置 |
| `status` | `VARCHAR(50)` | `NOT NULL` | 状态（如：active, inactive, maintenance等） |
| `purchase_date` | `DATE` | | 购买日期 |
| `warranty_end_date` | `DATE` | | 保修结束日期 |
| `owner` | `VARCHAR(255)` | | 负责人 |
| `department` | `VARCHAR(255)` | | 所属部门 |
| `description` | `TEXT` | | 描述 |
| `created_at` | `TIMESTAMP` | `NOT NULL DEFAULT CURRENT_TIMESTAMP` | 创建时间 |
| `updated_at` | `TIMESTAMP` | `NOT NULL DEFAULT CURRENT_TIMESTAMP` | 更新时间 |

### 1.2 漏洞表 (systems.vulnerability)

| 字段名 | 数据类型 | 约束 | 描述 |
| :--- | :--- | :--- | :--- |
| `vulnerability_id` | `UUID` | `PRIMARY KEY` | 漏洞唯一标识符 |
| `cve_id` | `VARCHAR(50)` | `UNIQUE` | CVE编号 |
| `nist_cve_id` | `VARCHAR(50)` | | NIST CVE编号 |
| `title` | `VARCHAR(255)` | `NOT NULL` | 漏洞标题 |
| `description` | `TEXT` | `NOT NULL` | 漏洞描述 |
| `severity` | `VARCHAR(20)` | `NOT NULL` | 严重程度（如：critical, high, medium, low） |
| `cvss_score` | `DECIMAL(3,1)` | | CVSS评分 |
| `affected_software` | `VARCHAR(255)` | | 受影响软件 |
| `affected_version` | `VARCHAR(255)` | | 受影响版本 |
| `publication_date` | `DATE` | | 发布日期 |
| `last_updated_date` | `DATE` | | 最后更新日期 |
| `mitigation` | `TEXT` | | 缓解措施 |
| `reference` | `TEXT` | | 参考链接 |
| `created_at` | `TIMESTAMP` | `NOT NULL DEFAULT CURRENT_TIMESTAMP` | 创建时间 |
| `updated_at` | `TIMESTAMP` | `NOT NULL DEFAULT CURRENT_TIMESTAMP` | 更新时间 |

### 1.3 资产漏洞关联表 (systems.asset_vulnerability)

| 字段名 | 数据类型 | 约束 | 描述 |
| :--- | :--- | :--- | :--- |
| `asset_vulnerability_id` | `UUID` | `PRIMARY KEY` | 关联唯一标识符 |
| `asset_id` | `UUID` | `NOT NULL REFERENCES systems.asset(asset_id)` | 资产ID |
| `vulnerability_id` | `UUID` | `NOT NULL REFERENCES systems.vulnerability(vulnerability_id)` | 漏洞ID |
| `detection_date` | `DATE` | `NOT NULL` | 检测日期 |
| `status` | `VARCHAR(50)` | `NOT NULL` | 状态（如：open, in_progress, closed） |
| `remediation_plan` | `TEXT` | | 修复计划 |
| `remediation_date` | `DATE` | | 修复日期 |
| `closed_date` | `DATE` | | 关闭日期 |
| `created_at` | `TIMESTAMP` | `NOT NULL DEFAULT CURRENT_TIMESTAMP` | 创建时间 |
| `updated_at` | `TIMESTAMP` | `NOT NULL DEFAULT CURRENT_TIMESTAMP` | 更新时间 |

### 1.4 扫描任务表 (systems.scan_task)

| 字段名 | 数据类型 | 约束 | 描述 |
| :--- | :--- | :--- | :--- |
| `scan_task_id` | `UUID` | `PRIMARY KEY` | 扫描任务唯一标识符 |
| `task_name` | `VARCHAR(255)` | `NOT NULL` | 任务名称 |
| `scan_type` | `VARCHAR(50)` | `NOT NULL` | 扫描类型（如：full, quick, targeted） |
| `target_assets` | `TEXT` | | 目标资产（如：IP范围、资产ID列表） |
| `status` | `VARCHAR(50)` | `NOT NULL` | 状态（如：pending, running, completed, failed） |
| `start_time` | `TIMESTAMP` | | 开始时间 |
| `end_time` | `TIMESTAMP` | | 结束时间 |
| `scan_tool` | `VARCHAR(255)` | | 扫描工具 |
| `scan_result` | `TEXT` | | 扫描结果摘要 |
| `created_at` | `TIMESTAMP` | `NOT NULL DEFAULT CURRENT_TIMESTAMP` | 创建时间 |
| `updated_at` | `TIMESTAMP` | `NOT NULL DEFAULT CURRENT_TIMESTAMP` | 更新时间 |

### 1.5 扫描结果表 (systems.scan_result)

| 字段名 | 数据类型 | 约束 | 描述 |
| :--- | :--- | :--- | :--- |
| `scan_result_id` | `UUID` | `PRIMARY KEY` | 扫描结果唯一标识符 |
| `scan_task_id` | `UUID` | `NOT NULL REFERENCES systems.scan_task(scan_task_id)` | 扫描任务ID |
| `asset_id` | `UUID` | `NOT NULL REFERENCES systems.asset(asset_id)` | 资产ID |
| `vulnerability_id` | `UUID` | `REFERENCES systems.vulnerability(vulnerability_id)` | 漏洞ID |
| `detection_method` | `VARCHAR(255)` | | 检测方法 |
| `confidence` | `VARCHAR(50)` | | 置信度（如：high, medium, low） |
| `evidence` | `TEXT` | | 证据 |
| `created_at` | `TIMESTAMP` | `NOT NULL DEFAULT CURRENT_TIMESTAMP` | 创建时间 |

## 2. API 设计

### 2.1 资产管理 API

| 端点 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/api/v1/assets` | `POST` | 创建资产 |
| `/api/v1/assets/:id` | `GET` | 获取资产详情 |
| `/api/v1/assets` | `GET` | 获取资产列表 |
| `/api/v1/assets/:id` | `PUT` | 更新资产 |
| `/api/v1/assets/:id` | `DELETE` | 删除资产 |
| `/api/v1/assets/bulk` | `POST` | 批量创建资产 |
| `/api/v1/assets/bulk` | `PUT` | 批量更新资产 |
| `/api/v1/assets/import` | `POST` | 导入资产（如：从CSV文件） |
| `/api/v1/assets/export` | `GET` | 导出资产（如：为CSV文件） |

### 2.2 漏洞管理 API

| 端点 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/api/v1/vulnerabilities` | `POST` | 创建漏洞 |
| `/api/v1/vulnerabilities/:id` | `GET` | 获取漏洞详情 |
| `/api/v1/vulnerabilities` | `GET` | 获取漏洞列表 |
| `/api/v1/vulnerabilities/:id` | `PUT` | 更新漏洞 |
| `/api/v1/vulnerabilities/:id` | `DELETE` | 删除漏洞 |
| `/api/v1/vulnerabilities/import` | `POST` | 导入漏洞（如：从CVE数据库） |
| `/api/v1/vulnerabilities/export` | `GET` | 导出漏洞（如：为CSV文件） |
| `/api/v1/vulnerabilities/cve/:cveId` | `GET` | 通过CVE ID获取漏洞 |

### 2.3 资产漏洞关联 API

| 端点 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/api/v1/asset-vulnerabilities` | `POST` | 创建资产漏洞关联 |
| `/api/v1/asset-vulnerabilities/:id` | `GET` | 获取资产漏洞关联详情 |
| `/api/v1/asset-vulnerabilities` | `GET` | 获取资产漏洞关联列表 |
| `/api/v1/asset-vulnerabilities/:id` | `PUT` | 更新资产漏洞关联 |
| `/api/v1/asset-vulnerabilities/:id` | `DELETE` | 删除资产漏洞关联 |
| `/api/v1/assets/:assetId/vulnerabilities` | `GET` | 获取资产的所有漏洞 |
| `/api/v1/vulnerabilities/:vulnerabilityId/assets` | `GET` | 获取受漏洞影响的所有资产 |

### 2.4 扫描管理 API

| 端点 | 方法 | 描述 |
| :--- | :--- | :--- |
| `/api/v1/scan-tasks` | `POST` | 创建扫描任务 |
| `/api/v1/scan-tasks/:id` | `GET` | 获取扫描任务详情 |
| `/api/v1/scan-tasks` | `GET` | 获取扫描任务列表 |
| `/api/v1/scan-tasks/:id` | `PUT` | 更新扫描任务 |
| `/api/v1/scan-tasks/:id` | `DELETE` | 删除扫描任务 |
| `/api/v1/scan-tasks/:id/start` | `POST` | 开始扫描任务 |
| `/api/v1/scan-tasks/:id/cancel` | `POST` | 取消扫描任务 |
| `/api/v1/scan-results` | `GET` | 获取扫描结果列表 |
| `/api/v1/scan-results/:id` | `GET` | 获取扫描结果详情 |
| `/api/v1/scan-tasks/:taskId/results` | `GET` | 获取扫描任务的所有结果 |

## 3. 系统功能设计

### 3.1 资产发现与 inventory

- **自动发现**：通过网络扫描自动发现网络中的资产
- **手动录入**：支持手动录入资产信息
- **批量导入**：支持从CSV、Excel等文件批量导入资产
- **资产分类**：根据类型、位置、部门等对资产进行分类
- **资产标签**：支持为资产添加标签，便于管理和搜索

### 3.2 漏洞管理

- **漏洞扫描**：集成漏洞扫描工具，定期或按需扫描资产漏洞
- **CVE 集成**：自动同步 CVE 数据库，获取最新漏洞信息
- **漏洞评估**：根据 CVSS 评分和业务影响评估漏洞风险
- **漏洞修复**：跟踪漏洞修复状态，提供修复建议
- **漏洞报告**：生成漏洞报告，包括趋势分析和风险评估

### 3.3 风险评估

- **资产风险评分**：基于资产的重要性、漏洞数量和严重程度计算风险评分
- **风险趋势分析**：分析风险变化趋势，识别风险增加的资产
- **风险报告**：生成风险报告，为管理层提供决策依据

### 3.4 合规管理

- **合规检查**：检查资产是否符合安全合规要求
- **合规报告**：生成合规报告，证明组织符合相关法规要求

### 3.5 配置管理

- **配置基线**：定义资产的安全配置基线
- **配置偏差**：检测资产配置与基线的偏差
- **配置变更**：跟踪资产配置变更，识别未授权变更

### 3.6 报告与分析

- **自定义报告**：支持创建自定义报告
- **定期报告**：自动生成定期报告（如：每周、每月）
- **仪表板**：提供实时的资产和漏洞状态仪表板
- **趋势分析**：分析资产和漏洞的变化趋势

## 4. 技术栈选择

### 4.1 后端

- **语言**：Go
- **Web 框架**：Gin
- **数据库**：PostgreSQL
- **ORM**：GORM
- **认证**：JWT
- **缓存**：Redis（可选）
- **消息队列**：RabbitMQ（可选，用于扫描任务）

### 4.2 前端

- **框架**：Angular
- **UI 库**：Angular Material
- **状态管理**：NgRx 或 Signal
- **图表库**：Chart.js 或 D3.js
- **HTTP 客户端**：Axios 或 Angular HttpClient

### 4.3 安全工具集成

- **漏洞扫描**：Nessus, OpenVAS, Qualys 等
- **配置管理**：Chef, Puppet, Ansible 等
- **日志管理**：ELK Stack, Splunk 等

## 5. 部署架构

### 5.1 单服务器部署

适用于小型组织，所有组件部署在同一台服务器上：

```
┌─────────────────────────────────────────────────────┐
│                   服务器                           │
├─────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐ │
│  │  Web 应用   │  │  数据库      │  │  扫描工具   │ │
│  │  (Go + Gin) │  │  (PostgreSQL)│  │  (集成)     │ │
│  └─────────────┘  └─────────────┘  └─────────────┘ │
└─────────────────────────────────────────────────────┘
```

### 5.2 分布式部署

适用于大型组织，组件分布在多台服务器上：

```
┌────────────────────┐     ┌────────────────────┐     ┌────────────────────┐
│  Web 应用服务器    │────>│  数据库服务器      │<────│  扫描服务器        │
│  (Go + Gin)        │     │  (PostgreSQL)      │     │  (漏洞扫描工具)     │
└────────────────────┘     └────────────────────┘     └────────────────────┘
          ^                           ^                           ^
          │                           │                           │
          └────────────────────────────────────────────────────────┘
                                网络
```

## 6. 安全考虑

### 6.1 数据安全

- **加密存储**：敏感数据（如：密码、证书）加密存储
- **访问控制**：基于角色的访问控制（RBAC）
- **审计日志**：记录所有敏感操作的审计日志
- **数据备份**：定期备份数据库，确保数据可恢复

### 6.2 网络安全

- **防火墙**：配置防火墙，限制网络访问
- **HTTPS**：使用 HTTPS 加密传输数据
- **VPN**：对于远程访问，使用 VPN

### 6.3 应用安全

- **输入验证**：验证所有用户输入，防止注入攻击
- **安全头部**：设置适当的安全 HTTP 头部
- **漏洞扫描**：定期扫描应用自身的漏洞
- **依赖管理**：定期更新依赖库，修复已知漏洞

## 7. 扩展性考虑

### 7.1 模块扩展

- **插件系统**：支持通过插件扩展功能
- **API 集成**：提供 RESTful API，便于与其他系统集成
- **第三方集成**：集成第三方工具和服务

### 7.2 性能扩展

- **水平扩展**：支持通过增加服务器实例水平扩展
- **缓存**：使用缓存减少数据库负载
- **异步处理**：对于耗时操作，使用异步处理

### 7.3 功能扩展

- **云资产管理**：支持管理云服务和资源
- **IoT 设备管理**：支持管理 IoT 设备
- **供应链风险**：管理供应链中的安全风险

## 8. 实施计划

### 8.1 阶段一：基础架构搭建

1. 搭建开发环境
2. 创建数据库模型
3. 实现基础 API
4. 搭建前端框架

### 8.2 阶段二：核心功能实现

1. 资产管理功能
2. 漏洞管理功能
3. 扫描管理功能
4. 报告与分析功能

### 8.3 阶段三：集成与优化

1. 集成漏洞扫描工具
2. 优化性能
3. 增强安全性
4. 测试与调试

### 8.4 阶段四：部署与维护

1. 部署到生产环境
2. 建立监控系统
3. 制定维护计划
4. 提供培训和文档

## 9. 总结

本设计提供了一个全面的网络资产管理系统架构，包括资产发现、漏洞管理、风险评估、合规管理等功能。系统采用现代技术栈，具有良好的扩展性和安全性，能够满足组织对网络资产的全面管理需求。

通过实施此系统，组织可以：

1. **提高资产可见性**：全面了解网络中的所有资产
2. **减少安全风险**：及时发现和修复漏洞
3. **提升合规水平**：确保符合相关法规要求
4. **优化资源配置**：基于资产信息合理分配资源
5. **降低运营成本**：自动化管理流程，减少人工操作

该设计可根据组织的具体需求进行调整和扩展，以满足不同规模和行业的网络资产管理需求。