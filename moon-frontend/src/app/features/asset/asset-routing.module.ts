import { Routes } from '@angular/router';
import { AssetManagementComponent } from './pages/asset-management/asset-management.component';
import { AssetDetailComponent } from './pages/asset-detail/asset-detail.component';
import { AssetFormComponent } from './pages/asset-form/asset-form.component';

export const assetRoutes: Routes = [
  {
    path: '',
    component: AssetManagementComponent
  },
  {
    path: 'detail/:id',
    component: AssetDetailComponent
  },
  {
    path: 'create',
    component: AssetFormComponent
  },
  {
    path: 'edit/:id',
    component: AssetFormComponent
  }
];
