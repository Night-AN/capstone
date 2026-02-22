import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatTreeModule } from '@angular/material/tree';
import { MatIconModule } from '@angular/material/icon';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { ActivatedRoute, Router } from '@angular/router';
import { OrganizationService } from '@core/services/organization.service';
import { OrganizationCreateRequest, OrganizationUpdateRequest } from '@models/organization.model';
import { NotificationService } from '@shared/service/notification/notification.service';
import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

interface OrganizationNode {
  organization_id: string;
  organization_name: string;
  children?: OrganizationNode[];
}

interface FlatOrganizationNode {
  expandable: boolean;
  organization_name: string;
  organization_id: string;
  level: number;
}

@Component({
  selector: 'app-organization-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatTreeModule,
    MatIconModule,
    MatCheckboxModule
  ],
  templateUrl: './organization-form.component.html',
  styleUrl: './organization-form.component.scss'
})
export class OrganizationFormComponent implements OnInit {
  organizationForm: FormGroup;
  organizationId: string = '';
  isEditMode: boolean = false;
  loading = signal<boolean>(false);
  isLoading = signal<boolean>(false);
  organizationTree = signal<any>(null);
  selectedParentId: string | null = null;

  // 树形控件配置
  private _transformer = (node: OrganizationNode, level: number) => {
    return {
      expandable: !!node.children && node.children.length > 0,
      organization_name: node.organization_name,
      organization_id: node.organization_id,
      level: level,
    };
  };

  treeControl = new FlatTreeControl<FlatOrganizationNode>(
    node => node.level,
    node => node.expandable
  );

  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private organizationService = inject(OrganizationService);
  private notificationService = inject(NotificationService);

  constructor() {
    this.organizationForm = this.fb.group({
      organization_name: ['', Validators.required],
      organization_description: [''],
      organization_code: ['', Validators.required],
      organization_flag: ['', Validators.required]
    });
    this.organizationId = '';
    this.isEditMode = false;
    this.loading.set(false);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.organizationId = this.route.snapshot.paramMap.get('id') || '';
    this.isEditMode = this.organizationId !== '';

    if (this.isEditMode) {
      this.loadOrganizationDetail();
    }

    // 加载组织树
    this.loadOrganizationTree();
  }

  loadOrganizationTree(): void {
    this.organizationService.getOrganizationTree().subscribe({
      next: (tree: any) => {
        console.log('Organization Tree:', tree);
        // 转换组织树数据为树形控件所需的格式
        const organizationNode: OrganizationNode = {
          organization_id: tree.organization_id,
          organization_name: tree.organization_name,
          children: this.convertToOrganizationNodes(tree.children || [])
        };
        this.dataSource.data = [organizationNode];
        this.organizationTree.set(tree);
      },
      error: (error: any) => {
        this.notificationService.error('加载组织树失败');
      }
    });
  }

  convertToOrganizationNodes(nodes: any[]): OrganizationNode[] {
    return nodes.map(node => ({
      organization_id: node.organization_id,
      organization_name: node.organization_name,
      children: this.convertToOrganizationNodes(node.children || [])
    }));
  }

  hasChild = (_: number, node: FlatOrganizationNode) => node.expandable;

  selectParent(node: FlatOrganizationNode): void {
    this.selectedParentId = node.organization_id;
    this.notificationService.success(`已选择父组织: ${node.organization_name}`);
  }

  loadOrganizationDetail(): void {
    if (this.isLoading()) {
      return;
    }
    
    this.isLoading.set(true);
    this.loading.set(true);
    
    this.organizationService.getOrganizationById(this.organizationId).subscribe({
      next: (organization: any) => {
        this.organizationForm.patchValue({
          organization_name: organization.organization_name,
          organization_description: organization.organization_description,
          organization_code: organization.organization_code,
          organization_flag: organization.organization_flag
        });
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error: any) => {
        this.notificationService.error('加载组织详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  onSubmit(): void {
    if (this.organizationForm.invalid || this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    if (this.isEditMode) {
      // 编辑模式
      const updateRequest: OrganizationUpdateRequest = {
        organization_id: this.organizationId,
        organization_name: this.organizationForm.value.organization_name,
        organization_description: this.organizationForm.value.organization_description,
        organization_code: this.organizationForm.value.organization_code,
        organization_flag: this.organizationForm.value.organization_flag,
        parent_id: this.selectedParentId
      };

      this.organizationService.updateOrganization(updateRequest).subscribe({
        next: (response: any) => {
          this.notificationService.success('更新组织成功');
          this.router.navigate(['/organizations']);
          this.loading.set(false);
          this.isLoading.set(false);
        },
        error: (error: any) => {
          this.notificationService.error('更新组织失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    } else {
      // 创建模式
      const createRequest: OrganizationCreateRequest = {
        organization_name: this.organizationForm.value.organization_name,
        organization_description: this.organizationForm.value.organization_description,
        organization_code: this.organizationForm.value.organization_code,
        organization_flag: this.organizationForm.value.organization_flag,
        parent_id: this.selectedParentId
      };

      this.organizationService.createOrganization(createRequest).subscribe({
        next: (response: any) => {
          this.notificationService.success('创建组织成功');
          this.router.navigate(['/organizations']);
          this.loading.set(false);
          this.isLoading.set(false);
        },
        error: (error: any) => {
          this.notificationService.error('创建组织失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    }
  }

  onCancel(): void {
    this.router.navigate(['/organizations']);
  }
}
