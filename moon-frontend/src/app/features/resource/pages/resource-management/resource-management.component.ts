import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { Router } from '@angular/router';
import { ResourceService } from '@core/services/resource.service';
import { ResourceListItem } from '@models/resource.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-resource-management',
  imports: [
    CommonModule,
    FormsModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatButtonModule,
    MatIconModule,
    MatFormFieldModule,
    MatInputModule
  ],
  templateUrl: './resource-management.component.html',
  styleUrl: './resource-management.component.scss'
})
export class ResourceManagementComponent implements OnInit {
  displayedColumns: string[] = ['name', 'description', 'created_at', 'actions'];
  dataSource = new MatTableDataSource<ResourceListItem>();
  searchKeyword: string = '';

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  private resourceService = inject(ResourceService);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() { }

  ngOnInit(): void {
    this.loadResources();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  loadResources(): void {
    console.log('Loading resources...');
    this.resourceService.getResources().subscribe({
      next: (resources: ResourceListItem[]) => {
        console.log('Received resources:', resources);
        this.dataSource.data = resources;
        console.log('DataSource data set:', this.dataSource.data);
      },
      error: (error: any) => {
        console.error('Error loading resources:', error);
        this.notificationService.error('加载资源失败');
      }
    });
  }

  applyFilter(): void {
    this.dataSource.filter = this.searchKeyword.trim().toLowerCase();
  }

  createResource(): void {
    this.router.navigate(['/resources/create']);
  }

  editResource(resourceId: string): void {
    this.router.navigate(['/resources/edit', resourceId]);
  }

  viewResource(resourceId: string): void {
    this.router.navigate(['/resources/detail', resourceId]);
  }

  deleteResource(resourceId: string): void {
    if (confirm('确定要删除这个资源吗？')) {
      this.resourceService.deleteResource(resourceId).subscribe({
        next: (success: boolean) => {
          if (success) {
            this.notificationService.success('删除资源成功');
            this.loadResources();
          } else {
            this.notificationService.error('删除资源失败');
          }
        },
        error: (error: any) => {
          this.notificationService.error('删除资源失败');
        }
      });
    }
  }
}
