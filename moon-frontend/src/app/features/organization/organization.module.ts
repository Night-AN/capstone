import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatTableModule } from '@angular/material/table';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDialogModule } from '@angular/material/dialog';
import { MatIconModule } from '@angular/material/icon';

// 路由配置
import { organizationRoutes } from './organization-routing.module';

// 组件
import { OrganizationManagementComponent } from './pages/organization-management/organization-management.component';
import { OrganizationDetailComponent } from './pages/organization-detail/organization-detail.component';
import { OrganizationFormComponent } from './pages/organization-form/organization-form.component';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forChild(organizationRoutes)
  ],
  providers: []
})
export class OrganizationModule { }
