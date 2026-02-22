import { ComponentFixture, TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { MatTableModule } from '@angular/material/table';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { FormsModule } from '@angular/forms';
import { of } from 'rxjs';
import { OrganizationManagementComponent } from './organization-management.component';
import { OrganizationService } from '@core/services/organization.service';
import { NotificationService } from '@shared/service/notification/notification.service';

class MockOrganizationService {
  getOrganizations() {
    return of([]);
  }

  deleteOrganization() {
    return of({});
  }
}

class MockNotificationService {
  success() {}
  error() {}
}

describe('OrganizationManagementComponent', () => {
  let component: OrganizationManagementComponent;
  let fixture: ComponentFixture<OrganizationManagementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        RouterTestingModule,
        HttpClientTestingModule,
        MatTableModule,
        MatPaginatorModule,
        MatSortModule,
        MatButtonModule,
        MatIconModule,
        MatFormFieldModule,
        MatInputModule,
        MatProgressSpinnerModule,
        FormsModule,
        OrganizationManagementComponent
      ],
      providers: [
        {
          provide: OrganizationService,
          useClass: MockOrganizationService
        },
        {
          provide: NotificationService,
          useClass: MockNotificationService
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(OrganizationManagementComponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should apply filter', () => {
    component.searchKeyword = 'test';
    component.applyFilter();
    expect(component.dataSource.filter).toBe('test');
  });
});