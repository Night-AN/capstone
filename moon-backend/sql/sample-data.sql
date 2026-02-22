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
