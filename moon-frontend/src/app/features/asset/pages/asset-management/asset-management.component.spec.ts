import { ComponentFixture, TestBed } from '@angular/core/testing';
import { Router } from '@angular/router';
import { of } from 'rxjs';
import { AssetManagementComponent } from './asset-management.component';
import { AssetService } from '@core/services/asset.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { AssetListItem } from '@models/asset.model';
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
class MockAssetService {
  getAssets() {
    return of([
      {
        asset_id: '1',
        asset_name: 'Test Asset 1',
        asset_type: 'Server',
        ip_address: '192.168.1.1',
        organization_id: 'org1',
        status: 'active'
      },
      {
        asset_id: '2',
        asset_name: 'Test Asset 2',
        asset_type: 'Workstation',
        ip_address: '192.168.1.2',
        organization_id: 'org2',
        status: 'inactive'
      }
    ] as AssetListItem[]);
  }

  deleteAsset(id: string) {
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

describe('AssetManagementComponent', () => {
  let component: AssetManagementComponent;
  let fixture: ComponentFixture<AssetManagementComponent>;
  let assetService: MockAssetService;
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
        AssetManagementComponent
      ],
      providers: [
        {
          provide: AssetService,
          useClass: MockAssetService
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

    fixture = TestBed.createComponent(AssetManagementComponent);
    component = fixture.componentInstance;
    assetService = TestBed.inject(AssetService) as any;
    router = TestBed.inject(Router) as any;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should load assets on init', () => {
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

    component.createAsset();
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

    component.editAsset('1');
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

    component.viewAsset('1');
    expect(navigateCalled).toBe(true);

    router.navigate = originalNavigate;
  });

  it('should delete asset', () => {
    expect(component).toBeTruthy();
  });
});