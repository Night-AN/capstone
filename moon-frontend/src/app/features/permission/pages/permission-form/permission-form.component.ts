import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { PermissionService } from '@core/services/permission.service';
import { PermissionCreateRequest, PermissionUpdateRequest } from '@models/permission.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-permission-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './permission-form.component.html',
  styleUrl: './permission-form.component.scss'
})
export class PermissionFormComponent implements OnInit {
  permissionForm: FormGroup;
  permissionId: string = '';
  isEditMode: boolean = false;
  loading = signal<boolean>(false);
  isLoading = signal<boolean>(false);
  isSensitive = signal<boolean>(false);

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private permissionService = inject(PermissionService);
  private notificationService = inject(NotificationService);

  constructor() {
    this.permissionForm = this.fb.group({
      name: ['', Validators.required],
      description: ['']
    });
    this.permissionId = '';
    this.isEditMode = false;
    this.loading.set(false);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.permissionId = this.route.snapshot.paramMap.get('id') || '';
    this.isEditMode = this.permissionId !== '';

    if (this.isEditMode) {
      this.loadPermissionDetail();
    }
  }

  loadPermissionDetail(): void {
    if (this.isLoading()) {
      return;
    }
    
    this.isLoading.set(true);
    this.loading.set(true);
    
    this.permissionService.getPermissionById(this.permissionId).subscribe({
      next: (permission: any) => {
        console.log('Loaded Permission Detail:', permission);
        this.permissionForm.patchValue({
          name: permission.name,
          description: permission.description
        });
        // 检查权限是否是敏感的
        this.isSensitive.set(permission.sensitive_flag || false);
        // 如果是敏感权限，禁用表单
        if (this.isSensitive()) {
          this.permissionForm.disable();
        }
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error: any) => {
        console.error('Error loading permission detail:', error);
        this.notificationService.error('加载权限详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  onSubmit(): void {
    if (this.permissionForm.invalid || this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    if (this.isEditMode) {
      // 编辑模式
      const updateRequest: PermissionUpdateRequest = {
        permission_id: this.permissionId,
        name: this.permissionForm.value.name,
        description: this.permissionForm.value.description
      };

      this.permissionService.updatePermission(updateRequest).subscribe({
        next: (response: any) => {
          this.notificationService.success('更新权限成功');
          this.router.navigate(['/permissions']);
          this.loading.set(false);
          this.isLoading.set(false);
        },
        error: (error: any) => {
          this.notificationService.error('更新权限失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    } else {
      // 创建模式
      const createRequest: PermissionCreateRequest = {
        name: this.permissionForm.value.name,
        description: this.permissionForm.value.description
      };

      this.permissionService.createPermission(createRequest).subscribe({
        next: (response: any) => {
          this.notificationService.success('创建权限成功');
          this.router.navigate(['/permissions']);
          this.loading.set(false);
          this.isLoading.set(false);
        },
        error: (error: any) => {
          this.notificationService.error('创建权限失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    }
  }

  onCancel(): void {
    this.router.navigate(['/permissions']);
  }
}
