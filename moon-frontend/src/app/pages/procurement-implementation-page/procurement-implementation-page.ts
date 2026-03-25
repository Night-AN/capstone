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
import { ProcurementImplementationService } from 'src/app/services/procurement-implementation.service';

interface ProcurementImplementation {
  id: string;
  implementation_name: string;
  implementation_code: string;
  implementation_description: string;
  implementation_flag: string;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建实施记录',
  Edit = '编辑实施记录'
}

@Component({
  selector: 'app-procurement-implementation-page',
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
  templateUrl: './procurement-implementation-page.html',
  styleUrl: './procurement-implementation-page.less',
  providers: [NzMessageService]
})
export class ProcurementImplementationPage implements OnInit {
  implementationForm: FormGroup;
  implementations: ProcurementImplementation[] = [];

  isEditing = false;
  currentImplementationId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService,
    private implementationService: ProcurementImplementationService
  ) {
    this.implementationForm = this.fb.group({
      implementation_name: ['', [Validators.required]],
      implementation_code: ['', [Validators.required]],
      implementation_description: ['', [Validators.required]],
      implementation_flag: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadImplementations();
  }

  loadImplementations(): void {
    this.implementationService.getList().subscribe(implementations => {
      this.implementations = implementations[0].edges.map((edge: any) => ({
        id: edge.node.id,
        implementation_name: edge.node.implementation_name,
        implementation_code: edge.node.implementation_code,
        implementation_description: edge.node.implementation_description,
        implementation_flag: edge.node.implementation_flag,
        created_at: edge.node.created_at,
        updated_at: edge.node.updated_at
      }));
    });
  }

  createImplementation(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editImplementation(implementation: ProcurementImplementation): void {
    this.isEditing = true;
    this.currentImplementationId = implementation.id;
    this.modalTitle = ModalTitle.Edit;
    this.implementationForm.patchValue({
      implementation_name: implementation.implementation_name,
      implementation_code: implementation.implementation_code,
      implementation_description: implementation.implementation_description,
      implementation_flag: implementation.implementation_flag
    });
    this.visible = true;
  }

  deleteImplementation(id: string): void {
    this.implementationService.delete(id).subscribe(() => {
      this.loadImplementations();
      this.message.success('实施记录删除成功');
    });
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.implementationForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentImplementationId) {
          this.updateImplementation();
        } else {
          this.createImplementationSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.implementationForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createImplementationSubmit(): void {
    this.implementationService.create(this.implementationForm.value).subscribe(() => {
      this.loadImplementations();
      this.message.success('实施记录创建成功');
      this.resetForm();
    });
  }

  updateImplementation(): void {
    if (this.currentImplementationId) {
      this.implementationService.update(this.currentImplementationId, this.implementationForm.value).subscribe(() => {
        this.loadImplementations();
        this.message.success('实施记录更新成功');
        this.resetForm();
      });
    }
  }

  resetForm(): void {
    this.implementationForm.reset();
    this.isEditing = false;
    this.currentImplementationId = null;
  }
}