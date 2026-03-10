import { CommonModule } from '@angular/common';
import { Component, OnInit, ViewChild, inject } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatPaginator, MatPaginatorModule } from '@angular/material/paginator';
import { MatSort, MatSortModule } from '@angular/material/sort';
import { MatTableDataSource, MatTableModule } from '@angular/material/table';
import { AIService } from '@core/services/ai.service';
import { APICallLog } from '@models/ai.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-ai-logs',
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
  templateUrl: './ai-logs.component.html',
  styleUrl: './ai-logs.component.scss'
})
export class AILogsComponent implements OnInit {
  displayedColumns: string[] = ['call_type', 'prompt_tokens', 'completion_tokens', 'total_tokens', 'status_code', 'latency_ms', 'success', 'created_at'];
  dataSource = new MatTableDataSource<APICallLog>();
  searchKeyword: string = '';
  totalLogs: number = 0;
  pageSize: number = 10;
  pageIndex: number = 0;

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  private aiService = inject(AIService);
  private notificationService = inject(NotificationService);

  ngOnInit(): void {
    this.loadLogs();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  loadLogs(): void {
    this.aiService.getAPICallLogs({
      limit: this.pageSize,
      offset: this.pageIndex * this.pageSize,
      call_type: ''
    }).subscribe({
      next: (response) => {
        this.dataSource.data = response.data?.logs || [];
        this.totalLogs = response.data?.logs?.length || 0;
      },
      error: (error) => {
        this.notificationService.error('加载 AI 日志失败');
      }
    });
  }

  onPageChange(event: any): void {
    this.pageIndex = event.pageIndex;
    this.pageSize = event.pageSize;
    this.loadLogs();
  }

  applyFilter(): void {
    this.dataSource.filter = this.searchKeyword.trim().toLowerCase();
  }

  getStatusClass(statusCode: number): string {
    if (statusCode >= 200 && statusCode < 300) {
      return 'text-green-600';
    } else if (statusCode >= 400) {
      return 'text-red-600';
    }
    return 'text-yellow-600';
  }

  getSuccessClass(success: boolean): string {
    return success ? 'text-green-600' : 'text-red-600';
  }
}
