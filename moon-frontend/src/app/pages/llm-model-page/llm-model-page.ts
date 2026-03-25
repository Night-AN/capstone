import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzMessageService } from 'ng-zorro-antd/message';
import { CommonModule } from '@angular/common';

interface LLMModel {
  id: string;
  model_name: string;
  provider: string;
  model_code: string;
  api_key: string;
  api_endpoint: string | null;
  enabled: boolean;
  max_tokens: number;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建LLM模型',
  Edit = '编辑LLM模型'
}

@Component({
  selector: 'app-llm-model-page',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NzButtonModule,
    NzCardModule,
    NzFormModule,
    NzInputModule,
    NzTableModule,
    NzModalModule
  ],
  templateUrl: './llm-model-page.html',
  styleUrl: './llm-model-page.less',
  providers: [NzMessageService]
})
export class LLMModelPage implements OnInit {
  modelForm: FormGroup;
  models: LLMModel[] = [];

  isEditing = false;
  currentModelId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService
  ) {
    this.modelForm = this.fb.group({
      model_name: ['', [Validators.required]],
      provider: ['', [Validators.required]],
      model_code: ['', [Validators.required]],
      api_key: ['', [Validators.required]],
      api_endpoint: ['', []],
      enabled: [true, []],
      max_tokens: [4096, [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadModels();
  }

  loadModels(): void {
    // Mock data for now
    this.models = [
      {
        id: '1',
        model_name: 'GPT-4',
        provider: 'OpenAI',
        model_code: 'gpt-4',
        api_key: 'sk-...',
        api_endpoint: 'https://api.openai.com/v1/chat/completions',
        enabled: true,
        max_tokens: 4096,
        created_at: '2024-01-01T00:00:00Z',
        updated_at: null
      }
    ];
  }

  createModel(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editModel(model: LLMModel): void {
    this.isEditing = true;
    this.currentModelId = model.id;
    this.modalTitle = ModalTitle.Edit;
    this.modelForm.patchValue({
      model_name: model.model_name,
      provider: model.provider,
      model_code: model.model_code,
      api_key: model.api_key,
      api_endpoint: model.api_endpoint,
      enabled: model.enabled,
      max_tokens: model.max_tokens
    });
    this.visible = true;
  }

  deleteModel(id: string): void {
    this.models = this.models.filter(model => model.id !== id);
    this.message.success('LLM模型删除成功');
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.modelForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentModelId) {
          this.updateModel();
        } else {
          this.createModelSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.modelForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createModelSubmit(): void {
    const newModel: LLMModel = {
      id: (this.models.length + 1).toString(),
      ...this.modelForm.value,
      created_at: new Date().toISOString(),
      updated_at: null
    };
    this.models.push(newModel);
    this.message.success('LLM模型创建成功');
    this.resetForm();
  }

  updateModel(): void {
    if (this.currentModelId) {
      const index = this.models.findIndex(model => model.id === this.currentModelId);
      if (index !== -1) {
        this.models[index] = {
          ...this.models[index],
          ...this.modelForm.value,
          updated_at: new Date().toISOString()
        };
        this.message.success('LLM模型更新成功');
        this.resetForm();
      }
    }
  }

  resetForm(): void {
    this.modelForm.reset();
    this.isEditing = false;
    this.currentModelId = null;
  }
}