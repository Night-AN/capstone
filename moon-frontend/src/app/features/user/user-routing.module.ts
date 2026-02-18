import { Routes } from '@angular/router';
import { UserManagementComponent } from './pages/user-management/user-management.component';
import { UserDetailComponent } from './pages/user-detail/user-detail.component';
import { UserFormComponent } from './pages/user-form/user-form.component';

export const userRoutes: Routes = [
  {
    path: '',
    component: UserManagementComponent
  },
  {
    path: 'detail/:id',
    component: UserDetailComponent
  },
  {
    path: 'create',
    component: UserFormComponent
  },
  {
    path: 'edit/:id',
    component: UserFormComponent
  }
];
