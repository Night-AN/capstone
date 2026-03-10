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
import { PromptTemplate } from '@models/ai.model';
import { NotificationService } from '@shared/service/notification/notification.service';

@Component({
  selector: 'app-prompt-template-management',
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
  templateUrl: './prompt-template-management.component.html',
  styleUrl: './prompt-template-management.component.scss'
})
export class PromptTemplateManagementComponent implements OnInit {
  displayedColumns: string[] = ['template_name', 'template_type', 'description', 'is_active', 'actions'];
  dataSource = new MatTableDataSource<PromptTemplate>();
  searchKeyword: string = '';

  @ViewChild(MatPaginator) paginator!: MatPaginator;
  @ViewChild(MatSort) sort!: MatSort;

  private aiService = inject(AIService);
  private router = inject(Router);
  private notificationService = inject(NotificationService);

  ngOnInit(): void {
    this.loadPromptTemplates();
  }

  ngAfterViewInit() {
    this.dataSource.paginator = this.paginator;
    this.dataSource.sort = this.sort;
  }

  loadPromptTemplates(): void {
    this.aiService.getPromptTemplates().subscribe({
      next: (response) => {
        this.dataSource.data = response.data?.templates || [];
      },
      error: (error) => {
        this.notificationService.error('加载提示模板失败');
      }
    });
  }

  applyFilter(): void {
    this.dataSource.filter = this.searchKeyword.trim().toLowerCase();
  }

  createPromptTemplate(): void {
    this.router.navigate(['/ai/prompt-template/create']);
  }

  editPromptTemplate(templateId: string): void {
    this.router.navigate(['/ai/prompt-template/edit', templateId]);
  }

  deletePromptTemplate(templateId: string): void {
    if (confirm('确定要删除这个提示模板吗？')) {
      this.aiService.deletePromptTemplate(templateId).subscribe({
        next: () => {
          this.notificationService.success('删除成功');
          this.loadPromptTemplates();
        },
        error: (error) => {
          this.notificationService.error('删除失败');
        }
      });
    }
  }
}
