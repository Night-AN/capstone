import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatTableModule } from '@angular/material/table';
import { MatIconModule } from '@angular/material/icon';
import { ActivatedRoute, Router } from '@angular/router';
import { UserService } from '@core/services/user.service';
import { User } from '@models/user.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-user-detail',
  imports: [
    CommonModule,
    MatButtonModule,
    MatCardModule,
    MatProgressSpinnerModule,
    MatTableModule,
    MatIconModule
  ],
  templateUrl: './user-detail.component.html',
  styleUrl: './user-detail.component.scss'
})
export class UserDetailComponent implements OnInit {
  userId: string = '';
  user = signal<User | null>(null);
  userRoles = signal<any[]>([]);
  loading = signal<boolean>(true);
  isLoading = signal<boolean>(false);

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private userService = inject(UserService);
  private notificationService = inject(NotificationService);

  constructor() {
    // 初始化状态
    this.userId = '';
    this.user.set(null);
    this.loading.set(true);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.userId = this.route.snapshot.paramMap.get('id') || '';
    this.loadUserDetail();
  }

  loadUserDetail(): void {
    // 防止重复调用
    if (this.isLoading()) {
      return;
    }
    
    // 设置加载标志位
    this.isLoading.set(true);
    // 设置加载状态为 true
    this.loading.set(true);
    
    // 检查userId是否为空
    if (!this.userId) {
      this.notificationService.error('用户ID不能为空');
      this.loading.set(false);
      this.isLoading.set(false);
      return;
    }
    
    // 调用 userService.getUserById 获取用户信息
    this.userService.getUserById(this.userId).subscribe({
      next: (response) => {
        // 检查响应是否包含code和message字段，并且值是否正确
        if (response.code === '200' || response.message === 'success') {
          if (response.data) {
            this.user.set({
              user_id: response.data.user_id,
              nickname: response.data.nickname,
              full_name: response.data.full_name,
              email: response.data.email,
              created_at: new Date().toISOString(), // 后端没有返回 created_at，暂时使用当前时间
              updated_at: new Date().toISOString()  // 后端没有返回 updated_at，暂时使用当前时间
            });
          }
        } else {
          this.notificationService.error('加载用户详情失败');
        }
        
        // 加载用户角色
        this.loadUserRoles();
        
        // 无论响应如何，只要请求成功，就停止加载状态
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

  editUser(): void {
    this.router.navigate(['/users/edit', this.userId]);
  }

  deleteUser(): void {
    if (confirm('确定要删除这个用户吗？')) {
      this.userService.deleteUser({ user_id: this.userId }).subscribe({
        next: () => {
          this.notificationService.success('删除用户成功');
          this.router.navigate(['/users']);
        },
        error: (error) => {
          console.error('Error deleting user:', error);
          this.notificationService.error('删除用户失败');
        }
      });
    }
  }

  loadUserRoles(): void {
    this.userService.getUserRoles(this.userId).subscribe({
      next: (response) => {
        if (response.code === '200' || response.message === 'success') {
          this.userRoles.set(response.data || []);
        } else {
          this.notificationService.error('加载用户角色失败');
        }
      },
      error: (error) => {
        console.error('Error loading user roles:', error);
        this.notificationService.error('加载用户角色失败');
      }
    });
  }

  backToList(): void {
    this.router.navigate(['/users']);
  }
}
