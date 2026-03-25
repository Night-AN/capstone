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
import { ProcurementPlanService } from 'src/app/services/procurement-plan.service';

interface ProcurementPlan {
  id: string;
  ProcurementPlanName: string;
  ProcurementPlanCode: string;
  ProcurementPlanDescription: string;
  ProcurementPlanFlag: string;
  ProcurementPlanQuantity: number;
  ProcurementPlanPrice: number;
  ProcurementPlanPurchaseDate: string | null;
  ProcurementPlanPurchaseType: string;
  OtherMetadata: any;
  CreatedAt: string;
  UpdatedAt: string | null;
  procurement_plan_type_id: string | null;
  organization_id: string | null;
  user_id: string | null;
}

enum ModalTitle {
  Create = '创建采购计划',
  Edit = '编辑采购计划'
}

@Component({
  selector: 'app-procurement-plan-page',
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
  templateUrl: './procurement-plan-page.html',
  styleUrl: './procurement-plan-page.less',
  providers: [NzMessageService]
})
export class ProcurementPlanPage implements OnInit {
  planForm: FormGroup;
  plans: ProcurementPlan[] = [];

  isEditing = false;
  currentPlanId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private planService: ProcurementPlanService
  ) {
    this.planForm = this.fb.group({
      ProcurementPlanName: ['', [Validators.required]],
      ProcurementPlanCode: ['', [Validators.required]],
      ProcurementPlanDescription: ['', [Validators.required]],
      ProcurementPlanFlag: ['', [Validators.required]],
      ProcurementPlanQuantity: [1, [Validators.required]],
      ProcurementPlanPrice: [0, [Validators.required]],
      ProcurementPlanPurchaseDate: [null, []],
      ProcurementPlanPurchaseType: ['', []],
      OtherMetadata: [{}, []],
      procurement_plan_type_id: [null, []],
      organization_id: [null, []],
      user_id: [null, []]
    });
  }

  ngOnInit(): void {
    this.loadPlans();
  }

  loadPlans(): void {
    this.planService.getList().subscribe(plans => {
      this.plans = plans[0].edges.map((edge: any) => ({
        id: edge.node.id,
        ProcurementPlanName: edge.node.ProcurementPlanName,
        ProcurementPlanCode: edge.node.ProcurementPlanCode,
        ProcurementPlanDescription: edge.node.ProcurementPlanDescription,
        ProcurementPlanFlag: edge.node.ProcurementPlanFlag,
        ProcurementPlanQuantity: edge.node.ProcurementPlanQuantity,
        ProcurementPlanPrice: edge.node.ProcurementPlanPrice,
        ProcurementPlanPurchaseDate: edge.node.ProcurementPlanPurchaseDate,
        ProcurementPlanPurchaseType: edge.node.ProcurementPlanPurchaseType,
        OtherMetadata: edge.node.OtherMetadata,
        CreatedAt: edge.node.CreatedAt,
        UpdatedAt: edge.node.UpdatedAt,
        procurement_plan_type_id: edge.node.procurement_plan_type_id,
        organization_id: edge.node.organization_id,
        user_id: edge.node.user_id
      }));
    });
  }

  createPlan(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editPlan(plan: ProcurementPlan): void {
    this.isEditing = true;
    this.currentPlanId = plan.id;
    this.modalTitle = ModalTitle.Edit;
    this.planForm.patchValue({
      ProcurementPlanName: plan.ProcurementPlanName,
      ProcurementPlanCode: plan.ProcurementPlanCode,
      ProcurementPlanDescription: plan.ProcurementPlanDescription,
      ProcurementPlanFlag: plan.ProcurementPlanFlag,
      ProcurementPlanQuantity: plan.ProcurementPlanQuantity,
      ProcurementPlanPrice: plan.ProcurementPlanPrice,
      ProcurementPlanPurchaseDate: plan.ProcurementPlanPurchaseDate,
      ProcurementPlanPurchaseType: plan.ProcurementPlanPurchaseType,
      OtherMetadata: plan.OtherMetadata,
      procurement_plan_type_id: plan.procurement_plan_type_id,
      organization_id: plan.organization_id,
      user_id: plan.user_id
    });
    this.visible = true;
  }

  deletePlan(id: string): void {
    this.planService.delete(id).subscribe(() => {
      this.loadPlans();
      this.message.success('采购计划删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.planForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentPlanId) {
          this.updatePlan();
        } else {
          this.createPlanSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.planForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createPlanSubmit(): void {
    this.planService.create(this.planForm.value).subscribe(() => {
      this.loadPlans();
      this.message.success('采购计划创建成功');
      this.resetForm();
    });
  }

  updatePlan(): void {
    if (this.currentPlanId) {
      this.planService.update(this.currentPlanId, this.planForm.value).subscribe(() => {
        this.loadPlans();
        this.message.success('采购计划更新成功');
        this.resetForm();
      });
    }
  }

  resetForm(): void {
    this.planForm.reset();
    this.isEditing = false;
    this.currentPlanId = null;
  }
}