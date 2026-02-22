import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ActivatedRoute, Router } from '@angular/router';
import { of, throwError } from 'rxjs';
import { RoleFormComponent } from './role-form.component';
import { RoleService } from '@core/services/role.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { Role } from '@models/role.model';

// Mock services
class MockRoleService {
  getRoleById = vi.fn();
  createRole = vi.fn();
  updateRole = vi.fn();
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

describe('RoleFormComponent', () => {
  let component: RoleFormComponent;
  let fixture: ComponentFixture<RoleFormComponent>;
  let mockRoleService: MockRoleService;
  let mockNotificationService: MockNotificationService;
  let mockRouter: MockRouter;
  let mockActivatedRoute: MockActivatedRoute;

  beforeEach(async () => {
    mockRoleService = new MockRoleService();
    mockNotificationService = new MockNotificationService();
    mockRouter = new MockRouter();
    mockActivatedRoute = new MockActivatedRoute();

    await TestBed.configureTestingModule({
      imports: [RoleFormComponent],
      providers: [
        {
          provide: RoleService,
          useValue: mockRoleService
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
    fixture = TestBed.createComponent(RoleFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });

  it('should initialize form with empty values for create mode', () => {
    mockActivatedRoute.snapshot.paramMap.get.mockReturnValue(null);
    fixture = TestBed.createComponent(RoleFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    
    expect(component.isEditMode).toBe(false);
    expect(component.roleId).toBe(null);
    expect(component.roleForm.get('role_name')?.value).toBe('');
    expect(component.roleForm.get('description')?.value).toBe('');
    expect(component.roleForm.get('sensitive_flag')?.value).toBe(false);
  });

  describe('Edit mode', () => {
    beforeEach(() => {
      // 设置为编辑模式
      mockActivatedRoute.snapshot.paramMap.get.mockReturnValue('1');
      
      // 模拟角色数据
      const mockRole: Role = {
        role_id: '1',
        role_name: 'Test Role',
        description: 'Test Description',
        sensitive_flag: true,
        created_at: '2023-01-01T00:00:00Z',
        updated_at: '2023-01-01T00:00:00Z'
      };
      
      mockRoleService.getRoleById.mockReturnValue(of(mockRole));
    });

    it('should initialize form with values for edit mode', () => {
      fixture = TestBed.createComponent(RoleFormComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
      
      expect(component.isEditMode).toBe(true);
      expect(component.roleId).toBe('1');
    });

    it('should submit form for update', () => {
      fixture = TestBed.createComponent(RoleFormComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
      
      // 准备表单数据
      component.roleForm.patchValue({
        role_name: 'Updated Test Role',
        description: 'Updated Test Description',
        sensitive_flag: true
      });
      
      // 模拟更新成功
      mockRoleService.updateRole.mockReturnValue(of({}));
      
      // 提交表单
      component.onSubmit();
      
      expect(mockRoleService.updateRole).toHaveBeenCalled();
      expect(mockRouter.navigate).toHaveBeenCalled();
    });

    it('should handle error when loading role data fails', () => {
      // 模拟加载失败
      mockRoleService.getRoleById.mockReturnValue(throwError(() => new Error('Failed to load role')));
      
      fixture = TestBed.createComponent(RoleFormComponent);
      component = fixture.componentInstance;
      fixture.detectChanges();
      
      expect(mockNotificationService.error).toHaveBeenCalled();
    });
  });

  it('should submit form for create', () => {
    mockActivatedRoute.snapshot.paramMap.get.mockReturnValue(null);
    fixture = TestBed.createComponent(RoleFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    
    // 准备表单数据
    component.roleForm.patchValue({
      role_name: 'New Test Role',
      description: 'New Test Description',
      sensitive_flag: false
    });
    
    // 模拟创建成功
    const mockCreatedRole: Role = {
      role_id: 'new-id',
      role_name: 'New Test Role',
      description: 'New Test Description',
      sensitive_flag: false,
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString()
    };
    mockRoleService.createRole.mockReturnValue(of(mockCreatedRole));
    
    // 提交表单
    component.onSubmit();
    
    expect(mockRoleService.createRole).toHaveBeenCalled();
    expect(mockRouter.navigate).toHaveBeenCalled();
  });

  it('should cancel and navigate back', () => {
    mockActivatedRoute.snapshot.paramMap.get.mockReturnValue(null);
    fixture = TestBed.createComponent(RoleFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    
    // 调用取消方法
    component.onCancel();
    
    expect(mockRouter.navigate).toHaveBeenCalled();
  });
});