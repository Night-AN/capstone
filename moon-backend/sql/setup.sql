-- 0x0001 Setup
BEGIN;

SET
    LOCAL statement_timeout = '5s';

SET
    LOCAL lock_timeout = '1s';

CREATE SCHEMA IF NOT EXISTS systems;

CREATE SCHEMA IF NOT EXISTS business;

CREATE SCHEMA IF NOT EXISTS logs;

CREATE
    OR REPLACE FUNCTION systems.set_updated_at() RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at
        = NOW();
    RETURN NEW;
END;
$$
    LANGUAGE plpgsql;

COMMIT;

BEGIN;

SET
    LOCAL statement_timeout = '5s';

SET
    LOCAL lock_timeout = '1s';

CREATE TABLE IF NOT EXISTS systems.users
(
    user_id       UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    nickname      TEXT        NOT NULL DEFAULT '',
    full_name     TEXT        NOT NULL DEFAULT '',
    email         TEXT        NOT NULL DEFAULT '',
    password_hash TEXT        NOT NULL DEFAULT '',
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS systems.permission
(
    permission_id          UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    permission_name        TEXT        NOT NULL DEFAULT '',
    permission_code        TEXT        NOT NULL DEFAULT '',
    permission_description TEXT        NOT NULL DEFAULT '',
    sensitive_flag         BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS systems.role
(
    role_id          UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    role_name        TEXT        NOT NULL DEFAULT '',
    role_code        TEXT        NOT NULL DEFAULT '',
    role_description TEXT        NOT NULL DEFAULT '',
    role_flag        TEXT        NOT NULL DEFAULT '',
    sensitive_flag   BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS systems.organization
(
    organization_id          UUID        NOT NULL DEFAULT  GEN_RANDOM_UUID(),
    organization_name        TEXT        NOT NULL DEFAULT '',
    organization_code        TEXT        NOT NULL DEFAULT '',
    organization_description TEXT        NOT NULL DEFAULT '',
    organization_flag        TEXT        NOT NULL DEFAULT '',
    sensitive_flag           BOOLEAN     NOT NULL DEFAULT FALSE,
    created_at               TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at               TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS systems.menu
(
    menu_id          UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    menu_name        TEXT        NOT NULL DEFAULT '',
    menu_code        TEXT        NOT NULL DEFAULT '',
    menu_description TEXT        NOT NULL DEFAULT '',
    menu_flag        TEXT        NOT NULL DEFAULT '',
    menu_type        TEXT        NOT NULL DEFAULT '',
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS systems.resource
(
    resource_id          UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    resource_name        TEXT        NOT NULL DEFAULT '',
    resource_code        TEXT        NOT NULL DEFAULT '',
    resource_description TEXT        NOT NULL DEFAULT '',
    resource_flag        TEXT        NOT NULL DEFAULT '',
    resource_type        TEXT        NOT NULL DEFAULT '',
    resource_path        TEXT        NOT NULL DEFAULT '',
    request_method       TEXT        NOT NULL DEFAULT '',
    created_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS
    systems.permission_role
(
    id            UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    permission_id UUID NOT NULL REFERENCES permission(permission_id),
    role_id       UUID NOT NULL REFERENCES role(role_id)
);

CREATE TABLE IF NOT EXISTS
    systems.user_role
(
    id      UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    user_id UUID NOT NULL REFERENCES  users(user_id),
    role_id UUID NOT NULL REFERENCES  role(role_id)
);

CREATE TABLE IF NOT EXISTS
    systems.role_menu
(
    id      UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    user_id UUID NOT NULL,
    menu_id UUID NOT NULL
);

CREATE TABLE IF NOT EXISTS
    systems.permission_resource
(
    id            UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    permission_id UUID NOT NULL,
    resource_id   UUID NOT NULL
);

COMMIT;

-- 0x0002 Index
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

-- 0x0003 Alter
-- 1.DefaultValue And Trigger
-- 2.Constraint
-- 3.Validate
BEGIN;

SET LOCAL statement_timeout = '5s';

SET LOCAL lock_timeout = '1s';

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
    BEFORE
        UPDATE
    ON systems.users
    FOR EACH ROW
EXECUTE FUNCTION systems.set_updated_at();

COMMIT;

BEGIN;

SET LOCAL statement_timeout = '5s';

SET LOCAL lock_timeout = '1s';


DROP TRIGGER IF EXISTS tr_sys_permission_set_updated_at ON systems.permission;

CREATE TRIGGER tr_sys_permission_set_updated_at
    BEFORE
        UPDATE
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

COMMIT;

BEGIN;

SET
    LOCAL statement_timeout = '5s';

SET
    LOCAL lock_timeout = '1s';

DROP TRIGGER IF EXISTS tr_sys_role_set_updated_at ON systems.role;

CREATE TRIGGER tr_sys_role_set_updated_at
    BEFORE
        UPDATE
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

COMMIT;

BEGIN;

SET
    LOCAL statement_timeout = '5s';

SET
    LOCAL lock_timeout = '1s';

DROP TRIGGER IF EXISTS tr_sys_organization_set_updated_at ON systems.role;

CREATE TRIGGER tr_sys_organization_set_updated_at
    BEFORE
        UPDATE
    ON systems.role
    FOR EACH ROW
EXECUTE FUNCTION systems.set_updated_at();

ALTER TABLE systems.organization
    ADD CONSTRAINT pk_organization_id PRIMARY KEY USING INDEX idx_sys_organization_id,
    ADD CONSTRAINT uk_organization_name UNIQUE USING INDEX idx_sys_organization_name,
    ADD CONSTRAINT uk_organization_code UNIQUE USING INDEX idx_sys_organization_code,
    ADD CONSTRAINT chk_organization_name CHECK (LENGTH(organization_name) <= 32),
    ADD CONSTRAINT chk_organization_code CHECK (LENGTH(organization_code) <= 64),
    ADD CONSTRAINT chk_organization_description CHECK (LENGTH(organization_description) <= 1024);
COMMIT;



INSERT INTO systems.role (role_name, role_code, role_description, role_flag)
VALUES ('SuperAdmin',
        'sa',
        'This is a superadmin role',
        'active'),
       ('OrganizationAdmin',
        'org-admin',
        'this is ad admin role',
        'active'),
       ('AssetManager',
        'am',
        'this is a asset-manage role',
        'active'),
       ('ThirdParty',
        'tp',
        'this is a ThirdParty role',
        'active');

INSERT INTO systems.permission (permission_name,
                                permission_code,
                                permission_description)
VALUES ('UserRead', 'user:read', 'user read'),
       ('UserWrite', 'user:write', 'user write'),
       ('UserDelete', 'user:delete', 'user delete'),
       ('RoleRead', 'role:read', 'role read'),
       ('RoleWrite', 'role:write', 'role write'),
       ('RoleDelete', 'role:delete', 'role delete'),
       ('PermissionRead', 'permission:read', 'permission read'),
       ('PermissionWrite',
        'permission:write',
        'permission write'),
       ('PermissionDelete',
        'permission:delete',
        'permission delete'),
       ('OrganizationRead',
        'organization:read',
        'organization read'),
       ('OrganizationWrite',
        'organization:write',
        'organization write'),
       ('OrganizationDelete',
        'organization:delete',
        'organization delete'),
       ('MenuRead', 'menu:read', 'menu read'),
       ('MenuWrite', 'menu:write', 'menu write'),
       ('MenuDelete', 'menu:delete', 'menu delete'),
       ('ResourceRead', 'resource:read', 'resource read'),
       ('ResourceWrite',
        'resource:write',
        'resource write'),
       ('ResourceDelete',
        'resource:delete',
        'resource delete');

INSERT INTO systems.menu (menu_name, menu_code, menu_description)
VALUES ('UserManage', 'menu:user_manage', 'User Manager'),
       ('RoleManage', 'menu:role_manage', 'Role Manager'),
       ('PermissionManage',
        'menu:perm_manage',
        'Permission Manager'),
       ('OrganizationManage',
        'menu:org_manage',
        'Organization Manager'),
       ('ResourceManage',
        'menu:resource_manage',
        'Resource Manager');

INSERT INTO systems.resource (resource_name,
                              resource_code,
                              resource_description,
                              resource_type,
                              resource_path,
                              request_method)
VALUES ('ReadUser',
        'user::read',
        'Read A user profile',
        'HTTP',
        '/api/v1/users/:id', 'GET'),
       ('CreateUser',
        'user::write',
        'Create A user',
        'HTTP',
        '/api/v1/users', 'POST'),
       ('UpdateUser',
        'user::write',
        'Update a user profile',
        'HTTP',
        '/api/v1/users', 'PUT'),
       ('DeleteUser',
        'user::delete',
        'Delete a user profile',
        'HTTP',
        '/api/v1/users', 'DELETE'),
       ('ListAllUser',
        'user::read',
        'ReadAllUserProfile',
        'HTTP',
        '/api/v1/users/list', 'GET'),
       ('ReadRole',
        'role::read',
        'Read A role profile',
        'HTTP',
        '/api/v1/roles', 'GET'),
       ('CreateRole',
        'role::write',
        'Create A role',
        'HTTP',
        '/api/v1/roles', 'POST'),
       ('UpdateRole',
        'role::write',
        'Update a role profile',
        'HTTP',
        '/api/v1/roles', 'PUT'),
       ('DeleteRole',
        'role::delete',
        'Delete a role profile',
        'HTTP',
        '/api/v1/roles', 'DELETE'),
       ('ListAllRole',
        'role::read',
        'ReadAllRoleProfile',
        'HTTP',
        '/api/v1/roles/list', 'GET'),
       ('ReadPermission',
        'permission::read',
        'Read A permission profile',
        'HTTP',
        '/api/v1/permissions/:id', 'GET'),
       ('CreatePermission',
        'permission::write',
        'Create A permission',
        'HTTP',
        '/api/v1/permissions', 'POST'),
       ('UpdatePermission',
        'permission::write',
        'Update a permission profile',
        'HTTP',
        '/api/v1/permissions', 'PUT'),
       ('DeletePermission',
        'permission::delete',
        'Delete a permission profile',
        'HTTP',
        '/api/v1/permissions', 'DELETE'),
       ('ListAllPermission',
        'permission::read',
        'ReadAllPermissionProfile',
        'HTTP',
        '/api/v1/permissions/list', 'GET'),
       ('ReadOrganization',
        'organization::read',
        'Read A organization profile',
        'HTTP',
        '/api/v1/organizations', 'GET'),
       ('CreateOrganization',
        'organization::write',
        'Create A organization',
        'HTTP',
        '/api/v1/organizations', 'POST'),
       ('UpdateOrganization',
        'organization::write',
        'Update a organization profile',
        'HTTP',
        '/api/v1/organizations', 'PUT'),
       ('DeleteOrganization',
        'organization::delete',
        'Delete a organization profile',
        'HTTP',
        '/api/v1/organizations', 'DELETE'),
       ('ListAllOrganization',
        'organization::read',
        'ReadAllOrganizationProfile',
        'HTTP',
        '/api/v1/organizations/list', 'GET'),
       ('ReadResource',
        'resource::read',
        'Read A resource profile',
        'HTTP',
        '/api/v1/resources', 'GET'),
       ('CreateResource',
        'resource::write',
        'Create A resource',
        'HTTP',
        '/api/v1/resources', 'POST'),
       ('UpdateResource',
        'resource::write',
        'Update a resource profile',
        'HTTP',
        '/api/v1/resources', 'PUT'),
       ('DeleteResource',
        'resource::delete',
        'Delete a resource profile',
        'HTTP',
        '/api/v1/resources', 'DELETE'),
       ('ListAllResource',
        'resource::read',
        'ReadAllResourceProfile',
        'HTTP',
        '/api/v1/resources/list', 'GET');