import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import {
  EXPORT_ASSET_LEDGER,
  EXPORT_ACCOUNTING_VOUCHERS,
  EXPORT_AUDIT_REPORT,
  EXPORT_CASHIER_RECORDS
} from './financial-export.graphql';

@Injectable({
  providedIn: 'root'
})
export class FinancialExportService {

  constructor(private apollo: Apollo) { }

  // 导出资产台账
  exportAssetLedger(format: string): Observable<any> {
    return this.apollo.mutate({
      mutation: EXPORT_ASSET_LEDGER,
      variables: { format }
    });
  }

  // 导出会计凭证
  exportAccountingVouchers(format: string, startDate?: string, endDate?: string): Observable<any> {
    return this.apollo.mutate({
      mutation: EXPORT_ACCOUNTING_VOUCHERS,
      variables: { format, startDate, endDate }
    });
  }

  // 导出审计报告
  exportAuditReport(format: string, startDate?: string, endDate?: string): Observable<any> {
    return this.apollo.mutate({
      mutation: EXPORT_AUDIT_REPORT,
      variables: { format, startDate, endDate }
    });
  }

  // 导出出纳流水
  exportCashierRecords(format: string, startDate?: string, endDate?: string): Observable<any> {
    return this.apollo.mutate({
      mutation: EXPORT_CASHIER_RECORDS,
      variables: { format, startDate, endDate }
    });
  }
}
