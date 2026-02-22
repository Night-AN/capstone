import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ActivatedRoute, Router } from '@angular/router';
import { of } from 'rxjs';
import { AssetFormComponent } from './asset-form.component';
import { AssetService } from '@core/services/asset.service';
import { OrganizationService } from '@core/services/organization.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { Asset } from '@models/asset.model';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';

// Mock services
class MockAssetService {
  getAssetById(id: string) {
    return of({
      asset_id: id,
      asset_name: 'Test Asset',
      asset_type: 'Server',
      ip_address: '192.168.1.1',
      organization_id: 'org1',
      status: 'active',
      created_at: '2023-01-01T00:00:00Z',
      updated_at: '2023-01-01T00:00:00Z'
    } as Asset);
  }

  updateAsset(asset: any) {
    return of({
      asset_id: asset.asset_id,
      asset_name: asset.asset_name,
      asset_type: asset.asset_type,
      ip_address: asset.ip_address,
      organization_id: asset.organization_id,
      status: asset.status,
      created_at: '2023-01-01T00:00:00Z',
      updated_at: new Date().toISOString()
    } as Asset);
  }

  createAsset(asset: any) {
    return of({
      asset_id: '1',
      asset_name: asset.asset_name,
      asset_type: asset.asset_type,
      ip_address: asset.ip_address,
      organization_id: asset.organization_id,
      status: asset.status,
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString()
    } as Asset);
  }
}

class MockOrganizationService {
  getOrganizations() {
    return of([
      { organization_id: 'org1', organization_name: 'Test Org 1' },
      { organization_id: 'org2', organization_name: 'Test Org 2' }
    ]);
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
      get: (key: string) => null
    }
  };
}

class MockActivatedRouteWithId {
  snapshot = {
    paramMap: {
      get: (key: string) => '1'
    }
  };
}

describe('AssetFormComponent', () => {
  let component: AssetFormComponent;
  let fixture: ComponentFixture<AssetFormComponent>;
  let assetService: MockAssetService;
  let router: MockRouter;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        CommonModule,
        ReactiveFormsModule,
        MatFormFieldModule,
        MatInputModule,
        MatSelectModule,
        MatButtonModule,
        MatCardModule,
        MatIconModule,
        MatProgressSpinnerModule,
        AssetFormComponent
      ],
      providers: [
        {
          provide: AssetService,
          useClass: MockAssetService
        },
        {
          provide: OrganizationService,
          useClass: MockOrganizationService
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
  });

  it('should create', () => {
    fixture = TestBed.createComponent(AssetFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    expect(component).toBeTruthy();
  });

  it('should initialize form with empty values for create mode', () => {
    fixture = TestBed.createComponent(AssetFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
    expect(component.isEditMode).toBe(false);
    expect(component.assetId).toBeNull();
    expect(component.assetForm.get('asset_name')?.value).toBe('');
    expect(component.assetForm.get('asset_type')?.value).toBe('');
  });

  it('should submit form for create', () => {
    fixture = TestBed.createComponent(AssetFormComponent);
    component = fixture.componentInstance;
    assetService = TestBed.inject(AssetService) as any;
    router = TestBed.inject(Router) as any;
    fixture.detectChanges();

    let navigateCalled = false;
    const originalNavigate = router.navigate;
    router.navigate = () => {
      navigateCalled = true;
      return Promise.resolve(true);
    };

    component.assetForm.patchValue({
      asset_name: 'New Test Asset',
      asset_code: 'ASSET-001',
      asset_type: 'Server',
      asset_class: 'Hardware',
      ip_address: '192.168.1.1',
      organization_id: 'org1',
      status: 'active'
    });

    component.onSubmit();
    expect(navigateCalled).toBe(true);

    router.navigate = originalNavigate;
  });

  it('should cancel and navigate back', () => {
    fixture = TestBed.createComponent(AssetFormComponent);
    component = fixture.componentInstance;
    router = TestBed.inject(Router) as any;
    fixture.detectChanges();

    let navigateCalled = false;
    const originalNavigate = router.navigate;
    router.navigate = () => {
      navigateCalled = true;
      return Promise.resolve(true);
    };

    component.onCancel();
    expect(navigateCalled).toBe(true);

    router.navigate = originalNavigate;
  });
});

// Test for edit mode
describe('AssetFormComponent (Edit Mode)', () => {
  let component: AssetFormComponent;
  let fixture: ComponentFixture<AssetFormComponent>;
  let assetService: MockAssetService;
  let router: MockRouter;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [
        CommonModule,
        ReactiveFormsModule,
        MatFormFieldModule,
        MatInputModule,
        MatSelectModule,
        MatButtonModule,
        MatCardModule,
        MatIconModule,
        MatProgressSpinnerModule,
        AssetFormComponent
      ],
      providers: [
        {
          provide: AssetService,
          useClass: MockAssetService
        },
        {
          provide: OrganizationService,
          useClass: MockOrganizationService
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
          useClass: MockActivatedRouteWithId
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(AssetFormComponent);
    component = fixture.componentInstance;
    assetService = TestBed.inject(AssetService) as any;
    router = TestBed.inject(Router) as any;
    fixture.detectChanges();
  });

  it('should initialize form with values for edit mode', () => {
    expect(component.isEditMode).toBe(true);
    expect(component.assetId).toBe('1');
  });

  it('should submit form for update', () => {
    let navigateCalled = false;
    const originalNavigate = router.navigate;
    router.navigate = () => {
      navigateCalled = true;
      return Promise.resolve(true);
    };

    component.assetForm.patchValue({
      asset_name: 'Updated Test Asset',
      asset_code: 'ASSET-001',
      asset_type: 'Server',
      asset_class: 'Hardware',
      ip_address: '192.168.1.1',
      organization_id: 'org1',
      status: 'active'
    });

    component.onSubmit();
    expect(navigateCalled).toBe(true);

    router.navigate = originalNavigate;
  });
});