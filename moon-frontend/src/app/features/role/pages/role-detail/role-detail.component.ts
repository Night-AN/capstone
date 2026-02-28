import { Component, OnInit, inject, signal } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { RoleService } from '@core/services/role.service';
import { UserService } from '@core/services/user.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { Role } from '@models/role.model';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatTableModule } from '@angular/material/table';

@Component({
  selector: 'app-role-detail',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatButtonModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatTableModule
  ],
  templateUrl: './role-detail.component.html',
  styleUrl: './role-detail.component.scss'
})
export class RoleDetailComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private roleService = inject(RoleService);
  private userService = inject(UserService);
  private notificationService = inject(NotificationService);

  roleId: string | null = null;
  role = signal<Role | null>(null);
  loading = signal<boolean>(false);
  error = signal<string | null>(null);
  
  // 用户管理相关
  allUsers: any[] = [];
  roleUsers: string[] = [];
  usersLoading = signal<boolean>(false);

  ngOnInit(): void {
    this.roleId = this.route.snapshot.paramMap.get('id');
    if (this.roleId) {
      this.loadRoleDetail();
      this.loadAllUsers();
      this.loadRoleUsers();
    }
  }

  loadRoleDetail(): void {
    if (!this.roleId) return;

    this.loading.set(true);
    this.error.set(null);

    this.roleService.getRoleById(this.roleId).subscribe({
      next: (role) => {
        this.role.set({
          role_id: role.role_id,
          role_name: role.role_name,
          description: role.description,
          sensitive_flag: role.sensitive_flag,
          created_at: role.created_at,
          updated_at: role.updated_at
        });
        this.loading.set(false);
      },
      error: (err) => {
        this.error.set('Failed to load role details');
        this.loading.set(false);
        this.notificationService.error('Failed to load role details');
      }
    });
  }

  loadAllUsers(): void {
    this.usersLoading.set(true);
    this.userService.getUsers().subscribe({
      next: (response) => {
        if (response.data && response.data.users) {
          this.allUsers = response.data.users.map((user: any) => ({
            id: user.user_id,
            label: user.nickname + ' (' + user.email + ')',
            checked: false
          }));
          // 更新用户选择状态
          this.updateUserSelection();
        }
        this.usersLoading.set(false);
      },
      error: (error) => {
        this.notificationService.error('加载用户列表失败');
        this.usersLoading.set(false);
      }
    });
  }

  loadRoleUsers(): void {
    if (!this.roleId) return;

    this.usersLoading.set(true);
    this.roleService.getRoleUsers(this.roleId).subscribe({
      next: (users: any[]) => {
        if (users) {
          this.roleUsers = users.map((user: any) => user.user_id);
          // 更新用户选择状态
          this.updateUserSelection();
        }
        this.usersLoading.set(false);
      },
      error: (error: any) => {
        console.error('Error loading role users:', error);
        this.notificationService.error('加载角色用户失败');
        this.usersLoading.set(false);
      }
    });
  }

  updateUserSelection(): void {
    this.allUsers.forEach(user => {
      user.checked = this.roleUsers.includes(user.id);
    });
  }

  onUserSelectionChange(selectedUsers: any[]): void {
    this.roleUsers = selectedUsers.map(user => user.id);
  }

  toggleUserSelection(userId: string): void {
    const index = this.roleUsers.indexOf(userId);
    if (index > -1) {
      // User is already selected, remove them
      this.roleUsers.splice(index, 1);
    } else {
      // User is not selected, add them
      this.roleUsers.push(userId);
    }
    // Update the checked status in allUsers
    this.updateUserSelection();
  }

  updateRoleUsers(): void {
    if (!this.roleId) return;

    // 先获取角色当前的用户
    this.roleService.getRoleUsers(this.roleId).subscribe({
      next: (users: any[]) => {
        let currentUsers: string[] = [];
        if (users) {
          currentUsers = users.map((user: any) => user.user_id);
        }

        // 计算需要添加和删除的用户
        const usersToAdd = this.roleUsers.filter(userId => !currentUsers.includes(userId));
        const usersToRemove = currentUsers.filter(userId => !this.roleUsers.includes(userId));

        // 执行用户更新操作
        this.executeUserUpdates(usersToAdd, usersToRemove);
      },
      error: (error: any) => {
        console.error('Error getting role users:', error);
        this.notificationService.error('获取角色用户失败');
      }
    });
  }

  executeUserUpdates(usersToAdd: string[], usersToRemove: string[]): void {
    let operationsCount = usersToAdd.length + usersToRemove.length;
    let completedOperations = 0;
    let hasError = false;

    // 添加用户
    usersToAdd.forEach(userId => {
      this.roleService.assignUserToRole(this.roleId!, userId).subscribe({
        next: () => {
          completedOperations++;
          this.checkUserOperationComplete(completedOperations, operationsCount, hasError);
        },
        error: () => {
          hasError = true;
          completedOperations++;
          this.checkUserOperationComplete(completedOperations, operationsCount, hasError);
        }
      });
    });

    // 移除用户
    usersToRemove.forEach(userId => {
      this.roleService.removeUserFromRole(this.roleId!, userId).subscribe({
        next: () => {
          completedOperations++;
          this.checkUserOperationComplete(completedOperations, operationsCount, hasError);
        },
        error: () => {
          hasError = true;
          completedOperations++;
          this.checkUserOperationComplete(completedOperations, operationsCount, hasError);
        }
      });
    });

    // 如果没有操作需要执行，直接完成
    if (operationsCount === 0) {
      this.checkUserOperationComplete(0, 0, false);
    }
  }

  checkUserOperationComplete(completed: number, total: number, hasError: boolean): void {
    if (completed === total) {
      if (hasError) {
        this.notificationService.warning('部分用户更新失败');
      } else {
        this.notificationService.success('用户更新成功');
      }
    }
  }

  editRole(): void {
    if (this.roleId) {
      this.router.navigate(['/role/edit', this.roleId]);
    }
  }

  deleteRole(): void {
    if (!this.roleId) return;

    if (confirm('Are you sure you want to delete this role?')) {
      this.roleService.deleteRole({ role_id: this.roleId }).subscribe({
        next: () => {
          this.notificationService.success('Role deleted successfully');
          this.router.navigate(['/role']);
        },
        error: (err) => {
          this.notificationService.error('Failed to delete role');
        }
      });
    }
  }

  backToList(): void {
    // 根据需求，backtolist 是无效的
    // 不执行任何操作
  }
}