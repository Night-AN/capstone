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
import { userRoutes } from './user-routing.module';

// 组件
import { UserManagementComponent } from './pages/user-management/user-management.component';
import { UserDetailComponent } from './pages/user-detail/user-detail.component';
import { UserFormComponent } from './pages/user-form/user-form.component';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forChild(userRoutes)
  ],
  providers: []
})
export class UserModule { }
