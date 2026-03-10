import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { Router } from '@angular/router';
import { AIService } from '@core/services/ai.service';
import { ModelConfig } from '@models/ai.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-model-config-management',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatButtonModule,
    MatIconModule
  ],
  templateUrl: './model-config-management.component.html',
  styleUrl: './model-config-management.component.scss'
})
export class ModelConfigManagementComponent implements OnInit {
  displayedColumns: string[] = ['provider_name', 'model_name', 'api_endpoint', 'is_active', 'priority', 'actions'];
  dataSource = new MatTableDataSource<ModelConfig>();
  searchKeyword: string = '';

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  private aiService = inject(AIService);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  ngOnInit(): void {
    this.loadModelConfigs();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  loadModelConfigs(): void {
    this.aiService.getModelConfigs().subscribe({
      next: (response) => {
        this.dataSource.data = response.data?.configs || [];
      },
      error: (error) => {
        this.notificationService.error('加载模型配置失败');
      }
    });
  }

  applyFilter(): void {
    this.dataSource.filter = this.searchKeyword.trim().toLowerCase();
  }

  createModelConfig(): void {
    this.router.navigate(['/ai/model-config/create']);
  }

  editModelConfig(configId: string): void {
    this.router.navigate(['/ai/model-config/edit', configId]);
  }

  deleteModelConfig(configId: string): void {
    if (confirm('确定要删除这个模型配置吗？')) {
      this.aiService.deleteModelConfig(configId).subscribe({
        next: () => {
          this.notificationService.success('删除成功');
          this.loadModelConfigs();
        },
        error: (error) => {
          this.notificationService.error('删除失败');
        }
      });
    }
  }
}
