import { Routes } from '@angular/router';
import { PermissionManagementComponent } from './pages/permission-management/permission-management.component';
import { PermissionDetailComponent } from './pages/permission-detail/permission-detail.component';
import { PermissionFormComponent } from './pages/permission-form/permission-form.component';

export const permissionRoutes: Routes = [
  {
    path: '',
    component: PermissionManagementComponent
  },
  {
    path: 'detail/:id',
    component: PermissionDetailComponent
  },
  {
    path: 'create',
    component: PermissionFormComponent
  },
  {
    path: 'edit/:id',
    component: PermissionFormComponent
  }
];
