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

interface LLMProcurementAnalysis {
  id: string;
  model_id: string | null;
  procurement_plan_id: string | null;
  procurement_implementation_id: string | null;
  analysis_result: string | null;
  token_usage_id: string | null;
  created_at: string;
}

enum ModalTitle {
  Create = '创建采购分析',
  Edit = '编辑采购分析'
}

@Component({
  selector: 'app-llm-procurement-analysis-page',
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
  templateUrl: './llm-procurement-analysis-page.html',
  styleUrl: './llm-procurement-analysis-page.less',
  providers: [NzMessageService]
})
export class LLMProcurementAnalysisPage implements OnInit {
  analysisForm: FormGroup;
  analyses: LLMProcurementAnalysis[] = [];

  isEditing = false;
  currentAnalysisId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService
  ) {
    this.analysisForm = this.fb.group({
      model_id: [null, []],
      procurement_plan_id: [null, []],
      procurement_implementation_id: [null, []],
      analysis_result: ['', []],
      token_usage_id: [null, []]
    });
  }

  ngOnInit(): void {
    this.loadAnalyses();
  }

  loadAnalyses(): void {
    // Mock data for now
    this.analyses = [
      {
        id: '1',
        model_id: '1',
        procurement_plan_id: '1',
        procurement_implementation_id: '1',
        analysis_result: '分析结果：供应商价格合理，可靠性高',
        token_usage_id: '1',
        created_at: '2024-01-01T00:00:00Z'
      }
    ];
  }

  createAnalysis(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editAnalysis(analysis: LLMProcurementAnalysis): void {
    this.isEditing = true;
    this.currentAnalysisId = analysis.id;
    this.modalTitle = ModalTitle.Edit;
    this.analysisForm.patchValue({
      model_id: analysis.model_id,
      procurement_plan_id: analysis.procurement_plan_id,
      procurement_implementation_id: analysis.procurement_implementation_id,
      analysis_result: analysis.analysis_result,
      token_usage_id: analysis.token_usage_id
    });
    this.visible = true;
  }

  deleteAnalysis(id: string): void {
    this.analyses = this.analyses.filter(analysis => analysis.id !== id);
    this.message.success('采购分析删除成功');
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.analysisForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentAnalysisId) {
          this.updateAnalysis();
        } else {
          this.createAnalysisSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.analysisForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createAnalysisSubmit(): void {
    const newAnalysis: LLMProcurementAnalysis = {
      id: (this.analyses.length + 1).toString(),
      ...this.analysisForm.value,
      created_at: new Date().toISOString(),
      updated_at: null
    };
    this.analyses.push(newAnalysis);
    this.message.success('采购分析创建成功');
    this.resetForm();
  }

  updateAnalysis(): void {
    if (this.currentAnalysisId) {
      const index = this.analyses.findIndex(analysis => analysis.id === this.currentAnalysisId);
      if (index !== -1) {
        this.analyses[index] = {
          ...this.analyses[index],
          ...this.analysisForm.value,
          updated_at: new Date().toISOString()
        };
        this.message.success('采购分析更新成功');
        this.resetForm();
      }
    }
  }

  resetForm(): void {
    this.analysisForm.reset();
    this.isEditing = false;
    this.currentAnalysisId = null;
  }
}