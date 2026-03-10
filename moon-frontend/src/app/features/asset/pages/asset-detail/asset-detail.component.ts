import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatChipsModule } from '@angular/material/chips';
import { ActivatedRoute, Router } from '@angular/router';
import { AssetService } from '@core/services/asset.service';
import { AIService } from '@core/services/ai.service';
import { Asset } from '@models/asset.model';
import { AssetClassification } from '@models/ai.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-asset-detail',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    MatProgressSpinnerModule,
    MatChipsModule
  ],
  templateUrl: './asset-detail.component.html',
  styleUrl: './asset-detail.component.scss'
})
export class AssetDetailComponent implements OnInit {
  asset = signal<Asset | null>(null);
  loading = signal<boolean>(true);
  assetId: string | null = null;

  classification = signal<AssetClassification | null>(null);
  classificationLoading = signal<boolean>(false);
  classifying = signal<boolean>(false);

  private assetService = inject(AssetService);
  private aiService = inject(AIService);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() { }

  ngOnInit(): void {
    this.assetId = this.route.snapshot.paramMap.get('id');
    if (this.assetId) {
      this.loadAssetDetail();
      this.loadClassification();
    } else {
      this.notificationService.error('资产ID不存在');
      this.router.navigate(['/assets']);
    }
  }

  loadAssetDetail(): void {
    if (!this.assetId) return;

    console.log('Loading asset with ID:', this.assetId);
    this.loading.set(true);
    this.assetService.getAssetById(this.assetId).subscribe({
      next: (asset: any) => {
        console.log('Received asset:', asset);
        this.asset.set(asset);
        this.loading.set(false);
        console.log('Asset loaded successfully');
      },
      error: (error: any) => {
        console.error('Error loading asset:', error);
        this.notificationService.error('加载资产详情失败');
        this.loading.set(false);
        console.log('Asset loading failed');
      }
    });
  }

  loadClassification(): void {
    if (!this.assetId) return;

    this.classificationLoading.set(true);
    this.aiService.getClassificationByAssetId(this.assetId).subscribe({
      next: (response: any) => {
        if (response.data) {
          this.classification.set(response.data);
        }
        this.classificationLoading.set(false);
      },
      error: () => {
        this.classificationLoading.set(false);
      }
    });
  }

  classifyAsset(): void {
    if (!this.assetId || this.classifying()) return;

    this.classifying.set(true);
    this.aiService.classifyAsset(this.assetId).subscribe({
      next: (response: any) => {
        this.notificationService.success('AI 分类成功');
        this.loadClassification();
        this.classifying.set(false);
      },
      error: (error: any) => {
        this.notificationService.error('AI 分类失败: ' + (error.message || '未知错误'));
        this.classifying.set(false);
      }
    });
  }

  approveClassification(approved: boolean): void {
    const currentClassification = this.classification();
    if (!currentClassification) return;

    this.aiService.approveClassification({
      classification_id: currentClassification.classification_id,
      approve: approved
    }).subscribe({
      next: () => {
        this.notificationService.success(approved ? '分类已批准' : '分类已拒绝');
        this.loadClassification();
      },
      error: () => {
        this.notificationService.error('操作失败');
      }
    });
  }

  editAsset(): void {
    if (this.assetId) {
      this.router.navigate(['/assets/edit', this.assetId]);
    }
  }

  backToList(): void {
    this.router.navigate(['/assets']);
  }

  getRiskLevelColor(level: string): string {
    const lowerLevel = level?.toLowerCase() || '';
    if (lowerLevel.includes('high') || lowerLevel.includes('高')) {
      return 'warn';
    } else if (lowerLevel.includes('medium') || lowerLevel.includes('中')) {
      return 'accent';
    }
    return 'primary';
  }
}
