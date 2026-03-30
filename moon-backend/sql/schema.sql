CREATE TABLE IF NOT EXISTS sys.organization (
    organization_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    organization_name text NOT NULL,
    organization_code text NOT NULL,
    organization_description text NOT NULL,
    organization_flag text NOT NULL,
    parent_id uuid DEFAULT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT now(),
    ts_vec vector(3)
);

INSERT INTO sys.organization (
    organization_id,
    organization_name,
    organization_code,
    organization_description,
    organization_flag,
    parent_id,
    created_at,
    updated_at
) VALUES
-- 1️⃣ 顶级：学校（parent_id = NULL）
('11111111-1111-1111-1111-111111111111', '岭南师范学院', '岭南师范学院', '广东省湛江市赤坎区跃进路3号', 'ACTIVE', NULL, NOW(), NOW()),

-- 2️⃣ 二级：校本部 → 父级=学校
('22222222-2222-2222-2222-222222222222', '校本部', '岭南师范学院::校本部', '主校区，位于市中心', 'ACTIVE',
 (SELECT organization_id FROM systems.organization WHERE organization_name = '岭南师范学院'),
 NOW(), NOW()),

-- 3️⃣ 二级：湖光校区 → 父级=学校
('33333333-3333-3333-3333-333333333333', '湖光校区', '岭南师范学院::湖光校区', '湖光校区', 'ACTIVE',
 (SELECT organization_id FROM systems.organization WHERE organization_name = '岭南师范学院'),
 NOW(), NOW()),

-- 4️⃣ 三级：计算机学院 → 父级=校本部
('44444444-4444-4444-4444-444444444444', '计算机与智能教育学院', '岭南师范学院::校本部::计算机与智能教育学院', '计算机与智能教育学院', 'ACTIVE',
 (SELECT organization_id FROM systems.organization WHERE organization_name = '校本部'),
 NOW(), NOW()),

-- 5️⃣ 三级：音乐与舞蹈学院 → 父级=校本部
('55555555-5555-5555-5555-555555555555', '音乐与舞蹈学院', '岭南师范学院::校本部::音乐与舞蹈学院', '音乐与舞蹈学院', 'ACTIVE',
 (SELECT organization_id FROM systems.organization WHERE organization_name = '校本部'),
 NOW(), NOW()),

-- 6️⃣ 三级：教室管理科 → 父级=校本部
('66666666-6666-6666-6666-666666666666', '校本部第五教学楼教室管理科', '岭南师范学院::校本部::第五教学楼教室管理科', '第五教学楼教室管理科', 'ACTIVE',
 (SELECT organization_id FROM systems.organization WHERE organization_name = '校本部'),
 NOW(), NOW()),

-- 7️⃣ 三级：法政学院 → 父级=湖光校区
('77777777-7777-7777-7777-777777777777', '法政学院', '岭南师范学院::湖光校区::法政学院', '法政学院', 'ACTIVE',
 (SELECT organization_id FROM systems.organization WHERE organization_name = '湖光校区'),
 NOW(), NOW()),

-- 8️⃣ 三级：信息工程学院 → 父级=湖光校区
('88888888-8888-8888-8888-888888888888', '信息工程学院', '岭南师范学院::湖光校区::信息工程学院', '信息工程工程学院', 'DEPRECATED',
 (SELECT organization_id FROM systems.organization WHERE organization_name = '湖光校区'),
 NOW(), NOW());

CREATE TABLE IF NOT EXISTS sys.app_user (
    user_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    nickname text NOT NULL,
    full_name text NOT NULL,
    email text NOT NULL,
    password_hash text NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL,
    organization_id uuid DEFAULT NULL REFERENCES systems.organization(organization_id)
);

INSERT INTO sys.app_user (
    nickname,
    full_name,
    email,
    password_hash,
    organization_id,
    created_at,
    updated_at
) VALUES
('admin', '管理员', 'admin@example.com', '123456', (select organization_id from systems.organization where organization_name = '岭南师范学院'), NOW(), NOW()),
('user1', '用户1', 'user1@example.com', 'user1', (select organization_id from systems.organization where organization_name = '校本部'), NOW(), NOW()),
('user3', '用户3', 'user3@example.com', 'user3', (select organization_id from systems.organization where organization_name = '计算机与智能教育学院'), NOW(), NOW()),
('user2', '用户2', 'user2@example.com', 'user2', (select organization_id from systems.organization where organization_name = '音乐与舞蹈学院'), NOW(), NOW()),
('user4', '用户4', 'user4@example.com', 'user4', (select organization_id from systems.organization where organization_name = '湖光校区'), NOW(), NOW()),
('user5', '用户5', 'user5@example.com', 'user5', (select organization_id from systems.organization where organization_name = '法政学院'), NOW(), NOW()),
('user6', '用户6', 'user6@example.com', 'user6', (select organization_id from systems.organization where organization_name = '信息工程学院'), NOW(), NOW());

CREATE TABLE IF NOT EXISTS systems.role(
    role_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    role_name text NOT NULL,
    role_code text NOT NULL,
    role_description text NOT NULL,
    role_flag text NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL
);

INSERT INTO systems.role (
    role_name,
    role_code,
    role_description,
    role_flag,
    created_at,
    updated_at
) VALUES
-- 资产管理系统是一个覆盖资产全生命周期（包括采购、入库、出库、验收、维护和报废等环节）的系统。
-- 系统管理员，负责管理系统的所有功能
('System Administrator', 'system::admin', '系统管理员', 'ACTIVE', NOW(), NOW()),
-- 资产管理员，负责管理学校内所有资产
('Asset', 'biz::asset-manager', '资产管理员', 'ACTIVE', NOW(), NOW()),
-- 仓库管理员，负责管理学校内所有仓库的资产
('Warehouse', 'biz::warehouse-manager', '仓库管理员', 'ACTIVE', NOW(), NOW()),
-- 采购为什么做采购，因为学校每年大约有200项10万以上的采购计划需要通过国资委审议，其余10万以下并不是不需要申报，只是合并统一上会处理
('Procure', 'biz::procure-manager', '采购管理员', 'ACTIVE', NOW(), NOW()),
-- 验收验收管理员，负责管理验收人员和验收流程
('Acceptance Manager', 'biz::acceptance-manager', '验收管理员', 'ACTIVE', NOW(), NOW()),
-- 采购人员，负责采购计划
('Procure Personnel', 'biz::procure-personnel', '采购人员', 'ACTIVE', NOW(), NOW()),
-- 采购专家，负责采购计划的评定
('Procure Expert', 'biz::procure-expert', '采购专家', 'ACTIVE', NOW(), NOW()),
-- 验收目前还是使用的纸质化，从资产的入场，安装到最终验收都需要手写签名并且打印水印相机照片留痕，会收集大量的纸质材料，如果丢失不好找
('Acceptance', 'biz::acceptance-personnel', '验收人员', 'ACTIVE', NOW(), NOW()),
-- 验收专家，负验收计划的评定
('Acceptance Expert', 'biz::acceptance-expert', '验收专家', 'ACTIVE', NOW(), NOW()),
-- 第三方服务供应商，包括但不限于：
-- 1. 维护供应商
-- 2. 服务供应商
-- 3. 其他供应商
('Provider', 'biz::provider', '供应商', 'ACTIVE', NOW(), NOW()),
-- 做资产维护，学校总计有五千多台电脑，数百个打印机，还有不计其数的小设备，物件需要维护更换，提交工单给负责人由他分配维护人员会方便很多
('Maintenance Manager', 'biz::maintenance-manager', '维护管理员', 'ACTIVE', NOW(), NOW()),
-- 维护人员，负责维护资产，如安装、维修、保养等
('Maintenance Personnel', 'biz::maintenance-personnel', '维护人员', 'ACTIVE', NOW(), NOW()),
-- 部门领导，负责协调部门内的资产维护和使用
('Leader', 'biz::leader', '负责人', 'ACTIVE', NOW(), NOW()),
-- 普通用户，负责使用资产，如学生、教师、员工等
('User', 'biz::user', '普通用户', 'ACTIVE', NOW(), NOW()),
-- 会计，负责资产的会计核算和管理
('Accountant', 'biz::accountant', '会计', 'ACTIVE', NOW(), NOW()),
-- 审计，负责资产的审计和管理
('Audit', 'biz::audit', '审计', 'ACTIVE', NOW(), NOW()),
-- 出纳，负责处理资产的现金交易
('Cashier', 'biz::cashier', '出纳', 'ACTIVE', NOW(), NOW());


CREATE TABLE IF NOT EXISTS systems.permission (
    permission_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    permission_name text NOT NULL,
    permission_code text NOT NULL,
    permission_description text NOT NULL,
    permission_flag text NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz NULL
);

INSERT INTO systems.permission (
    permission_name,
    permission_code,
    permission_description,
    permission_flag,
    created_at,
    updated_at 
) VALUES
('Any Permission', 'any::any', '任意权限', 'ACTIVE', NOW(), NOW()),
-- 组织机构相关权限
('查看组织机构', 'organization::view', '查看组织机构信息', 'ACTIVE', NOW(), NOW()),
('新增组织机构', 'organization::add', '新增组织机构', 'ACTIVE', NOW(), NOW()),
('修改组织机构', 'organization::edit', '修改组织机构信息', 'ACTIVE', NOW(), NOW()),
('删除组织机构', 'organization::delete', '删除组织机构', 'ACTIVE', NOW(), NOW()),

-- 用户相关权限
('查看用户', 'user::view', '查看用户信息', 'ACTIVE', NOW(), NOW()),
('新增用户', 'user::add', '新增用户', 'ACTIVE', NOW(), NOW()),
('修改用户', 'user::edit', '修改用户信息', 'ACTIVE', NOW(), NOW()),
('删除用户', 'user::delete', '删除用户', 'ACTIVE', NOW(), NOW()),

-- 角色相关权限
('查看角色', 'role::view', '查看角色信息', 'ACTIVE', NOW(), NOW()),
('新增角色', 'role::add', '新增角色', 'ACTIVE', NOW(), NOW()),
('修改角色', 'role::edit', '修改角色信息', 'ACTIVE', NOW(), NOW()),
('删除角色', 'role::delete', '删除角色', 'ACTIVE', NOW(), NOW()),

-- 权限相关权限
('查看权限', 'permission::view', '查看权限信息', 'ACTIVE', NOW(), NOW()),
('新增权限', 'permission::add', '新增权限', 'ACTIVE', NOW(), NOW()),
('修改权限', 'permission::edit', '修改权限信息', 'ACTIVE', NOW(), NOW()),
('删除权限', 'permission::delete', '删除权限', 'ACTIVE', NOW(), NOW()),

-- 资产类别权限
('查看资产类别', 'asset-category::view', '查看资产类别信息', 'ACTIVE', NOW(), NOW()),
('新增资产类别', 'asset-category::add', '新增资产类别', 'ACTIVE', NOW(), NOW()),
('修改资产类别', 'asset-category::edit', '修改资产类别信息', 'ACTIVE', NOW(), NOW()),
('删除资产类别', 'asset-category::delete', '删除资产类别', 'ACTIVE', NOW(), NOW()),

-- 资产类型权限
('查看资产类型', 'asset-type::view', '查看资产类型信息', 'ACTIVE', NOW(), NOW()),
('新增资产类型', 'asset-type::add', '新增资产类型', 'ACTIVE', NOW(), NOW()),
('修改资产类型', 'asset-type::edit', '修改资产类型信息', 'ACTIVE', NOW(), NOW()),
('删除资产类型', 'asset-type::delete', '删除资产类型', 'ACTIVE', NOW(), NOW()),

-- 资产明细权限
('查看资产明细', 'assets::view', '查看资产明细信息', 'ACTIVE', NOW(), NOW()),
('新增资产明细', 'assets::add', '新增资产明细', 'ACTIVE', NOW(), NOW()),
('修改资产明细', 'assets::edit', '修改资产明细信息', 'ACTIVE', NOW(), NOW()),
('删除资产明细', 'assets::delete', '删除资产明细', 'ACTIVE', NOW(), NOW()),

-- 采购相关权限
('查看采购计划类型', 'procurement-plan-type::view', '查看采购计划类型', 'ACTIVE', NOW(), NOW()),
('新增采购计划类型', 'procurement-plan-type::add', '新增采购计划类型', 'ACTIVE', NOW(), NOW()),
('修改采购计划类型', 'procurement-plan-type::edit', '修改采购计划类型', 'ACTIVE', NOW(), NOW()),
('删除采购计划类型', 'procurement-plan-type::delete', '删除采购计划类型', 'ACTIVE', NOW(), NOW()),
('查看采购计划', 'procurement-plan::view', '查看采购计划', 'ACTIVE', NOW(), NOW()),
('新增采购计划', 'procurement-plan::add', '新增采购计划', 'ACTIVE', NOW(), NOW()),
('修改采购计划', 'procurement-plan::edit', '修改采购计划', 'ACTIVE', NOW(), NOW()),
('删除采购计划', 'procurement-plan::delete', '删除采购计划', 'ACTIVE', NOW(), NOW()),
('查看采购实施', 'procurement-implementation::view', '查看采购实施', 'ACTIVE', NOW(), NOW()),
('新增采购实施', 'procurement-implementation::add', '新增采购实施', 'ACTIVE', NOW(), NOW()),
('修改采购实施', 'procurement-implementation::edit', '修改采购实施', 'ACTIVE', NOW(), NOW()),
('删除采购实施', 'procurement-implementation::delete', '删除采购实施', 'ACTIVE', NOW(), NOW()),
('查看采购专家', 'procurement-expert::view', '查看采购专家', 'ACTIVE', NOW(), NOW()),
('新增采购专家', 'procurement-expert::add', '新增采购专家', 'ACTIVE', NOW(), NOW()),
('修改采购专家', 'procurement-expert::edit', '修改采购专家', 'ACTIVE', NOW(), NOW()),
('删除采购专家', 'procurement-expert::delete', '删除采购专家', 'ACTIVE', NOW(), NOW()),

-- 漏洞相关权限
('查看漏洞', 'vulnerability::view', '查看漏洞信息', 'ACTIVE', NOW(), NOW()),
('新增漏洞', 'vulnerability::add', '新增漏洞', 'ACTIVE', NOW(), NOW()),
('修改漏洞', 'vulnerability::edit', '修改漏洞信息', 'ACTIVE', NOW(), NOW()),
('删除漏洞', 'vulnerability::delete', '删除漏洞', 'ACTIVE', NOW(), NOW()),
('查看漏洞软件', 'vulnerability-software::view', '查看漏洞软件', 'ACTIVE', NOW(), NOW()),
('新增漏洞软件', 'vulnerability-software::add', '新增漏洞软件', 'ACTIVE', NOW(), NOW()),
('修改漏洞软件', 'vulnerability-software::edit', '修改漏洞软件', 'ACTIVE', NOW(), NOW()),
('删除漏洞软件', 'vulnerability-software::delete', '删除漏洞软件', 'ACTIVE', NOW(), NOW()),

-- 会计、出纳、审计角色专用权限
('查看资产台账', 'asset-ledger::view', '查看资产台账', 'ACTIVE', NOW(), NOW()),
('导出资产台账', 'asset-ledger::export', '导出资产台账', 'ACTIVE', NOW(), NOW()),
('资产折旧', 'asset::depreciate', '执行资产折旧', 'ACTIVE', NOW(), NOW()),
('资产盘点', 'asset::inventory', '执行资产盘点', 'ACTIVE', NOW(), NOW()),
('资产报废', 'asset::retire', '资产报废处理', 'ACTIVE', NOW(), NOW()),
('查看财务报表', 'financial-report::view', '查看财务报表', 'ACTIVE', NOW(), NOW()),
('生成财务报表', 'financial-report::generate', '生成财务报表', 'ACTIVE', NOW(), NOW()),
('审核采购付款', 'procurement-payment::audit', '审核采购付款申请', 'ACTIVE', NOW(), NOW()),
('确认采购付款', 'procurement-payment::confirm', '确认采购付款', 'ACTIVE', NOW(), NOW()),
('查看付款记录', 'payment-record::view', '查看付款记录', 'ACTIVE', NOW(), NOW()),
('导出付款记录', 'payment-record::export', '导出付款记录', 'ACTIVE', NOW(), NOW()),

-- 菜单相关权限('采购管理菜单', 'menu::procurement', '查看采购管理菜单项', 'ACTIVE', NOW(), NOW()),
('组织机构菜单', 'menu::organization', '查看组织机构菜单项', 'ACTIVE', NOW(), NOW()),
('用户菜单', 'menu::user', '查看用户菜单项', 'ACTIVE', NOW(), NOW()),
('角色菜单', 'menu::role', '查看角色菜单项', 'ACTIVE', NOW(), NOW()),

('资产类别菜单', 'menu::asset-category', '查看资产类别菜单项', 'ACTIVE', NOW(), NOW()),
('资产类型菜单', 'menu::asset-type', '查看资产类型菜单项', 'ACTIVE', NOW(), NOW()),
('资产明细菜单', 'menu::assets', '查看资产明细菜单项', 'ACTIVE', NOW(), NOW()),

('漏洞管理菜单', 'menu::vulnerability', '查看漏洞管理菜单项', 'ACTIVE', NOW(), NOW()),
('财务资产菜单', 'menu::financial-asset', '查看财务资产菜单项', 'ACTIVE', NOW(), NOW()),
('会计出纳菜单', 'menu::accounting-cashier', '查看会计出纳菜单项', 'ACTIVE', NOW(), NOW()),
('审计菜单', 'menu::audit', '查看审计菜单项', 'ACTIVE', NOW(), NOW());

CREATE TABLE IF NOT EXISTS systems.role_permission (
    role_id UUID NOT NULL REFERENCES systems.role(role_id) ON DELETE CASCADE,
    permission_id UUID NOT NULL REFERENCES systems.permission(permission_id) ON DELETE CASCADE,
    PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE IF NOT EXISTS systems.files(
    file_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    file_name text NOT NULL,
    file_type text NOT NULL,
    file_size bigint NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz NULL
);

CREATE TABLE IF NOT EXISTS systems.file_refs(
    file_ref_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    file_id uuid NOT NULL REFERENCES systems.files(file_id) ON DELETE CASCADE,
);

INSERT INTO systems.role_permission (role_id, permission_id)
SELECT
    r.role_id,
    p.permission_id
FROM systems.role r
CROSS JOIN systems.permission p
WHERE r.role_code = 'system::admin';

-- 资产主类型表：区分 固定资产/无形资产
CREATE TABLE IF NOT EXISTS biz.asset_category (
    asset_category_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    category_name text NOT NULL UNIQUE, -- 固定资产、无形资产
    category_code text NOT NULL UNIQUE, -- FIXED、INTANGIBLE
    category_flag text NOT NULL DEFAULT 'ACTIVE',
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz NULL
);

-- 资产子类型表（原固定资产类型、无形资产类型合并）
CREATE TABLE IF NOT EXISTS biz.asset_type (
    asset_type_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_type_name text NOT NULL,
    asset_type_code text NOT NULL,
    asset_type_flag text NOT NULL DEFAULT 'ACTIVE',
    asset_category_id uuid NOT NULL REFERENCES biz.asset_category(asset_category_id),
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz NULL,
    UNIQUE(asset_type_name, asset_category_id)
);

CREATE TABLE IF NOT EXISTS biz.assets (
    asset_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    -- 基础信息
    asset_name text NOT NULL,
    asset_code text NOT NULL UNIQUE, -- 资产编码唯一
    asset_description text NOT NULL DEFAULT '',
    asset_flag text NOT NULL, -- 库存状态 IN-STOCK/PENDING-STOCK
    -- 通用资产属性
    quantity integer NOT NULL DEFAULT 1, -- 数量（无形资产默认1）
    location text NOT NULL DEFAULT '', -- 存放位置（无形资产可为空）
    purchase_price decimal NOT NULL, -- 原值
    depreciation_price decimal NOT NULL, -- 折旧/摊销值
    purchase_date date NULL, -- 购置日期
    manufacturer text NOT NULL DEFAULT '', -- 生产商（无形资产可为空）
    model text NOT NULL DEFAULT '', -- 型号（无形资产可为空）
    -- 扩展信息
    other_metadata jsonb NOT NULL DEFAULT '{}',
    -- 关联关系
    asset_type_id uuid NULL REFERENCES biz.asset_type(asset_type_id),
    asset_category_id uuid NOT NULL REFERENCES biz.asset_category(asset_category_id),
    organization_id uuid NULL REFERENCES systems.organization(organization_id),
    -- 审计字段
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz NULL,
    -- 索引优化
    CONSTRAINT uk_asset_code UNIQUE (asset_code)
);

INSERT INTO biz.asset_category (category_name, category_code, category_flag)
VALUES 
('固定资产', 'FIXED', 'ACTIVE'),
('无形资产', 'INTANGIBLE', 'ACTIVE');

-- 固定资产子类型 批量插入
INSERT INTO biz.asset_type (asset_type_name, asset_type_code, asset_type_flag, asset_category_id)
VALUES
('办公楼','房屋::办公楼','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('厂房','房屋::厂房','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('仓库','房屋::仓库','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('员工宿舍','房屋::员工宿舍','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('食堂','房屋::食堂','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('围墙','房屋::围墙','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('烟囱','房屋::烟囱','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('停车场','房屋::停车场','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('生产流水线','机器设备::生产流水线','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('加工机床','机器设备::加工机床','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('包装机械','机器设备::包装机械','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('注塑机','机器设备::注塑机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('纺织机','机器设备::纺织机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('印刷机','机器设备::印刷机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('专用生产工具','机器设备::专用生产工具','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('公司车辆','运输工具::公司车辆','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('叉车','运输工具::叉车','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('拖车','运输工具::拖车','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('船舶','运输工具::船舶','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('飞机','运输工具::飞机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('电脑','电子设备::电脑','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('服务器','电子设备::服务器','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('打印机','电子设备::打印机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('复印机','电子设备::复印机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('投影仪','电子设备::投影仪','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('路由器','电子设备::路由器','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('交换机','电子设备::交换机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('企业级交换机','电子设备::企业级交换机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('企业级路由器','电子设备::企业级路由器','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('监控设备','电子设备::监控设备','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('收银机','电子设备::收银机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('办公桌椅','办公设备::办公桌椅','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('文件柜','办公设备::文件柜','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('保险柜','办公设备::保险柜','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('空调','办公设备::空调','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('碎纸机','办公设备::碎纸机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('饮水机','办公设备::饮水机','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('沙发','家具及器具::沙发','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('茶几','家具及器具::茶几','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('会议桌','家具及器具::会议桌','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('会议椅','家具及器具::会议椅','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('书架','家具及器具::书架','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1)),
('床','家具及器具::床','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'FIXED' LIMIT 1));

-- 无形资产子类型 批量插入
INSERT INTO biz.asset_type (asset_type_name, asset_type_code, asset_type_flag, asset_category_id)
VALUES
('发明专利','知识产权::发明专利','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('实用新型','知识产权::实用新型','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('外观设计','知识产权::外观设计','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('商标权','知识产权::商标权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('著作权','知识产权::著作权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('软件著作权','知识产权::软件著作权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('集成电路布图设计','知识产权::集成电路布图设计','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('植物新品种','知识产权::植物新品种','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('特许经营权','市场权利::特许经营权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('商标使用权','市场权利::商标使用权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('土地使用权','市场权利::土地使用权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('采矿权','市场权利::采矿权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('探矿权','市场权利::探矿权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('排污权','市场权利::排污权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('碳排放权','市场权利::碳排放权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('非专利技术','技术信息::非专利技术','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('技术秘密','技术信息::技术秘密','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('计算机软件','技术信息::计算机软件','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('客户关系','技术信息::客户关系','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('客户名单','技术信息::客户名单','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('销售网络','技术信息::销售网络','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('分销渠道','技术信息::分销渠道','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('合同权益','其他::合同权益','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('影视版权','其他::影视版权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1)),
('音乐版权','其他::音乐版权','ACTIVE',(SELECT asset_category_id FROM biz.asset_category WHERE category_code = 'INTANGIBLE' LIMIT 1));


-- 1. 插入 固定资产 10条
INSERT INTO biz.assets (
    asset_name, asset_code, asset_description, asset_flag,
    quantity, location, purchase_price, depreciation_price, purchase_date,
    manufacturer, model, other_metadata, asset_type_id, asset_category_id, organization_id
)
VALUES
-- 1
('办公楼','FIXED001','教学办公楼','IN-STOCK',1,'校本部',5000000,5000000,'2023-01-01','城建集团','OFFICE-BUILD-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='办公楼' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='校本部' LIMIT 1)),
-- 2
('厂房','FIXED002','实训厂房','IN-STOCK',1,'工程实训中心',2000000,2000000,'2023-02-01','重工机械','FACTORY-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='厂房' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='工程实训中心' LIMIT 1)),
-- 3
('仓库','FIXED003','物资仓库','IN-STOCK',1,'后勤园区',800000,800000,'2023-03-01','仓储设备','WAREHOUSE-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='仓库' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='后勤管理处' LIMIT 1)),
-- 4
('员工宿舍','FIXED004','教师公寓','IN-STOCK',1,'教工生活区',1500000,1500000,'2023-04-01','建设集团','DORM-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='员工宿舍' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='人事处' LIMIT 1)),
-- 5
('食堂','FIXED005','学生食堂','IN-STOCK',1,'生活区',1200000,1200000,'2023-05-01','餐饮设备','CANTEEN-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='食堂' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='后勤服务中心' LIMIT 1)),
-- 6
('生产流水线','FIXED006','实训流水线','IN-STOCK',1,'工学院大楼',800000,800000,'2023-06-01','工业设备','LINE-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='生产流水线' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='工学院' LIMIT 1)),
-- 7
('公司车辆','FIXED007','公务用车','IN-STOCK',2,'行政楼',300000,300000,'2023-07-01','汽车品牌','CAR-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='公司车辆' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='党政办公室' LIMIT 1)),
-- 8
('电脑','FIXED008','办公电脑','IN-STOCK',50,'计算机学院',4000,4000,'2023-08-01','电脑品牌','PC-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='电脑' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='计算机与智能教育学院' LIMIT 1)),
-- 9
('办公桌椅','FIXED009','办公桌椅套装','IN-STOCK',100,'行政办公楼',800,800,'2023-09-01','家具厂','DESK-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='办公桌椅' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='行政处' LIMIT 1)),
-- 10
('沙发','FIXED10','会议室沙发','IN-STOCK',10,'学术交流中心',2000,2000,'2023-10-01','家具品牌','SOFA-01','{}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='沙发' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='FIXED' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='学术交流中心' LIMIT 1));

-- 2. 插入 无形资产 10条
INSERT INTO biz.assets (
    asset_name, asset_code, asset_description, asset_flag,
    quantity, location, purchase_price, depreciation_price, purchase_date,
    manufacturer, model, other_metadata, asset_type_id, asset_category_id, organization_id
)
VALUES
-- 1
('发明专利','INTANG001','教学方法发明专利','IN-STOCK',1,'',1000000,1000000,'2023-01-01','','','{"author":"教授A"}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='发明专利' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='科研处' LIMIT 1)),
-- 2
('实用新型','INTANG002','实验设备实用新型','IN-STOCK',1,'',800000,800000,'2023-02-01','','','{"author":"教授B"}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='实用新型' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='实验室管理处' LIMIT 1)),
-- 3
('外观设计','INTANG003','产品外观设计专利','IN-STOCK',1,'',500000,500000,'2023-03-01','','','{"author":"设计师C"}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='外观设计' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='艺术学院' LIMIT 1)),
-- 4
('商标权','INTANG004','学校品牌商标权','IN-STOCK',1,'',600000,600000,'2023-04-01','','','{"reg_no":"TM-2023-001"}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='商标权' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='品牌管理中心' LIMIT 1)),
-- 5
('著作权','INTANG005','教材著作权','IN-STOCK',1,'',300000,300000,'2023-05-01','','','{"isbn":"978-7-xxxx"}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='著作权' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='出版社' LIMIT 1)),
-- 6
('软件著作权','INTANG006','教学管理系统','IN-STOCK',1,'',400000,400000,'2023-06-01','','','{"version":"v2.0"}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='软件著作权' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='信息化中心' LIMIT 1)),
-- 7
('特许经营权','INTANG007','校园服务特许经营权','IN-STOCK',1,'',1500000,1500000,'2023-07-01','','','{"period":"10年"}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='特许经营权' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='资产管理处' LIMIT 1)),
-- 8
('土地使用权','INTANG008','校区土地使用权','IN-STOCK',1,'',3000000,3000000,'2023-08-01','','','{"area":"100亩"}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='土地使用权' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='规划处' LIMIT 1)),
-- 9
('非专利技术','INTANG009','实验专有技术','IN-STOCK',1,'',700000,700000,'2023-09-01','','','{"field":"新材料"}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='非专利技术' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='材料学院' LIMIT 1)),
-- 10
('客户关系','INTANG10','校企合作客户资源','IN-STOCK',1,'',500000,500000,'2023-10-01','','','{"partner_count":50}',
 (SELECT asset_type_id FROM biz.asset_type WHERE asset_type_name='客户关系' LIMIT 1),
 (SELECT asset_category_id FROM biz.asset_category WHERE category_code='INTANGIBLE' LIMIT 1),
 (SELECT organization_id FROM systems.organization WHERE organization_name='校企合作办' LIMIT 1));

CREATE TABLE IF NOT EXISTS biz.procurement_plan_type(
    procurement_plan_type_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    procurement_plan_type_name text NOT NULL,
    procurement_plan_type_code text NOT NULL,
    procurement_plan_type_flag text NOT NULL,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL
);

INSERT INTO biz.procurement_plan_type(
    procurement_plan_type_name,
    procurement_plan_type_code,
    procurement_plan_type_flag
) VALUES
-- 按组织形式分类
('自行招标', '组织形式::自行招标', 'ACTIVE'),
('委托招标', '组织形式::委托招标', 'ACTIVE'),
-- 按采购内容分类
('货物类采购', '采购内容::货物类采购', 'ACTIVE'),
('工程类采购', '采购内容::工程类采购', 'ACTIVE'),
('服务类采购', '采购内容::服务类采购', 'ACTIVE'),
-- 按采购方式分类
('公开招标', '采购方式::公开招标', 'ACTIVE'),
('邀请招标', '采购方式::邀请招标', 'ACTIVE'),
('竞争性谈判', '采购方式::竞争性谈判', 'ACTIVE'),
('竞争性磋商', '采购方式::竞争性磋商', 'ACTIVE'),
('询价采购', '采购方式::询价采购', 'ACTIVE'),
('单一来源采购', '采购方式::单一来源采购', 'ACTIVE'),
('框架协议采购', '采购方式::框架协议采购', 'ACTIVE'),
('电子采购', '采购方式::电子采购', 'ACTIVE'),
('本地化采购', '采购方式::本地化采购', 'ACTIVE'),
('全球采购', '采购方式::全球采购', 'ACTIVE');

CREATE TABLE IF NOT EXISTS biz.procurement_plan(
    procurement_plan_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    procurement_plan_name text NOT NULL,
    procurement_plan_code text NOT NULL,
    procurement_plan_description text NOT NULL,
    procurement_plan_flag text NOT NULL,
    procurement_plan_quantity integer NOT NULL DEFAULT 1,
    procurement_plan_price decimal NOT NULL,
    procurement_plan_purchase_date timestamptz DEFAULT NULL,
    procurement_plan_purchase_type text NOT NULL DEFAULT '',
    other_metadata jsonb NOT NULL DEFAULT '{}',
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL,
    procurement_plan_type_id uuid DEFAULT NULL REFERENCES biz.procurement_plan_type(procurement_plan_type_id),
    organization_id uuid DEFAULT NULL REFERENCES systems.organization(organization_id),
    user_id uuid DEFAULT NULL REFERENCES systems.user(user_id)
);
INSERT INTO biz.procurement_plan(
    procurement_plan_name,
    procurement_plan_code,
    procurement_plan_description,
    procurement_plan_flag,
    procurement_plan_quantity,
    procurement_plan_price,
    procurement_plan_purchase_date,
    procurement_plan_purchase_type,
    other_metadata,
    procurement_plan_type_id,
    organization_id,
    user_id
) VALUES
('采购计划1','G100001','采购计划1','TENDER',100,1000000,'2023-01-01','软件采购','{}',
(select procurement_plan_type_id from biz.procurement_plan_type where procurement_plan_type_name = '自行招标'),
(select organization_id from systems.organization where organization_name = '校本部第五教学楼教室管理科'),
(select user_id from systems.user where full_name = '用户2')),

('采购计划2','G100002','采购计划2','TENDER',50,800000,'2023-02-01','硬件采购','{}',
(select procurement_plan_type_id from biz.procurement_plan_type where procurement_plan_type_name = '委托招标'),
(select organization_id from systems.organization where organization_name = '湖光校区'),
(select user_id from systems.user where full_name = '用户4')),

('采购计划3','G100003','采购计划3','BID',200,600000,'2023-03-01','服务采购','{}',
(select procurement_plan_type_id from biz.procurement_plan_type where procurement_plan_type_name = '货物类采购'),
(select organization_id from systems.organization where organization_name = '法政学院'),
(select user_id from systems.user where full_name = '用户5')),

('采购计划4','G100004','采购计划4','TENDER',30,1200000,'2023-04-01','软件升级','{}',
(select procurement_plan_type_id from biz.procurement_plan_type where procurement_plan_type_name = '竞争性谈判'),
(select organization_id from systems.organization where organization_name = '音乐与舞蹈学院'),
(select user_id from systems.user where full_name = '用户2')),

('采购计划5','G100005','采购计划5','TENDER',80,400000,'2023-05-01','办公用品','{}',
(select procurement_plan_type_id from biz.procurement_plan_type where procurement_plan_type_name = '询价采购'),
(select organization_id from systems.organization where organization_name = '计算机与智能教育学院'),
(select user_id from systems.user where full_name = '用户3')),

('采购计划6','G100006','采购计划6','BID',150,950000,'2023-06-01','实验设备','{}',
(select procurement_plan_type_id from biz.procurement_plan_type where procurement_plan_type_name = '公开招标'),
(select organization_id from systems.organization where organization_name = '校本部'),
(select user_id from systems.user where full_name = '用户1'));

CREATE TABLE IF NOT EXISTS biz.procurement_implementation(
    procurement_implementation_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    procurement_implementation_name text NOT NULL,
    procurement_implementation_code text NOT NULL,
    procurement_implementation_description text NOT NULL,
    procurement_implementation_flag text NOT NULL,
    other_metadata jsonb NOT NULL DEFAULT '{}',
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL,
    procurement_plan_id uuid DEFAULT NULL REFERENCES biz.procurement_plan(procurement_plan_id),
    organization_id uuid DEFAULT NULL REFERENCES systems.organization(organization_id),
    user_id uuid DEFAULT NULL REFERENCES systems.user(user_id)
);

INSERT INTO biz.procurement_implementation(
    procurement_implementation_name,
    procurement_implementation_code,
    procurement_implementation_description,
    procurement_implementation_flag,
    other_metadata,
    organization_id,
    user_id
) VALUES
('采购计划1实施','G100001','采购计划1实施','STARTING','[{"date":"2023-01-01","status":"到场","acceptance":"通过","现场":"张三","photo":"https://example.com/photo1.jpg"},{"date":"2023-01-02","status":" 安装","acceptance":"通过","现场":"李四","photo":"https://example.com/photo2.jpg"},{"date":"2023-01-03","status":"项目验收专家入场","acceptance":"通过","现场":"王五","photo":"https://example.com/photo3.jpg"}]',
(select organization_id from systems.organization where organization_name = '校本部第五教学楼教室管理科'),
(select user_id from systems.user where full_name = '用户2')),
('采购计划2实施','G100002','采购计划2实施','DOING','[{"date":"2023-02-01","status":"到场","acceptance":"通过","现场":"王五","photo":"https://example.com/photo1.jpg"},{"date":"2023-02-02","status":" 安装","acceptance":"通过","现场":"赵六","photo":"https://example.com/photo2.jpg"},{"date":"2023-02-03","status":"项目验收专家入场","acceptance":"通过","现场":"王五","photo":"https://example.com/photo3.jpg"}]',
(select organization_id from systems.organization where organization_name = '湖光校区'),
(select user_id from systems.user where full_name = '用户4')),
('采购计划3实施','G100003','采购计划3实施','FINISH','[{"date":"2023-03-01","status":"到场","acceptance":"通过","现场":"钱七","photo":"https://example.com/photo1.jpg"},{"date":"2023-03-02","status":" 安装","acceptance":"通过","现场":"王八","photo":"https://example.com/photo2.jpg"},{"date":"2023-03-03","status":"项目验收专家入场","acceptance":"通过","现场":"王五","photo":"https://example.com/photo3.jpg"}]',
(select organization_id from systems.organization where organization_name = '法政学院'),
(select user_id from systems.user where full_name = '用户5')),
('采购计划4实施','G100004','采购计划4实施','FINISH','[{"date":"2023-04-01","status":"到场","acceptance":"通过","现场":"赵六","photo":"https://example.com/photo1.jpg"},{"date":"2023-04-02","status":" 安装","acceptance":"通过","现场":"王五","photo":"https://example.com/photo2.jpg"},{"date":"2023-04-03","status":"项目验收专家入场","acceptance":"通过","现场":"王五","photo":"https://example.com/photo3.jpg"}]',
(select organization_id from systems.organization where organization_name = '音乐与舞蹈学院'),
(select user_id from systems.user where full_name = '用户2'));

CREATE TABLE IF NOT EXISTS biz.procurement_expert(
    procurement_expert_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    procurement_expert_name text NOT NULL,
    procurement_expert_code text NOT NULL,
    procurement_expert_description text NOT NULL,
    procurement_expert_flag text NOT NULL,
    procurement_expert_job_grade text NOT NULL DEFAULT '',
    procurement_expert_bank_name text NOT NULL DEFAULT '',
    procurement_expert_bank_account text NOT NULL DEFAULT '',
    other_metadata jsonb NOT NULL DEFAULT '{}',
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL,
    organization_id uuid DEFAULT NULL REFERENCES systems.organization(organization_id),
    user_id uuid DEFAULT NULL REFERENCES systems.user(user_id)
);

INSERT INTO biz.procurement_expert(
    procurement_expert_name,
    procurement_expert_code,
    procurement_expert_description,
    procurement_expert_flag,
    procurement_expert_job_grade,
    procurement_expert_bank_name,
    procurement_expert_bank_account,
    other_metadata,
    organization_id,
    user_id
) VALUES
('专家1','ZJ100001','专家1','ACTIVE','高级工程师','中国银行','6100000000000000000','{}',
(select organization_id from systems.organization where organization_name = '校本部第五教学楼教室管理科'),
(select user_id from systems.user where full_name = '用户2')),
('专家2','ZJ100002','专家2','ACTIVE','高级工程师','中国银行','6200000000000000000','{}',
(select organization_id from systems.organization where organization_name = '校本部第五教学楼教室管理科'),
(select user_id from systems.user where full_name = '用户3')),
('专家3','ZJ100003','专家3','ACTIVE','高级工程师','中国银行','6300000000000000000','{}',
(select organization_id from systems.organization where organization_name = '校本部第五教学楼教室管理科'),
(select user_id from systems.user where full_name = '用户4')),
('专家4','ZJ100004','专家4','ACTIVE','高级工程师','中国银行','6400000000000000000','{}',
(select organization_id from systems.organization where organization_name = '校本部第五教学楼教室管理科'),
(select user_id from systems.user where full_name = '用户5'));

CREATE TABLE IF NOT EXISTS biz.procurement_review (
    procurement_review_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    procurement_review_name text NOT NULL,
    procurement_review_code text NOT NULL,
    procurement_review_result text NOT NULL,
    procurement_review_opinion text NOT NULL DEFAULT '',
    other_metadata jsonb NOT NULL DEFAULT '{}',
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz NULL,
    procurement_implementation_id uuid NULL REFERENCES biz.procurement_implementation(procurement_implementation_id),
    organization_id uuid NULL REFERENCES systems.organization(organization_id),
    user_id uuid NULL REFERENCES systems.user(user_id)
);

CREATE TABLE IF NOT EXISTS biz.procurement_acceptance (
    procurement_acceptance_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    procurement_acceptance_name text NOT NULL,
    procurement_acceptance_code text NOT NULL,
    procurement_acceptance_result text NOT NULL,
    procurement_acceptance_description text NOT NULL DEFAULT '',
    other_metadata jsonb NOT NULL DEFAULT '{}',
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz NULL,
    procurement_implementation_id uuid NULL REFERENCES biz.procurement_implementation(procurement_implementation_id),
    organization_id uuid NULL REFERENCES systems.organization(organization_id),
    user_id uuid NULL REFERENCES systems.user(user_id)
);

CREATE TABLE IF NOT EXISTS biz.procurement_review_procurement_expert (
    procurement_review_id uuid NOT NULL REFERENCES biz.procurement_review(procurement_review_id) ON DELETE CASCADE,
    procurement_expert_id uuid NOT NULL REFERENCES biz.procurement_expert(procurement_expert_id) ON DELETE CASCADE,
    PRIMARY KEY (procurement_review_id, procurement_expert_id)
);

-- 1. 多大模型配置表（兼容GPT、DeepSeek、通义千问、火山方舟）
CREATE TABLE IF NOT EXISTS ai.llm_model (
    model_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    model_name text NOT NULL,
    provider text NOT NULL,
    model_code text UNIQUE NOT NULL,
    api_key text NOT NULL,
    api_endpoint text,
    enabled bool DEFAULT true,
    max_tokens int DEFAULT 4096,
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz
);

-- 2. Token 消耗统计表（计费、监控、用量统计）
CREATE TABLE IF NOT EXISTS ai.llm_token_usage (
    usage_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    model_id uuid NOT NULL REFERENCES ai.llm_model(model_id),
    user_id uuid NULL,
    prompt_tokens int DEFAULT 0,
    completion_tokens int DEFAULT 0,
    total_tokens int DEFAULT 0,
    cost_amount decimal(18,6) DEFAULT 0.0,
    created_at timestamptz DEFAULT now()
);

-- 3. AI 采购舞弊分析结果表
CREATE TABLE IF NOT EXISTS ai.procurement_fraud_risk (
    risk_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    procurement_plan_id uuid NULL REFERENCES biz.procurement_plan(procurement_plan_id),
    procurement_implementation_id uuid NULL REFERENCES biz.procurement_implementation(procurement_implementation_id),
    risk_level text NOT NULL,
    risk_reason text,
    risk_score int DEFAULT 0,
    created_at timestamptz DEFAULT now()
);

-- 4. AI 采购分析日志表
CREATE TABLE IF NOT EXISTS ai.llm_procurement_analysis (
    log_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    model_id uuid REFERENCES ai.llm_model(model_id),
    procurement_plan_id uuid NULL REFERENCES biz.procurement_plan(procurement_plan_id),
    procurement_implementation_id uuid NULL REFERENCES biz.procurement_implementation(procurement_implementation_id),
    analysis_result text,
    token_usage_id uuid REFERENCES ai.llm_token_usage(usage_id),
    created_at timestamptz DEFAULT now()
);

CREATE TABLE IF NOT EXISTS biz.vulnerability(
    vulnerability_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    vulnerability_name text NOT NULL,
    vulnerability_code text NOT NULL,
    vulnerability_description text NOT NULL,
    vulnerability_flag text NOT NULL,
    vulnerability_type text NOT NULL DEFAULT '',
    other_metadata jsonb NOT NULL DEFAULT '{}',
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL,
);
INSERT INTO biz.vulnerability(
    vulnerability_name,
    vulnerability_code,
    vulnerability_description,
    vulnerability_flag,
    vulnerability_type,
    other_metadata,
) VALUES
('漏洞1','V100001','漏洞1','ACTIVE','SQL注入','{"cveid":"CVE-2023-1234","cve_description":"SQL注入漏洞1","cve_score":9.0,"cvss_score":9.0"}'),
('漏洞2','V100002','漏洞2','ACTIVE','SQL注入','{"cveid":"CVE-2023-1235","cve_description":"SQL注入漏洞2","cve_score":9.0,"cvss_score":9.0"}'),
('漏洞3','V100003','漏洞3','ACTIVE','SQL注入','{"cveid":"CVE-2023-1236","cve_description":"SQL注入漏洞3","cve_score":9.0,"cvss_score":9.0"}');

CREATE TABLE IF NOT EXISTS biz.vulnerability_software(
    vulnerability_software_id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    vulnerability_software_name text NOT NULL,
    vulnerability_software_code text NOT NULL,
    vulnerability_software_description text NOT NULL,
    vulnerability_software_flag text NOT NULL,
    other_metadata jsonb NOT NULL DEFAULT '{}',
    created_at timestamptz DEFAULT now(),
    updated_at timestamptz DEFAULT NULL,

);

INSERT INTO biz.vulnerability_software(
    vulnerability_software_name,
    vulnerability_software_code,
    vulnerability_software_description,
    vulnerability_software_flag,
    other_metadata,
)VALUES
('漏洞1软件1','V100001软件1','漏洞1软件1','ACTIVE','{"version":[{"version":"v1.0","release_date":"2023-01-01"},{"version":"v1.1","release_date":"2023-02-01"}]}'),
('漏洞2软件1','V100002软件1','漏洞2软件1','ACTIVE','{"version":[{"version":"v1.0","release_date":"2023-01-01"},{"version":"v1.1","release_date":"2023-02-01"}]}'),
('漏洞3软件1','V100003软件1','漏洞3软件1','ACTIVE','{"version":[{"version":"v1.0","release_date":"2023-01-01"},{"version":"v1.1","release_date":"2023-02-01"}]}');






