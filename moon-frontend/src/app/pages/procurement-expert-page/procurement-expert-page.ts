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
import { ProcurementExpertService } from 'src/app/services/procurement-expert.service';

interface ProcurementExpert {
  id: string;
  expert_name: string;
  expert_code: string;
  expert_description: string;
  expert_flag: string;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建专家',
  Edit = '编辑专家'
}

@Component({
  selector: 'app-procurement-expert-page',
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
  templateUrl: './procurement-expert-page.html',
  styleUrl: './procurement-expert-page.less',
  providers: [NzMessageService]
})
export class ProcurementExpertPage implements OnInit {
  expertForm: FormGroup;
  experts: ProcurementExpert[] = [];

  isEditing = false;
  currentExpertId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;
  searchKeyword: string = '';

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private expertService: ProcurementExpertService
  ) {
    this.expertForm = this.fb.group({
      expert_name: ['', [Validators.required]],
      expert_code: ['', [Validators.required]],
      expert_description: ['', [Validators.required]],
      expert_flag: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadExperts();
  }

  loadExperts(): void {
    this.expertService.getList().subscribe(experts => {
      this.experts = experts[0].edges.map((edge: any) => ({
        id: edge.node.id,
        expert_name: edge.node.expert_name,
        expert_code: edge.node.expert_code,
        expert_description: edge.node.expert_description,
        expert_flag: edge.node.expert_flag,
        created_at: edge.node.created_at,
        updated_at: edge.node.updated_at
      }));
    });
  }

  searchExperts(): void {
    if (this.searchKeyword) {
      this.expertService.getList({ 
        expert_name: { contains: this.searchKeyword },
        expert_code: { contains: this.searchKeyword },
        expert_description: { contains: this.searchKeyword }
      }).subscribe(experts => {
        this.experts = experts[0].edges.map((edge: any) => ({
          id: edge.node.id,
          expert_name: edge.node.expert_name,
          expert_code: edge.node.expert_code,
          expert_description: edge.node.expert_description,
          expert_flag: edge.node.expert_flag,
          created_at: edge.node.created_at,
          updated_at: edge.node.updated_at
        }));
      });
    } else {
      this.loadExperts();
    }
  }

  createExpert(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editExpert(expert: ProcurementExpert): void {
    this.isEditing = true;
    this.currentExpertId = expert.id;
    this.modalTitle = ModalTitle.Edit;
    this.expertForm.patchValue({
      expert_name: expert.expert_name,
      expert_code: expert.expert_code,
      expert_description: expert.expert_description,
      expert_flag: expert.expert_flag
    });
    this.visible = true;
  }

  deleteExpert(id: string): void {
    this.expertService.delete(id).subscribe(() => {
      this.loadExperts();
      this.message.success('专家删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.expertForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentExpertId) {
          this.updateExpert();
        } else {
          this.createExpertSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.expertForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createExpertSubmit(): void {
    this.expertService.create(this.expertForm.value).subscribe(() => {
      this.loadExperts();
      this.message.success('专家创建成功');
      this.resetForm();
    });
  }

  updateExpert(): void {
    if (this.currentExpertId) {
      this.expertService.update(this.currentExpertId, this.expertForm.value).subscribe(() => {
        this.loadExperts();
        this.message.success('专家更新成功');
        this.resetForm();
      });
    }
  }

  resetForm(): void {
    this.expertForm.reset();
    this.isEditing = false;
    this.currentExpertId = null;
  }
}