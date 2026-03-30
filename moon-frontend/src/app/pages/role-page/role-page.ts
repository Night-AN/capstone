import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators, FormsModule } from '@angular/forms';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzMessageService } from 'ng-zorro-antd/message';
import { CommonModule } from '@angular/common';
import { RoleService } from 'src/app/services/role.service';

interface Role {
  id: string;
  role_name: string;
  role_code: string;
  role_description: string;
  role_flag: string;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建角色',
  Edit = '编辑角色'
}

@Component({
  selector: 'app-role-page',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    NzButtonModule,
    NzCardModule,
    NzFormModule,
    NzInputModule,
    NzTableModule,
    NzModalModule
  ],
  templateUrl: './role-page.html',
  styleUrl: './role-page.scss',
  providers: [NzMessageService]
})
export class RolePage implements OnInit {
  roleForm: FormGroup;
  roles: Role[] = [];

  isEditing = false;
  currentRoleId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;
  loading = false;

  // 搜索相关
  searchName = '';
  searchCode = '';
  searchFlag = '';

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private roleService: RoleService
  ) {
    this.roleForm = this.fb.group({
      role_name: ['', [Validators.required]],
      role_code: ['', [Validators.required]],
      role_description: ['', [Validators.required]],
      role_flag: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadRoles();
  }

  loadRoles(): void {
    this.loading = true;
    this.roleService.getList().subscribe(roles => {
      this.roles = roles[0].edges.map((edge: any) => ({
        id: edge.node.id,
        role_name: edge.node.role_name,
        role_code: edge.node.role_code,
        role_description: edge.node.role_description,
        role_flag: edge.node.role_flag,
        created_at: edge.node.created_at,
        updated_at: edge.node.updated_at
      }));
      this.loading = false;
    });
  }

  // 搜索角色
  searchRoles(): void {
    this.loading = true;
    // 这里可以根据搜索条件调用后端API
    // 暂时使用前端过滤
    setTimeout(() => {
      this.loading = false;
    }, 500);
  }

  // 重置搜索
  resetSearch(): void {
    this.searchName = '';
    this.searchCode = '';
    this.searchFlag = '';
  }

  createRole(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editRole(role: Role): void {
    this.isEditing = true;
    this.currentRoleId = role.id;
    this.modalTitle = ModalTitle.Edit;
    this.roleForm.patchValue({
      role_name: role.role_name,
      role_code: role.role_code,
      role_description: role.role_description,
      role_flag: role.role_flag
    });
    this.visible = true;
  }

  deleteRole(id: string): void {
    this.roleService.delete(id).subscribe(() => {
      this.loadRoles();
      this.message.success('角色删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.roleForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentRoleId) {
          this.updateRole();
        } else {
          this.createRoleSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.roleForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createRoleSubmit(): void {
    this.roleService.create(this.roleForm.value).subscribe(() => {
      this.loadRoles();
      this.message.success('角色创建成功');
      this.resetForm();
    });
  }

  updateRole(): void {
    if (this.currentRoleId) {
      this.roleService.update(this.currentRoleId, this.roleForm.value).subscribe(() => {
        this.loadRoles();
        this.message.success('角色更新成功');
        this.resetForm();
      });
    }
  }

  resetForm(): void {
    this.roleForm.reset();
    this.isEditing = false;
    this.currentRoleId = null;
  }
}