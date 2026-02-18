import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from '@core/services/user.service';
import { UserCreateRequest, UserUpdateRequest } from '@models/user.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-user-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule
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

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private userService = inject(UserService);
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
          this.notificationService.success('更新用户成功');
          this.router.navigate(['/users']);
          this.loading.set(false);
          this.isLoading.set(false);
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
          this.notificationService.success('创建用户成功');
          this.router.navigate(['/users']);
          this.loading.set(false);
          this.isLoading.set(false);
        },
        error: (error) => {
          this.notificationService.error('创建用户失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    }
  }

  onCancel(): void {
    this.router.navigate(['/users']);
  }
}
