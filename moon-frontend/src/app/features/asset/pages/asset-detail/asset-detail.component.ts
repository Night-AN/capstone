import { CommonModule } from '@angular/common';
import { Component, OnInit, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { ActivatedRoute, Router } from '@angular/router';
import { AssetService } from '@core/services/asset.service';
import { Asset } from '@models/asset.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-asset-detail',
  imports: [
    CommonModule,
    FormsModule,
    MatCardModule,
    MatButtonModule,
    MatIconModule,
    MatProgressSpinnerModule
  ],
  templateUrl: './asset-detail.component.html',
  styleUrl: './asset-detail.component.scss'
})
export class AssetDetailComponent implements OnInit {
  asset = signal<Asset | null>(null);
  loading = signal<boolean>(true);
  assetId: string | null = null;

  private assetService = inject(AssetService);
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() { }

  ngOnInit(): void {
    this.assetId = this.route.snapshot.paramMap.get('id');
    if (this.assetId) {
      this.loadAssetDetail();
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

  editAsset(): void {
    if (this.assetId) {
      this.router.navigate(['/assets/edit', this.assetId]);
    }
  }

  backToList(): void {
    this.router.navigate(['/assets']);
  }
}
