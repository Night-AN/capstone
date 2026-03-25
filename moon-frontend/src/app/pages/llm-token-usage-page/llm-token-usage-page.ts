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

interface LLMTokenUsage {
  id: string;
  model_id: string;
  user_id: string | null;
  prompt_tokens: number;
  completion_tokens: number;
  total_tokens: number;
  cost_amount: number;
  created_at: string;
}

enum ModalTitle {
  Create = '创建Token使用记录',
  Edit = '编辑Token使用记录'
}

@Component({
  selector: 'app-llm-token-usage-page',
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
  templateUrl: './llm-token-usage-page.html',
  styleUrl: './llm-token-usage-page.less',
  providers: [NzMessageService]
})
export class LLMTokenUsagePage implements OnInit {
  usageForm: FormGroup;
  usages: LLMTokenUsage[] = [];

  isEditing = false;
  currentUsageId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService
  ) {
    this.usageForm = this.fb.group({
      model_id: ['', [Validators.required]],
      user_id: [null, []],
      prompt_tokens: [0, [Validators.required]],
      completion_tokens: [0, [Validators.required]],
      total_tokens: [0, [Validators.required]],
      cost_amount: [0, [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadUsages();
  }

  loadUsages(): void {
    // Mock data for now
    this.usages = [
      {
        id: '1',
        model_id: '1',
        user_id: '1',
        prompt_tokens: 100,
        completion_tokens: 200,
        total_tokens: 300,
        cost_amount: 0.015,
        created_at: '2024-01-01T00:00:00Z'
      }
    ];
  }

  createUsage(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editUsage(usage: LLMTokenUsage): void {
    this.isEditing = true;
    this.currentUsageId = usage.id;
    this.modalTitle = ModalTitle.Edit;
    this.usageForm.patchValue({
      model_id: usage.model_id,
      user_id: usage.user_id,
      prompt_tokens: usage.prompt_tokens,
      completion_tokens: usage.completion_tokens,
      total_tokens: usage.total_tokens,
      cost_amount: usage.cost_amount
    });
    this.visible = true;
  }

  deleteUsage(id: string): void {
    this.usages = this.usages.filter(usage => usage.id !== id);
    this.message.success('Token使用记录删除成功');
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.usageForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentUsageId) {
          this.updateUsage();
        } else {
          this.createUsageSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.usageForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createUsageSubmit(): void {
    const newUsage: LLMTokenUsage = {
      id: (this.usages.length + 1).toString(),
      ...this.usageForm.value,
      created_at: new Date().toISOString(),
      updated_at: null
    };
    this.usages.push(newUsage);
    this.message.success('Token使用记录创建成功');
    this.resetForm();
  }

  updateUsage(): void {
    if (this.currentUsageId) {
      const index = this.usages.findIndex(usage => usage.id === this.currentUsageId);
      if (index !== -1) {
        this.usages[index] = {
          ...this.usages[index],
          ...this.usageForm.value,
          updated_at: new Date().toISOString()
        };
        this.message.success('Token使用记录更新成功');
        this.resetForm();
      }
    }
  }

  resetForm(): void {
    this.usageForm.reset();
    this.isEditing = false;
    this.currentUsageId = null;
  }
}