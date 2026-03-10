import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { ActivatedRoute, Router } from '@angular/router';
import { AIService } from '@core/services/ai.service';
import { PromptTemplateCreateRequest, PromptTemplateUpdateRequest } from '@models/ai.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-prompt-template-form',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatSelectModule,
    MatSlideToggleModule
  ],
  templateUrl: './prompt-template-form.component.html',
  styleUrl: './prompt-template-form.component.scss'
})
export class PromptTemplateFormComponent implements OnInit {
  templateForm: FormGroup;
  templateId: string = '';
  isEditMode: boolean = false;
  loading = signal<boolean>(false);
  isLoading = signal<boolean>(false);

  templateTypeOptions = [
    { value: 'asset_classification', label: '资产分类' },
    { value: 'risk_assessment', label: '风险评估' },
    { value: 'security_recommendation', label: '安全建议' },
    { value: 'vulnerability_analysis', label: '漏洞分析' }
  ];

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private aiService = inject(AIService);
  private notificationService = inject(NotificationService);

  constructor() {
    this.templateForm = this.fb.group({
      template_name: ['', Validators.required],
      template_type: ['', Validators.required],
      template_content: ['', Validators.required],
      description: [''],
      is_active: [true]
    });
  }

  ngOnInit(): void {
    this.templateId = this.route.snapshot.paramMap.get('id') || '';
    this.isEditMode = this.templateId !== '';

    if (this.isEditMode) {
      this.loadTemplateDetail();
    }
  }

  loadTemplateDetail(): void {
    if (this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    this.aiService.getPromptTemplate(this.templateId).subscribe({
      next: (response) => {
        if (response.data) {
          this.templateForm.patchValue({
            template_name: response.data.template_name,
            template_type: response.data.template_type,
            template_content: response.data.template_content,
            description: response.data.description,
            is_active: response.data.is_active
          });
        }
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error) => {
        this.notificationService.error('加载模板详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  onSubmit(): void {
    if (this.templateForm.invalid || this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    if (this.isEditMode) {
      const updateRequest: PromptTemplateUpdateRequest = {
        template_id: this.templateId,
        ...this.templateForm.value
      };

      this.aiService.updatePromptTemplate(updateRequest).subscribe({
        next: () => {
          this.notificationService.success('更新成功');
          this.router.navigate(['/ai/prompt-template']);
        },
        error: (error) => {
          this.notificationService.error('更新失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    } else {
      const createRequest: PromptTemplateCreateRequest = this.templateForm.value;

      this.aiService.createPromptTemplate(createRequest).subscribe({
        next: () => {
          this.notificationService.success('创建成功');
          this.router.navigate(['/ai/prompt-template']);
        },
        error: (error) => {
          this.notificationService.error('创建失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    }
  }

  onCancel(): void {
    this.router.navigate(['/ai/prompt-template']);
  }
}
