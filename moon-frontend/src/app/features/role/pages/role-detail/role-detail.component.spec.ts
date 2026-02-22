import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { RoleDetailComponent } from './role-detail.component';
import { RoleService } from '@core/services/role.service';
import { NotificationService } from '@shared/service/notification/notification.service';

class MockRoleService {
  getRoleById() {
    return of({
      role_id: '1',
      role_name: 'Test Role',
      role_code: 'test',
      role_flag: 'active',
      sensitive_flag: 'false',
      created_at: '2026-02-17T15:31:36.464358+08:00'
    });
  }
}

class MockNotificationService {
  error() {}
}

describe('RoleDetailComponent', () => {
  let component: RoleDetailComponent;
  let fixture: ComponentFixture<RoleDetailComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
        HttpClientTestingModule,
        MatButtonModule,
        MatCardModule,
        MatListModule
      ],
      providers: [
        {
          provide: RoleService,
          useClass: MockRoleService
        },
        {
          provide: NotificationService,
          useClass: MockNotificationService
        },
        {
          provide: ActivatedRoute,
          useValue: {
            snapshot: {
              paramMap: {
                get: () => '1'
              }
            }
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(RoleDetailComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});