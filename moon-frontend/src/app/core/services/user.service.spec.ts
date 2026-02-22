import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { UserService } from './user.service';
import { User } from '@models/user.model';

describe('UserService', () => {
  let service: UserService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule]
    });
    service = TestBed.inject(UserService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should get users', () => {
    const mockUsers: User[] = [
      {
        user_id: 'b23b9127-226e-4de6-aeb9-0df5cfcd2f52',
        nickname: 'Test User',
        full_name: 'Test Userf',
        email: 'test@example.com',
        created_at: '2026-02-17T15:31:36.000000+08:00',
        updated_at: '2026-02-17T15:31:36.000000+08:00'
      },
      {
        user_id: '00000000-0000-0000-0000-000000000000',
        nickname: 'test user',
        full_name: '111',
        email: 'test@test.com',
        created_at: '2026-02-17T15:31:36.000000+08:00',
        updated_at: '2026-02-17T15:31:36.000000+08:00'
      }
    ];

    service.getUsers().subscribe(users => {
      expect(users).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/users/list');
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: {
        users: mockUsers
      }
    });
  });

  it('should get user by id', () => {
    const mockUser: User = {
      user_id: 'b23b9127-226e-4de6-aeb9-0df5cfcd2f52',
      nickname: 'Test User',
      full_name: 'Test Userf',
      email: 'test@example.com',
      created_at: '2026-02-17T15:31:36.000000+08:00',
      updated_at: '2026-02-17T15:31:36.000000+08:00'
    };

    service.getUserById('b23b9127-226e-4de6-aeb9-0df5cfcd2f52').subscribe(user => {
      expect(user).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/users/b23b9127-226e-4de6-aeb9-0df5cfcd2f52');
    expect(req.request.method).toBe('GET');
    req.flush({
      code: '200',
      message: 'success',
      data: mockUser
    });
  });

  it('should create user', () => {
    const mockUserCreate = {
      nickname: 'New User',
      full_name: 'New User Full Name',
      email: 'newuser@example.com',
      password: 'password123'
    };

    const mockUser: User = {
      user_id: 'new-id',
      nickname: 'New User',
      full_name: 'New User Full Name',
      email: 'newuser@example.com',
      created_at: '2026-02-22T00:00:00Z',
      updated_at: '2026-02-22T00:00:00Z'
    };

    service.createUser(mockUserCreate).subscribe(user => {
      expect(user).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/users');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual(mockUserCreate);
    req.flush({
      code: '200',
      message: 'success',
      data: mockUser
    });
  });

  it('should update user', () => {
    const mockUserUpdate = {
      user_id: 'b23b9127-226e-4de6-aeb9-0df5cfcd2f52',
      nickname: 'Updated User',
      full_name: 'Updated User Full Name',
      email: 'updateduser@example.com'
    };

    const mockUser: User = {
      user_id: 'b23b9127-226e-4de6-aeb9-0df5cfcd2f52',
      nickname: 'Updated User',
      full_name: 'Updated User Full Name',
      email: 'updateduser@example.com',
      created_at: '2026-02-17T15:31:36.000000+08:00',
      updated_at: '2026-02-22T00:00:00Z'
    };

    service.updateUser(mockUserUpdate).subscribe(user => {
      expect(user).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/users');
    expect(req.request.method).toBe('PUT');
    expect(req.request.body).toEqual(mockUserUpdate);
    req.flush({
      code: '200',
      message: 'success',
      data: mockUser
    });
  });

  it('should delete user', () => {
    const mockUserDelete = {
      user_id: 'b23b9127-226e-4de6-aeb9-0df5cfcd2f52'
    };

    service.deleteUser(mockUserDelete).subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/users');
    expect(req.request.method).toBe('DELETE');
    expect(req.request.body).toEqual(mockUserDelete);
    req.flush({
      code: '200',
      message: 'success'
    });
  });

  it('should get user roles', () => {
    service.getUserRoles('b23b9127-226e-4de6-aeb9-0df5cfcd2f52').subscribe(roles => {
      expect(roles).toBeTruthy();
    });

    const req = httpMock.expectOne((req) => {
      return req.url === '/api/v1/users/roles' && req.params.get('user_id') === 'b23b9127-226e-4de6-aeb9-0df5cfcd2f52';
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

  it('should assign role to user', () => {
    const userId = 'b23b9127-226e-4de6-aeb9-0df5cfcd2f52';
    const roleId = '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0';

    service.assignRoleToUser(userId, roleId).subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/users/assign-role');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ user_id: userId, role_id: roleId });
    req.flush({
      code: '200',
      message: 'success'
    });
  });

  it('should remove role from user', () => {
    const userId = 'b23b9127-226e-4de6-aeb9-0df5cfcd2f52';
    const roleId = '0ae3c63a-7ee9-44d9-8daf-aa70b992d8e0';

    service.removeRoleFromUser(userId, roleId).subscribe(success => {
      expect(success).toBeTruthy();
    });

    const req = httpMock.expectOne('/api/v1/users/remove-role');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual({ user_id: userId, role_id: roleId });
    req.flush({
      code: '200',
      message: 'success'
    });
  });
});