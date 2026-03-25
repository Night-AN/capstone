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

interface ProcurementFraudRisk {
  id: string;
  procurement_plan_id: string | null;
  procurement_implementation_id: string | null;
  risk_level: string;
  risk_reason: string | null;
  risk_score: number;
  created_at: string;
}

enum ModalTitle {
  Create = '创建欺诈风险记录',
  Edit = '编辑欺诈风险记录'
}

@Component({
  selector: 'app-procurement-fraud-risk-page',
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
  templateUrl: './procurement-fraud-risk-page.html',
  styleUrl: './procurement-fraud-risk-page.less',
  providers: [NzMessageService]
})
export class ProcurementFraudRiskPage implements OnInit {
  riskForm: FormGroup;
  risks: ProcurementFraudRisk[] = [];

  isEditing = false;
  currentRiskId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService
  ) {
    this.riskForm = this.fb.group({
      procurement_plan_id: [null, []],
      procurement_implementation_id: [null, []],
      risk_level: ['', [Validators.required]],
      risk_reason: ['', []],
      risk_score: [0, [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadRisks();
  }

  loadRisks(): void {
    // Mock data for now
    this.risks = [
      {
        id: '1',
        procurement_plan_id: '1',
        procurement_implementation_id: '1',
        risk_level: 'high',
        risk_reason: '供应商提供虚假信息',
        risk_score: 85,
        created_at: '2024-01-01T00:00:00Z'
      }
    ];
  }

  createRisk(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editRisk(risk: ProcurementFraudRisk): void {
    this.isEditing = true;
    this.currentRiskId = risk.id;
    this.modalTitle = ModalTitle.Edit;
    this.riskForm.patchValue({
      procurement_plan_id: risk.procurement_plan_id,
      procurement_implementation_id: risk.procurement_implementation_id,
      risk_level: risk.risk_level,
      risk_reason: risk.risk_reason,
      risk_score: risk.risk_score
    });
    this.visible = true;
  }

  deleteRisk(id: string): void {
    this.risks = this.risks.filter(risk => risk.id !== id);
    this.message.success('欺诈风险记录删除成功');
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.riskForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentRiskId) {
          this.updateRisk();
        } else {
          this.createRiskSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.riskForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createRiskSubmit(): void {
    const newRisk: ProcurementFraudRisk = {
      id: (this.risks.length + 1).toString(),
      ...this.riskForm.value,
      created_at: new Date().toISOString(),
      updated_at: null
    };
    this.risks.push(newRisk);
    this.message.success('欺诈风险记录创建成功');
    this.resetForm();
  }

  updateRisk(): void {
    if (this.currentRiskId) {
      const index = this.risks.findIndex(risk => risk.id === this.currentRiskId);
      if (index !== -1) {
        this.risks[index] = {
          ...this.risks[index],
          ...this.riskForm.value,
          updated_at: new Date().toISOString()
        };
        this.message.success('欺诈风险记录更新成功');
        this.resetForm();
      }
    }
  }

  resetForm(): void {
    this.riskForm.reset();
    this.isEditing = false;
    this.currentRiskId = null;
  }
}