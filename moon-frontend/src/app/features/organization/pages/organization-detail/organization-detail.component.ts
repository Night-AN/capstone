import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { OrganizationService } from '@core/services/organization.service';
import { Organization } from '@models/organization.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-organization-detail',
  imports: [
    CommonModule,
    MatButtonModule,
    MatCardModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './organization-detail.component.html',
  styleUrl: './organization-detail.component.scss'
})
export class OrganizationDetailComponent implements OnInit {
  organizationId: string = '';
  organization = signal<Organization | null>(null);
  users = signal<any[]>([]);
  loading = signal<boolean>(true);
  isLoading = signal<boolean>(false);
  loadingUsers = signal<boolean>(false);

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private organizationService = inject(OrganizationService);
  private notificationService = inject(NotificationService);

  constructor() {
    // 初始化状态
    this.organizationId = '';
    this.organization.set(null);
    this.users.set([]);
    this.loading.set(true);
    this.isLoading.set(false);
    this.loadingUsers.set(false);
  }

  ngOnInit(): void {
    this.organizationId = this.route.snapshot.paramMap.get('id') || '';
    this.loadOrganizationDetail();
    this.loadOrganizationUsers();
  }

  loadOrganizationDetail(): void {
    // 防止重复调用
    if (this.isLoading()) {
      return;
    }
    
    // 设置加载标志位
    this.isLoading.set(true);
    // 设置加载状态为 true
    this.loading.set(true);
    
    // 检查organizationId是否为空
    if (!this.organizationId) {
      this.notificationService.error('组织ID不能为空');
      this.loading.set(false);
      this.isLoading.set(false);
      return;
    }
    
    // 调用 organizationService.getOrganizationById 获取组织信息
    this.organizationService.getOrganizationById(this.organizationId).subscribe({
      next: (organization) => {
        // 直接使用返回的组织对象
        this.organization.set({
          organization_id: organization.organization_id,
          organization_name: organization.organization_name,
          organization_code: organization.organization_code,
          organization_description: organization.organization_description,
          organization_flag: organization.organization_flag,
          created_at: organization.created_at || new Date().toISOString(),
          updated_at: organization.updated_at || new Date().toISOString()
        });
        
        // 无论响应如何，只要请求成功，就停止加载状态
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error) => {
        this.notificationService.error('加载组织详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  loadOrganizationUsers(): void {
    // 防止重复调用
    if (this.loadingUsers()) {
      return;
    }
    
    // 设置加载标志位
    this.loadingUsers.set(true);
    
    // 检查organizationId是否为空
    if (!this.organizationId) {
      this.notificationService.error('组织ID不能为空');
      this.loadingUsers.set(false);
      return;
    }
    
    // 调用 organizationService.getOrganizationUsers 获取组织用户列表
    this.organizationService.getOrganizationUsers(this.organizationId).subscribe({
      next: (users) => {
        // 直接使用返回的用户列表
        this.users.set(users);
        
        // 无论响应如何，只要请求成功，就停止加载状态
        this.loadingUsers.set(false);
      },
      error: (error) => {
        this.notificationService.error('加载组织用户列表失败');
        this.loadingUsers.set(false);
      }
    });
  }

  editOrganization(): void {
    this.router.navigate(['/organizations/edit', this.organizationId]);
  }

  deleteOrganization(): void {
    if (confirm('确定要删除这个组织吗？')) {
      this.organizationService.deleteOrganization({ organization_id: this.organizationId }).subscribe({
        next: () => {
          this.notificationService.success('删除组织成功');
          this.router.navigate(['/organizations']);
        },
        error: (error) => {
          this.notificationService.error('删除组织失败');
        }
      });
    }
  }

  backToList(): void {
    this.router.navigate(['/organizations']);
  }
}
