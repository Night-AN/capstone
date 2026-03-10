import { Component, Inject, OnInit, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSelectModule } from '@angular/material/select';
import { MatTableModule } from '@angular/material/table';
import { AssetService } from '@core/services/asset.service';
import { AIService } from '@core/services/ai.service';
import { PromptTemplate } from '@models/ai.model';

export interface BatchImportDialogData {
  assets: any[];
}

export interface BatchImportResult {
  total_count: number;
  success_count: number;
  failed_count: number;
  results: {
    asset_id: string;
    asset_name: string;
    asset_code: string;
    success: boolean;
    error?: string;
  }[];
  classifications: {
    asset_id: string;
    predicted_category: string;
    confidence: number;
    reasoning: string;
    success: boolean;
  }[];
}

@Component({
  selector: 'app-batch-import-dialog',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatButtonModule,
    MatCardModule,
    MatCheckboxModule,
    MatDialogModule,
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatProgressSpinnerModule,
    MatSelectModule,
    MatTableModule
  ],
  template: `
    <h2 mat-dialog-title>
      <mat-icon>upload_file</mat-icon>
      批量导入资产
    </h2>
    <mat-dialog-content>
      <div class="import-options">
        <mat-checkbox [(ngModel)]="enableAIClassification" [disabled]="importing()">
          <mat-icon>psychology</mat-icon>
          启用AI自动分类
        </mat-checkbox>
      </div>

      <div class="template-section" *ngIf="enableAIClassification()">
        <mat-form-field appearance="outline" class="full-width">
          <mat-label>选择提示词模板</mat-label>
          <mat-select [(ngModel)]="selectedPromptTemplateId" [disabled]="importing()">
            <mat-option [value]="null">默认模板</mat-option>
            <mat-option *ngFor="let template of promptTemplates()" [value]="template.template_id">
              {{ template.template_name }}
            </mat-option>
          </mat-select>
          <mat-hint *ngIf="selectedPromptTemplateId">
            {{ getSelectedTemplateDescription() }}
          </mat-hint>
        </mat-form-field>

        <button mat-button color="primary" (click)="createPromptTemplate()" class="create-template-btn">
          <mat-icon>add</mat-icon>
          创建新模板
        </button>
      </div>

      <div class="preview-section" *ngIf="previewAssets().length > 0">
        <h4>预览 (共 {{ previewAssets().length }} 条)</h4>
        <div class="table-container">
          <table mat-table [dataSource]="previewAssets().slice(0, 5)" class="preview-table">
            <ng-container matColumnDef="asset_name">
              <th mat-header-cell *matHeaderCellDef>资产名称</th>
              <td mat-cell *matCellDef="let element">{{ element.asset_name }}</td>
            </ng-container>

            <ng-container matColumnDef="asset_code">
              <th mat-header-cell *matHeaderCellDef>资产代码</th>
              <td mat-cell *matCellDef="let element">{{ element.asset_code }}</td>
            </ng-container>

            <ng-container matColumnDef="asset_type">
              <th mat-header-cell *matHeaderCellDef>资产类型</th>
              <td mat-cell *matCellDef="let element">{{ element.asset_type }}</td>
            </ng-container>

            <ng-container matColumnDef="ip_address">
              <th mat-header-cell *matHeaderCellDef>IP地址</th>
              <td mat-cell *matCellDef="let element">{{ element.ip_address }}</td>
            </ng-container>

            <tr mat-header-row *matHeaderRowDef="previewColumns"></tr>
            <tr mat-row *matRowDef="let row; columns: previewColumns;"></tr>
          </table>
        </div>
        <p class="more-hint" *ngIf="previewAssets().length > 5">
          还有 {{ previewAssets().length - 5 }} 条数据...
        </p>
      </div>

      <div *ngIf="importing()" class="importing">
        <mat-spinner diameter="40"></mat-spinner>
        <p>正在导入资产{{ enableAIClassification() ? '并进行AI分类...' : '...' }}</p>
        <p class="progress-text">{{ importProgress() }} / {{ totalCount() }}</p>
      </div>

      <div *ngIf="importResult()" class="result-section">
        <mat-card class="result-card">
          <mat-card-header>
            <mat-card-title>导入结果</mat-card-title>
          </mat-card-header>
          <mat-card-content>
            <div class="result-stats">
              <div class="stat success">
                <mat-icon>check_circle</mat-icon>
                <span class="count">{{ importResult()?.success_count }}</span>
                <span class="label">成功</span>
              </div>
              <div class="stat failed">
                <mat-icon>error</mat-icon>
                <span class="count">{{ importResult()?.failed_count }}</span>
                <span class="label">失败</span>
              </div>
              <div class="stat total">
                <mat-icon>inventory</mat-icon>
                <span class="count">{{ importResult()?.total_count }}</span>
                <span class="label">总计</span>
              </div>
            </div>

            <div *ngIf="importResult()?.classifications?.length" class="classifications">
              <h4>AI分类结果</h4>
              <div class="classification-list">
                <div *ngFor="let cls of importResult()?.classifications" class="classification-item">
                  <div class="cls-header">
                    <span class="asset-name">{{ getAssetName(cls.asset_id) }}</span>
                    <span class="category">{{ cls.predicted_category }}</span>
                    <span class="confidence">{{ (cls.confidence * 100).toFixed(1) }}%</span>
                  </div>
                  <p class="reasoning">{{ cls.reasoning }}</p>
                </div>
              </div>
            </div>

            <div *ngIf="failedResults().length" class="errors">
              <h4>失败记录</h4>
              <div *ngFor="let result of failedResults()" class="error-item">
                <strong>{{ result.asset_name }}</strong>: {{ result.error }}
              </div>
            </div>
          </mat-card-content>
        </mat-card>
      </div>
    </mat-dialog-content>
    <mat-dialog-actions align="end">
      <button mat-button (click)="onCancel()" [disabled]="importing()">取消</button>
      <button mat-raised-button color="primary" (click)="onImport()" 
              [disabled]="importing() || previewAssets().length === 0">
        <mat-icon>upload</mat-icon>
        {{ importing() ? '导入中...' : '开始导入' }}
      </button>
    </mat-dialog-actions>
  `,
  styles: [`
    .import-options {
      margin-bottom: 16px;
      display: flex;
      align-items: center;
      gap: 16px;
    }
    .template-section {
      margin-bottom: 16px;
      padding: 16px;
      background: #f5f5f5;
      border-radius: 8px;
    }
    .full-width {
      width: 100%;
    }
    .create-template-btn {
      margin-top: 8px;
    }
    .preview-section {
      margin-bottom: 16px;
    }
    .preview-section h4 {
      margin-bottom: 8px;
      color: #666;
    }
    .table-container {
      max-height: 200px;
      overflow: auto;
    }
    .preview-table {
      width: 100%;
    }
    .more-hint {
      color: #999;
      font-size: 12px;
      margin-top: 8px;
    }
    .importing {
      text-align: center;
      padding: 32px;
    }
    .importing mat-spinner {
      margin: 0 auto 16px;
    }
    .progress-text {
      color: #666;
      font-size: 14px;
    }
    .result-section {
      max-height: 400px;
      overflow: auto;
    }
    .result-card {
      margin-bottom: 16px;
    }
    .result-stats {
      display: flex;
      justify-content: space-around;
      padding: 16px 0;
    }
    .stat {
      text-align: center;
    }
    .stat mat-icon {
      font-size: 32px;
      width: 32px;
      height: 32px;
    }
    .stat.success mat-icon {
      color: #4caf50;
    }
    .stat.failed mat-icon {
      color: #f44336;
    }
    .stat.total mat-icon {
      color: #2196f3;
    }
    .stat .count {
      display: block;
      font-size: 24px;
      font-weight: bold;
    }
    .stat .label {
      font-size: 12px;
      color: #666;
    }
    .classifications {
      margin-top: 16px;
      padding-top: 16px;
      border-top: 1px solid #eee;
    }
    .classifications h4 {
      margin-bottom: 12px;
    }
    .classification-list {
      max-height: 200px;
      overflow: auto;
    }
    .classification-item {
      background: #f5f5f5;
      padding: 12px;
      margin-bottom: 8px;
      border-radius: 4px;
    }
    .cls-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      margin-bottom: 8px;
    }
    .asset-name {
      font-weight: 500;
    }
    .category {
      background: #e3f2fd;
      padding: 2px 8px;
      border-radius: 12px;
      font-size: 12px;
    }
    .confidence {
      color: #4caf50;
      font-weight: 500;
    }
    .reasoning {
      font-size: 12px;
      color: #666;
      margin: 0;
    }
    .errors {
      margin-top: 16px;
      padding-top: 16px;
      border-top: 1px solid #eee;
    }
    .errors h4 {
      color: #f44336;
      margin-bottom: 8px;
    }
    .error-item {
      color: #f44336;
      font-size: 12px;
      margin-bottom: 4px;
    }
  `]
})
export class BatchImportDialogComponent implements OnInit {
  previewAssets = signal<any[]>([]);
  enableAIClassification = signal<boolean>(false);
  selectedPromptTemplateId: string | null = null;
  promptTemplates = signal<PromptTemplate[]>([]);
  importing = signal<boolean>(false);
  importProgress = signal<number>(0);
  totalCount = signal<number>(0);
  importResult = signal<BatchImportResult | null>(null);
  failedResults = signal<any[]>([]);

  previewColumns = ['asset_name', 'asset_code', 'asset_type', 'ip_address'];

  constructor(
    public dialogRef: MatDialogRef<BatchImportDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: BatchImportDialogData,
    private assetService: AssetService,
    private aiService: AIService
  ) {
    this.previewAssets.set(data.assets || []);
    this.totalCount.set(data.assets?.length || 0);
  }

  ngOnInit(): void {
    this.loadPromptTemplates();
  }

  loadPromptTemplates(): void {
    this.aiService.getPromptTemplates().subscribe({
      next: (response) => {
        const templates = response.data?.templates || [];
        this.promptTemplates.set(templates.filter((t: PromptTemplate) => 
          t.template_type === 'asset_classification' && t.is_active
        ));
      },
      error: (error) => {
        console.error('Error loading prompt templates:', error);
      }
    });
  }

  getSelectedTemplateDescription(): string {
    if (!this.selectedPromptTemplateId) return '';
    const template = this.promptTemplates().find(t => t.template_id === this.selectedPromptTemplateId);
    return template?.description || '';
  }

  createPromptTemplate(): void {
    this.dialogRef.close({ openPromptTemplateEditor: true });
  }

  getAssetName(assetId: string): string {
    const result = this.importResult()?.results?.find(r => r.asset_id === assetId);
    return result?.asset_name || assetId;
  }

  onImport(): void {
    if (this.previewAssets().length === 0) return;

    this.importing.set(true);
    this.importProgress.set(0);
    this.importResult.set(null);

    this.assetService.batchImportAssets(
      this.previewAssets(),
      this.enableAIClassification(),
      this.selectedPromptTemplateId || undefined
    ).subscribe({
      next: (result: BatchImportResult) => {
        this.importing.set(false);
        this.importResult.set(result);
        this.failedResults.set(
          result.results?.filter(r => !r.success) || []
        );
      },
      error: (error: any) => {
        this.importing.set(false);
        console.error('Batch import error:', error);
      }
    });
  }

  onCancel(): void {
    this.dialogRef.close();
  }
}
