import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ActivatedRoute, Router } from '@angular/router';
import { of } from 'rxjs';
import { AssetDetailComponent } from './asset-detail.component';
import { AssetService } from '@core/services/asset.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { Asset } from '@models/asset.model';
import { MatCardModule } from '@angular/material/card';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
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

describe('AssetDetailComponent', () => {
  let component: AssetDetailComponent;
  let fixture: ComponentFixture<AssetDetailComponent>;
  let assetService: MockAssetService;
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
        AssetDetailComponent
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
        },
        {
          provide: ActivatedRoute,
          useClass: MockActivatedRoute
        }
      ]
    }).compileComponents();

    fixture = TestBed.createComponent(AssetDetailComponent);
    component = fixture.componentInstance;
    assetService = TestBed.inject(AssetService) as any;
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

    component.editAsset();
    expect(navigateCalled).toBe(true);

    router.navigate = originalNavigate;
  });
});