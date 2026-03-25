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

interface AssetType {
  id: string;
  asset_type_name: string;
  asset_type_code: string;
  asset_type_flag: string;
  asset_category_id: string;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建资产类型',
  Edit = '编辑资产类型'
}

@Component({
  selector: 'app-asset-type-page',
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
  templateUrl: './asset-type-page.html',
  styleUrl: './asset-type-page.less',
  providers: [NzMessageService]
})
export class AssetTypePage implements OnInit {
  typeForm: FormGroup;
  types: AssetType[] = [];

  isEditing = false;
  currentTypeId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService
  ) {
    this.typeForm = this.fb.group({
      asset_type_name: ['', [Validators.required]],
      asset_type_code: ['', [Validators.required]],
      asset_type_flag: ['', [Validators.required]],
      asset_category_id: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadTypes();
  }

  loadTypes(): void {
    // Mock data for now
    this.types = [
      {
        id: '1',
        asset_type_name: '服务器',
        asset_type_code: 'server',
        asset_type_flag: 'ACTIVE',
        asset_category_id: '1',
        created_at: '2024-01-01T00:00:00Z',
        updated_at: null
      }
    ];
  }

  createType(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editType(type: AssetType): void {
    this.isEditing = true;
    this.currentTypeId = type.id;
    this.modalTitle = ModalTitle.Edit;
    this.typeForm.patchValue({
      asset_type_name: type.asset_type_name,
      asset_type_code: type.asset_type_code,
      asset_type_flag: type.asset_type_flag,
      asset_category_id: type.asset_category_id
    });
    this.visible = true;
  }

  deleteType(id: string): void {
    this.types = this.types.filter(type => type.id !== id);
    this.message.success('资产类型删除成功');
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.typeForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentTypeId) {
          this.updateType();
        } else {
          this.createTypeSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.typeForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createTypeSubmit(): void {
    const newType: AssetType = {
      id: (this.types.length + 1).toString(),
      ...this.typeForm.value,
      created_at: new Date().toISOString(),
      updated_at: null
    };
    this.types.push(newType);
    this.message.success('资产类型创建成功');
    this.resetForm();
  }

  updateType(): void {
    if (this.currentTypeId) {
      const index = this.types.findIndex(type => type.id === this.currentTypeId);
      if (index !== -1) {
        this.types[index] = {
          ...this.types[index],
          ...this.typeForm.value,
          updated_at: new Date().toISOString()
        };
        this.message.success('资产类型更新成功');
        this.resetForm();
      }
    }
  }

  resetForm(): void {
    this.typeForm.reset();
    this.isEditing = false;
    this.currentTypeId = null;
  }
}