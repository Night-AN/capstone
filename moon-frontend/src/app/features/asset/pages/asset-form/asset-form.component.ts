import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormBuilder, FormGroup, ReactiveFormsModule, Validators } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { ActivatedRoute, Router } from '@angular/router';
import { AssetService } from '@core/services/asset.service';
import { Asset } from '@models/asset.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-asset-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatCardModule,
    MatFormFieldModule,
    MatInputModule,
    MatSelectModule,
    MatButtonModule,
    MatIconModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './asset-form.component.html',
  styleUrl: './asset-form.component.scss'
})
export class AssetFormComponent implements OnInit {
  assetForm: FormGroup;
  isEditMode = false;
  assetId: string | null = null;
  loading = signal<boolean>(false);

  private formBuilder = inject(FormBuilder);
  private assetService = inject(AssetService);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() {
    this.assetForm = this.formBuilder.group({
      asset_name: ['', Validators.required],
      asset_code: ['', Validators.required],
      asset_description: [''],
      asset_type: ['', Validators.required],
      asset_class: ['', Validators.required],
      manufacturer: [''],
      model: [''],
      serial_number: [''],
      ip_address: [''],
      mac_address: [''],
      location: [''],
      department: [''],
      owner: [''],
      contact_info: [''],
      status: ['', Validators.required],
      purchase_date: [''],
      warranty_end_date: [''],
      value: [''],
      notes: ['']
    });
  }

  ngOnInit(): void {
    this.assetId = this.route.snapshot.paramMap.get('id');
    if (this.assetId) {
      this.isEditMode = true;
      this.loadAssetData();
    }
  }

  loadAssetData(): void {
    if (!this.assetId) return;

    this.loading.set(true);
    this.assetService.getAssetById(this.assetId).subscribe({
      next: (asset: Asset) => {
        this.assetForm.patchValue(asset);
        this.loading.set(false);
      },
      error: (error: any) => {
        console.error('Error loading asset:', error);
        this.notificationService.error('加载资产数据失败');
        this.loading.set(false);
      }
    });
  }

  onSubmit(): void {
    if (this.assetForm.invalid) {
      return;
    }

    this.loading.set(true);
    const formData = this.assetForm.value;

    if (this.isEditMode && this.assetId) {
      // 更新资产
      const updateData = {
        asset_id: this.assetId,
        ...formData
      };

      this.assetService.updateAsset(updateData).subscribe({
        next: () => {
          this.notificationService.success('更新资产成功');
          this.router.navigate(['/assets/detail', this.assetId]);
          this.loading.set(false);
        },
        error: (error: any) => {
          console.error('Error updating asset:', error);
          this.notificationService.error('更新资产失败');
          this.loading.set(false);
        }
      });
    } else {
      // 创建资产
      this.assetService.createAsset(formData).subscribe({
        next: (asset: Asset) => {
          this.notificationService.success('创建资产成功');
          this.router.navigate(['/assets/detail', asset.asset_id]);
          this.loading.set(false);
        },
        error: (error: any) => {
          console.error('Error creating asset:', error);
          this.notificationService.error('创建资产失败');
          this.loading.set(false);
        }
      });
    }
  }

  onCancel(): void {
    if (this.isEditMode && this.assetId) {
      this.router.navigate(['/assets/detail', this.assetId]);
    } else {
      this.router.navigate(['/assets']);
    }
  }
}
