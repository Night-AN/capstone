import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from '@core/services/user.service';
import { RoleService } from '@core/services/role.service';
import { UserCreateRequest, UserUpdateRequest } from '@models/user.model';
import { NotificationService } from '@shared/service/notification/notification.service';
import { TreeSelectComponent } from '@shared/components/tree-select/tree-select.component';

@Component({
  selector: 'app-user-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    TreeSelectComponent
  ],
  templateUrl: './user-form.component.html',
  styleUrl: './user-form.component.scss'
})
export class UserFormComponent implements OnInit {
  userForm: FormGroup;
  userId: string = '';
  isEditMode: boolean = false;
  loading = signal<boolean>(false);
  isLoading = signal<boolean>(false);
  roles: any[] = [];
  userRoles: string[] = [];

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private userService = inject(UserService);
  private roleService = inject(RoleService);
  private notificationService = inject(NotificationService);

  constructor() {
    this.userForm = this.fb.group({
      nickname: ['', Validators.required],
      full_name: ['', Validators.required],
      email: ['', [Validators.required, Validators.email]],
      password: ['', Validators.required]
    });
    this.userId = '';
    this.isEditMode = false;
    this.loading.set(false);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.userId = this.route.snapshot.paramMap.get('id') || '';
    this.isEditMode = this.userId !== '';

    if (this.isEditMode) {
      // 编辑模式下，移除密码字段的必填验证
      this.userForm.get('password')?.clearValidators();
      this.userForm.get('password')?.updateValueAndValidity();
      this.loadUserDetail();
    }

    // 获取角色列表
    this.loadRoles();
  }

  loadRoles(): void {
    this.roleService.getRoles().subscribe({
      next: (roles) => {
        if (roles) {
          this.roles = roles.map((role: any) => ({
            id: role.role_id,
            label: role.role_name,
            checked: false
          }));

          // 如果是编辑模式，加载用户当前角色
          if (this.isEditMode) {
            this.loadUserRoles();
          }
        }
      },
      error: (error) => {
        this.notificationService.error('加载角色列表失败');
      }
    });
  }

  loadUserRoles(): void {
    this.userService.getUserRoles(this.userId).subscribe({
      next: (response) => {
        if (response && response.data) {
          this.userRoles = response.data.map((role: any) => role.role_id);
          // 更新角色选择状态
          this.updateRoleSelection();
        }
      },
      error: (error) => {
        console.error('Error loading user roles:', error);
        this.notificationService.error('加载用户角色失败');
      }
    });
  }

  updateRoleSelection(): void {
    this.roles.forEach(role => {
      role.checked = this.userRoles.includes(role.id);
    });
  }

  onRoleSelectionChange(selectedRoles: any[]): void {
    this.userRoles = selectedRoles.map(role => role.id);
  }

  loadUserDetail(): void {
    if (this.isLoading()) {
      return;
    }
    
    this.isLoading.set(true);
    this.loading.set(true);
    
    this.userService.getUserById(this.userId).subscribe({
      next: (response) => {
        if (response.data) {
          this.userForm.patchValue({
            nickname: response.data.nickname,
            full_name: response.data.full_name,
            email: response.data.email
          });
        }
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error) => {
        this.notificationService.error('加载用户详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  onSubmit(): void {
    if (this.userForm.invalid || this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    if (this.isEditMode) {
      // 编辑模式
      const updateRequest: UserUpdateRequest = {
        user_id: this.userId,
        nickname: this.userForm.value.nickname,
        full_name: this.userForm.value.full_name,
        email: this.userForm.value.email
      };

      this.userService.updateUser(updateRequest).subscribe({
        next: (response) => {
          // 更新用户成功后，更新角色
          this.updateUserRoles();
        },
        error: (error) => {
          this.notificationService.error('更新用户失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    } else {
      // 创建模式
      const createRequest: UserCreateRequest = {
        nickname: this.userForm.value.nickname,
        full_name: this.userForm.value.full_name,
        email: this.userForm.value.email,
        password: this.userForm.value.password
      };

      this.userService.createUser(createRequest).subscribe({
        next: (response) => {
          if (response.data && response.data.user_id) {
            // 创建用户成功后，保存用户ID并更新角色
            this.userId = response.data.user_id;
            this.updateUserRoles();
          } else {
            this.notificationService.success('创建用户成功');
            this.router.navigate(['/users']);
            this.loading.set(false);
            this.isLoading.set(false);
          }
        },
        error: (error) => {
          this.notificationService.error('创建用户失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    }
  }

  updateUserRoles(): void {
    // 先获取用户当前的角色
    this.userService.getUserRoles(this.userId).subscribe({
      next: (response) => {
        let currentRoles: string[] = [];
        if (response && response.data) {
          currentRoles = response.data.map((role: any) => role.role_id);
        }

        // 计算需要添加和删除的角色
        const rolesToAdd = this.userRoles.filter(roleId => !currentRoles.includes(roleId));
        const rolesToRemove = currentRoles.filter(roleId => !this.userRoles.includes(roleId));

        // 执行角色更新操作
        this.executeRoleUpdates(rolesToAdd, rolesToRemove);
      },
      error: (error) => {
        console.error('Error getting user roles:', error);
        this.notificationService.error('获取用户角色失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  executeRoleUpdates(rolesToAdd: string[], rolesToRemove: string[]): void {
    let operationsCount = rolesToAdd.length + rolesToRemove.length;
    let completedOperations = 0;
    let hasError = false;

    // 添加角色
    rolesToAdd.forEach(roleId => {
      this.userService.assignRoleToUser(this.userId, roleId).subscribe({
        next: () => {
          completedOperations++;
          this.checkOperationComplete(completedOperations, operationsCount, hasError);
        },
        error: () => {
          hasError = true;
          completedOperations++;
          this.checkOperationComplete(completedOperations, operationsCount, hasError);
        }
      });
    });

    // 移除角色
    rolesToRemove.forEach(roleId => {
      this.userService.removeRoleFromUser(this.userId, roleId).subscribe({
        next: () => {
          completedOperations++;
          this.checkOperationComplete(completedOperations, operationsCount, hasError);
        },
        error: () => {
          hasError = true;
          completedOperations++;
          this.checkOperationComplete(completedOperations, operationsCount, hasError);
        }
      });
    });

    // 如果没有操作需要执行，直接完成
    if (operationsCount === 0) {
      this.notificationService.success(this.isEditMode ? '更新用户成功' : '创建用户成功');
      this.router.navigate(['/users']);
      this.loading.set(false);
      this.isLoading.set(false);
    }
  }

  checkOperationComplete(completed: number, total: number, hasError: boolean): void {
    if (completed === total) {
      if (hasError) {
        this.notificationService.warning('用户信息保存成功，但部分角色更新失败');
      } else {
        this.notificationService.success(this.isEditMode ? '更新用户成功' : '创建用户成功');
      }
      this.router.navigate(['/users']);
      this.loading.set(false);
      this.isLoading.set(false);
    }
  }

  onCancel(): void {
    this.router.navigate(['/users']);
  }
}
