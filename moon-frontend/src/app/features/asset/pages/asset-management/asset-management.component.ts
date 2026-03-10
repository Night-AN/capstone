import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild, inject, signal } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatDialog, MatDialogModule } from '@angular/material/dialog';
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
import { BatchImportDialogComponent } from '../../components/batch-import-dialog/batch-import-dialog.component';

@Component({
  selector: 'app-asset-management',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatButtonModule,
    MatIconModule,
    MatFormFieldModule,
    MatInputModule,
    MatDialogModule
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
  private dialog = inject(MatDialog);

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

  openBatchImportDialog(): void {
    const sampleAssets = [
      {
        asset_name: 'Web Server 01',
        asset_code: 'WS-001',
        asset_description: '公司Web服务器',
        asset_type: 'server',
        manufacturer: 'Dell',
        model: 'PowerEdge R740',
        ip_address: '192.168.1.100',
        mac_address: '00:1A:2B:3C:4D:5E',
        location: '机房1',
        department: 'IT部门',
        owner: '张三',
        status: 'running'
      },
      {
        asset_name: 'Database Server 01',
        asset_code: 'DB-001',
        asset_description: 'MySQL数据库服务器',
        asset_type: 'database',
        manufacturer: 'HP',
        model: 'ProLiant DL380',
        ip_address: '192.168.1.101',
        mac_address: '00:1A:2B:3C:4D:5F',
        location: '机房1',
        department: 'IT部门',
        owner: '李四',
        status: 'running'
      },
      {
        asset_name: 'Workstation-PC-001',
        asset_code: 'PC-001',
        asset_description: '员工办公电脑',
        asset_type: 'workstation',
        manufacturer: 'Lenovo',
        model: 'ThinkPad T490',
        ip_address: '192.168.1.200',
        mac_address: '00:1A:2B:3C:4D:60',
        location: '办公楼3层',
        department: '财务部',
        owner: '王五',
        status: 'running'
      },
      {
        asset_name: 'Cisco Switch 01',
        asset_code: 'SW-001',
        asset_description: '核心交换机',
        asset_type: 'network_device',
        manufacturer: 'Cisco',
        model: 'Catalyst 9300',
        ip_address: '192.168.1.1',
        mac_address: '00:1A:2B:3C:4D:61',
        location: '机房1',
        department: 'IT部门',
        owner: '张三',
        status: 'running'
      },
      {
        asset_name: 'Firewall 01',
        asset_code: 'FW-001',
        asset_description: '企业防火墙',
        asset_type: 'network_device',
        manufacturer: 'Fortinet',
        model: 'FortiGate 100F',
        ip_address: '192.168.1.2',
        mac_address: '00:1A:2B:3C:4D:62',
        location: '机房1',
        department: 'IT部门',
        owner: '张三',
        status: 'running'
      }
    ];

    const dialogRef = this.dialog.open(BatchImportDialogComponent, {
      width: '700px',
      maxHeight: '80vh',
      data: { assets: sampleAssets }
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        this.loadAssets();
      }
    });
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
