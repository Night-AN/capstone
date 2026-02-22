import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { ResourceService } from './resource.service';
import { Resource } from '@models/resource.model';

describe('ResourceService', () => {
  let service: ResourceService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(ResourceService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should get resources', () => {
    const mockResources: Resource[] = [
      {
        resource_id: '7571e7cf-42b0-4471-9e6f-2f628e073f6f',
        name: 'ReadUser',
        description: 'Read A user profile',
        sensitive_flag: false,
        created_at: '2026-02-18T11:15:54Z',
        updated_at: '2026-02-18T11:15:54Z'
      },
      {
        resource_id: '9b446bde-d357-4ed6-8b36-c28b09ab6961',
        name: 'CreateUser',
        description: 'Create A user',
        sensitive_flag: false,
        created_at: '2026-02-18T11:15:54Z',
        updated_at: '2026-02-18T11:15:54Z'
      }
    ];

    service.getResources().subscribe(resources => {
      expect(resources).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/resources/list');
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: {
        resources: mockResources
      }
    });
  });

  it('should get resource by id', () => {
    const mockResource: Resource = {
      resource_id: '7571e7cf-42b0-4471-9e6f-2f628e073f6f',
      name: 'ReadUser',
      description: 'Read A user profile',
      sensitive_flag: false,
      created_at: '2026-02-18T11:15:54Z',
      updated_at: '2026-02-18T11:15:54Z'
    };

    service.getResourceById('7571e7cf-42b0-4471-9e6f-2f628e073f6f').subscribe(resource => {
      expect(resource).toBeTruthy();
    });

    const req = httpMock.expectOne((req) => {
      return req.url === '/api/v1/resources' && req.params.get('resource_id') === '7571e7cf-42b0-4471-9e6f-2f628e073f6f';
    });
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: mockResource
    });
  });

  it('should create resource', () => {
    const mockResourceCreate = {
      name: 'New Resource',
      description: 'New resource description'
    };

    const mockResource: Resource = {
      resource_id: 'new-id',
      name: 'New Resource',
      description: 'New resource description',
      sensitive_flag: false,
      created_at: '2026-02-22T00:00:00Z',
      updated_at: '2026-02-22T00:00:00Z'
    };

    service.createResource(mockResourceCreate).subscribe(resource => {
      expect(resource).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/resources');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual(mockResourceCreate);
    req.flush({
      code: '200',
      message: 'success',
      data: mockResource
    });
  });

  it('should update resource', () => {
    const mockResourceUpdate = {
      resource_id: '7571e7cf-42b0-4471-9e6f-2f628e073f6f',
      name: 'Updated Resource',
      description: 'Updated resource description'
    };

    const mockResource: Resource = {
      resource_id: '7571e7cf-42b0-4471-9e6f-2f628e073f6f',
      name: 'Updated Resource',
      description: 'Updated resource description',
      sensitive_flag: false,
      created_at: '2026-02-18T11:15:54Z',
      updated_at: '2026-02-22T00:00:00Z'
    };

    service.updateResource(mockResourceUpdate).subscribe(resource => {
      expect(resource).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/resources');
    expect(req.request.method).toBe('PUT');
    expect(req.request.body).toEqual(mockResourceUpdate);
    req.flush({
      code: '200',
      message: 'success',
      data: mockResource
    });
  });

  it('should delete resource', () => {
    service.deleteResource('7571e7cf-42b0-4471-9e6f-2f628e073f6f').subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/resources');
    expect(req.request.method).toBe('DELETE');
    req.flush({
      code: '200',
      message: 'success'
    });
  });
});