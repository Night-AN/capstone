import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { PermissionService } from './permission.service';
import { Permission } from '@models/permission.model';

describe('PermissionService', () => {
  let service: PermissionService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(PermissionService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should get permissions', () => {
    const mockPermissions = [
      {
        permission_id: '9f16dfe8-2869-4b88-9dec-7e5ddd96a2a3',
        permission_name: 'UserRead',
        permission_code: 'user:read',
        sensitive_flag: true,
        created_at: '2026-02-17T15:31:36.000000+08:00',
        updated_at: '2026-02-17T15:31:36.000000+08:00'
      },
      {
        permission_id: 'b7146960-ec40-47a0-bb48-841b15698625',
        permission_name: 'UserWrite',
        permission_code: 'user:write',
        sensitive_flag: true,
        created_at: '2026-02-17T15:31:36.000000+08:00',
        updated_at: '2026-02-17T15:31:36.000000+08:00'
      }
    ];

    service.getPermissions().subscribe(permissions => {
      expect(permissions).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/permissions/list');
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: {
        permissions: mockPermissions
      }
    });
  });

  it('should get permission by id', () => {
    const mockPermissionData = {
      permission_id: '9f16dfe8-2869-4b88-9dec-7e5ddd96a2a3',
      permission_name: 'UserRead',
      permission_code: 'user:read',
      sensitive_flag: true,
      created_at: '2026-02-17T15:31:36.000000+08:00',
      updated_at: '2026-02-17T15:31:36.000000+08:00'
    };

    service.getPermissionById('9f16dfe8-2869-4b88-9dec-7e5ddd96a2a3').subscribe(permission => {
      expect(permission).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/permissions/9f16dfe8-2869-4b88-9dec-7e5ddd96a2a3');
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: mockPermissionData
    });
  });

  it('should create permission', () => {
    const mockPermissionCreate = {
      name: 'New Permission',
      description: 'New permission description'
    };

    const mockPermissionData = {
      permission_id: 'new-id',
      permission_name: 'New Permission',
      permission_code: 'new_permission',
      sensitive_flag: false,
      created_at: '2026-02-22T00:00:00Z',
      updated_at: '2026-02-22T00:00:00Z'
    };

    service.createPermission(mockPermissionCreate).subscribe(permission => {
      expect(permission).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/permissions');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({
      permission_name: 'New Permission',
      description: 'New permission description',
      permission_code: 'new_permission',
      sensitive_flag: false
    });
    req.flush({
      code: '200',
      message: 'success',
      data: mockPermissionData
    });
  });

  it('should update permission', () => {
    const mockPermissionUpdate = {
      permission_id: '9f16dfe8-2869-4b88-9dec-7e5ddd96a2a3',
      name: 'Updated Permission',
      description: 'Updated permission description'
    };

    const mockPermissionData = {
      permission_id: '9f16dfe8-2869-4b88-9dec-7e5ddd96a2a3',
      permission_name: 'Updated Permission',
      permission_code: 'updated_permission',
      sensitive_flag: true,
      created_at: '2026-02-17T15:31:36.000000+08:00',
      updated_at: '2026-02-22T00:00:00Z'
    };

    service.updatePermission(mockPermissionUpdate).subscribe(permission => {
      expect(permission).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/permissions');
    expect(req.request.method).toBe('PUT');
    expect(req.request.body).toEqual({
      permission_id: '9f16dfe8-2869-4b88-9dec-7e5ddd96a2a3',
      permission_name: 'Updated Permission',
      description: 'Updated permission description',
      permission_code: 'updated_permission',
      sensitive_flag: false
    });
    req.flush({
      code: '200',
      message: 'success',
      data: mockPermissionData
    });
  });

  it('should delete permission', () => {
    service.deletePermission('9f16dfe8-2869-4b88-9dec-7e5ddd96a2a3').subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/permissions');
    expect(req.request.method).toBe('DELETE');
    expect(req.request.body).toEqual({ permission_id: '9f16dfe8-2869-4b88-9dec-7e5ddd96a2a3' });
    req.flush({
      code: '200',
      message: 'success',
      data: {
        success: true
      }
    });
  });
});