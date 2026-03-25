import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzMessageService } from 'ng-zorro-antd/message';
import { CommonModule } from '@angular/common';
import { PermissionService } from 'src/app/services/permission.service';

interface Permission {
  id: string;
  permission_name: string;
  permission_code: string;
  permission_description: string;
  permission_flag: string;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建权限',
  Edit = '编辑权限'
}

@Component({
  selector: 'app-permission-page',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NzButtonModule,
    NzCardModule,
    NzFormModule,
    NzInputModule,
    NzTableModule,
    NzModalModule
  ],
  templateUrl: './permission-page.html',
  styleUrl: './permission-page.less',
  providers: [NzMessageService]
})
export class PermissionPage implements OnInit {
  permissionForm: FormGroup;
  permissions: Permission[] = [];

  isEditing = false;
  currentPermissionId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private permissionService: PermissionService
  ) {
    this.permissionForm = this.fb.group({
      permission_name: ['', [Validators.required]],
      permission_code: ['', [Validators.required]],
      permission_description: ['', [Validators.required]],
      permission_flag: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadPermissions();
  }

  loadPermissions(): void {
    this.permissionService.getList().subscribe(permissions => {
      this.permissions = permissions[0].edges.map((edge: any) => ({
        id: edge.node.id,
        permission_name: edge.node.permission_name,
        permission_code: edge.node.permission_code,
        permission_description: edge.node.permission_description,
        permission_flag: edge.node.permission_flag,
        created_at: edge.node.created_at,
        updated_at: edge.node.updated_at
      }));
    });
  }

  createPermission(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editPermission(permission: Permission): void {
    this.isEditing = true;
    this.currentPermissionId = permission.id;
    this.modalTitle = ModalTitle.Edit;
    this.permissionForm.patchValue({
      permission_name: permission.permission_name,
      permission_code: permission.permission_code,
      permission_description: permission.permission_description,
      permission_flag: permission.permission_flag
    });
    this.visible = true;
  }

  deletePermission(id: string): void {
    this.permissionService.delete(id).subscribe(() => {
      this.loadPermissions();
      this.message.success('权限删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.permissionForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentPermissionId) {
          this.updatePermission();
        } else {
          this.createPermissionSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.permissionForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createPermissionSubmit(): void {
    this.permissionService.create(this.permissionForm.value).subscribe(() => {
      this.loadPermissions();
      this.message.success('权限创建成功');
      this.resetForm();
    });
  }

  updatePermission(): void {
    if (this.currentPermissionId) {
      this.permissionService.update(this.currentPermissionId, this.permissionForm.value).subscribe(() => {
        this.loadPermissions();
        this.message.success('权限更新成功');
        this.resetForm();
      });
    }
  }

  resetForm(): void {
    this.permissionForm.reset();
    this.isEditing = false;
    this.currentPermissionId = null;
  }
}