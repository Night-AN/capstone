import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { PermissionService } from '@core/services/permission.service';
import { Permission } from '@models/permission.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-permission-detail',
  imports: [
    CommonModule,
    MatButtonModule,
    MatCardModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './permission-detail.component.html',
  styleUrl: './permission-detail.component.scss'
})
export class PermissionDetailComponent implements OnInit {
  permissionId: string = '';
  permission = signal<Permission | null>(null);
  loading = signal<boolean>(true);
  isLoading = signal<boolean>(false);

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private permissionService = inject(PermissionService);
  private notificationService = inject(NotificationService);

  constructor() {
    // 初始化状态
    this.permissionId = '';
    this.permission.set(null);
    this.loading.set(true);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.permissionId = this.route.snapshot.paramMap.get('id') || '';
    this.loadPermissionDetail();
  }

  loadPermissionDetail(): void {
    // 防止重复调用
    if (this.isLoading()) {
      return;
    }
    
    // 设置加载标志位
    this.isLoading.set(true);
    // 设置加载状态为 true
    this.loading.set(true);
    
    // 检查permissionId是否为空
    if (!this.permissionId) {
      this.notificationService.error('权限ID不能为空');
      this.loading.set(false);
      this.isLoading.set(false);
      return;
    }
    
    // 调用 permissionService.getPermissionById 获取权限信息
    this.permissionService.getPermissionById(this.permissionId).subscribe({
      next: (permission: any) => {
        console.log('Loaded Permission Detail:', permission);
        // 直接使用返回的permission对象
        this.permission.set({
          permission_id: permission.permission_id,
          name: permission.name,
          description: permission.description,
          sensitive_flag: permission.sensitive_flag || false,
          created_at: permission.created_at || new Date().toISOString(),
          updated_at: permission.updated_at || new Date().toISOString()
        });
        
        // 无论响应如何，只要请求成功，就停止加载状态
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

  editPermission(): void {
    this.router.navigate(['/permissions/edit', this.permissionId]);
  }

  deletePermission(): void {
    if (confirm('确定要删除这个权限吗？')) {
      this.permissionService.deletePermission(this.permissionId).subscribe({
        next: (success: boolean) => {
          if (success) {
            this.notificationService.success('删除权限成功');
            this.router.navigate(['/permissions']);
          } else {
            this.notificationService.error('删除权限失败');
          }
        },
        error: (error: any) => {
          this.notificationService.error('删除权限失败');
        }
      });
    }
  }

  backToList(): void {
    this.router.navigate(['/permissions']);
  }
}
