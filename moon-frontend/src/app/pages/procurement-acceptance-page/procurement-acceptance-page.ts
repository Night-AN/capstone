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
import { ProcurementAcceptanceService } from 'src/app/services/procurement-acceptance.service';

interface ProcurementAcceptance {
  id: string;
  acceptance_name: string;
  acceptance_code: string;
  acceptance_description: string;
  acceptance_flag: string;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建验收记录',
  Edit = '编辑验收记录'
}

@Component({
  selector: 'app-procurement-acceptance-page',
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
  templateUrl: './procurement-acceptance-page.html',
  styleUrl: './procurement-acceptance-page.less',
  providers: [NzMessageService]
})
export class ProcurementAcceptancePage implements OnInit {
  acceptanceForm: FormGroup;
  acceptances: ProcurementAcceptance[] = [];

  isEditing = false;
  currentAcceptanceId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private acceptanceService: ProcurementAcceptanceService
  ) {
    this.acceptanceForm = this.fb.group({
      acceptance_name: ['', [Validators.required]],
      acceptance_code: ['', [Validators.required]],
      acceptance_description: ['', [Validators.required]],
      acceptance_flag: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadAcceptances();
  }

  loadAcceptances(): void {
    this.acceptanceService.getList().subscribe(acceptances => {
      this.acceptances = acceptances[0].edges.map((edge: any) => ({
        id: edge.node.id,
        acceptance_name: edge.node.acceptance_name,
        acceptance_code: edge.node.acceptance_code,
        acceptance_description: edge.node.acceptance_description,
        acceptance_flag: edge.node.acceptance_flag,
        created_at: edge.node.created_at,
        updated_at: edge.node.updated_at
      }));
    });
  }

  createAcceptance(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editAcceptance(acceptance: ProcurementAcceptance): void {
    this.isEditing = true;
    this.currentAcceptanceId = acceptance.id;
    this.modalTitle = ModalTitle.Edit;
    this.acceptanceForm.patchValue({
      acceptance_name: acceptance.acceptance_name,
      acceptance_code: acceptance.acceptance_code,
      acceptance_description: acceptance.acceptance_description,
      acceptance_flag: acceptance.acceptance_flag
    });
    this.visible = true;
  }

  deleteAcceptance(id: string): void {
    this.acceptanceService.delete(id).subscribe(() => {
      this.loadAcceptances();
      this.message.success('验收记录删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.acceptanceForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentAcceptanceId) {
          this.updateAcceptance();
        } else {
          this.createAcceptanceSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.acceptanceForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createAcceptanceSubmit(): void {
    this.acceptanceService.create(this.acceptanceForm.value).subscribe(() => {
      this.loadAcceptances();
      this.message.success('验收记录创建成功');
      this.resetForm();
    });
  }

  updateAcceptance(): void {
    if (this.currentAcceptanceId) {
      this.acceptanceService.update(this.currentAcceptanceId, this.acceptanceForm.value).subscribe(() => {
        this.loadAcceptances();
        this.message.success('验收记录更新成功');
        this.resetForm();
      });
    }
  }

  resetForm(): void {
    this.acceptanceForm.reset();
    this.isEditing = false;
    this.currentAcceptanceId = null;
  }
}