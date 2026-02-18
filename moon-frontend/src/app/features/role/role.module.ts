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
import { roleRoutes } from './role-routing.module';

// 组件
import { RoleManagementComponent } from './pages/role-management/role-management.component';
import { RoleDetailComponent } from './pages/role-detail/role-detail.component';
import { RoleFormComponent } from './pages/role-form/role-form.component';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forChild(roleRoutes)
  ],
  providers: []
})
export class RoleModule { }
