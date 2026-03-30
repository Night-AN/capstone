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

interface Asset {
  id: string;
  asset_name: string;
  asset_code: string;
  asset_description: string;
  asset_flag: string;
  quantity: number;
  location: string;
  purchase_price: number;
  depreciation_price: number;
  purchase_date: string | null;
  manufacturer: string;
  model: string;
  other_metadata: any;
  asset_type_id: string | null;
  asset_category_id: string | null;
  organization_id: string | null;
  created_at: string;
  updated_at: string | null;
}

enum ModalTitle {
  Create = '创建资产',
  Edit = '编辑资产'
}

@Component({
  selector: 'app-assets-page',
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
  templateUrl: './assets-page.html',
  styleUrl: './assets-page.less',
  providers: [NzMessageService]
})
export class AssetsPage implements OnInit {
  assetForm: FormGroup;
  assets: Asset[] = [];
  originalAssets: Asset[] = [];

  isEditing = false;
  currentAssetId: string | null = null;

  visible = false;
  modalTitle = ModalTitle.Create;
  isLoading = false;
  searchKeyword: string = '';

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService
  ) {
    this.assetForm = this.fb.group({
      asset_name: ['', [Validators.required]],
      asset_code: ['', [Validators.required]],
      asset_description: ['', [Validators.required]],
      asset_flag: ['', [Validators.required]],
      quantity: [1, [Validators.required]],
      location: ['', []],
      purchase_price: [0, [Validators.required]],
      depreciation_price: [0, [Validators.required]],
      purchase_date: [null, []],
      manufacturer: ['', []],
      model: ['', []],
      other_metadata: [{}, []],
      asset_type_id: [null, []],
      asset_category_id: [null, []],
      organization_id: [null, []]
    });
  }

  ngOnInit(): void {
    this.loadAssets();
  }

  loadAssets(): void {
    // Mock data for now
    this.assets = [
      {
        id: '1',
        asset_name: '服务器A',
        asset_code: 'server-001',
        asset_description: '主服务器',
        asset_flag: 'ACTIVE',
        quantity: 1,
        location: '机房A',
        purchase_price: 10000,
        depreciation_price: 8000,
        purchase_date: '2024-01-01T00:00:00Z',
        manufacturer: 'Dell',
        model: 'PowerEdge R740',
        other_metadata: {},
        asset_type_id: '1',
        asset_category_id: '1',
        organization_id: '1',
        created_at: '2024-01-01T00:00:00Z',
        updated_at: null
      }
    ];
    this.originalAssets = [...this.assets];
  }

  searchAssets(): void {
    if (this.searchKeyword) {
      const keyword = this.searchKeyword.toLowerCase();
      this.assets = this.originalAssets.filter(asset => 
        asset.asset_name.toLowerCase().includes(keyword) ||
        asset.asset_code.toLowerCase().includes(keyword) ||
        asset.asset_description.toLowerCase().includes(keyword)
      );
    } else {
      this.assets = [...this.originalAssets];
    }
  }

  createAsset(): void {
    this.resetForm();
    this.modalTitle = ModalTitle.Create;
    this.visible = true;
  }

  editAsset(asset: Asset): void {
    this.isEditing = true;
    this.currentAssetId = asset.id;
    this.modalTitle = ModalTitle.Edit;
    this.assetForm.patchValue({
      asset_name: asset.asset_name,
      asset_code: asset.asset_code,
      asset_description: asset.asset_description,
      asset_flag: asset.asset_flag,
      quantity: asset.quantity,
      location: asset.location,
      purchase_price: asset.purchase_price,
      depreciation_price: asset.depreciation_price,
      purchase_date: asset.purchase_date,
      manufacturer: asset.manufacturer,
      model: asset.model,
      other_metadata: asset.other_metadata,
      asset_type_id: asset.asset_type_id,
      asset_category_id: asset.asset_category_id,
      organization_id: asset.organization_id
    });
    this.visible = true;
  }

  deleteAsset(id: string): void {
    this.assets = this.assets.filter(asset => asset.id !== id);
    this.message.success('资产删除成功');
  }

  handleCancel(): void {
    this.visible = false;
    this.resetForm();
  }

  handleOk(): void {
    if (this.assetForm.valid) {
      this.isLoading = true;
      setTimeout(() => {
        if (this.isEditing && this.currentAssetId) {
          this.updateAsset();
        } else {
          this.createAssetSubmit();
        }
        this.isLoading = false;
        this.visible = false;
      }, 500);
    } else {
      Object.values(this.assetForm.controls).forEach(control => {
        if (control.invalid) {
          control.markAsDirty();
          control.updateValueAndValidity({ onlySelf: true });
        }
      });
    }
  }

  createAssetSubmit(): void {
    const newAsset: Asset = {
      id: (this.assets.length + 1).toString(),
      ...this.assetForm.value,
      created_at: new Date().toISOString(),
      updated_at: null
    };
    this.assets.push(newAsset);
    this.message.success('资产创建成功');
    this.resetForm();
  }

  updateAsset(): void {
    if (this.currentAssetId) {
      const index = this.assets.findIndex(asset => asset.id === this.currentAssetId);
      if (index !== -1) {
        this.assets[index] = {
          ...this.assets[index],
          ...this.assetForm.value,
          updated_at: new Date().toISOString()
        };
        this.message.success('资产更新成功');
        this.resetForm();
      }
    }
  }

  resetForm(): void {
    this.assetForm.reset();
    this.isEditing = false;
    this.currentAssetId = null;
  }
}