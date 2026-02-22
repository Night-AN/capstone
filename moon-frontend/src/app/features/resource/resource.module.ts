import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';

// 路由配置
import { resourceRoutes } from './resource-routing.module';

@NgModule({
  declarations: [],
  imports: [
    CommonModule,
    RouterModule.forChild(resourceRoutes)
  ],
  providers: []
})
export class ResourceModule { }
