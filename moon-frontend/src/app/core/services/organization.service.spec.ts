import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { OrganizationService } from './organization.service';
import { Organization } from '@models/organization.model';

describe('OrganizationService', () => {
  let service: OrganizationService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(OrganizationService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should get organizations', () => {
    const mockOrganizations = [
      {
        OrganizationID: 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3',
        OrganizationName: 'Test Organization',
        OrganizationCode: 'test',
        OrganizationFlag: 'active',
        CreatedAt: '2026-02-17T15:31:36.424362+08:00'
      },
      {
        OrganizationID: '11111111-1111-1111-1111-111111111111',
        OrganizationName: 'test',
        OrganizationCode: 'test',
        OrganizationFlag: 'active',
        CreatedAt: '2026-02-17T15:31:36.000000+08:00'
      }
    ];

    service.getOrganizations().subscribe(organizations => {
      expect(organizations).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/organizations/list');
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: {
        Organizations: mockOrganizations
      }
    });
  });

  it('should get organization by id', () => {
    const mockOrganization = {
      OrganizationID: 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3',
      OrganizationName: 'Test Organization',
      OrganizationCode: 'test',
      OrganizationDescription: '',
      OrganizationFlag: 'active',
      CreatedAt: '2026-02-17T15:31:36.424362+08:00',
      UpdatedAt: '2026-02-17T15:31:36.424362+08:00'
    };

    service.getOrganizationById('d5933bf5-d5b3-4546-9663-b7cd1f6c74f3').subscribe(organization => {
      expect(organization).toBeTruthy();
    });

    const req = httpMock.expectOne((req) => {
      return req.url === '/api/v1/organizations' && req.params.get('organization_id') === 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3';
    });
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: mockOrganization
    });
  });

  it('should create organization', () => {
    const mockOrganizationCreate = {
      organization_name: 'New Organization',
      organization_code: 'new-org',
      organization_description: 'New organization description'
    };

    const mockResponse = {
      OrganizationID: 'new-id',
      OrganizationName: 'New Organization',
      OrganizationCode: 'new-org',
      OrganizationDescription: 'New organization description',
      OrganizationFlag: '',
      CreatedAt: '2026-02-22T00:00:00Z',
      UpdatedAt: '2026-02-22T00:00:00Z'
    };

    service.createOrganization(mockOrganizationCreate).subscribe(organization => {
      expect(organization).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/organizations');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({
      OrganizationName: mockOrganizationCreate.organization_name,
      OrganizationCode: mockOrganizationCreate.organization_code,
      OrganizationDescription: mockOrganizationCreate.organization_description,
      OrganizationFlag: undefined,
      ParentID: undefined
    });
    req.flush({
      code: '200',
      message: 'success',
      data: mockResponse
    });
  });

  it('should update organization', () => {
    const mockOrganizationUpdate = {
      organization_id: 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3',
      organization_name: 'Updated Organization',
      organization_code: 'updated-org',
      organization_description: 'Updated organization description'
    };

    const mockResponse = {
      OrganizationID: 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3',
      OrganizationName: 'Updated Organization',
      OrganizationCode: 'updated-org',
      OrganizationDescription: 'Updated organization description',
      OrganizationFlag: '',
      CreatedAt: '2026-02-17T15:31:36.424362+08:00',
      UpdatedAt: '2026-02-22T00:00:00Z'
    };

    service.updateOrganization(mockOrganizationUpdate).subscribe(organization => {
      expect(organization).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/organizations');
    expect(req.request.method).toBe('PUT');
    expect(req.request.body).toEqual({
      OrganizationID: mockOrganizationUpdate.organization_id,
      OrganizationName: mockOrganizationUpdate.organization_name,
      OrganizationCode: mockOrganizationUpdate.organization_code,
      OrganizationDescription: mockOrganizationUpdate.organization_description,
      OrganizationFlag: undefined,
      ParentID: undefined
    });
    req.flush({
      code: '200',
      message: 'success',
      data: mockResponse
    });
  });

  it('should delete organization', () => {
    const mockOrganizationDelete = {
      organization_id: 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3'
    };

    service.deleteOrganization(mockOrganizationDelete).subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/organizations');
    expect(req.request.method).toBe('DELETE');
    expect(req.request.body).toEqual({
      OrganizationID: mockOrganizationDelete.organization_id
    });
    req.flush({
      code: '200',
      message: 'success'
    });
  });

  it('should assign role to organization', () => {
    const organizationId = 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3';
    const roleId = '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0';

    service.assignRoleToOrganization(organizationId, roleId).subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/organizations/assign-role');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ OrganizationID: organizationId, RoleID: roleId });
    req.flush({
      code: '200',
      message: 'success'
    });
  });

  it('should remove role from organization', () => {
    const organizationId = 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3';
    const roleId = '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0';

    service.removeRoleFromOrganization(organizationId, roleId).subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/organizations/remove-role');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ OrganizationID: organizationId, RoleID: roleId });
    req.flush({
      code: '200',
      message: 'success'
    });
  });

  it('should get organization roles', () => {
    service.getOrganizationRoles('d5933bf5-d5b3-4546-9663-b7cd1f6c74f3').subscribe(roles => {
      expect(roles).toBeTruthy();
    });

    const req = httpMock.expectOne((req) => {
      return req.url === '/api/v1/organizations/roles' && req.params.get('organization_id') === 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3';
    });
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: {
        roles: []
      }
    });
  });

  it('should move organization', () => {
    const organizationId = 'd5933bf5-d5b3-4546-9663-b7cd1f6c74f3';
    const newParentId = '11111111-1111-1111-1111-111111111111';

    service.moveOrganization(organizationId, newParentId).subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/organizations/move');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ OrganizationID: organizationId, NewParentID: newParentId });
    req.flush({
      code: '200',
      message: 'success'
    });
  });
});