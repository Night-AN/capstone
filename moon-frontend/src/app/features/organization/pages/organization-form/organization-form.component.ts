import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatTreeModule } from '@angular/material/tree';
import { MatIconModule } from '@angular/material/icon';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { ActivatedRoute, Router } from '@angular/router';
import { OrganizationService } from '@core/services/organization.service';
import { RoleService } from '@core/services/role.service';
import { UserService } from '@core/services/user.service';
import { OrganizationCreateRequest, OrganizationUpdateRequest } from '@models/organization.model';
import { NotificationService } from '@shared/service/notification/notification.service';
import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';
import { TreeSelectComponent } from '@shared/components/tree-select/tree-select.component';

interface OrganizationNode {
  organization_id: string;
  organization_name: string;
  children?: OrganizationNode[];
}

interface FlatOrganizationNode {
  expandable: boolean;
  organization_name: string;
  organization_id: string;
  level: number;
}

@Component({
  selector: 'app-organization-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatTreeModule,
    MatIconModule,
    MatCheckboxModule,
    TreeSelectComponent
  ],
  templateUrl: './organization-form.component.html',
  styleUrl: './organization-form.component.scss'
})
export class OrganizationFormComponent implements OnInit {
  organizationForm: FormGroup;
  organizationId: string = '';
  isEditMode: boolean = false;
  loading = signal<boolean>(false);
  isLoading = signal<boolean>(false);
  organizationTree = signal<any>(null);
  selectedParentId: string | null = null;
  
  // 角色管理相关
  roles: any[] = [];
  organizationRoles: string[] = [];
  
  // 人员管理相关
  users: any[] = [];
  organizationUsers: string[] = [];

  // 树形控件配置
  private _transformer = (node: OrganizationNode, level: number) => {
    return {
      expandable: !!node.children && node.children.length > 0,
      organization_name: node.organization_name,
      organization_id: node.organization_id,
      level: level,
    };
  };

  treeControl = new FlatTreeControl<FlatOrganizationNode>(
    node => node.level,
    node => node.expandable
  );

  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private organizationService = inject(OrganizationService);
  private roleService = inject(RoleService);
  private userService = inject(UserService);
  private notificationService = inject(NotificationService);

  constructor() {
    this.organizationForm = this.fb.group({
      organization_name: ['', Validators.required],
      organization_description: [''],
      organization_code: ['', Validators.required],
      organization_flag: ['', Validators.required]
    });
    this.organizationId = '';
    this.isEditMode = false;
    this.loading.set(false);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.organizationId = this.route.snapshot.paramMap.get('id') || '';
    this.isEditMode = this.organizationId !== '';

    if (this.isEditMode) {
      this.loadOrganizationDetail();
    }

    // 加载组织树
    this.loadOrganizationTree();
    
    // 加载角色列表
    this.loadRoles();
    
    // 加载用户列表
    this.loadUsers();
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

          // 如果是编辑模式，获取组织当前角色
          if (this.isEditMode) {
            this.loadOrganizationRoles();
          }
        }
      },
      error: (error) => {
        console.error('Error loading roles:', error);
        this.notificationService.error('加载角色列表失败');
      }
    });
  }

  loadOrganizationRoles(): void {
    this.organizationService.getOrganizationRoles(this.organizationId).subscribe({
      next: (response) => {
        if (response && response.data) {
          this.organizationRoles = response.data.map((role: any) => role.role_id);
          // 更新角色选择状态
          this.updateRoleSelection();
        }
      },
      error: (error) => {
        console.error('Error loading organization roles:', error);
        this.notificationService.error('加载组织角色失败');
      }
    });
  }

  updateRoleSelection(): void {
    this.roles.forEach(role => {
      role.checked = this.organizationRoles.includes(role.id);
    });
  }

  onRoleSelectionChange(selectedRoles: any[]): void {
    this.organizationRoles = selectedRoles.map(role => role.id);
  }

  loadUsers(): void {
    this.userService.getUsers().subscribe({
      next: (response) => {
        if (response && response.data && response.data.users) {
          this.users = response.data.users.map((user: any) => ({
            id: user.user_id,
            label: user.nickname + ' (' + user.email + ')',
            checked: false
          }));

          // 如果是编辑模式，获取组织当前用户
          if (this.isEditMode) {
            this.loadOrganizationUsers();
          }
        }
      },
      error: (error) => {
        console.error('Error loading users:', error);
        this.notificationService.error('加载用户列表失败');
      }
    });
  }

  loadOrganizationUsers(): void {
    this.organizationService.getOrganizationUsers(this.organizationId).subscribe({
      next: (users) => {
        if (users) {
          this.organizationUsers = users.map((user: any) => user.user_id);
          // 更新用户选择状态
          this.updateUserSelection();
        }
      },
      error: (error) => {
        console.error('Error loading organization users:', error);
        this.notificationService.error('加载组织用户失败');
      }
    });
  }

  updateUserSelection(): void {
    this.users.forEach(user => {
      user.checked = this.organizationUsers.includes(user.id);
    });
  }

  onUserSelectionChange(selectedUsers: any[]): void {
    this.organizationUsers = selectedUsers.map(user => user.id);
  }

  loadOrganizationTree(): void {
    this.organizationService.getOrganizationTree().subscribe({
      next: (tree: any) => {
        console.log('Organization Tree Response:', tree);
        if (tree) {
          // 转换组织树数据为树形控件所需的格式
          const organizationNode: OrganizationNode = {
            organization_id: tree.organization_id,
            organization_name: tree.organization_name,
            children: this.convertToOrganizationNodes(tree.children || [])
          };
          this.dataSource.data = [organizationNode];
          this.organizationTree.set(tree);
        }
      },
      error: (error: any) => {
        console.error('Error loading organization tree:', error);
        this.notificationService.error('加载组织树失败');
      }
    });
  }

  convertToOrganizationNodes(nodes: any[]): OrganizationNode[] {
    return nodes.map(node => ({
      organization_id: node.organization_id,
      organization_name: node.organization_name,
      children: this.convertToOrganizationNodes(node.children || [])
    }));
  }

  hasChild = (_: number, node: FlatOrganizationNode) => node.expandable;

  selectParent(node: FlatOrganizationNode): void {
    this.selectedParentId = node.organization_id;
    this.notificationService.success(`已选择父组织: ${node.organization_name}`);
  }

  loadOrganizationDetail(): void {
    if (this.isLoading()) {
      return;
    }
    
    this.isLoading.set(true);
    this.loading.set(true);
    
    this.organizationService.getOrganizationById(this.organizationId).subscribe({
      next: (organization: any) => {
        this.organizationForm.patchValue({
          organization_name: organization.organization_name,
          organization_description: organization.organization_description,
          organization_code: organization.organization_code,
          organization_flag: organization.organization_flag
        });
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error: any) => {
        this.notificationService.error('加载组织详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  onSubmit(): void {
    if (this.organizationForm.invalid || this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    if (this.isEditMode) {
      // 编辑模式
      const updateRequest: OrganizationUpdateRequest = {
        organization_id: this.organizationId,
        organization_name: this.organizationForm.value.organization_name,
        organization_description: this.organizationForm.value.organization_description,
        organization_code: this.organizationForm.value.organization_code,
        organization_flag: this.organizationForm.value.organization_flag,
        parent_id: this.selectedParentId
      };

      this.organizationService.updateOrganization(updateRequest).subscribe({
        next: (response: any) => {
          // 更新组织成功后，更新角色和用户
          this.updateOrganizationRoles();
          this.updateOrganizationUsers();
        },
        error: (error: any) => {
          this.notificationService.error('更新组织失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    } else {
      // 创建模式
      const createRequest: OrganizationCreateRequest = {
        organization_name: this.organizationForm.value.organization_name,
        organization_description: this.organizationForm.value.organization_description,
        organization_code: this.organizationForm.value.organization_code,
        organization_flag: this.organizationForm.value.organization_flag,
        parent_id: this.selectedParentId
      };

      this.organizationService.createOrganization(createRequest).subscribe({
        next: (response: any) => {
          if (response.data && response.data.organization_id) {
            // 创建组织成功后，保存组织ID并更新角色和用户
            this.organizationId = response.data.organization_id;
            this.updateOrganizationRoles();
            this.updateOrganizationUsers();
          } else {
            this.notificationService.success('创建组织成功');
            this.router.navigate(['/organizations']);
            this.loading.set(false);
            this.isLoading.set(false);
          }
        },
        error: (error: any) => {
          this.notificationService.error('创建组织失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    }
  }

  updateOrganizationRoles(): void {
    // 先获取组织当前的角色
    this.organizationService.getOrganizationRoles(this.organizationId).subscribe({
      next: (response) => {
        let currentRoles: string[] = [];
        if (response && response.data) {
          currentRoles = response.data.map((role: any) => role.role_id);
        }

        // 计算需要添加和删除的角色
        const rolesToAdd = this.organizationRoles.filter(roleId => !currentRoles.includes(roleId));
        const rolesToRemove = currentRoles.filter(roleId => !this.organizationRoles.includes(roleId));

        // 执行角色更新操作
        this.executeRoleUpdates(rolesToAdd, rolesToRemove);
      },
      error: (error) => {
        console.error('Error getting organization roles:', error);
        this.notificationService.error('获取组织角色失败');
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
      this.organizationService.assignRoleToOrganization(this.organizationId, roleId).subscribe({
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
      this.organizationService.removeRoleFromOrganization(this.organizationId, roleId).subscribe({
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
      this.checkOperationComplete(0, 0, false);
    }
  }

  updateOrganizationUsers(): void {
    // 这里可以添加更新组织用户的逻辑
    // 由于当前的 API 可能不支持直接的组织-用户关联管理
    // 我们可以在后续的迭代中实现
  }

  checkOperationComplete(completed: number, total: number, hasError: boolean): void {
    if (completed === total) {
      if (hasError) {
        this.notificationService.warning('组织信息保存成功，但部分角色更新失败');
      } else {
        this.notificationService.success(this.isEditMode ? '更新组织成功' : '创建组织成功');
      }
      this.router.navigate(['/organizations']);
      this.loading.set(false);
      this.isLoading.set(false);
    }
  }

  onCancel(): void {
    this.router.navigate(['/organizations']);
  }
}
