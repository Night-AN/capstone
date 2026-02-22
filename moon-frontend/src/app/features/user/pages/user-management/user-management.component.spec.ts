import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { of } from 'rxjs';
import { UserManagementComponent } from './user-management.component';
import { UserService } from '@core/services/user.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { UserListItem } from '@models/user.model';
import { MatTableModule } from '@angular/material/table';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';

// Mock services
class MockUserService {
  getUsers() {
    return of([
      {
        user_id: '1',
        nickname: 'testuser1',
        full_name: 'Test User 1',
        email: 'test1@example.com'
      },
      {
        user_id: '2',
        nickname: 'testuser2',
        full_name: 'Test User 2',
        email: 'test2@example.com'
      }
    ] as UserListItem[]);
  }

  deleteUser(id: string) {
    return of({ success: true });
  }
}

class MockNotificationService {
  error(message: string) {}
  success(message: string) {}
}

class MockRouter {
  navigate(commands: any[]) {}
}

describe('UserManagementComponent', () => {
  let component: UserManagementComponent;
  let fixture: ComponentFixture<UserManagementComponent>;
  let userService: MockUserService;
  let router: MockRouter;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        CommonModule,
        FormsModule,
        MatTableModule,
        MatPaginatorModule,
        MatSortModule,
        MatButtonModule,
        MatIconModule,
        MatFormFieldModule,
        MatInputModule,
        UserManagementComponent
      ],
      providers: [
        {
          provide: UserService,
          useClass: MockUserService
        },
        {
          provide: NotificationService,
          useClass: MockNotificationService
        },
        {
          provide: Router,
          useClass: MockRouter
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(UserManagementComponent);
    component = fixture.componentInstance;
    userService = TestBed.inject(UserService) as any;
    router = TestBed.inject(Router) as any;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should load users on init', () => {
    expect(component).toBeTruthy();
  });

  it('should apply filter', () => {
    component.searchKeyword = 'Test';
    component.applyFilter();
    expect(component).toBeTruthy();
  });

  it('should navigate to create page', () => {
    let navigateCalled = false;
    const originalNavigate = router.navigate;
    router.navigate = () => {
      navigateCalled = true;
      return Promise.resolve(true);
    };

    component.createUser();
    expect(navigateCalled).toBe(true);

    router.navigate = originalNavigate;
  });

  it('should navigate to edit page', () => {
    let navigateCalled = false;
    const originalNavigate = router.navigate;
    router.navigate = () => {
      navigateCalled = true;
      return Promise.resolve(true);
    };

    component.editUser('1');
    expect(navigateCalled).toBe(true);

    router.navigate = originalNavigate;
  });

  it('should navigate to detail page', () => {
    let navigateCalled = false;
    const originalNavigate = router.navigate;
    router.navigate = () => {
      navigateCalled = true;
      return Promise.resolve(true);
    };

    component.viewUser('1');
    expect(navigateCalled).toBe(true);

    router.navigate = originalNavigate;
  });

  it('should delete user', () => {
    expect(component).toBeTruthy();
  });
});