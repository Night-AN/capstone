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
import { ModelConfigCreateRequest, ModelConfigUpdateRequest } from '@models/ai.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-model-config-form',
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
  templateUrl: './model-config-form.component.html',
  styleUrl: './model-config-form.component.scss'
})
export class ModelConfigFormComponent implements OnInit {
  configForm: FormGroup;
  configId: string = '';
  isEditMode: boolean = false;
  loading = signal<boolean>(false);
  isLoading = signal<boolean>(false);

  providerOptions = [
    { value: 'openai', label: 'OpenAI' },
    { value: 'anthropic', label: 'Anthropic' },
    { value: 'qianwen', label: '阿里云通义千问' },
    { value: 'wenxin', label: '百度文心一言' },
    { value: 'moonshot', label: '月之暗面 Moonshot' },
    { value: 'zhipu', label: '智谱 GLM' }
  ];

  private fb = inject(FormBuilder);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private aiService = inject(AIService);
  private notificationService = inject(NotificationService);

  constructor() {
    this.configForm = this.fb.group({
      provider_name: ['', Validators.required],
      model_name: ['', Validators.required],
      api_key: [''],
      api_endpoint: [''],
      api_version: ['v1'],
      max_tokens: [4096, [Validators.required, Validators.min(1)]],
      temperature: [0.7, [Validators.required, Validators.min(0), Validators.max(2)]],
      timeout_seconds: [30, [Validators.required, Validators.min(1)]],
      is_active: [true],
      priority: [1, [Validators.required, Validators.min(0)]]
    });
  }

  ngOnInit(): void {
    this.configId = this.route.snapshot.paramMap.get('id') || '';
    this.isEditMode = this.configId !== '';

    if (this.isEditMode) {
      this.loadConfigDetail();
    }
  }

  loadConfigDetail(): void {
    if (this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    this.aiService.getModelConfig(this.configId).subscribe({
      next: (response) => {
        if (response.data) {
          this.configForm.patchValue({
            provider_name: response.data.provider_name,
            model_name: response.data.model_name,
            api_endpoint: response.data.api_endpoint,
            api_version: response.data.api_version,
            max_tokens: response.data.max_tokens,
            temperature: response.data.temperature,
            timeout_seconds: response.data.timeout_seconds,
            is_active: response.data.is_active,
            priority: response.data.priority
          });
        }
        this.loading.set(false);
        this.isLoading.set(false);
      },
      error: (error) => {
        this.notificationService.error('加载配置详情失败');
        this.loading.set(false);
        this.isLoading.set(false);
      }
    });
  }

  onSubmit(): void {
    if (this.configForm.invalid || this.isLoading()) {
      return;
    }

    this.isLoading.set(true);
    this.loading.set(true);

    if (this.isEditMode) {
      const updateRequest: ModelConfigUpdateRequest = {
        config_id: this.configId,
        ...this.configForm.value
      };

      this.aiService.updateModelConfig(updateRequest).subscribe({
        next: () => {
          this.notificationService.success('更新成功');
          this.router.navigate(['/ai/model-config']);
        },
        error: (error) => {
          this.notificationService.error('更新失败');
          this.loading.set(false);
          this.isLoading.set(false);
        }
      });
    } else {
      const createRequest: ModelConfigCreateRequest = this.configForm.value;

      this.aiService.createModelConfig(createRequest).subscribe({
        next: () => {
          this.notificationService.success('创建成功');
          this.router.navigate(['/ai/model-config']);
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
    this.router.navigate(['/ai/model-config']);
  }
}
