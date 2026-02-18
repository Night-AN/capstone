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
import { resourceRoutes } from './resource-routing.module';

// 组件
import { ResourceManagementComponent } from './pages/resource-management/resource-management.component';
import { ResourceDetailComponent } from './pages/resource-detail/resource-detail.component';
import { ResourceFormComponent } from './pages/resource-form/resource-form.component';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forChild(resourceRoutes)
  ],
  providers: []
})
export class ResourceModule { }
