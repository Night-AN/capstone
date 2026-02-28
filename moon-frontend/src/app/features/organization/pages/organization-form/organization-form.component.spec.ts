import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { ReactiveFormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatCardModule } from '@angular/material/card';
import { ActivatedRoute } from '@angular/router';
import { of, throwError } from 'rxjs';
import { OrganizationFormComponent } from './organization-form.component';
import { OrganizationService } from '@core/services/organization.service';
import { NotificationService } from '@shared/service/notification/notification.service';

class MockOrganizationService {
  createOrganization() {
    return of({});
  }

  updateOrganization() {
    return of({});
  }

  getOrganization() {
    return of({
      organization_id: '1',
      organization_name: 'Test Organization',
      organization_code: 'test',
      organization_flag: 'active'
    });
  }

  getOrganizationTree() {
    return of([]);
  }

  getOrganizationRoles() {
    return of([]);
  }

  assignRoleToOrganization() {
    return of({});
  }

  removeRoleFromOrganization() {
    return of({});
  }

  getOrganizationUsers() {
    return of([]);
  }

  moveOrganization() {
    return of({});
  }
}

class MockNotificationService {
  success() {}
  error() {}
}

describe('OrganizationFormComponent', () => {
  let component: OrganizationFormComponent;
  let fixture: ComponentFixture<OrganizationFormComponent>;
  let organizationService: OrganizationService;
  let notificationService: NotificationService;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
        HttpClientTestingModule,
        ReactiveFormsModule,
        MatButtonModule,
        MatFormFieldModule,
        MatInputModule,
        MatCardModule
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
                get: () => null
              }
            }
          }
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(OrganizationFormComponent);
    component = fixture.componentInstance;
    organizationService = TestBed.inject(OrganizationService);
    notificationService = TestBed.inject(NotificationService);
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should initialize form with default values', () => {
    expect(component.organizationForm).toBeTruthy();
    expect(component.organizationForm.controls['organization_name']).toBeTruthy();
    expect(component.organizationForm.controls['organization_code']).toBeTruthy();
    expect(component.organizationForm.controls['organization_flag']).toBeTruthy();
  });

  it('should mark form as invalid when required fields are empty', () => {
    component.organizationForm.controls['organization_name'].setValue('');
    component.organizationForm.controls['organization_code'].setValue('');
    component.organizationForm.controls['organization_flag'].setValue('');
    expect(component.organizationForm.invalid).toBeTruthy();
  });

  it('should mark form as valid when required fields are filled', () => {
    component.organizationForm.controls['organization_name'].setValue('Test');
    component.organizationForm.controls['organization_code'].setValue('test');
    component.organizationForm.controls['organization_flag'].setValue('active');
    expect(component.organizationForm.valid).toBeTruthy();
  });
});