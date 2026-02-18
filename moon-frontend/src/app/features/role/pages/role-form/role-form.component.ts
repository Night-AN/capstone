import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { RoleService } from '@core/services/role.service';
import { RoleCreateRequest, RoleUpdateRequest } from '@models/role.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-role-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './role-form.component.html',
  styleUrl: './role-form.component.scss'
})
export class RoleFormComponent implements OnInit {
  roleForm: FormGroup;
  roleId: string = '';
  isEditMode: boolean = false;
  loading = signal<boolean>(false);
  isLoading = signal<boolean>(false);

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private roleService = inject(RoleService);
  private notificationService = inject(NotificationService);

  constructor() {
    this.roleForm = this.fb.group({
      name: ['', Validators.required],
      description: ['']
    });
    this.roleId = '';
    this.isEditMode = false;
    this.loading.set(false);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.roleId = this.route.snapshot.paramMap.get('id') || '';
    this.isEditMode = this.roleId !== '';

    if (this.isEditMode) {
      this.loadRoleDetail();
    }
  }

  loadRoleDetail(): void {
    if (this.isLoading()) {
      return;
    }
    
    this.isLoading.set(true);
    this.loading.set(true);
    
    this.roleService.getRoleById(this.roleId).subscribe({
      next: (response: any) => {
        if (response.data) {
          this.roleForm.patchValue({
            name: response.data.name,
            description: response.data.description
          });
        }
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error: any) => {
        this.notificationService.error('加载角色详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  onSubmit(): void {
    if (this.roleForm.invalid || this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    if (this.isEditMode) {
      // 编辑模式
      const updateRequest: RoleUpdateRequest = {
        role_id: this.roleId,
        name: this.roleForm.value.name,
        description: this.roleForm.value.description
      };

      this.roleService.updateRole(updateRequest).subscribe({
        next: (response: any) => {
          this.notificationService.success('更新角色成功');
          this.router.navigate(['/roles']);
          this.loading.set(false);
          this.isLoading.set(false);
        },
        error: (error: any) => {
          this.notificationService.error('更新角色失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    } else {
      // 创建模式
      const createRequest: RoleCreateRequest = {
        name: this.roleForm.value.name,
        description: this.roleForm.value.description
      };

      this.roleService.createRole(createRequest).subscribe({
        next: (response: any) => {
          this.notificationService.success('创建角色成功');
          this.router.navigate(['/roles']);
          this.loading.set(false);
          this.isLoading.set(false);
        },
        error: (error: any) => {
          this.notificationService.error('创建角色失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    }
  }

  onCancel(): void {
    this.router.navigate(['/roles']);
  }
}
