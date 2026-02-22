-- 创建 biz schema（如果不存在）
CREATE SCHEMA IF NOT EXISTS biz;

-- 创建资产表
CREATE TABLE IF NOT EXISTS biz.asset
(
    asset_id          UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    asset_name        TEXT        NOT NULL DEFAULT '',
    asset_code        TEXT        NOT NULL DEFAULT '',
    asset_description TEXT        NOT NULL DEFAULT '',
    asset_type        TEXT        NOT NULL DEFAULT '',
    asset_class       TEXT        NOT NULL DEFAULT '',
    manufacturer     TEXT        NOT NULL DEFAULT '',
    model            TEXT        NOT NULL DEFAULT '',
    serial_number     TEXT        NOT NULL DEFAULT '',
    ip_address        TEXT        NOT NULL DEFAULT '',
    mac_address       TEXT        NOT NULL DEFAULT '',
    location         TEXT        NOT NULL DEFAULT '',
    department       TEXT        NOT NULL DEFAULT '',
    owner            TEXT        NOT NULL DEFAULT '',
    contact_info      TEXT        NOT NULL DEFAULT '',
    status           TEXT        NOT NULL DEFAULT '',
    purchase_date     TEXT        NOT NULL DEFAULT '',
    warranty_end_date  TEXT        NOT NULL DEFAULT '',
    value            TEXT        NOT NULL DEFAULT '',
    notes            TEXT        NOT NULL DEFAULT '',
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 创建漏洞表
CREATE TABLE IF NOT EXISTS biz.vulnerability
(
    vulnerability_id       UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    cve_id                 TEXT        NOT NULL DEFAULT '',
    nist_cve_id             TEXT        NOT NULL DEFAULT '',
    title                 TEXT        NOT NULL DEFAULT '',
    description           TEXT        NOT NULL DEFAULT '',
    severity              TEXT        NOT NULL DEFAULT '',
    cvss_score             FLOAT       NOT NULL DEFAULT 0.0,
    cvss_vector            TEXT        NOT NULL DEFAULT '',
    affected_software      TEXT        NOT NULL DEFAULT '',
    affected_versions      TEXT        NOT NULL DEFAULT '',
    attack_vector          TEXT        NOT NULL DEFAULT '',
    attack_complexity      TEXT        NOT NULL DEFAULT '',
    privileges_required    TEXT        NOT NULL DEFAULT '',
    user_interaction       TEXT        NOT NULL DEFAULT '',
    scope                 TEXT        NOT NULL DEFAULT '',
    confidentiality_impact TEXT        NOT NULL DEFAULT '',
    integrity_impact       TEXT        NOT NULL DEFAULT '',
    availability_impact    TEXT        NOT NULL DEFAULT '',
    reference_urls         TEXT        NOT NULL DEFAULT '',
    solution              TEXT        NOT NULL DEFAULT '',
    status                TEXT        NOT NULL DEFAULT '',
    published_date         TEXT        NOT NULL DEFAULT '',
    last_modified_date      TEXT        NOT NULL DEFAULT '',
    created_at             TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at             TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 创建资产漏洞关联表
CREATE TABLE IF NOT EXISTS biz.asset_vulnerability
(
    asset_vulnerability_id UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    asset_id             UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    vulnerability_id     UUID        NOT NULL DEFAULT GEN_RANDOM_UUID(),
    detection_date       TEXT        NOT NULL DEFAULT '',
    status              TEXT        NOT NULL DEFAULT '',
    risk_level           TEXT        NOT NULL DEFAULT '',
    remediation_plan     TEXT        NOT NULL DEFAULT '',
    assigned_to          TEXT        NOT NULL DEFAULT '',
    due_date             TEXT        NOT NULL DEFAULT '',
    notes               TEXT        NOT NULL DEFAULT '',
    created_at           TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at           TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- 创建索引
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_asset_id ON biz.asset (asset_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_asset_code ON biz.asset (asset_code);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_vulnerability_id ON biz.vulnerability (vulnerability_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_vulnerability_cve_id ON biz.vulnerability (cve_id);
CREATE UNIQUE INDEX CONCURRENTLY IF NOT EXISTS idx_biz_asset_vulnerability_id ON biz.asset_vulnerability (asset_vulnerability_id);

-- 创建触发器函数（如果不存在）
CREATE OR REPLACE FUNCTION biz.set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = NOW();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- 创建触发器
DROP TRIGGER IF EXISTS tr_biz_asset_set_updated_at ON biz.asset;
CREATE TRIGGER tr_biz_asset_set_updated_at
    BEFORE
        UPDATE
    ON biz.asset
    FOR EACH ROW
EXECUTE FUNCTION biz.set_updated_at();

DROP TRIGGER IF EXISTS tr_biz_vulnerability_set_updated_at ON biz.vulnerability;
CREATE TRIGGER tr_biz_vulnerability_set_updated_at
    BEFORE
        UPDATE
    ON biz.vulnerability
    FOR EACH ROW
EXECUTE FUNCTION biz.set_updated_at();

DROP TRIGGER IF EXISTS tr_biz_asset_vulnerability_set_updated_at ON biz.asset_vulnerability;
CREATE TRIGGER tr_biz_asset_vulnerability_set_updated_at
    BEFORE
        UPDATE
    ON biz.asset_vulnerability
    FOR EACH ROW
EXECUTE FUNCTION biz.set_updated_at();

-- 添加约束
ALTER TABLE biz.asset
    ADD CONSTRAINT pk_biz_asset_id PRIMARY KEY USING INDEX idx_biz_asset_id,
    ADD CONSTRAINT uk_biz_asset_code UNIQUE USING INDEX idx_biz_asset_code,
    ADD CONSTRAINT chk_biz_asset_name CHECK (LENGTH(asset_name) <= 128),
    ADD CONSTRAINT chk_biz_asset_code CHECK (LENGTH(asset_code) <= 64),
    ADD CONSTRAINT chk_biz_asset_description CHECK (LENGTH(asset_description) <= 1024);

ALTER TABLE biz.vulnerability
    ADD CONSTRAINT pk_biz_vulnerability_id PRIMARY KEY USING INDEX idx_biz_vulnerability_id,
    ADD CONSTRAINT uk_biz_vulnerability_cve_id UNIQUE USING INDEX idx_biz_vulnerability_cve_id,
    ADD CONSTRAINT chk_biz_vulnerability_title CHECK (LENGTH(title) <= 256),
    ADD CONSTRAINT chk_biz_vulnerability_cve_id CHECK (LENGTH(cve_id) <= 64);

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
