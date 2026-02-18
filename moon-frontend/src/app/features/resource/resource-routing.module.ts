import { Routes } from '@angular/router';
import { ResourceManagementComponent } from './pages/resource-management/resource-management.component';
import { ResourceDetailComponent } from './pages/resource-detail/resource-detail.component';
import { ResourceFormComponent } from './pages/resource-form/resource-form.component';

export const resourceRoutes: Routes = [
  {
    path: '',
    component: ResourceManagementComponent
  },
  {
    path: 'detail/:id',
    component: ResourceDetailComponent
  },
  {
    path: 'create',
    component: ResourceFormComponent
  },
  {
    path: 'edit/:id',
    component: ResourceFormComponent
  }
];
