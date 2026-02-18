import { Routes } from '@angular/router';
import { OrganizationManagementComponent } from './pages/organization-management/organization-management.component';
import { OrganizationDetailComponent } from './pages/organization-detail/organization-detail.component';
import { OrganizationFormComponent } from './pages/organization-form/organization-form.component';

export const organizationRoutes: Routes = [
  {
    path: '',
    component: OrganizationManagementComponent
  },
  {
    path: 'detail/:id',
    component: OrganizationDetailComponent
  },
  {
    path: 'create',
    component: OrganizationFormComponent
  },
  {
    path: 'edit/:id',
    component: OrganizationFormComponent
  }
];
