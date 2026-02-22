import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ActivatedRoute, Router } from '@angular/router';
import { of, throwError } from 'rxjs';
import { UserFormComponent } from './user-form.component';
import { UserService } from '@core/services/user.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { User, UserCreateRequest } from '@models/user.model';

// Mock services
class MockUserService {
  getUserById = vi.fn();
  createUser = vi.fn();
  updateUser = vi.fn();
}

class MockNotificationService {
  success = vi.fn();
  error = vi.fn();
}

class MockRouter {
  navigate = vi.fn();
}

class MockActivatedRoute {
  snapshot = {
    paramMap: {
      get: vi.fn()
    }
  };
}

describe('UserFormComponent', () => {
  let component: UserFormComponent;
  let fixture: ComponentFixture<UserFormComponent>;
  let mockUserService: MockUserService;
  let mockNotificationService: MockNotificationService;
  let mockRouter: MockRouter;
  let mockActivatedRoute: MockActivatedRoute;

  beforeEach(async () => {
    mockUserService = new MockUserService();
    mockNotificationService = new MockNotificationService();
    mockRouter = new MockRouter();
    mockActivatedRoute = new MockActivatedRoute();

    await TestBed.configureTestingModule({
      imports: [UserFormComponent],
      providers: [
        {
          provide: UserService,
          useValue: mockUserService
        },
        {
          provide: NotificationService,
          useValue: mockNotificationService
        },
        {
          provide: Router,
          useValue: mockRouter
        },
        {
          provide: ActivatedRoute,
          useValue: mockActivatedRoute
        }
      ]
    }).compileComponents();
  });

  it('should create', () => {
    mockActivatedRoute.snapshot.paramMap.get.mockReturnValue(null);
    fixture = TestBed.createComponent(UserFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });

  it('should initialize form with empty values for create mode', () => {
    mockActivatedRoute.snapshot.paramMap.get.mockReturnValue(null);
    fixture = TestBed.createComponent(UserFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    
    expect(component.isEditMode).toBe(false);
    expect(component.userId).toBe('');
    expect(component.userForm.get('nickname')?.value).toBe('');
    expect(component.userForm.get('full_name')?.value).toBe('');
    expect(component.userForm.get('email')?.value).toBe('');
    expect(component.userForm.get('password')?.value).toBe('');
  });

  describe('Edit mode', () => {
    beforeEach(() => {
      // 设置为编辑模式
      mockActivatedRoute.snapshot.paramMap.get.mockReturnValue('1');
      
      // 模拟用户数据
      const mockUser: User = {
        user_id: '1',
        nickname: 'testuser',
        full_name: 'Test User',
        email: 'test@example.com',
        created_at: '2023-01-01T00:00:00Z',
        updated_at: '2023-01-01T00:00:00Z'
      };
      
      mockUserService.getUserById.mockReturnValue(of(mockUser));
    });

    it('should initialize form with values for edit mode', () => {
      fixture = TestBed.createComponent(UserFormComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
      
      expect(component.isEditMode).toBe(true);
      expect(component.userId).toBe('1');
    });

    it('should submit form for update', () => {
      fixture = TestBed.createComponent(UserFormComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
      
      // 准备表单数据
      component.userForm.patchValue({
        nickname: 'updateduser',
        full_name: 'Updated User',
        email: 'updated@example.com',
        password: 'newpassword'
      });
      
      // 模拟更新成功
      mockUserService.updateUser.mockReturnValue(of({}));
      
      // 提交表单
      component.onSubmit();
      
      expect(mockUserService.updateUser).toHaveBeenCalled();
      expect(mockRouter.navigate).toHaveBeenCalled();
    });

    it('should handle error when loading user data fails', () => {
      // 模拟加载失败
      mockUserService.getUserById.mockReturnValue(throwError(() => new Error('Failed to load user')));
      
      fixture = TestBed.createComponent(UserFormComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
      
      expect(mockNotificationService.error).toHaveBeenCalled();
    });
  });

  it('should submit form for create', () => {
    mockActivatedRoute.snapshot.paramMap.get.mockReturnValue(null);
    fixture = TestBed.createComponent(UserFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    
    // 准备表单数据
    component.userForm.patchValue({
      nickname: 'newuser',
      full_name: 'New User',
      email: 'new@example.com',
      password: 'password123'
    });
    
    // 模拟创建成功
    const mockCreatedUser: User = {
      user_id: 'new-id',
      nickname: 'newuser',
      full_name: 'New User',
      email: 'new@example.com',
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString()
    };
    mockUserService.createUser.mockReturnValue(of(mockCreatedUser));
    
    // 提交表单
    component.onSubmit();
    
    expect(mockUserService.createUser).toHaveBeenCalled();
    expect(mockRouter.navigate).toHaveBeenCalled();
  });

  it('should cancel and navigate back', () => {
    mockActivatedRoute.snapshot.paramMap.get.mockReturnValue(null);
    fixture = TestBed.createComponent(UserFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    
    // 调用取消方法
    component.onCancel();
    
    expect(mockRouter.navigate).toHaveBeenCalled();
  });
});