import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IonHeader, IonToolbar, IonTitle, IonContent, IonList, IonItem, IonBadge, IonButton } from '@ionic/angular/standalone';

interface ProcurementRecord {
  id: string;
  projectName: string;
  date: string;
  status: string;
  supplier: string;
  amount: number;
  remarks: string;
  expanded?: boolean;
}

@Component({
  selector: 'app-procurement',
  templateUrl: './procurement.page.html',
  styleUrls: ['./procurement.page.scss'],
  imports: [
    CommonModule,
    IonHeader, IonToolbar, IonTitle, IonContent,
    IonList, IonItem,
    IonBadge, IonButton
  ],
})
export class ProcurementPage {
  procurementRecords: ProcurementRecord[] = [
    {
      id: '2',
      projectName: '网络设备采购',
      date: '2024-03-10',
      status: '进行中',
      supplier: '思科系统公司',
      amount: 85000,
      remarks: '采购网络交换机和路由器',
      expanded: false
    },
    {
      id: '1',
      projectName: '服务器采购项目',
      date: '2024-03-15',
      status: '已完成',
      supplier: '华为技术有限公司',
      amount: 120000,
      remarks: '采购2台服务器，用于数据中心扩容',
      expanded: false
    },
    {
      id: '3',
      projectName: '办公设备采购',
      date: '2024-03-05',
      status: '已完成',
      supplier: '联想集团',
      amount: 35000,
      remarks: '采购10台办公电脑和5台打印机',
      expanded: false
    }
  ];

  constructor() {}

  getStatusColor(status: string): string {
    switch (status) {
      case '已完成':
        return 'success';
      case '进行中':
        return 'warning';
      default:
        return 'primary';
    }
  }

  toggleExpand(record: ProcurementRecord): void {
    record.expanded = !record.expanded;
  }

  approveAcceptance(record: ProcurementRecord): void {
    record.status = '已完成';
    record.expanded = false;
    console.log('验收通过:', record);
  }

  rejectAcceptance(record: ProcurementRecord): void {
    record.status = '已完成';
    record.expanded = false;
    console.log('验收不通过:', record);
  }

  viewDetails(record: ProcurementRecord): void {
    console.log('View details for:', record);
  }
}
