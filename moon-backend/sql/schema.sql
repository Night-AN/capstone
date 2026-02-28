-- 数据库整合文件
-- 包含所有schema、表、索引、触发器、约束和示例数据

-- 开始事务
BEGIN;

-- 设置超时
SET LOCAL statement_timeout = '5s';
SET LOCAL lock_timeout = '1s';

-- 创建schema
CREATE SCHEMA IF NOT EXISTS systems;
CREATE SCHEMA IF NOT EXISTS biz;
CREATE SCHEMA IF NOT EXISTS logs;

-- 创建触发器函数
CREATE OR REPLACE FUNCTION systems.set_updated_at() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION biz.set_updated_at() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 提交事务
COMMIT;

-- 创建系统相关表
BEGIN;

SET LOCAL statement_timeout = '5s';
SET LOCAL lock_timeout = '1s';

-- 用户表
CREATE TABLE IF NOT EXISTS systems.users (
    user_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    nickname TEXT NOT NULL DEFAULT '',
    full_name TEXT NOT NULL DEFAULT '',
    email TEXT NOT NULL DEFAULT '',
    password_hash TEXT NOT NULL DEFAULT '',
    organization_id UUID REFERENCES systems.organization(organization_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 组织表
CREATE TABLE IF NOT EXISTS systems.organization (
    organization_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    organization_name TEXT NOT NULL DEFAULT '',
    organization_code TEXT NOT NULL DEFAULT '',
    organization_description TEXT NOT NULL DEFAULT '',
    parent_id UUID NULL REFERENCES systems.organization(organization_id),
    organization_flag TEXT NOT NULL DEFAULT '',
    sensitive_flag BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 权限表
CREATE TABLE IF NOT EXISTS systems.permission (
    permission_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    permission_name TEXT NOT NULL DEFAULT '',
    permission_code TEXT NOT NULL DEFAULT '',
    permission_description TEXT NOT NULL DEFAULT '',
    parent_id UUID NULL REFERENCES systems.permission(permission_id),
    sensitive_flag BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 角色表
CREATE TABLE IF NOT EXISTS systems.role (
    role_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    role_name TEXT NOT NULL DEFAULT '',
    role_code TEXT NOT NULL DEFAULT '',
    role_description TEXT NOT NULL DEFAULT '',
    role_flag TEXT NOT NULL DEFAULT '',
    sensitive_flag BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 菜单表
CREATE TABLE IF NOT EXISTS systems.menu (
    menu_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    menu_name TEXT NOT NULL DEFAULT '',
    menu_code TEXT NOT NULL DEFAULT '',
    menu_description TEXT NOT NULL DEFAULT '',
    menu_flag TEXT NOT NULL DEFAULT '',
    menu_type TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 资源表
CREATE TABLE IF NOT EXISTS systems.resource (
    resource_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    resource_name TEXT NOT NULL DEFAULT '',
    resource_code TEXT NOT NULL DEFAULT '',
    resource_description TEXT NOT NULL DEFAULT '',
    resource_flag TEXT NOT NULL DEFAULT '',
    resource_type TEXT NOT NULL DEFAULT '',
    resource_path TEXT NOT NULL DEFAULT '',
    request_method TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 用户-组织关联表
CREATE TABLE IF NOT EXISTS systems.user_organization (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    user_id UUID NOT NULL REFERENCES systems.users(user_id),
    organization_id UUID NOT NULL REFERENCES systems.organization(organization_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 用户-角色关联表
CREATE TABLE IF NOT EXISTS systems.user_role (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    user_id UUID NOT NULL REFERENCES systems.users(user_id),
    role_id UUID NOT NULL REFERENCES systems.role(role_id)
);

-- 组织-角色关联表
CREATE TABLE IF NOT EXISTS systems.organization_role (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    organization_id UUID NOT NULL REFERENCES systems.organization(organization_id),
    role_id UUID NOT NULL REFERENCES systems.role(role_id)
);

-- 角色-权限关联表
CREATE TABLE IF NOT EXISTS systems.permission_role (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    permission_id UUID NOT NULL REFERENCES systems.permission(permission_id),
    role_id UUID NOT NULL REFERENCES systems.role(role_id)
);

-- 角色-菜单关联表
CREATE TABLE IF NOT EXISTS systems.role_menu (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    user_id UUID NOT NULL,
    menu_id UUID NOT NULL
);

-- 权限-资源关联表
CREATE TABLE IF NOT EXISTS systems.permission_resource (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    permission_id UUID NOT NULL,
    resource_id UUID NOT NULL
);

-- 提交事务
COMMIT;

-- 创建业务相关表
BEGIN;

SET LOCAL statement_timeout = '5s';
SET LOCAL lock_timeout = '1s';

-- 资产表
CREATE TABLE IF NOT EXISTS biz.asset (
    asset_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    asset_name TEXT NOT NULL DEFAULT '',
    asset_code TEXT NOT NULL DEFAULT '',
    asset_description TEXT NOT NULL DEFAULT '',
    asset_type TEXT NOT NULL DEFAULT '',
    asset_class TEXT NOT NULL DEFAULT '',
    manufacturer TEXT NOT NULL DEFAULT '',
    model TEXT NOT NULL DEFAULT '',
    serial_number TEXT NOT NULL DEFAULT '',
    ip_address TEXT NOT NULL DEFAULT '',
    mac_address TEXT NOT NULL DEFAULT '',
    location TEXT NOT NULL DEFAULT '',
    department TEXT NOT NULL DEFAULT '',
    owner TEXT NOT NULL DEFAULT '',
    contact_info TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT '',
    purchase_date TEXT NOT NULL DEFAULT '',
    warranty_end_date TEXT NOT NULL DEFAULT '',
    value TEXT NOT NULL DEFAULT '',
    notes TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 漏洞表
CREATE TABLE IF NOT EXISTS biz.vulnerability (
    vulnerability_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    cve_id TEXT NOT NULL DEFAULT '',
    nist_cve_id TEXT NOT NULL DEFAULT '',
    title TEXT NOT NULL DEFAULT '',
    description TEXT NOT NULL DEFAULT '',
    severity TEXT NOT NULL DEFAULT '',
    cvss_score FLOAT NOT NULL DEFAULT 0.0,
    cvss_vector TEXT NOT NULL DEFAULT '',
    affected_software TEXT NOT NULL DEFAULT '',
    affected_versions TEXT NOT NULL DEFAULT '',
    attack_vector TEXT NOT NULL DEFAULT '',
    attack_complexity TEXT NOT NULL DEFAULT '',
    privileges_required TEXT NOT NULL DEFAULT '',
    user_interaction TEXT NOT NULL DEFAULT '',
    scope TEXT NOT NULL DEFAULT '',
    confidentiality_impact TEXT NOT NULL DEFAULT '',
    integrity_impact TEXT NOT NULL DEFAULT '',
    availability_impact TEXT NOT NULL DEFAULT '',
    reference_urls TEXT NOT NULL DEFAULT '',
    solution TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT '',
    published_date TEXT NOT NULL DEFAULT '',
    last_modified_date TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 资产漏洞关联表
CREATE TABLE IF NOT EXISTS biz.asset_vulnerability (
    asset_vulnerability_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    asset_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    vulnerability_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    detection_date TEXT NOT NULL DEFAULT '',
    status TEXT NOT NULL DEFAULT '',
    risk_level TEXT NOT NULL DEFAULT '',
    remediation_plan TEXT NOT NULL DEFAULT '',
    assigned_to TEXT NOT NULL DEFAULT '',
    due_date TEXT NOT NULL DEFAULT '',
    notes TEXT NOT NULL DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 提交事务
COMMIT;

-- 创建系统相关索引
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_user_id ON systems.users (user_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_user_email ON systems.users (email);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_user_nickname ON systems.users (nickname);

CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_permission_id ON systems.permission (permission_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_permission_name ON systems.permission (permission_name);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_permission_code ON systems.permission (permission_code);

CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_role_id ON systems.role (role_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_role_name ON systems.role (role_name);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_role_code ON systems.role (role_code);

CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_organization_id ON systems.organization (organization_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_organization_name ON systems.organization (organization_name);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_sys_organization_code ON systems.organization (organization_code);

-- 创建业务相关索引
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_asset_id ON biz.asset (asset_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_asset_code ON biz.asset (asset_code);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_vulnerability_id ON biz.vulnerability (vulnerability_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_vulnerability_cve_id ON biz.vulnerability (cve_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_asset_vulnerability_id ON biz.asset_vulnerability (asset_vulnerability_id);

-- 添加系统相关约束和触发器
BEGIN;

SET LOCAL statement_timeout = '5s';
SET LOCAL lock_timeout = '1s';

-- 用户表
ALTER TABLE systems.users
    ADD CONSTRAINT pk_sys_user PRIMARY KEY USING INDEX idx_sys_user_id,
    ADD CONSTRAINT uk_sys_user_email UNIQUE USING INDEX idx_sys_user_email,
    ADD CONSTRAINT uk_sys_user_nickname UNIQUE USING INDEX idx_sys_user_nickname,
    ADD CONSTRAINT chk_nickname_size CHECK (LENGTH(nickname) <= 64) NOT VALID,
    ADD CONSTRAINT chk_full_name_size CHECK (LENGTH(full_name) <= 128) NOT VALID,
    ADD CONSTRAINT chk_email_size CHECK (LENGTH(email) <= 256) NOT VALID,
    ADD CONSTRAINT chk_password_hash_size CHECK (LENGTH(password_hash) <= 256) NOT VALID;

DROP TRIGGER IF EXISTS tr_sys_user_set_updated_at ON systems.users;

CREATE TRIGGER tr_sys_user_set_updated_at
    BEFORE UPDATE
    ON systems.users
    FOR EACH ROW
EXECUTE FUNCTION systems.set_updated_at();

-- 权限表
DROP TRIGGER IF EXISTS tr_sys_permission_set_updated_at ON systems.permission;

CREATE TRIGGER tr_sys_permission_set_updated_at
    BEFORE UPDATE
    ON systems.permission
    FOR EACH ROW
EXECUTE FUNCTION systems.set_updated_at();

ALTER TABLE systems.permission
    ADD CONSTRAINT pk_permission_id PRIMARY KEY USING INDEX idx_sys_permission_id,
    ADD CONSTRAINT uk_permission_name UNIQUE USING INDEX idx_sys_permission_name,
    ADD CONSTRAINT uk_permission_code UNIQUE USING INDEX idx_sys_permission_code,
    ADD CONSTRAINT chk_permission_name CHECK (LENGTH(permission_name) <= 32),
    ADD CONSTRAINT chk_permission_code CHECK (LENGTH(permission_code) <= 64),
    ADD CONSTRAINT chk_permission_description CHECK (LENGTH(permission_description) <= 1024);

-- 角色表
DROP TRIGGER IF EXISTS tr_sys_role_set_updated_at ON systems.role;

CREATE TRIGGER tr_sys_role_set_updated_at
    BEFORE UPDATE
    ON systems.role
    FOR EACH ROW
EXECUTE FUNCTION systems.set_updated_at();

ALTER TABLE systems.role
    ADD CONSTRAINT pk_role_id PRIMARY KEY USING INDEX idx_sys_role_id,
    ADD CONSTRAINT uk_role_name UNIQUE USING INDEX idx_sys_role_name,
    ADD CONSTRAINT uk_role_code UNIQUE USING INDEX idx_sys_role_code,
    ADD CONSTRAINT chk_role_name CHECK (LENGTH(role_name) <= 32),
    ADD CONSTRAINT chk_role_code CHECK (LENGTH(role_code) <= 64),
    ADD CONSTRAINT chk_role_description CHECK (LENGTH(role_description) <= 1024);

-- 组织表
DROP TRIGGER IF EXISTS tr_sys_organization_set_updated_at ON systems.organization;

CREATE TRIGGER tr_sys_organization_set_updated_at
    BEFORE UPDATE
    ON systems.organization
    FOR EACH ROW
EXECUTE FUNCTION systems.set_updated_at();

ALTER TABLE systems.organization
    ADD CONSTRAINT pk_organization_id PRIMARY KEY USING INDEX idx_sys_organization_id,
    ADD CONSTRAINT uk_organization_name UNIQUE USING INDEX idx_sys_organization_name,
    ADD CONSTRAINT uk_organization_code UNIQUE USING INDEX idx_sys_organization_code,
    ADD CONSTRAINT chk_organization_name CHECK (LENGTH(organization_name) <= 32),
    ADD CONSTRAINT chk_organization_code CHECK (LENGTH(organization_code) <= 64),
    ADD CONSTRAINT chk_organization_description CHECK (LENGTH(organization_description) <= 1024);

-- 组织-角色关联表约束
ALTER TABLE systems.organization_role
    ADD CONSTRAINT pk_organization_role_id PRIMARY KEY (id),
    ADD CONSTRAINT fk_organization_role_organization_id FOREIGN KEY (organization_id) REFERENCES systems.organization(organization_id),
    ADD CONSTRAINT fk_organization_role_role_id FOREIGN KEY (role_id) REFERENCES systems.role(role_id);

-- 角色-权限关联表约束
ALTER TABLE systems.permission_role
    ADD CONSTRAINT pk_permission_role_id PRIMARY KEY (id),
    ADD CONSTRAINT fk_permission_role_permission_id FOREIGN KEY (permission_id) REFERENCES systems.permission(permission_id),
    ADD CONSTRAINT fk_permission_role_role_id FOREIGN KEY (role_id) REFERENCES systems.role(role_id);

-- 权限-资源关联表约束
ALTER TABLE systems.permission_resource
    ADD CONSTRAINT pk_permission_resource_id PRIMARY KEY (id),
    ADD CONSTRAINT fk_permission_resource_permission_id FOREIGN KEY (permission_id) REFERENCES systems.permission(permission_id),
    ADD CONSTRAINT fk_permission_resource_resource_id FOREIGN KEY (resource_id) REFERENCES systems.resource(resource_id);

-- 提交事务
COMMIT;

-- 添加业务相关约束和触发器
BEGIN;

SET LOCAL statement_timeout = '5s';
SET LOCAL lock_timeout = '1s';

-- 资产表
DROP TRIGGER IF EXISTS tr_biz_asset_set_updated_at ON biz.asset;
CREATE TRIGGER tr_biz_asset_set_updated_at
    BEFORE UPDATE
    ON biz.asset
    FOR EACH ROW
EXECUTE FUNCTION biz.set_updated_at();

ALTER TABLE biz.asset
    ADD CONSTRAINT pk_biz_asset_id PRIMARY KEY USING INDEX idx_biz_asset_id,
    ADD CONSTRAINT uk_biz_asset_code UNIQUE USING INDEX idx_biz_asset_code,
    ADD CONSTRAINT chk_biz_asset_name CHECK (LENGTH(asset_name) <= 128),
    ADD CONSTRAINT chk_biz_asset_code CHECK (LENGTH(asset_code) <= 64),
    ADD CONSTRAINT chk_biz_asset_description CHECK (LENGTH(asset_description) <= 1024);

-- 漏洞表
DROP TRIGGER IF EXISTS tr_biz_vulnerability_set_updated_at ON biz.vulnerability;
CREATE TRIGGER tr_biz_vulnerability_set_updated_at
    BEFORE UPDATE
    ON biz.vulnerability
    FOR EACH ROW
EXECUTE FUNCTION biz.set_updated_at();

ALTER TABLE biz.vulnerability
    ADD CONSTRAINT pk_biz_vulnerability_id PRIMARY KEY USING INDEX idx_biz_vulnerability_id,
    ADD CONSTRAINT uk_biz_vulnerability_cve_id UNIQUE USING INDEX idx_biz_vulnerability_cve_id,
    ADD CONSTRAINT chk_biz_vulnerability_title CHECK (LENGTH(title) <= 256),
    ADD CONSTRAINT chk_biz_vulnerability_cve_id CHECK (LENGTH(cve_id) <= 64);

-- 资产漏洞关联表
DROP TRIGGER IF EXISTS tr_biz_asset_vulnerability_set_updated_at ON biz.asset_vulnerability;
CREATE TRIGGER tr_biz_asset_vulnerability_set_updated_at
    BEFORE UPDATE
    ON biz.asset_vulnerability
    FOR EACH ROW
EXECUTE FUNCTION biz.set_updated_at();

ALTER TABLE biz.asset_vulnerability
    ADD CONSTRAINT pk_biz_asset_vulnerability_id PRIMARY KEY USING INDEX idx_biz_asset_vulnerability_id,
    ADD CONSTRAINT fk_biz_asset_vulnerability_asset_id FOREIGN KEY (asset_id) REFERENCES biz.asset(asset_id),
    ADD CONSTRAINT fk_biz_asset_vulnerability_vulnerability_id FOREIGN KEY (vulnerability_id) REFERENCES biz.vulnerability(vulnerability_id);

-- 为 capstone 用户授予权限
GRANT ALL PRIVILEGES ON SCHEMA biz TO capstone;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA biz TO capstone;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA biz TO capstone;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA biz TO capstone;

-- 提交事务
COMMIT;

-- 创建 AI 相关表
BEGIN;

SET LOCAL statement_timeout = '5s';
SET LOCAL lock_timeout = '1s';

-- 创建 AI schema 触发器函数
CREATE OR REPLACE FUNCTION ai.set_updated_at() RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 表 4: ai.model_config 模型配置表
CREATE TABLE IF NOT EXISTS ai.model_config (
    config_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    provider_name TEXT NOT NULL,
    model_name TEXT NOT NULL,
    api_key TEXT NOT NULL,
    api_endpoint TEXT,
    api_version TEXT,
    max_tokens INTEGER NOT NULL DEFAULT 4096,
    temperature FLOAT NOT NULL DEFAULT 0.7,
    timeout_seconds INTEGER NOT NULL DEFAULT 30,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    priority INTEGER NOT NULL DEFAULT 1,
    config_metadata JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 表 5: ai.api_call_log AI调用日志
CREATE TABLE IF NOT EXISTS ai.api_call_log (
    log_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    config_id UUID,
    call_type TEXT NOT NULL,
    prompt_tokens INTEGER,
    completion_tokens INTEGER,
    total_tokens INTEGER,
    request_payload JSONB,
    response_payload JSONB,
    status_code INTEGER,
    error_message TEXT,
    latency_ms INTEGER,
    success BOOLEAN NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 表 6: ai.asset_classification 资产分类表
CREATE TABLE IF NOT EXISTS ai.asset_classification (
    classification_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    asset_id UUID NOT NULL,
    log_id UUID NOT NULL,
    predicted_category TEXT NOT NULL,
    confidence FLOAT,
    reasoning TEXT,
    manual_category TEXT,
    is_approved BOOLEAN NOT NULL DEFAULT FALSE,
    approved_by UUID,
    approved_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 表 7: ai.risk_assessment AI风险评估表
CREATE TABLE IF NOT EXISTS ai.risk_assessment (
    assessment_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    vulnerability_id UUID NOT NULL,
    asset_id UUID,
    log_id UUID NOT NULL,
    risk_score FLOAT NOT NULL,
    risk_level TEXT NOT NULL,
    analysis TEXT,
    factor_weights JSONB,
    provider TEXT NOT NULL,
    model TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 表 8: ai.security_recommendation AI安全建议表
CREATE TABLE IF NOT EXISTS ai.security_recommendation (
    recommendation_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    vulnerability_id UUID NOT NULL,
    log_id UUID NOT NULL,
    summary TEXT NOT NULL,
    vulnerability_analysis TEXT,
    remediation_steps JSONB,
    recommended_patches JSONB,
    mitigation_measures JSONB,
    prevention_tips JSONB,
    references JSONB,
    provider TEXT NOT NULL,
    model TEXT NOT NULL,
    is_useful BOOLEAN,
    feedback TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 表 9: ai.prompt_template AI提示词模板表
CREATE TABLE IF NOT EXISTS ai.prompt_template (
    template_id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    template_name TEXT NOT NULL UNIQUE,
    template_type TEXT NOT NULL,
    template_content TEXT NOT NULL,
    variables JSONB,
    description TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    version INTEGER NOT NULL DEFAULT 1,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 创建索引
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_model_config_id ON ai.model_config (config_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_model_config_active ON ai.model_config (is_active) WHERE is_active = TRUE;

CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_api_call_log_id ON ai.api_call_log (log_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_api_call_log_type ON ai.api_call_log (call_type);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_api_call_log_created ON ai.api_call_log (created_at DESC);

CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_asset_classification_id ON ai.asset_classification (classification_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_asset_classification_asset_id ON ai.asset_classification (asset_id);

CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_risk_assessment_id ON ai.risk_assessment (assessment_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_risk_assessment_vuln_id ON ai.risk_assessment (vulnerability_id);

CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_security_recommendation_id ON ai.security_recommendation (recommendation_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_security_recommendation_vuln_id ON ai.security_recommendation (vulnerability_id);

CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_prompt_template_id ON ai.prompt_template (template_id);
CREATE INDEX CONCURRENTLY IF NOT EXISTS idx_ai_prompt_template_type ON ai.prompt_template (template_type);

-- 添加触发器
DROP TRIGGER IF EXISTS tr_ai_model_config_set_updated_at ON ai.model_config;
CREATE TRIGGER tr_ai_model_config_set_updated_at
    BEFORE UPDATE ON ai.model_config
    FOR EACH ROW EXECUTE FUNCTION ai.set_updated_at();

DROP TRIGGER IF EXISTS tr_ai_asset_classification_set_updated_at ON ai.asset_classification;
CREATE TRIGGER tr_ai_asset_classification_set_updated_at
    BEFORE UPDATE ON ai.asset_classification
    FOR EACH ROW EXECUTE FUNCTION ai.set_updated_at();

DROP TRIGGER IF EXISTS tr_ai_prompt_template_set_updated_at ON ai.prompt_template;
CREATE TRIGGER tr_ai_prompt_template_set_updated_at
    BEFORE UPDATE ON ai.prompt_template
    FOR EACH ROW EXECUTE FUNCTION ai.set_updated_at();

-- 添加主键和外键约束
ALTER TABLE ai.model_config
    ADD CONSTRAINT pk_ai_model_config_id PRIMARY KEY USING INDEX idx_ai_model_config_id;

ALTER TABLE ai.api_call_log
    ADD CONSTRAINT pk_ai_api_call_log_id PRIMARY KEY USING INDEX idx_ai_api_call_log_id,
    ADD CONSTRAINT fk_ai_api_call_log_config_id FOREIGN KEY (config_id) REFERENCES ai.model_config(config_id);

ALTER TABLE ai.asset_classification
    ADD CONSTRAINT pk_ai_asset_classification_id PRIMARY KEY USING INDEX idx_ai_asset_classification_id,
    ADD CONSTRAINT fk_ai_asset_classification_asset_id FOREIGN KEY (asset_id) REFERENCES biz.asset(asset_id),
    ADD CONSTRAINT fk_ai_asset_classification_log_id FOREIGN KEY (log_id) REFERENCES ai.api_call_log(log_id),
    ADD CONSTRAINT fk_ai_asset_classification_approved_by FOREIGN KEY (approved_by) REFERENCES systems.users(user_id);

ALTER TABLE ai.risk_assessment
    ADD CONSTRAINT pk_ai_risk_assessment_id PRIMARY KEY USING INDEX idx_ai_risk_assessment_id,
    ADD CONSTRAINT fk_ai_risk_assessment_vuln_id FOREIGN KEY (vulnerability_id) REFERENCES biz.vulnerability(vulnerability_id),
    ADD CONSTRAINT fk_ai_risk_assessment_asset_id FOREIGN KEY (asset_id) REFERENCES biz.asset(asset_id),
    ADD CONSTRAINT fk_ai_risk_assessment_log_id FOREIGN KEY (log_id) REFERENCES ai.api_call_log(log_id);

ALTER TABLE ai.security_recommendation
    ADD CONSTRAINT pk_ai_security_recommendation_id PRIMARY KEY USING INDEX idx_ai_security_recommendation_id,
    ADD CONSTRAINT fk_ai_security_recommendation_vuln_id FOREIGN KEY (vulnerability_id) REFERENCES biz.vulnerability(vulnerability_id),
    ADD CONSTRAINT fk_ai_security_recommendation_log_id FOREIGN KEY (log_id) REFERENCES ai.api_call_log(log_id);

ALTER TABLE ai.prompt_template
    ADD CONSTRAINT pk_ai_prompt_template_id PRIMARY KEY USING INDEX idx_ai_prompt_template_id;

-- 为 capstone 用户授予权限
GRANT ALL PRIVILEGES ON SCHEMA ai TO capstone;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA ai TO capstone;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA ai TO capstone;
GRANT ALL PRIVILEGES ON ALL FUNCTIONS IN SCHEMA ai TO capstone;

-- 提交事务
COMMIT;

-- 插入系统初始数据
BEGIN;

SET LOCAL statement_timeout = '5s';
SET LOCAL lock_timeout = '1s';

-- 插入角色数据
INSERT INTO systems.role (role_name, role_code, role_description, role_flag)
VALUES ('SuperAdmin', 'sa', 'This is a superadmin role', 'active'),
       ('OrganizationAdmin', 'org-admin', 'this is ad admin role', 'active'),
       ('AssetManager', 'am', 'this is a asset-manage role', 'active'),
       ('ThirdParty', 'tp', 'this is a ThirdParty role', 'active');

-- 插入权限数据
INSERT INTO systems.permission (permission_name, permission_code, permission_description)
VALUES ('UserRead', 'user:read', 'user read'),
       ('UserWrite', 'user:write', 'user write'),
       ('UserDelete', 'user:delete', 'user delete'),
       ('RoleRead', 'role:read', 'role read'),
       ('RoleWrite', 'role:write', 'role write'),
       ('RoleDelete', 'role:delete', 'role delete'),
       ('PermissionRead', 'permission:read', 'permission read'),
       ('PermissionWrite', 'permission:write', 'permission write'),
       ('PermissionDelete', 'permission:delete', 'permission delete'),
       ('OrganizationRead', 'organization:read', 'organization read'),
       ('OrganizationWrite', 'organization:write', 'organization write'),
       ('OrganizationDelete', 'organization:delete', 'organization delete'),
       ('MenuRead', 'menu:read', 'menu read'),
       ('MenuWrite', 'menu:write', 'menu write'),
       ('MenuDelete', 'menu:delete', 'menu delete'),
       ('ResourceRead', 'resource:read', 'resource read'),
       ('ResourceWrite', 'resource:write', 'resource write'),
       ('ResourceDelete', 'resource:delete', 'resource delete');

-- 插入菜单数据
INSERT INTO systems.menu (menu_name, menu_code, menu_description)
VALUES ('UserManage', 'menu:user_manage', 'User Manager'),
       ('RoleManage', 'menu:role_manage', 'Role Manager'),
       ('PermissionManage', 'menu:perm_manage', 'Permission Manager'),
       ('OrganizationManage', 'menu:org_manage', 'Organization Manager'),
       ('ResourceManage', 'menu:resource_manage', 'Resource Manager');

-- 插入资源数据
INSERT INTO systems.resource (resource_name, resource_code, resource_description, resource_type, resource_path, request_method)
VALUES ('ReadUser', 'user::read', 'Read A user profile', 'HTTP', '/api/v1/users/:id', 'GET'),
       ('CreateUser', 'user::write', 'Create A user', 'HTTP', '/api/v1/users', 'POST'),
       ('UpdateUser', 'user::write', 'Update a user profile', 'HTTP', '/api/v1/users', 'PUT'),
       ('DeleteUser', 'user::delete', 'Delete a user profile', 'HTTP', '/api/v1/users', 'DELETE'),
       ('ListAllUser', 'user::read', 'ReadAllUserProfile', 'HTTP', '/api/v1/users/list', 'GET'),
       ('ReadRole', 'role::read', 'Read A role profile', 'HTTP', '/api/v1/roles', 'GET'),
       ('CreateRole', 'role::write', 'Create A role', 'HTTP', '/api/v1/roles', 'POST'),
       ('UpdateRole', 'role::write', 'Update a role profile', 'HTTP', '/api/v1/roles', 'PUT'),
       ('DeleteRole', 'role::delete', 'Delete a role profile', 'HTTP', '/api/v1/roles', 'DELETE'),
       ('ListAllRole', 'role::read', 'ReadAllRoleProfile', 'HTTP', '/api/v1/roles/list', 'GET'),
       ('ReadPermission', 'permission::read', 'Read A permission profile', 'HTTP', '/api/v1/permissions/:id', 'GET'),
       ('CreatePermission', 'permission::write', 'Create A permission', 'HTTP', '/api/v1/permissions', 'POST'),
       ('UpdatePermission', 'permission::write', 'Update a permission profile', 'HTTP', '/api/v1/permissions', 'PUT'),
       ('DeletePermission', 'permission::delete', 'Delete a permission profile', 'HTTP', '/api/v1/permissions', 'DELETE'),
       ('ListAllPermission', 'permission::read', 'ReadAllPermissionProfile', 'HTTP', '/api/v1/permissions/list', 'GET'),
       ('ReadOrganization', 'organization::read', 'Read A organization profile', 'HTTP', '/api/v1/organizations', 'GET'),
       ('CreateOrganization', 'organization::write', 'Create A organization', 'HTTP', '/api/v1/organizations', 'POST'),
       ('UpdateOrganization', 'organization::write', 'Update a organization profile', 'HTTP', '/api/v1/organizations', 'PUT'),
       ('DeleteOrganization', 'organization::delete', 'Delete a organization profile', 'HTTP', '/api/v1/organizations', 'DELETE'),
       ('ListAllOrganization', 'organization::read', 'ReadAllOrganizationProfile', 'HTTP', '/api/v1/organizations/list', 'GET'),
       ('ReadResource', 'resource::read', 'Read A resource profile', 'HTTP', '/api/v1/resources', 'GET'),
       ('CreateResource', 'resource::write', 'Create A resource', 'HTTP', '/api/v1/resources', 'POST'),
       ('UpdateResource', 'resource::write', 'Update a resource profile', 'HTTP', '/api/v1/resources', 'PUT'),
       ('DeleteResource', 'resource::delete', 'Delete a resource profile', 'HTTP', '/api/v1/resources', 'DELETE'),
       ('ListAllResource', 'resource::read', 'ReadAllResourceProfile', 'HTTP', '/api/v1/resources/list', 'GET');

-- 提交事务
COMMIT;

-- 插入业务示例数据
BEGIN;

SET LOCAL statement_timeout = '5s';
SET LOCAL lock_timeout = '1s';

-- 插入示例资产数据
INSERT INTO biz.asset (asset_id, asset_name, asset_code, asset_description, asset_type, asset_class, manufacturer, model, serial_number, ip_address, mac_address, location, department, owner, contact_info, status, purchase_date, warranty_end_date, value, notes)
VALUES 
  (GEN_RANDOM_UUID(), 'Web Server 1', 'WEB-001', 'Main web server for the company website', 'server', 'hardware', 'Dell', 'PowerEdge R740', 'S123456789', '192.168.1.100', '00:11:22:33:44:55', 'Data Center', 'IT Department', 'John Doe', 'john.doe@example.com', 'active', '2023-01-15', '2026-01-14', '5000', 'Production web server'),
  (GEN_RANDOM_UUID(), 'Database Server', 'DB-001', 'Main database server', 'server', 'hardware', 'HP', 'ProLiant DL380', 'S987654321', '192.168.1.101', '00:11:22:33:44:56', 'Data Center', 'IT Department', 'Jane Smith', 'jane.smith@example.com', 'active', '2023-02-20', '2026-02-19', '8000', 'Production database server'),
  (GEN_RANDOM_UUID(), 'Workstation 1', 'WS-001', 'Developer workstation', 'workstation', 'hardware', 'Lenovo', 'ThinkCentre M70t', 'S135792468', '192.168.1.200', '00:11:22:33:44:57', 'Office', 'Development', 'Bob Johnson', 'bob.johnson@example.com', 'active', '2023-03-10', '2026-03-09', '2000', 'Developer workstation'),
  (GEN_RANDOM_UUID(), 'Network Switch', 'NW-001', 'Core network switch', 'network_device', 'hardware', 'Cisco', 'Catalyst 9300', 'S246813579', '192.168.1.1', '00:11:22:33:44:58', 'Data Center', 'IT Department', 'Alice Brown', 'alice.brown@example.com', 'active', '2023-04-05', '2026-04-04', '3000', 'Core network switch'),
  (GEN_RANDOM_UUID(), 'Firewall', 'FW-001', 'Corporate firewall', 'network_device', 'hardware', 'Palo Alto', 'PA-220', 'S975310864', '192.168.1.2', '00:11:22:33:44:59', 'Data Center', 'IT Department', 'Charlie Davis', 'charlie.davis@example.com', 'active', '2023-05-12', '2026-05-11', '4000', 'Corporate firewall');

-- 插入示例漏洞数据
INSERT INTO biz.vulnerability (cve_id, nist_cve_id, title, description, severity, cvss_score, cvss_vector, affected_software, affected_versions, attack_vector, attack_complexity, privileges_required, user_interaction, scope, confidentiality_impact, integrity_impact, availability_impact, reference_urls, solution, status, published_date, last_modified_date)
VALUES 
  ('CVE-2023-21706', 'CVE-2023-21706', 'Windows Kerberos Elevation of Privilege Vulnerability', 'An elevation of privilege vulnerability exists in Windows Kerberos.', 'Critical', 9.8, 'CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H', 'Windows Server 2019, Windows Server 2022', '10.0.17763, 10.0.20348', 'Network', 'Low', 'None', 'None', 'Unchanged', 'High', 'High', 'High', 'https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-21706', 'Apply security update KB5022282', 'open', '2023-01-10', '2023-01-10'),
  ('CVE-2023-22515', 'CVE-2023-22515', 'Atlassian Confluence Server and Data Center Remote Code Execution Vulnerability', 'A remote code execution vulnerability exists in Atlassian Confluence Server and Data Center.', 'Critical', 10.0, 'CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H', 'Atlassian Confluence Server, Atlassian Confluence Data Center', '7.18.0 - 7.19.16, 7.20.0 - 7.20.12, 7.21.0 - 7.21.8, 7.22.0 - 7.22.3', 'Network', 'Low', 'None', 'None', 'Unchanged', 'High', 'High', 'High', 'https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-22515', 'Update to Confluence 7.19.17, 7.20.13, 7.21.9, or 7.22.4', 'open', '2023-03-23', '2023-03-23'),
  ('CVE-2023-1389', 'CVE-2023-1389', 'Microsoft Exchange Server Remote Code Execution Vulnerability', 'A remote code execution vulnerability exists in Microsoft Exchange Server.', 'Critical', 9.8, 'CVSS:3.1/AV:N/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H', 'Microsoft Exchange Server 2013, Microsoft Exchange Server 2016, Microsoft Exchange Server 2019', '15.0.1497.0 - 15.0.1497.32, 15.1.2308.0 - 15.1.2308.24, 15.2.986.0 - 15.2.986.22', 'Network', 'Low', 'None', 'None', 'Unchanged', 'High', 'High', 'High', 'https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-1389', 'Apply security update KB5024204', 'open', '2023-04-11', '2023-04-11'),
  ('CVE-2023-20198', 'CVE-2023-20198', 'Cisco IOS XE Software Web UI Privilege Escalation Vulnerability', 'A privilege escalation vulnerability exists in the web UI feature of Cisco IOS XE Software.', 'High', 8.8, 'CVSS:3.1/AV:A/AC:L/PR:N/UI:N/S:U/C:H/I:H/A:H', 'Cisco IOS XE Software', '17.6.1 - 17.6.3, 17.7.1 - 17.7.2, 17.8.1', 'Adjacent', 'Low', 'None', 'None', 'Unchanged', 'High', 'High', 'High', 'https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-20198', 'Update to Cisco IOS XE Software 17.6.4, 17.7.3, or 17.8.2', 'open', '2023-07-12', '2023-07-12'),
  ('CVE-2023-3519', 'CVE-2023-3519', 'Microsoft Windows Secure Boot Security Feature Bypass Vulnerability', 'A security feature bypass vulnerability exists in Microsoft Windows Secure Boot.', 'High', 8.2, 'CVSS:3.1/AV:L/AC:L/PR:N/UI:N/S:U/C:H/I:N/A:H', 'Microsoft Windows 10, Microsoft Windows 11, Microsoft Windows Server 2019, Microsoft Windows Server 2022', '10.0.19044, 10.0.19045, 10.0.20348', 'Local', 'Low', 'None', 'None', 'Unchanged', 'High', 'None', 'High', 'https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2023-3519', 'Apply security update KB5028185', 'open', '2023-07-11', '2023-07-11');

-- 插入示例资产漏洞关联数据
INSERT INTO biz.asset_vulnerability (asset_id, vulnerability_id, detection_date, status, risk_level, remediation_plan, assigned_to, due_date, notes)
VALUES 
  ((SELECT asset_id FROM biz.asset WHERE asset_code = 'WEB-001'), 
   (SELECT vulnerability_id FROM biz.vulnerability WHERE cve_id = 'CVE-2023-21706'), 
   '2023-01-15', 'open', 'high', 'Apply security update KB5022282', 'John Doe', '2023-01-22', 'Critical vulnerability, needs immediate attention'),
  ((SELECT asset_id FROM biz.asset WHERE asset_code = 'DB-001'), 
   (SELECT vulnerability_id FROM biz.vulnerability WHERE cve_id = 'CVE-2023-21706'), 
   '2023-01-15', 'open', 'high', 'Apply security update KB5022282', 'Jane Smith', '2023-01-22', 'Critical vulnerability, needs immediate attention'),
  ((SELECT asset_id FROM biz.asset WHERE asset_code = 'WEB-001'), 
   (SELECT vulnerability_id FROM biz.vulnerability WHERE cve_id = 'CVE-2023-1389'), 
   '2023-04-12', 'open', 'high', 'Apply security update KB5024204', 'John Doe', '2023-04-19', 'Critical vulnerability, needs immediate attention'),
  ((SELECT asset_id FROM biz.asset WHERE asset_code = 'NW-001'), 
   (SELECT vulnerability_id FROM biz.vulnerability WHERE cve_id = 'CVE-2023-20198'), 
   '2023-07-15', 'open', 'medium', 'Update to Cisco IOS XE Software 17.6.4', 'Alice Brown', '2023-07-22', 'High severity vulnerability, needs attention'),
  ((SELECT asset_id FROM biz.asset WHERE asset_code = 'WS-001'), 
   (SELECT vulnerability_id FROM biz.vulnerability WHERE cve_id = 'CVE-2023-3519'), 
   '2023-07-12', 'open', 'medium', 'Apply security update KB5028185', 'Bob Johnson', '2023-07-19', 'High severity vulnerability, needs attention');

-- 提交事务
COMMIT;