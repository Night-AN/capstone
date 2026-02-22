import { Asset, AssetListItem, AssetCreateRequest, AssetUpdateRequest } from './asset.model';

describe('Asset Models', () => {
  describe('Asset interface', () => {
    it('should create a valid Asset object', () => {
      const asset: Asset = {
        asset_id: '1',
        asset_name: 'Test Asset',
        asset_code: 'ASSET-001',
        asset_description: 'Test Asset Description',
        organization_id: 'org1',
        asset_type: 'server',
        asset_class: 'hardware',
        manufacturer: 'Dell',
        model: 'PowerEdge R740',
        serial_number: 'ABC123',
        ip_address: '192.168.1.100',
        mac_address: '00:11:22:33:44:55',
        location: 'Data Center',
        department: 'IT',
        owner: 'John Doe',
        contact_info: 'john.doe@example.com',
        status: 'active',
        purchase_date: '2023-01-01',
        warranty_end_date: '2024-01-01',
        value: '5000',
        notes: 'Test Notes',
        created_at: '2023-01-01T00:00:00Z',
        updated_at: '2023-01-01T00:00:00Z'
      };

      expect(asset.asset_id).toBe('1');
      expect(asset.asset_name).toBe('Test Asset');
      expect(asset.asset_code).toBe('ASSET-001');
      expect(asset.organization_id).toBe('org1');
      expect(asset.asset_type).toBe('server');
    });

    it('should handle optional fields', () => {
      const minimalAsset: Asset = {
        asset_id: '1',
        asset_name: 'Test Asset',
        asset_code: 'ASSET-001',
        asset_description: '',
        organization_id: '',
        asset_type: '',
        asset_class: '',
        manufacturer: '',
        model: '',
        serial_number: '',
        ip_address: '',
        mac_address: '',
        location: '',
        department: '',
        owner: '',
        contact_info: '',
        status: '',
        purchase_date: '',
        warranty_end_date: '',
        value: '',
        notes: '',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      };

      expect(minimalAsset).toBeTruthy();
      expect(typeof minimalAsset.asset_id).toBe('string');
      expect(typeof minimalAsset.asset_name).toBe('string');
    });
  });

  describe('AssetListItem interface', () => {
    it('should create a valid AssetListItem object', () => {
      const assetListItem: AssetListItem = {
        asset_id: '1',
        asset_name: 'Test Asset',
        asset_code: 'ASSET-001',
        organization_id: 'org1',
        asset_type: 'server',
        asset_class: 'hardware',
        ip_address: '192.168.1.100',
        status: 'active',
        created_at: '2023-01-01T00:00:00Z'
      };

      expect(assetListItem.asset_id).toBe('1');
      expect(assetListItem.asset_name).toBe('Test Asset');
      expect(assetListItem.organization_id).toBe('org1');
      expect(assetListItem.asset_type).toBe('server');
    });

    it('should handle optional fields', () => {
      const minimalAssetListItem: AssetListItem = {
        asset_id: '1',
        asset_name: 'Test Asset',
        asset_code: 'ASSET-001',
        organization_id: '',
        asset_type: '',
        asset_class: '',
        ip_address: '',
        status: '',
        created_at: new Date().toISOString()
      };

      expect(minimalAssetListItem).toBeTruthy();
      expect(typeof minimalAssetListItem.asset_id).toBe('string');
      expect(typeof minimalAssetListItem.asset_name).toBe('string');
    });
  });

  describe('AssetCreateRequest interface', () => {
    it('should create a valid AssetCreateRequest object', () => {
      const assetCreateRequest: AssetCreateRequest = {
        asset_name: 'Test Asset',
        asset_code: 'ASSET-001',
        asset_description: 'Test Asset Description',
        organization_id: 'org1',
        asset_type: 'server',
        asset_class: 'hardware',
        manufacturer: 'Dell',
        model: 'PowerEdge R740',
        serial_number: 'ABC123',
        ip_address: '192.168.1.100',
        mac_address: '00:11:22:33:44:55',
        location: 'Data Center',
        department: 'IT',
        owner: 'John Doe',
        contact_info: 'john.doe@example.com',
        status: 'active',
        purchase_date: '2023-01-01',
        warranty_end_date: '2024-01-01',
        value: '5000',
        notes: 'Test Notes'
      };

      expect(assetCreateRequest.asset_name).toBe('Test Asset');
      expect(assetCreateRequest.asset_code).toBe('ASSET-001');
      expect(assetCreateRequest.organization_id).toBe('org1');
      expect(assetCreateRequest.asset_type).toBe('server');
    });

    it('should handle optional fields', () => {
      const minimalAssetCreateRequest: AssetCreateRequest = {
        asset_name: 'Test Asset',
        asset_code: 'ASSET-001',
        asset_description: '',
        organization_id: '',
        asset_type: '',
        asset_class: '',
        manufacturer: '',
        model: '',
        serial_number: '',
        ip_address: '',
        mac_address: '',
        location: '',
        department: '',
        owner: '',
        contact_info: '',
        status: '',
        purchase_date: '',
        warranty_end_date: '',
        value: '',
        notes: ''
      };

      expect(minimalAssetCreateRequest).toBeTruthy();
      expect(typeof minimalAssetCreateRequest.asset_name).toBe('string');
      expect(typeof minimalAssetCreateRequest.asset_code).toBe('string');
    });
  });

  describe('AssetUpdateRequest interface', () => {
    it('should create a valid AssetUpdateRequest object', () => {
      const assetUpdateRequest: AssetUpdateRequest = {
        asset_id: '1',
        asset_name: 'Updated Test Asset',
        asset_code: 'ASSET-001',
        asset_description: 'Updated Test Asset Description',
        organization_id: 'org1',
        asset_type: 'server',
        asset_class: 'hardware',
        manufacturer: 'Dell',
        model: 'PowerEdge R740',
        serial_number: 'ABC123',
        ip_address: '192.168.1.100',
        mac_address: '00:11:22:33:44:55',
        location: 'Data Center',
        department: 'IT',
        owner: 'John Doe',
        contact_info: 'john.doe@example.com',
        status: 'active',
        purchase_date: '2023-01-01',
        warranty_end_date: '2024-01-01',
        value: '5000',
        notes: 'Test Notes'
      };

      expect(assetUpdateRequest.asset_id).toBe('1');
      expect(assetUpdateRequest.asset_name).toBe('Updated Test Asset');
      expect(assetUpdateRequest.organization_id).toBe('org1');
      expect(assetUpdateRequest.asset_type).toBe('server');
    });

    it('should handle optional fields', () => {
      const minimalAssetUpdateRequest: AssetUpdateRequest = {
        asset_id: '1',
        asset_name: 'Updated Test Asset',
        asset_code: 'ASSET-001',
        asset_description: '',
        organization_id: '',
        asset_type: '',
        asset_class: '',
        manufacturer: '',
        model: '',
        serial_number: '',
        ip_address: '',
        mac_address: '',
        location: '',
        department: '',
        owner: '',
        contact_info: '',
        status: '',
        purchase_date: '',
        warranty_end_date: '',
        value: '',
        notes: ''
      };

      expect(minimalAssetUpdateRequest).toBeTruthy();
      expect(typeof minimalAssetUpdateRequest.asset_id).toBe('string');
      expect(typeof minimalAssetUpdateRequest.asset_name).toBe('string');
    });
  });
});