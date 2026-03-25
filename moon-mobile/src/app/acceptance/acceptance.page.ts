import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { IonHeader, IonToolbar, IonTitle, IonContent, IonList, IonItem, IonBadge, IonButton } from '@ionic/angular/standalone';

interface AcceptanceRecord {
  id: string;
  projectName: string;
  date: string;
  status: string;
  expertName: string;
  acceptanceResult: string;
  remarks: string;
  expanded?: boolean;
}

@Component({
  selector: 'app-acceptance',
  templateUrl: './acceptance.page.html',
  styleUrls: ['./acceptance.page.scss'],
  imports: [
    CommonModule,
    IonHeader, IonToolbar, IonTitle, IonContent,
    IonList, IonItem,
    IonBadge, IonButton
  ],
})
export class AcceptancePage {
  acceptanceRecords: AcceptanceRecord[] = [
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

  toggleExpand(record: AcceptanceRecord): void {
    record.expanded = !record.expanded;
  }

  approveAcceptance(record: AcceptanceRecord): void {
    record.acceptanceResult = '通过';
    record.status = '已完成';
    record.expanded = false;
    console.log('验收通过:', record);
  }

  rejectAcceptance(record: AcceptanceRecord): void {
    record.acceptanceResult = '不通过';
    record.status = '已完成';
    record.expanded = false;
    console.log('验收不通过:', record);
  }

  viewDetails(record: AcceptanceRecord): void {
    console.log('View details for:', record);
  }
}
