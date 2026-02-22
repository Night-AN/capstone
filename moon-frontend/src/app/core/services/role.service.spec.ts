import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { RoleService } from './role.service';
import { Role } from '@models/role.model';

describe('RoleService', () => {
  let service: RoleService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(RoleService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should get roles', () => {
    const mockRoles = [
      {
        RoleID: '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0',
        RoleName: 'Test Role',
        Description: '',
        RoleCode: 'TEST_ROLE',
        RoleFlag: 'test',
        SensitiveFlag: false,
        CreatedAt: '2026-02-17T15:31:36.46541+08:00',
        UpdatedAt: '2026-02-17T15:31:36.46541+08:00'
      },
      {
        RoleID: 'c6d5c22e-1da6-4922-ad8d-011b991b221c',
        RoleName: 'Updated Test Role',
        Description: '',
        RoleCode: 'UPDATED_TEST_ROLE',
        RoleFlag: 'updated',
        SensitiveFlag: false,
        CreatedAt: '2026-02-17T15:31:36.467002+08:00',
        UpdatedAt: '2026-02-17T15:31:36.467002+08:00'
      }
    ];

    service.getRoles().subscribe(roles => {
      expect(roles).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/roles/list');
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: {
        Roles: mockRoles
      }
    });
  });

  it('should get role by id', () => {
    const mockRole = {
      RoleID: '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0',
      RoleName: 'Test Role',
      Description: '',
      RoleCode: 'TEST_ROLE',
      RoleFlag: 'test',
      SensitiveFlag: false,
      CreatedAt: '2026-02-17T15:31:36.46541+08:00',
      UpdatedAt: '2026-02-17T15:31:36.46541+08:00'
    };

    service.getRoleById('0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0').subscribe(role => {
      expect(role).toBeTruthy();
    });

    const req = httpMock.expectOne((request) => {
      return request.url === '/api/v1/roles' && request.params.get('role_id') === '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0';
    });
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: mockRole
    });
  });

  it('should create role', () => {
    const mockRoleCreate = {
      role_name: 'New Role',
      description: 'New role description',
      role_code: 'NEW_ROLE',
      role_flag: 'new',
      sensitive_flag: false
    };

    const mockResponse = {
      RoleID: 'new-id',
      RoleName: 'New Role',
      Description: 'New role description',
      RoleCode: 'NEW_ROLE',
      RoleFlag: 'new',
      SensitiveFlag: false,
      CreatedAt: '2026-02-22T00:00:00Z',
      UpdatedAt: '2026-02-22T00:00:00Z'
    };

    service.createRole(mockRoleCreate).subscribe(role => {
      expect(role).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/roles');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({
      RoleName: mockRoleCreate.role_name,
      Description: mockRoleCreate.description,
      RoleCode: mockRoleCreate.role_code,
      RoleFlag: mockRoleCreate.role_flag,
      SensitiveFlag: mockRoleCreate.sensitive_flag
    });
    req.flush({
      code: '200',
      message: 'success',
      data: mockResponse
    });
  });

  it('should update role', () => {
    const mockRoleUpdate = {
      role_id: '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0',
      role_name: 'Updated Role',
      description: 'Updated role description',
      role_code: 'UPDATED_ROLE',
      role_flag: 'updated',
      sensitive_flag: true
    };

    const mockResponse = {
      RoleID: '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0',
      RoleName: 'Updated Role',
      Description: 'Updated role description',
      RoleCode: 'UPDATED_ROLE',
      RoleFlag: 'updated',
      SensitiveFlag: true,
      CreatedAt: '2026-02-17T15:31:36.46541+08:00',
      UpdatedAt: '2026-02-22T00:00:00Z'
    };

    service.updateRole(mockRoleUpdate).subscribe(role => {
      expect(role).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/roles');
    expect(req.request.method).toBe('PUT');
    expect(req.request.body).toEqual({
      RoleID: mockRoleUpdate.role_id,
      RoleName: mockRoleUpdate.role_name,
      Description: mockRoleUpdate.description,
      RoleCode: mockRoleUpdate.role_code,
      RoleFlag: mockRoleUpdate.role_flag,
      SensitiveFlag: mockRoleUpdate.sensitive_flag
    });
    req.flush({
      code: '200',
      message: 'success',
      data: mockResponse
    });
  });

  it('should delete role', () => {
    const mockRoleDelete = {
      role_id: '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0'
    };

    service.deleteRole(mockRoleDelete).subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/roles');
    expect(req.request.method).toBe('DELETE');
    expect(req.request.body).toEqual({
      RoleID: mockRoleDelete.role_id
    });
    req.flush({
      code: '200',
      message: 'success',
      data: {
        Success: true
      }
    });
  });
});