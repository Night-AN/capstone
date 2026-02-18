import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { OrganizationService } from '@core/services/organization.service';
import { OrganizationCreateRequest, OrganizationUpdateRequest } from '@models/organization.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-organization-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule
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

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private organizationService = inject(OrganizationService);
  private notificationService = inject(NotificationService);

  constructor() {
    this.organizationForm = this.fb.group({
      name: ['', Validators.required],
      description: ['']
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
  }

  loadOrganizationDetail(): void {
    if (this.isLoading()) {
      return;
    }
    
    this.isLoading.set(true);
    this.loading.set(true);
    
    this.organizationService.getOrganizationById(this.organizationId).subscribe({
      next: (response: any) => {
        if (response.data) {
          this.organizationForm.patchValue({
            name: response.data.name,
            description: response.data.description
          });
        }
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
        name: this.organizationForm.value.name,
        description: this.organizationForm.value.description
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
        name: this.organizationForm.value.name,
        description: this.organizationForm.value.description
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
