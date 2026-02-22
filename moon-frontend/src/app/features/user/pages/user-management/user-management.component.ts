import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { Router } from '@angular/router';
import { UserService } from '@core/services/user.service';
import { UserListItem } from '@models/user.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-user-management',
  imports: [
    CommonModule,
    FormsModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatButtonModule,
    MatIconModule
  ],
  templateUrl: './user-management.component.html',
  styleUrl: './user-management.component.scss'
})
export class UserManagementComponent implements OnInit {
  displayedColumns: string[] = ['nickname', 'full_name', 'email', 'created_at', 'actions'];
  dataSource = new MatTableDataSource<UserListItem>();
  searchKeyword: string = '';

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  private userService = inject(UserService);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() { }

  ngOnInit(): void {
    this.loadUsers();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  loadUsers(): void {
    this.userService.getUsers().subscribe({
      next: (response) => {
        this.dataSource.data = response.data?.users || [];
      },
      error: (error) => {
        this.notificationService.error('加载用户失败');
      }
    });
  }

  applyFilter(): void {
    this.dataSource.filter = this.searchKeyword.trim().toLowerCase();
  }

  createUser(): void {
    this.router.navigate(['/users/create']);
  }

  editUser(userId: string): void {
    this.router.navigate(['/users/edit', userId]);
  }

  viewUser(userId: string): void {
    this.router.navigate(['/users/detail', userId]);
  }

  deleteUser(userId: string): void {
    if (confirm('确定要删除这个用户吗？')) {
      this.userService.deleteUser({ user_id: userId }).subscribe({
        next: (response) => {
          this.notificationService.success('删除用户成功');
          this.loadUsers();
        },
        error: (error) => {
          this.notificationService.error('删除用户失败');
        }
      });
    }
  }

  // 分配角色给用户
  assignRoleToUser(userId: string): void {
    // 这里需要打开角色选择对话框
    // 暂时使用alert模拟
    const roleId = prompt('请输入角色ID:');
    if (roleId) {
      this.userService.assignRoleToUser(userId, roleId).subscribe({
        next: (response) => {
          this.notificationService.success('分配角色成功');
        },
        error: (error) => {
          this.notificationService.error('分配角色失败');
        }
      });
    }
  }

  // 查看用户角色
  viewUserRoles(userId: string): void {
    this.userService.getUserRoles(userId).subscribe({
      next: (response) => {
        console.log('User roles:', response);
        // 这里需要显示用户角色列表
        alert('用户角色: ' + JSON.stringify(response.data));
      },
      error: (error) => {
        this.notificationService.error('获取用户角色失败');
      }
    });
  }
}
