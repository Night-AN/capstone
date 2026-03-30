import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzStatisticModule } from 'ng-zorro-antd/statistic';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzGridModule } from 'ng-zorro-antd/grid';
import { NzProgressModule } from 'ng-zorro-antd/progress';
import { NzTableModule } from 'ng-zorro-antd/table';
import { NzTagModule } from 'ng-zorro-antd/tag';
import { NzButtonModule } from 'ng-zorro-antd/button';

@Component({
  selector: 'app-dashboard',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    NzCardModule,
    NzStatisticModule,
    NzIconModule,
    NzGridModule,
    NzProgressModule,
    NzTableModule,
    NzTagModule,
    NzButtonModule
  ],
  templateUrl: './dashboard.html',
  styleUrls: ['./dashboard.scss']
})
export class DashboardComponent implements OnInit {
  // 资产统计数据
  assetStatistics = {
    totalAssets: 128,
    totalValue: 5280000,
    depreciationRate: 32.5,
    highValueAssets: 23
  };

  // 资产类型分布
  assetTypeDistribution = [
    { name: '电子设备', value: 45, color: '#1890ff' },
    { name: '办公家具', value: 25, color: '#52c41a' },
    { name: '交通工具', value: 15, color: '#faad14' },
    { name: '房屋建筑', value: 10, color: '#f5222d' },
    { name: '其他资产', value: 5, color: '#722ed1' }
  ];

  // 资产状态分布
  assetStatusDistribution = [
    { name: '正常使用', value: 85, color: '#52c41a' },
    { name: '待维修', value: 10, color: '#faad14' },
    { name: '闲置', value: 3, color: '#1890ff' },
    { name: '已报废', value: 2, color: '#f5222d' }
  ];

  // 最近添加的资产
  recentAssets = [
    {
      id: '1',
      name: '笔记本电脑',
      type: '电子设备',
      value: 8999,
      purchaseDate: '2026-03-25',
      status: '正常使用'
    },
    {
      id: '2',
      name: '办公桌椅',
      type: '办公家具',
      value: 2400,
      purchaseDate: '2026-03-24',
      status: '正常使用'
    },
    {
      id: '3',
      name: '打印机',
      type: '电子设备',
      value: 3200,
      purchaseDate: '2026-03-23',
      status: '待维修'
    },
    {
      id: '4',
      name: '投影仪',
      type: '电子设备',
      value: 5600,
      purchaseDate: '2026-03-22',
      status: '正常使用'
    },
    {
      id: '5',
      name: '会议桌',
      type: '办公家具',
      value: 4800,
      purchaseDate: '2026-03-21',
      status: '正常使用'
    }
  ];

  // 资产价值趋势
  assetValueTrend = [
    { month: '1月', value: 4800000 },
    { month: '2月', value: 4950000 },
    { month: '3月', value: 5280000 }
  ];

  constructor() { }

  ngOnInit(): void {
  }

  // 获取状态标签颜色
  getStatusColor(status: string): string {
    switch (status) {
      case '正常使用':
        return 'success';
      case '待维修':
        return 'warning';
      case '闲置':
        return 'primary';
      case '已报废':
        return 'danger';
      default:
        return 'default';
    }
  }

  // 格式化金额
  formatCurrency(value: number): string {
    return new Intl.NumberFormat('zh-CN', { style: 'currency', currency: 'CNY' }).format(value);
  }
}
