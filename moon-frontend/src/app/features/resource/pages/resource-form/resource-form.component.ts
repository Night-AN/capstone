import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { ResourceService } from '@core/services/resource.service';
import { ResourceCreateRequest, ResourceUpdateRequest } from '@models/resource.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-resource-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './resource-form.component.html',
  styleUrl: './resource-form.component.scss'
})
export class ResourceFormComponent implements OnInit {
  resourceForm: FormGroup;
  resourceId: string = '';
  isEditMode: boolean = false;
  loading = signal<boolean>(false);
  isLoading = signal<boolean>(false);

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private resourceService = inject(ResourceService);
  private notificationService = inject(NotificationService);

  constructor() {
    this.resourceForm = this.fb.group({
      name: ['', Validators.required],
      description: ['']
    });
    this.resourceId = '';
    this.isEditMode = false;
    this.loading.set(false);
    this.isLoading.set(false);
  }

  ngOnInit(): void {
    this.resourceId = this.route.snapshot.paramMap.get('id') || '';
    this.isEditMode = this.resourceId !== '';

    if (this.isEditMode) {
      this.loadResourceDetail();
    }
  }

  loadResourceDetail(): void {
    if (this.isLoading()) {
      return;
    }
    
    this.isLoading.set(true);
    this.loading.set(true);
    
    this.resourceService.getResourceById(this.resourceId).subscribe({
      next: (resource) => {
        console.log('Loaded Resource Detail:', resource);
        this.resourceForm.patchValue({
          name: resource.name,
          description: resource.description
        });
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error) => {
        console.error('Error loading resource detail:', error);
        this.notificationService.error('加载资源详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  onSubmit(): void {
    if (this.resourceForm.invalid || this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    if (this.isEditMode) {
      // 编辑模式
      const updateRequest: ResourceUpdateRequest = {
        resource_id: this.resourceId,
        name: this.resourceForm.value.name,
        description: this.resourceForm.value.description
      };

      this.resourceService.updateResource(updateRequest).subscribe({
        next: (response) => {
          this.notificationService.success('更新资源成功');
          this.router.navigate(['/resources']);
          this.loading.set(false);
          this.isLoading.set(false);
        },
        error: (error) => {
          this.notificationService.error('更新资源失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    } else {
      // 创建模式
      const createRequest: ResourceCreateRequest = {
        name: this.resourceForm.value.name,
        description: this.resourceForm.value.description
      };

      this.resourceService.createResource(createRequest).subscribe({
        next: (response) => {
          this.notificationService.success('创建资源成功');
          this.router.navigate(['/resources']);
          this.loading.set(false);
          this.isLoading.set(false);
        },
        error: (error) => {
          this.notificationService.error('创建资源失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    }
  }

  onCancel(): void {
    this.router.navigate(['/resources']);
  }
}
