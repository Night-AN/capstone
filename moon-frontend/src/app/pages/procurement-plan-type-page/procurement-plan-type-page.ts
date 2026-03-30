import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, FormsModule, ReactiveFormsModule, Validators } from '@angular/forms';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzInputModule } from 'ng-zorro-antd/input';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzModalModule } from 'ng-zorro-antd/modal';
import { NzMessageService } from 'ng-zorro-antd/message';
import { CommonModule } from '@angular/common';
import { ProcurementPlanTypeService } from 'src/app/services/procurement-plan-type.service';

interface ProcurementPlanType {
  id: string;
  ProcurementPlanTypeName: string;
  ProcurementPlanTypeCode: string;
  ProcurementPlanTypeFlag: string;
  CreatedAt: string;
  UpdatedAt: string | null;
}

enum ModalTitle {
  Create = '创建计划类型',
  Edit = '编辑计划类型'
}

@Component({
  selector: 'app-procurement-plan-type-page',
  imports: [
    CommonModule,
    FormsModule,
    ReactiveFormsModule,
    NzButtonModule,
    NzCardModule,
    NzFormModule,
    NzInputModule,
    NzTableModule,
    NzModalModule
  ],
  templateUrl: './procurement-plan-type-page.html',
  styleUrl: './procurement-plan-type-page.less',
  providers: [NzMessageService]
})
export class ProcurementPlanTypePage implements OnInit {
  planTypeForm: FormGroup;
  planTypes: ProcurementPlanType[] = [];

  isEditing = false;
  currentPlanTypeId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;
  searchKeyword: string = '';

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private planTypeService: ProcurementPlanTypeService
  ) {
    this.planTypeForm = this.fb.group({
      ProcurementPlanTypeName: ['', [Validators.required]],
      ProcurementPlanTypeCode: ['', [Validators.required]],
      ProcurementPlanTypeFlag: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadPlanTypes();
  }

  loadPlanTypes(): void {
    this.planTypeService.getList().subscribe(planTypes => {
      this.planTypes = planTypes[0].edges.map((edge: any) => ({
        id: edge.node.id,
        ProcurementPlanTypeName: edge.node.ProcurementPlanTypeName,
        ProcurementPlanTypeCode: edge.node.ProcurementPlanTypeCode,
        ProcurementPlanTypeFlag: edge.node.ProcurementPlanTypeFlag,
        CreatedAt: edge.node.CreatedAt,
        UpdatedAt: edge.node.UpdatedAt
      }));
    });
  }

  searchPlanTypes(): void {
    if (this.searchKeyword) {
      this.planTypeService.getList({ 
        ProcurementPlanTypeName: { contains: this.searchKeyword },
        ProcurementPlanTypeCode: { contains: this.searchKeyword },
        ProcurementPlanTypeFlag: { contains: this.searchKeyword }
      }).subscribe(planTypes => {
        this.planTypes = planTypes[0].edges.map((edge: any) => ({
          id: edge.node.id,
          ProcurementPlanTypeName: edge.node.ProcurementPlanTypeName,
          ProcurementPlanTypeCode: edge.node.ProcurementPlanTypeCode,
          ProcurementPlanTypeFlag: edge.node.ProcurementPlanTypeFlag,
          CreatedAt: edge.node.CreatedAt,
          UpdatedAt: edge.node.UpdatedAt
        }));
      });
    } else {
      this.loadPlanTypes();
    }
  }

  createPlanType(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editPlanType(planType: ProcurementPlanType): void {
    this.isEditing = true;
    this.currentPlanTypeId = planType.id;
    this.modalTitle = ModalTitle.Edit;
    this.planTypeForm.patchValue({
      ProcurementPlanTypeName: planType.ProcurementPlanTypeName,
      ProcurementPlanTypeCode: planType.ProcurementPlanTypeCode,
      ProcurementPlanTypeFlag: planType.ProcurementPlanTypeFlag
    });
    this.visible = true;
  }

  deletePlanType(id: string): void {
    this.planTypeService.delete(id).subscribe(() => {
      this.loadPlanTypes();
      this.message.success('计划类型删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.planTypeForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentPlanTypeId) {
          this.updatePlanType();
        } else {
          this.createPlanTypeSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.planTypeForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createPlanTypeSubmit(): void {
    this.planTypeService.create(this.planTypeForm.value).subscribe(() => {
      this.loadPlanTypes();
      this.message.success('计划类型创建成功');
      this.resetForm();
    });
  }

  updatePlanType(): void {
    if (this.currentPlanTypeId) {
      this.planTypeService.update(this.currentPlanTypeId, this.planTypeForm.value).subscribe(() => {
        this.loadPlanTypes();
        this.message.success('计划类型更新成功');
        this.resetForm();
      });
    }
  }

  resetForm(): void {
    this.planTypeForm.reset();
    this.isEditing = false;
    this.currentPlanTypeId = null;
  }
}