import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatListModule } from '@angular/material/list';
import { ActivatedRoute } from '@angular/router';
import { of } from 'rxjs';
import { OrganizationDetailComponent } from './organization-detail.component';
import { OrganizationService } from '@core/services/organization.service';
import { NotificationService } from '@shared/service/notification/notification.service';

class MockOrganizationService {
  getOrganizationById() {
    return of({
      organization_id: '1',
      organization_name: 'Test Organization',
      organization_code: 'test',
      organization_flag: 'active',
      created_at: '2026-02-17T15:31:36.424362+08:00'
    });
  }

  getOrganizationUsers() {
    return of([
      {
        user_id: '1',
        nickname: 'testuser',
        full_name: 'Test User',
        email: 'test@example.com'
      }
    ]);
  }
}

class MockNotificationService {
  error() {}
}

describe('OrganizationDetailComponent', () => {
  let component: OrganizationDetailComponent;
  let fixture: ComponentFixture<OrganizationDetailComponent>;

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
          provide: OrganizationService,
          useClass: MockOrganizationService
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

    fixture = TestBed.createComponent(OrganizationDetailComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});