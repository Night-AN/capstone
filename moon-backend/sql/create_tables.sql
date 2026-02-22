-- 创建schema
CREATE SCHEMA IF NOT EXISTS systems;

-- 创建用户表
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

-- 创建组织表
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

-- 创建权限表
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

-- 创建角色表
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

-- 创建用户-组织关联表
CREATE TABLE IF NOT EXISTS systems.user_organization (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    user_id UUID NOT NULL REFERENCES systems.users(user_id),
    organization_id UUID NOT NULL REFERENCES systems.organization(organization_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 创建用户-角色关联表
CREATE TABLE IF NOT EXISTS systems.user_role (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    user_id UUID NOT NULL REFERENCES systems.users(user_id),
    role_id UUID NOT NULL REFERENCES systems.role(role_id)
);

-- 创建组织-角色关联表
CREATE TABLE IF NOT EXISTS systems.organization_role (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    organization_id UUID NOT NULL REFERENCES systems.organization(organization_id),
    role_id UUID NOT NULL REFERENCES systems.role(role_id)
);

-- 创建角色-权限关联表
CREATE TABLE IF NOT EXISTS systems.permission_role (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    permission_id UUID NOT NULL REFERENCES systems.permission(permission_id),
    role_id UUID NOT NULL REFERENCES systems.role(role_id)
);
