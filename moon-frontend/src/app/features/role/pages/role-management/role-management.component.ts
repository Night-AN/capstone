import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { Router } from '@angular/router';
import { RoleService } from '@core/services/role.service';
import { RoleListItem } from '@models/role.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-role-management',
  imports: [
    CommonModule,
    FormsModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatButtonModule,
    MatIconModule,
    MatFormFieldModule,
    MatInputModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './role-management.component.html',
  styleUrl: './role-management.component.scss'
})
export class RoleManagementComponent implements OnInit {
  displayedColumns: string[] = ['role_name', 'role_code', 'role_flag', 'sensitive_flag', 'created_at', 'actions'];
  dataSource = new MatTableDataSource<RoleListItem>();
  searchKeyword: string = '';
  loading = signal<boolean>(false);

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  private roleService = inject(RoleService);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() { }

  ngOnInit(): void {
    this.loadRoles();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  loadRoles(): void {
    this.loading.set(true);
    this.roleService.getRoles().subscribe({
      next: (roles: any[]) => {
        this.dataSource.data = roles;
        this.loading.set(false);
      },
      error: (error: any) => {
        this.notificationService.error('加载角色失败');
        this.loading.set(false);
      }
    });
  }

  applyFilter(): void {
    this.dataSource.filter = this.searchKeyword.trim().toLowerCase();
  }

  createRole(): void {
    this.router.navigate(['/roles/create']);
  }

  editRole(roleId: string): void {
    this.router.navigate(['/roles/edit', roleId]);
  }

  viewRole(roleId: string): void {
    this.router.navigate(['/roles/detail', roleId]);
  }

  deleteRole(roleId: string): void {
    if (confirm('确定要删除这个角色吗？')) {
      this.roleService.deleteRole({ role_id: roleId }).subscribe({
        next: (response: any) => {
          this.notificationService.success('删除角色成功');
          this.loadRoles();
        },
        error: (error: any) => {
          this.notificationService.error('删除角色失败');
        }
      });
    }
  }
}
