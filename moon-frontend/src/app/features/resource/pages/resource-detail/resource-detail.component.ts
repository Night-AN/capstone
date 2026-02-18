import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { ResourceService } from '@core/services/resource.service';
import { Resource } from '@models/resource.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-resource-detail',
  imports: [
    CommonModule,
    MatButtonModule,
    MatCardModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './resource-detail.component.html',
  styleUrl: './resource-detail.component.scss'
})
export class ResourceDetailComponent implements OnInit {
  resourceId: string = '';
  resource = signal<Resource | null>(null);
  loading = signal<boolean>(true);
  isLoading = signal<boolean>(false);

  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private resourceService = inject(ResourceService);
  private notificationService = inject(NotificationService);

  constructor() {
    // 初始化状态
    this.resourceId = '';
    this.resource.set(null);
    this.loading.set(true);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.resourceId = this.route.snapshot.paramMap.get('id') || '';
    this.loadResourceDetail();
  }

  loadResourceDetail(): void {
    // 防止重复调用
    if (this.isLoading()) {
      return;
    }
    
    // 设置加载标志位
    this.isLoading.set(true);
    // 设置加载状态为 true
    this.loading.set(true);
    
    // 检查resourceId是否为空
    if (!this.resourceId) {
      this.notificationService.error('资源ID不能为空');
      this.loading.set(false);
      this.isLoading.set(false);
      return;
    }
    
    // 调用 resourceService.getResourceById 获取资源信息
    this.resourceService.getResourceById(this.resourceId).subscribe({
      next: (resource: Resource) => {
        console.log('Loaded Resource Detail:', resource);
        this.resource.set({
          resource_id: resource.resource_id,
          name: resource.name,
          description: resource.description,
          sensitive_flag: resource.sensitive_flag || false,
          created_at: resource.created_at || new Date().toISOString(),
          updated_at: resource.updated_at || new Date().toISOString()
        });
        
        // 无论响应如何，只要请求成功，就停止加载状态
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error: any) => {
        console.error('Error loading resource detail:', error);
        this.notificationService.error('加载资源详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  editResource(): void {
    this.router.navigate(['/resources/edit', this.resourceId]);
  }

  deleteResource(): void {
    if (confirm('确定要删除这个资源吗？')) {
      this.resourceService.deleteResource(this.resourceId).subscribe({
        next: (success: boolean) => {
          if (success) {
            this.notificationService.success('删除资源成功');
            this.router.navigate(['/resources']);
          } else {
            this.notificationService.error('删除资源失败');
          }
        },
        error: (error: any) => {
          this.notificationService.error('删除资源失败');
        }
      });
    }
  }

  backToList(): void {
    this.router.navigate(['/resources']);
  }
}
