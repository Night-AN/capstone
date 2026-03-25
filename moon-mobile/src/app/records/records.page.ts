import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IonHeader, IonToolbar, IonTitle, IonContent, IonList, IonItem, IonBadge, IonButton } from '@ionic/angular/standalone';

interface ProcurementRecord {
  id: string;
  projectName: string;
  date: string;
  status: string;
  supplier?: string;
  amount?: number;
  expertName?: string;
  acceptanceResult?: string;
  remarks: string;
  expanded?: boolean;
}

@Component({
  selector: 'app-records',
  templateUrl: './records.page.html',
  styleUrls: ['./records.page.scss'],
  imports: [
    CommonModule,
    IonHeader, IonToolbar, IonTitle, IonContent,
    IonList, IonItem,
    IonBadge, IonButton
  ],
})
export class RecordsPage {
  activeTab = 'procurement';

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

  acceptanceRecords: ProcurementRecord[] = [
    {
      id: '2',
      projectName: '网络设备采购',
      date: '2024-03-15',
      status: '进行中',
      expertName: '李四',
      acceptanceResult: '待确认',
      remarks: '正在进行设备测试',
      expanded: false
    },
    {
      id: '1',
      projectName: '服务器采购项目',
      date: '2024-03-20',
      status: '已完成',
      expertName: '张三',
      acceptanceResult: '通过',
      remarks: '设备运行正常，符合采购要求',
      expanded: false
    },
    {
      id: '3',
      projectName: '办公设备采购',
      date: '2024-03-10',
      status: '已完成',
      expertName: '王五',
      acceptanceResult: '通过',
      remarks: '所有设备已验收合格',
      expanded: false
    }
  ];

  constructor() {}

  getStatusColor(status: string, acceptanceResult?: string): string {
    if (status === '已完成') {
      return acceptanceResult === '不通过' ? 'danger' : 'success';
    } else if (status === '进行中') {
      return 'warning';
    } else {
      return 'primary';
    }
  }

  toggleExpand(record: ProcurementRecord): void {
    record.expanded = !record.expanded;
  }

  approveAcceptance(record: ProcurementRecord): void {
    record.acceptanceResult = '通过';
    record.status = '已完成';
    record.expanded = false;
    console.log('验收通过:', record);
  }

  rejectAcceptance(record: ProcurementRecord): void {
    record.acceptanceResult = '不通过';
    record.status = '已完成';
    record.expanded = false;
    console.log('验收不通过:', record);
  }

  viewDetails(record: ProcurementRecord): void {
    console.log('View details for:', record);
  }

  setActiveTab(tab: string): void {
    this.activeTab = tab;
  }
}
