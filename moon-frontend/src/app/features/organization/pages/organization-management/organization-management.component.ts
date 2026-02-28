import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { Router } from '@angular/router';
import { OrganizationService } from '@core/services/organization.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { OrganizationTreeNodeComponent } from '../../components/organization-tree-node/organization-tree-node.component';

@Component({
  selector: 'app-organization-management',
  imports: [
    CommonModule,
    FormsModule,
    MatButtonModule,
    MatIconModule,
    MatFormFieldModule,
    MatInputModule,
    MatProgressSpinnerModule,
    OrganizationTreeNodeComponent
  ],
  templateUrl: './organization-management.component.html',
  styleUrl: './organization-management.component.scss'
})
export class OrganizationManagementComponent implements OnInit {
  searchKeyword: string = '';
  loading = signal<boolean>(false);

  private organizationService = inject(OrganizationService);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() { }

  ngOnInit(): void {
    this.loadOrganizations();
  }

  organizations: any[] = [];

  loadOrganizations(): void {
	this.loading.set(true);
	this.organizationService.getOrganizationTree('').subscribe({
		next: (organizations: any[]) => {
			this.organizations = organizations;
			this.loading.set(false);
		},
		error: (error: any) => {
			this.notificationService.error('加载组织失败');
			this.loading.set(false);
		}
	});
}

  applyFilter(): void {
    const filterValue = this.searchKeyword.trim().toLowerCase();
    // Filter logic would go here if we were using a filtered array
    // For now, we'll just reload the organizations
    this.loadOrganizations();
  }

  createOrganization(): void {
    this.router.navigate(['/organizations/edit']);
  }

  editOrganization(organizationId: string): void {
    this.router.navigate(['/organizations/edit', organizationId]);
  }

  viewOrganization(organizationId: string): void {
    this.router.navigate(['/organizations/detail', organizationId]);
  }

  deleteOrganization(organizationId: string): void {
    if (confirm('确定要删除这个组织吗？')) {
      this.organizationService.deleteOrganization({ organization_id: organizationId }).subscribe({
        next: (response: any) => {
          this.notificationService.success('删除组织成功');
          this.loadOrganizations();
        },
        error: (error: any) => {
          this.notificationService.error('删除组织失败');
        }
      });
    }
  }

  // 分配角色给组织
  assignRoleToOrganization(organizationId: string): void {
    // 这里需要打开角色选择对话框
    // 暂时使用prompt模拟
    const roleId = prompt('请输入角色ID:');
    if (roleId) {
      this.organizationService.assignRoleToOrganization(organizationId, roleId).subscribe({
        next: (response) => {
          this.notificationService.success('分配角色成功');
        },
        error: (error) => {
          this.notificationService.error('分配角色失败');
        }
      });
    }
  }

  // 查看组织角色
  viewOrganizationRoles(organizationId: string): void {
    this.organizationService.getOrganizationRoles(organizationId).subscribe({
      next: (response) => {
        console.log('Organization roles:', response);
        // 这里需要显示组织角色列表
        alert('组织角色: ' + JSON.stringify(response.data));
      },
      error: (error) => {
        this.notificationService.error('获取组织角色失败');
      }
    });
  }

  // 移动组织
  moveOrganization(organizationId: string): void {
    // 这里需要打开组织选择对话框
    // 暂时使用prompt模拟
    const newParentId = prompt('请输入新父组织ID（留空表示设置为顶级组织）:');
    // 将null转换为undefined
    const parentId = newParentId || undefined;
    this.organizationService.moveOrganization(organizationId, parentId).subscribe({
      next: (response) => {
        this.notificationService.success('移动组织成功');
        this.loadOrganizations();
      },
      error: (error) => {
        this.notificationService.error('移动组织失败');
      }
    });
  }
}
