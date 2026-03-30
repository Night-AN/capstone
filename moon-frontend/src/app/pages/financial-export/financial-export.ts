import { Component } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzDatePickerModule } from 'ng-zorro-antd/date-picker';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzMessageService } from 'ng-zorro-antd/message';
import { FinancialExportService } from '../../services/financial-export.service';

@Component({
  selector: 'app-financial-export',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    NzCardModule,
    NzFormModule,
    NzSelectModule,
    NzDatePickerModule,
    NzButtonModule,
    NzIconModule
  ],
  templateUrl: './financial-export.html',
  styleUrls: ['./financial-export.scss']
})
export class FinancialExportComponent {
  // 加载状态
  loading: boolean = false;

  // 资产台账导出
  assetFormat: string = 'excel';

  // 会计凭证导出
  voucherFormat: string = 'excel';
  voucherStartDate: Date | null = null;
  voucherEndDate: Date | null = null;

  // 审计报告导出
  auditFormat: string = 'pdf';
  auditStartDate: Date | null = null;
  auditEndDate: Date | null = null;

  // 出纳流水导出
  cashierFormat: string = 'excel';
  cashierStartDate: Date | null = null;
  cashierEndDate: Date | null = null;

  constructor(
    private financialExportService: FinancialExportService,
    private message: NzMessageService
  ) { }

  // 导出资产台账
  exportAssetLedger() {
    this.loading = true;
    this.financialExportService.exportAssetLedger(this.assetFormat).subscribe({
      next: (response: any) => {
        const { downloadUrl, fileName } = response.data.exportAssetLedger;
        this.downloadFile(downloadUrl, fileName);
        this.message.success('资产台账导出成功');
        this.loading = false;
      },
      error: (error) => {
        console.error('导出资产台账失败:', error);
        this.message.error('导出资产台账失败');
        this.loading = false;
      }
    });
  }

  // 导出会计凭证
  exportAccountingVouchers() {
    this.loading = true;
    const startDate = this.voucherStartDate ? this.formatDate(this.voucherStartDate) : undefined;
    const endDate = this.voucherEndDate ? this.formatDate(this.voucherEndDate) : undefined;

    this.financialExportService.exportAccountingVouchers(this.voucherFormat, startDate, endDate).subscribe({
      next: (response: any) => {
        const { downloadUrl, fileName } = response.data.exportAccountingVouchers;
        this.downloadFile(downloadUrl, fileName);
        this.message.success('会计凭证导出成功');
        this.loading = false;
      },
      error: (error) => {
        console.error('导出会计凭证失败:', error);
        this.message.error('导出会计凭证失败');
        this.loading = false;
      }
    });
  }

  // 导出审计报告
  exportAuditReport() {
    this.loading = true;
    const startDate = this.auditStartDate ? this.formatDate(this.auditStartDate) : undefined;
    const endDate = this.auditEndDate ? this.formatDate(this.auditEndDate) : undefined;

    this.financialExportService.exportAuditReport(this.auditFormat, startDate, endDate).subscribe({
      next: (response: any) => {
        const { downloadUrl, fileName } = response.data.exportAuditReport;
        this.downloadFile(downloadUrl, fileName);
        this.message.success('审计报告导出成功');
        this.loading = false;
      },
      error: (error) => {
        console.error('导出审计报告失败:', error);
        this.message.error('导出审计报告失败');
        this.loading = false;
      }
    });
  }

  // 导出出纳流水
  exportCashierRecords() {
    this.loading = true;
    const startDate = this.cashierStartDate ? this.formatDate(this.cashierStartDate) : undefined;
    const endDate = this.cashierEndDate ? this.formatDate(this.cashierEndDate) : undefined;

    this.financialExportService.exportCashierRecords(this.cashierFormat, startDate, endDate).subscribe({
      next: (response: any) => {
        const { downloadUrl, fileName } = response.data.exportCashierRecords;
        this.downloadFile(downloadUrl, fileName);
        this.message.success('出纳流水导出成功');
        this.loading = false;
      },
      error: (error) => {
        console.error('导出出纳流水失败:', error);
        this.message.error('导出出纳流水失败');
        this.loading = false;
      }
    });
  }

  // 下载文件
  downloadFile(url: string, fileName: string) {
    const link = document.createElement('a');
    link.href = url;
    link.download = fileName;
    link.click();
  }

  // 格式化日期
  formatDate(date: Date): string {
    return date.toISOString().split('T')[0];
  }
}
