import { Routes } from '@angular/router';
import { RoleManagementComponent } from './pages/role-management/role-management.component';
import { RoleDetailComponent } from './pages/role-detail/role-detail.component';
import { RoleFormComponent } from './pages/role-form/role-form.component';

export const roleRoutes: Routes = [
  {
    path: '',
    component: RoleManagementComponent
  },
  {
    path: 'detail/:id',
    component: RoleDetailComponent
  },
  {
    path: 'create',
    component: RoleFormComponent
  },
  {
    path: 'edit/:id',
    component: RoleFormComponent
  }
];
