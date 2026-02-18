import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { RoleService } from '@core/services/role.service';
import { Role } from '@models/role.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-role-detail',
  imports: [
    CommonModule,
    MatButtonModule,
    MatCardModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './role-detail.component.html',
  styleUrl: './role-detail.component.scss'
})
export class RoleDetailComponent implements OnInit {
  roleId: string = '';
  role = signal<Role | null>(null);
  loading = signal<boolean>(true);
  isLoading = signal<boolean>(false);

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private roleService = inject(RoleService);
  private notificationService = inject(NotificationService);

  constructor() {
    // 初始化状态
    this.roleId = '';
    this.role.set(null);
    this.loading.set(true);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.roleId = this.route.snapshot.paramMap.get('id') || '';
    this.loadRoleDetail();
  }

  loadRoleDetail(): void {
    // 防止重复调用
    if (this.isLoading()) {
      return;
    }
    
    // 设置加载标志位
    this.isLoading.set(true);
    // 设置加载状态为 true
    this.loading.set(true);
    
    // 检查roleId是否为空
    if (!this.roleId) {
      this.notificationService.error('角色ID不能为空');
      this.loading.set(false);
      this.isLoading.set(false);
      return;
    }
    
    // 调用 roleService.getRoleById 获取角色信息
    this.roleService.getRoleById(this.roleId).subscribe({
      next: (response: any) => {
        // 检查响应是否包含code和message字段，并且值是否正确
        if (response.code === '200' || response.message === 'success') {
          if (response.data) {
            this.role.set({
              role_id: response.data.role_id,
              name: response.data.name,
              description: response.data.description,
              created_at: response.data.created_at || new Date().toISOString(),
              updated_at: response.data.updated_at || new Date().toISOString()
            });
          }
        } else {
          this.notificationService.error('加载角色详情失败');
        }
        
        // 无论响应如何，只要请求成功，就停止加载状态
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

  editRole(): void {
    this.router.navigate(['/roles/edit', this.roleId]);
  }

  deleteRole(): void {
    if (confirm('确定要删除这个角色吗？')) {
      this.roleService.deleteRole({ role_id: this.roleId }).subscribe({
        next: () => {
          this.notificationService.success('删除角色成功');
          this.router.navigate(['/roles']);
        },
        error: (error: any) => {
          this.notificationService.error('删除角色失败');
        }
      });
    }
  }

  backToList(): void {
    this.router.navigate(['/roles']);
  }
}
