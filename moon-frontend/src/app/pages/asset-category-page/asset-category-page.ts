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

interface AssetCategory {
  id: string;
  category_name: string;
  category_code: string;
  category_flag: string;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建资产分类',
  Edit = '编辑资产分类'
}

@Component({
  selector: 'app-asset-category-page',
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
  templateUrl: './asset-category-page.html',
  styleUrl: './asset-category-page.less',
  providers: [NzMessageService]
})
export class AssetCategoryPage implements OnInit {
  categoryForm: FormGroup;
  categories: AssetCategory[] = [];

  isEditing = false;
  currentCategoryId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService
  ) {
    this.categoryForm = this.fb.group({
      category_name: ['', [Validators.required]],
      category_code: ['', [Validators.required]],
      category_flag: ['', [Validators.required]]
    });
  }

  ngOnInit(): void {
    this.loadCategories();
  }

  loadCategories(): void {
    // Mock data for now
    this.categories = [
      {
        id: '1',
        category_name: 'IT设备',
        category_code: 'it-equipment',
        category_flag: 'ACTIVE',
        created_at: '2024-01-01T00:00:00Z',
        updated_at: null
      }
    ];
  }

  createCategory(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editCategory(category: AssetCategory): void {
    this.isEditing = true;
    this.currentCategoryId = category.id;
    this.modalTitle = ModalTitle.Edit;
    this.categoryForm.patchValue({
      category_name: category.category_name,
      category_code: category.category_code,
      category_flag: category.category_flag
    });
    this.visible = true;
  }

  deleteCategory(id: string): void {
    this.categories = this.categories.filter(category => category.id !== id);
    this.message.success('资产分类删除成功');
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.categoryForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentCategoryId) {
          this.updateCategory();
        } else {
          this.createCategorySubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.categoryForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createCategorySubmit(): void {
    const newCategory: AssetCategory = {
      id: (this.categories.length + 1).toString(),
      ...this.categoryForm.value,
      created_at: new Date().toISOString(),
      updated_at: null
    };
    this.categories.push(newCategory);
    this.message.success('资产分类创建成功');
    this.resetForm();
  }

  updateCategory(): void {
    if (this.currentCategoryId) {
      const index = this.categories.findIndex(category => category.id === this.currentCategoryId);
      if (index !== -1) {
        this.categories[index] = {
          ...this.categories[index],
          ...this.categoryForm.value,
          updated_at: new Date().toISOString()
        };
        this.message.success('资产分类更新成功');
        this.resetForm();
      }
    }
  }

  resetForm(): void {
    this.categoryForm.reset();
    this.isEditing = false;
    this.currentCategoryId = null;
  }
}