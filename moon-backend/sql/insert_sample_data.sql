-- 创建user_organization表（如果不存在）
CREATE TABLE IF NOT EXISTS systems.user_organization (
    id UUID NOT NULL DEFAULT GEN_RANDOM_UUID(),
    user_id UUID NOT NULL REFERENCES systems.users(user_id),
    organization_id UUID NOT NULL REFERENCES systems.organization(organization_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

-- 插入示例数据，关联用户和组织
-- 假设我们有以下用户和组织
-- 用户1: user_id = 'some-user-id-1'
-- 用户2: user_id = 'some-user-id-2'
-- 组织1: organization_id = 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3'

-- 首先查询现有的用户和组织，获取实际的ID
SELECT user_id, nickname, email FROM systems.users LIMIT 5;
SELECT organization_id, organization_name FROM systems.organization LIMIT 5;

-- 插入示例关联数据（请根据实际查询结果修改ID）
-- 示例：关联用户到组织
-- INSERT INTO systems.user_organization (user_id, organization_id) VALUES
-- ('actual-user-id-1', 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3'),
-- ('actual-user-id-2', 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3');

-- 查询插入的关联数据
SELECT u.nickname, u.email, o.organization_name 
FROM systems.user_organization uo
JOIN systems.users u ON uo.user_id = u.user_id
JOIN systems.organization o ON uo.organization_id = o.organization_id;