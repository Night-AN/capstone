import { gql } from 'apollo-angular';

// 导出资产台账
export const EXPORT_ASSET_LEDGER = gql`
  mutation ExportAssetLedger($format: String!) {
    exportAssetLedger(format: $format) {
      downloadUrl
      fileName
    }
  }
`;

// 导出会计凭证
export const EXPORT_ACCOUNTING_VOUCHERS = gql`
  mutation ExportAccountingVouchers($format: String!, $startDate: String, $endDate: String) {
    exportAccountingVouchers(format: $format, startDate: $startDate, endDate: $endDate) {
      downloadUrl
      fileName
    }
  }
`;

// 导出审计报告
export const EXPORT_AUDIT_REPORT = gql`
  mutation ExportAuditReport($format: String!, $startDate: String, $endDate: String) {
    exportAuditReport(format: $format, startDate: $startDate, endDate: $endDate) {
      downloadUrl
      fileName
    }
  }
`;

// 导出出纳流水
export const EXPORT_CASHIER_RECORDS = gql`
  mutation ExportCashierRecords($format: String!, $startDate: String, $endDate: String) {
    exportCashierRecords(format: $format, startDate: $startDate, endDate: $endDate) {
      downloadUrl
      fileName
    }
  }
`;
