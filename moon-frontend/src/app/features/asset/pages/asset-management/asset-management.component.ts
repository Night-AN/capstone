import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { Router } from '@angular/router';
import { AssetService } from '@core/services/asset.service';
import { AssetListItem } from '@models/asset.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-asset-management',
  imports: [
    CommonModule,
    FormsModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatButtonModule,
    MatIconModule,
    MatFormFieldModule,
    MatInputModule
  ],
  templateUrl: './asset-management.component.html',
  styleUrl: './asset-management.component.scss'
})
export class AssetManagementComponent implements OnInit {
  displayedColumns: string[] = ['asset_name', 'asset_code', 'organization_id', 'asset_type', 'ip_address', 'status', 'created_at', 'actions'];
  dataSource = new MatTableDataSource<AssetListItem>();
  searchKeyword: string = '';
  organizationId: string = '';

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  private assetService = inject(AssetService);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  constructor() { }

  ngOnInit(): void {
    this.loadAssets();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  loadAssets(): void {
    console.log('Loading assets...');
    this.assetService.getAssets().subscribe({
      next: (assets: AssetListItem[]) => {
        console.log('Received assets:', assets);
        this.dataSource.data = assets;
        console.log('DataSource data set:', this.dataSource.data);
      },
      error: (error: any) => {
        console.error('Error loading assets:', error);
        this.notificationService.error('加载资产失败');
      }
    });
  }

  // 按组织加载资产
  loadAssetsByOrganization(): void {
    if (!this.organizationId) {
      this.notificationService.error('请输入组织ID');
      return;
    }
    console.log('Loading assets by organization:', this.organizationId);
    this.assetService.getAssetsByOrganization(this.organizationId).subscribe({
      next: (assets: AssetListItem[]) => {
        console.log('Received assets by organization:', assets);
        this.dataSource.data = assets;
        console.log('DataSource data set:', this.dataSource.data);
      },
      error: (error: any) => {
        console.error('Error loading assets by organization:', error);
        this.notificationService.error('加载组织资产失败');
      }
    });
  }

  applyFilter(): void {
    this.dataSource.filter = this.searchKeyword.trim().toLowerCase();
  }

  createAsset(): void {
    this.router.navigate(['/assets/create']);
  }

  editAsset(assetId: string): void {
    this.router.navigate(['/assets/edit', assetId]);
  }

  viewAsset(assetId: string): void {
    this.router.navigate(['/assets/detail', assetId]);
  }

  deleteAsset(assetId: string): void {
    if (confirm('确定要删除这个资产吗？')) {
      this.assetService.deleteAsset(assetId).subscribe({
        next: (success: boolean) => {
          if (success) {
            this.notificationService.success('删除资产成功');
            this.loadAssets();
          } else {
            this.notificationService.error('删除资产失败');
          }
        },
        error: (error: any) => {
          this.notificationService.error('删除资产失败');
        }
      });
    }
  }
}
