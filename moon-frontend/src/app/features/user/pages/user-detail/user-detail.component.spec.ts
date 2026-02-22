import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ActivatedRoute, Router } from '@angular/router';
import { of } from 'rxjs';
import { UserDetailComponent } from './user-detail.component';
import { UserService } from '@core/services/user.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { User } from '@models/user.model';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { CommonModule } from '@angular/common';

// Mock services
class MockUserService {
  getUserById(id: string) {
    return of({
      user_id: id,
      nickname: 'testuser',
      full_name: 'Test User',
      email: 'test@example.com',
      organization_id: 'org1',
      status: 'active',
      created_at: '2023-01-01T00:00:00Z',
      updated_at: '2023-01-01T00:00:00Z'
    } as User);
  }

  getUserRoles(userId: string) {
    return of([]);
  }
}

class MockNotificationService {
  error(message: string) {}
  success(message: string) {}
}

class MockRouter {
  navigate(commands: any[]) {}
}

class MockActivatedRoute {
  snapshot = {
    paramMap: {
      get: (key: string) => '1'
    }
  };
}

describe('UserDetailComponent', () => {
  let component: UserDetailComponent;
  let fixture: ComponentFixture<UserDetailComponent>;
  let userService: MockUserService;
  let router: MockRouter;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        CommonModule,
        MatCardModule,
        MatButtonModule,
        MatIconModule,
        MatListModule,
        MatProgressSpinnerModule,
        UserDetailComponent
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
        },
        {
          provide: ActivatedRoute,
          useClass: MockActivatedRoute
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(UserDetailComponent);
    component = fixture.componentInstance;
    userService = TestBed.inject(UserService) as any;
    router = TestBed.inject(Router) as any;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should navigate back', () => {
    let navigateCalled = false;
    const originalNavigate = router.navigate;
    router.navigate = () => {
      navigateCalled = true;
      return Promise.resolve(true);
    };

    component.backToList();
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

    component.editUser();
    expect(navigateCalled).toBe(true);

    router.navigate = originalNavigate;
  });
});