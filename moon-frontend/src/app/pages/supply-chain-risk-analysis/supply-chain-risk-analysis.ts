import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { NzCardModule } from 'ng-zorro-antd/card';
import { NzFormModule } from 'ng-zorro-antd/form';
import { NzSelectModule } from 'ng-zorro-antd/select';
import { NzDatePickerModule } from 'ng-zorro-antd/date-picker';
import { NzButtonModule } from 'ng-zorro-antd/button';
import { NzIconModule } from 'ng-zorro-antd/icon';
import { NzMessageService } from 'ng-zorro-antd/message';

interface RiskDetail {
  name: string;
  level: 'high' | 'medium' | 'low';
  description: string;
}

interface AnalysisResult {
  securityScore: number;
  financialScore: number;
  operationalScore: number;
  totalScore: number;
  riskDetails: RiskDetail[];
  suggestions: string[];
}

@Component({
  selector: 'app-supply-chain-risk-analysis',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    NzCardModule,
    NzFormModule,
    NzSelectModule,
    NzDatePickerModule,
    NzButtonModule,
    NzIconModule
  ],
  templateUrl: './supply-chain-risk-analysis.html',
  styleUrls: ['./supply-chain-risk-analysis.scss']
})
export class SupplyChainRiskAnalysisComponent {
  // 加载状态
  loading: boolean = false;
  
  // 分析结果
  analysisResult: AnalysisResult | null = null;
  
  // 分析表单
  analysisForm: FormGroup;

  constructor(
    private fb: FormBuilder,
    private message: NzMessageService
  ) {
    this.analysisForm = this.fb.group({
      analysisType: ['security', Validators.required],
      analysisPeriod: ['monthly', Validators.required],
      startDate: [null, Validators.required],
      endDate: [null, Validators.required]
    });
  }

  // 分析风险
  analyzeRisk() {
    if (this.analysisForm.invalid) {
      this.message.error('请填写完整的分析参数');
      return;
    }

    this.loading = true;
    
    // 模拟API调用
    setTimeout(() => {
      this.analysisResult = {
        securityScore: 65,
        financialScore: 45,
        operationalScore: 30,
        totalScore: 47,
        riskDetails: [
          {
            name: '资产安全漏洞',
            level: 'high',
            description: '资产相关系统存在多个高危安全漏洞，可能导致数据泄露'
          },
          {
            name: '资产合规性',
            level: 'medium',
            description: '部分资产的合规性认证即将到期，需要及时更新'
          },
          {
            name: '资产依赖风险',
            level: 'medium',
            description: '部分资产依赖于过时的第三方库或服务'
          },
          {
            name: '资产配置风险',
            level: 'low',
            description: '部分资产的安全配置未达到最佳实践标准'
          }
        ],
        suggestions: [
          '立即修复资产相关系统的安全漏洞',
          '建立资产合规性定期检查机制',
          '更新过时的第三方库和服务依赖',
          '优化资产的安全配置，达到最佳实践标准',
          '建立资产风险定期评估机制'
        ]
      };
      this.message.success('风险分析完成');
      this.loading = false;
    }, 1500);
  }

  // 导出报告
  exportReport() {
    if (!this.analysisResult) {
      this.message.error('请先进行风险分析');
      return;
    }

    this.loading = true;
    
    // 模拟导出过程
    setTimeout(() => {
      this.message.success('报告导出成功');
      this.loading = false;
    }, 1000);
  }
}
