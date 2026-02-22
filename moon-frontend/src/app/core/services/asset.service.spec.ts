import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { AssetService } from './asset.service';
import { Asset, AssetCreateRequest, AssetUpdateRequest, AssetListItem } from '@models/asset.model';

describe('AssetService', () => {
  let service: AssetService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(AssetService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should get assets', () => {
    const mockAssets: AssetListItem[] = [
      {
        asset_id: 'a1',
        asset_name: 'Server 1',
        asset_code: 'SRV-001',
        organization_id: 'org1',
        asset_type: 'Server',
        asset_class: 'Hardware',
        ip_address: '192.168.1.100',
        status: 'Active',
        created_at: '2026-01-01T00:00:00Z'
      },
      {
        asset_id: 'a2',
        asset_name: 'Server 2',
        asset_code: 'SRV-002',
        organization_id: 'org1',
        asset_type: 'Server',
        asset_class: 'Hardware',
        ip_address: '192.168.1.101',
        status: 'Active',
        created_at: '2026-01-02T00:00:00Z'
      }
    ];

    service.getAssets().subscribe(assets => {
      expect(assets).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/assets/list');
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: {
        assets: mockAssets
      }
    });
  });

  it('should get asset by id', () => {
    const mockAsset: Asset = {
      asset_id: 'a1',
      asset_name: 'Server 1',
      asset_code: 'SRV-001',
      asset_description: 'Main server',
      organization_id: 'org1',
      asset_type: 'Server',
      asset_class: 'Hardware',
      manufacturer: 'Dell',
      model: 'PowerEdge R740',
      serial_number: 'ABC123',
      ip_address: '192.168.1.100',
      mac_address: '00:11:22:33:44:55',
      location: 'Datacenter A',
      department: 'IT',
      owner: 'John Doe',
      contact_info: 'john.doe@example.com',
      status: 'Active',
      purchase_date: '2026-01-01T00:00:00Z',
      warranty_end_date: '2027-01-01T00:00:00Z',
      value: '10000',
      notes: 'Production server',
      created_at: '2026-01-01T00:00:00Z',
      updated_at: '2026-01-01T00:00:00Z'
    };

    service.getAssetById('a1').subscribe(asset => {
      expect(asset).toBeTruthy();
    });

    const req = httpMock.expectOne((request) => {
      return request.url === '/api/v1/assets' && request.params.get('asset_id') === 'a1';
    });
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: mockAsset
    });
  });

  it('should create asset', () => {
    const mockAssetCreate: AssetCreateRequest = {
      asset_name: 'Server 1',
      asset_code: 'SRV-001',
      asset_description: 'Main server',
      organization_id: 'org1',
      asset_type: 'Server',
      asset_class: 'Hardware',
      manufacturer: 'Dell',
      model: 'PowerEdge R740',
      serial_number: 'ABC123',
      ip_address: '192.168.1.100',
      mac_address: '00:11:22:33:44:55',
      location: 'Datacenter A',
      department: 'IT',
      owner: 'John Doe',
      contact_info: 'john.doe@example.com',
      status: 'Active',
      purchase_date: '2026-01-01T00:00:00Z',
      warranty_end_date: '2027-01-01T00:00:00Z',
      value: '10000',
      notes: 'Production server'
    };

    const mockAsset: Asset = {
      asset_id: 'a1',
      asset_name: 'Server 1',
      asset_code: 'SRV-001',
      asset_description: 'Main server',
      organization_id: 'org1',
      asset_type: 'Server',
      asset_class: 'Hardware',
      manufacturer: 'Dell',
      model: 'PowerEdge R740',
      serial_number: 'ABC123',
      ip_address: '192.168.1.100',
      mac_address: '00:11:22:33:44:55',
      location: 'Datacenter A',
      department: 'IT',
      owner: 'John Doe',
      contact_info: 'john.doe@example.com',
      status: 'Active',
      purchase_date: '2026-01-01T00:00:00Z',
      warranty_end_date: '2027-01-01T00:00:00Z',
      value: '10000',
      notes: 'Production server',
      created_at: '2026-01-01T00:00:00Z',
      updated_at: '2026-01-01T00:00:00Z'
    };

    service.createAsset(mockAssetCreate).subscribe(asset => {
      expect(asset).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/assets');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual(mockAssetCreate);
    req.flush({
      code: '200',
      message: 'success',
      data: mockAsset
    });
  });

  it('should update asset', () => {
    const mockAssetUpdate: AssetUpdateRequest = {
      asset_id: 'a1',
      asset_name: 'Updated Server 1',
      asset_code: 'SRV-001',
      asset_description: 'Main server',
      organization_id: 'org1',
      asset_type: 'Server',
      asset_class: 'Hardware',
      manufacturer: 'Dell',
      model: 'PowerEdge R740',
      serial_number: 'ABC123',
      ip_address: '192.168.1.100',
      mac_address: '00:11:22:33:44:55',
      location: 'Datacenter A',
      department: 'IT',
      owner: 'John Doe',
      contact_info: 'john.doe@example.com',
      status: 'Active',
      purchase_date: '2026-01-01T00:00:00Z',
      warranty_end_date: '2027-01-01T00:00:00Z',
      value: '10000',
      notes: 'Production server'
    };

    const mockAsset: Asset = {
      asset_id: 'a1',
      asset_name: 'Updated Server 1',
      asset_code: 'SRV-001',
      asset_description: 'Main server',
      organization_id: 'org1',
      asset_type: 'Server',
      asset_class: 'Hardware',
      manufacturer: 'Dell',
      model: 'PowerEdge R740',
      serial_number: 'ABC123',
      ip_address: '192.168.1.100',
      mac_address: '00:11:22:33:44:55',
      location: 'Datacenter A',
      department: 'IT',
      owner: 'John Doe',
      contact_info: 'john.doe@example.com',
      status: 'Active',
      purchase_date: '2026-01-01T00:00:00Z',
      warranty_end_date: '2027-01-01T00:00:00Z',
      value: '10000',
      notes: 'Production server',
      created_at: '2026-01-01T00:00:00Z',
      updated_at: '2026-01-02T00:00:00Z'
    };

    service.updateAsset(mockAssetUpdate).subscribe(asset => {
      expect(asset).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/assets');
    expect(req.request.method).toBe('PUT');
    expect(req.request.body).toEqual(mockAssetUpdate);
    req.flush({
      code: '200',
      message: 'success',
      data: mockAsset
    });
  });

  it('should delete asset', () => {
    service.deleteAsset('a1').subscribe(success => {
      expect(success).toBe(true);
    });

    const req = httpMock.expectOne('/api/v1/assets');
    expect(req.request.method).toBe('DELETE');
    expect(req.request.body).toEqual({ asset_id: 'a1' });
    req.flush({
      code: '200',
      message: 'success'
    });
  });
});