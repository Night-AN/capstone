import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';

import { NzBreadCrumbModule } from 'ng-zorro-antd/breadcrumb';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzLayoutModule } from 'ng-zorro-antd/layout';
import { NzMenuModule } from 'ng-zorro-antd/menu';

@Component({
  selector: 'nz-demo-layout-custom-trigger',
  imports: [CommonModule, RouterModule, NzBreadCrumbModule, NzIconModule, NzMenuModule, NzLayoutModule],
  templateUrl: './backend-layout.html',
  styleUrl:  './backend-layout.scss'
})
export class BackendLayout {
  isCollapsed = false;
  breadcrumbItems: string[] = [];
  constructor() {
    this.breadcrumbItems = ['User', 'Bill'];
  }
  protected readonly date = new Date();
}