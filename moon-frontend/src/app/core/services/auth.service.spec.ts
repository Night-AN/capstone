import { TestBed } from '@angular/core/testing';
import { HttpClientTestingModule, HttpTestingController } from '@angular/common/http/testing';
import { AuthService } from './auth.service';
import { AuthResponse, LoginCredentials, RegisterData } from '@models/auth.data';
import { StorageService } from './storage.service';
import { Router } from '@angular/router';

class MockStorageService {
  set(key: string, value: any): void {}
  get(key: string): any { return null; }
  remove(key: string): void {}
}

class MockRouter {
  navigate(commands: any[]): Promise<boolean> { return Promise.resolve(true); }
}

describe('AuthService', () => {
  let service: AuthService;
  let httpMock: HttpTestingController;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [
        {
          provide: StorageService,
          useClass: MockStorageService
        },
        {
          provide: Router,
          useClass: MockRouter
        }
      ]
    });
    service = TestBed.inject(AuthService);
    httpMock = TestBed.inject(HttpTestingController);
  });

  afterEach(() => {
    httpMock.verify();
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });

  it('should login successfully', () => {
    const loginData: LoginCredentials = {
      email: 'test@example.com',
      password: 'password123'
    };

    const mockResponse: AuthResponse = {
      token: 'mock-token-123',
      user: {
        UserID: 'user1',
        Username: 'testuser',
        Organization: 'org1',
        Role: ['Admin']
      }
    };

    service.login(loginData).subscribe(response => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpMock.expectOne('/api/v1/login');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual(loginData);
    req.flush(mockResponse);
  });

  it('should register successfully', () => {
    const registerData: RegisterData = {
      nickname: 'testuser',
      fullName: 'Test User',
      email: 'test@example.com',
      password: 'password123'
    };

    const mockResponse: AuthResponse = {
      token: 'mock-token-123',
      user: {
        UserID: 'user1',
        Username: 'testuser',
        Organization: 'org1',
        Role: ['User']
      }
    };

    service.register(registerData).subscribe(response => {
      expect(response).toEqual(mockResponse);
    });

    const req = httpMock.expectOne('/api/v1/register');
    expect(req.request.method).toBe('POST');
    expect(req.request.body).toEqual(registerData);
    req.flush(mockResponse);
  });

  it('should logout successfully', () => {
    // 直接调用logout方法，因为它是void返回类型
    expect(() => service.logout()).not.toThrow();
  });

  it('should check if user is authenticated', () => {
    const isAuthenticated = service.isAuthenticated();
    expect(typeof isAuthenticated).toBe('boolean');
  });
});