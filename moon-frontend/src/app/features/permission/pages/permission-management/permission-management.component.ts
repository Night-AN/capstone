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
import { PermissionService } from '@core/services/permission.service';
import { PermissionListItem } from '@models/permission.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-permission-management',
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
  templateUrl: './permission-management.component.html',
  styleUrl: './permission-management.component.scss'
})
export class PermissionManagementComponent implements OnInit {
  displayedColumns: string[] = ['name', 'description', 'created_at', 'actions'];
  dataSource = new MatTableDataSource<PermissionListItem>();
  searchKeyword: string = '';

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  private permissionService = inject(PermissionService);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() { }

  ngOnInit(): void {
    this.loadPermissions();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  loadPermissions(): void {
    console.log('Loading permissions...');
    this.permissionService.getPermissions().subscribe({
      next: (permissions: PermissionListItem[]) => {
        console.log('Received permissions:', permissions);
        this.dataSource.data = permissions;
        console.log('DataSource data set:', this.dataSource.data);
      },
      error: (error: any) => {
        console.error('Error loading permissions:', error);
        this.notificationService.error('加载权限失败');
      }
    });
  }

  applyFilter(): void {
    this.dataSource.filter = this.searchKeyword.trim().toLowerCase();
  }

  createPermission(): void {
    this.router.navigate(['/permissions/create']);
  }

  editPermission(permissionId: string): void {
    this.router.navigate(['/permissions/edit', permissionId]);
  }

  viewPermission(permissionId: string): void {
    this.router.navigate(['/permissions/detail', permissionId]);
  }

  deletePermission(permissionId: string): void {
    if (confirm('确定要删除这个权限吗？')) {
      this.permissionService.deletePermission(permissionId).subscribe({
        next: (success: boolean) => {
          if (success) {
            this.notificationService.success('删除权限成功');
            this.loadPermissions();
          } else {
            this.notificationService.error('删除权限失败');
          }
        },
        error: (error: any) => {
          this.notificationService.error('删除权限失败');
        }
      });
    }
  }
}
