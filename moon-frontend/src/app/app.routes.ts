import { Routes } from '@angular/router';
import { LoginPage } from '@pages/login-page/login-page';
import { BackendLayout } from '@layout/backend-layout/backend-layout';
import { UserPage } from '@pages/user-page/user-page';
import { OrganizationPage } from '@pages/organization-page/organization-page';
import { RolePage } from '@pages/role-page/role-page';
import { PermissionPage } from '@pages/permission-page/permission-page';
import { ProcurementAcceptancePage } from '@pages/procurement-acceptance-page/procurement-acceptance-page';
import { ProcurementExpertPage } from '@pages/procurement-expert-page/procurement-expert-page';
import { ProcurementImplementationPage } from '@pages/procurement-implementation-page/procurement-implementation-page';
import { ProcurementPlanPage } from '@pages/procurement-plan-page/procurement-plan-page';
import { ProcurementPlanTypePage } from '@pages/procurement-plan-type-page/procurement-plan-type-page';
import { ProcurementReviewPage } from '@pages/procurement-review-page/procurement-review-page';
import { LLMModelPage } from '@pages/llm-model-page/llm-model-page';
import { LLMProcurementAnalysisPage } from '@pages/llm-procurement-analysis-page/llm-procurement-analysis-page';
import { LLMTokenUsagePage } from '@pages/llm-token-usage-page/llm-token-usage-page';
import { ProcurementFraudRiskPage } from '@pages/procurement-fraud-risk-page/procurement-fraud-risk-page';
import { AssetsPage } from '@pages/assets-page/assets-page';
import { AssetTypePage } from '@pages/asset-type-page/asset-type-page';
import { AssetCategoryPage } from '@pages/asset-category-page/asset-category-page';
import { FilePageComponent } from '@pages/file-page/file-page';
import { FinancialExportComponent } from '@pages/financial-export/financial-export';
import { SupplyChainRiskAnalysisComponent } from '@pages/supply-chain-risk-analysis/supply-chain-risk-analysis';
import { DashboardComponent } from '@pages/dashboard/dashboard';

export const routes: Routes = [
  { path: '', redirectTo: 'login', pathMatch: 'full' },
  { path: 'login', component: LoginPage },
  { path: 'backend'
    , component: BackendLayout
    , children: [
      { path: '', component: DashboardComponent },
      { path: 'dashboard', component: DashboardComponent },
      { path: 'user-page', component: UserPage },
      { path: 'organization-page', component: OrganizationPage },
      { path: 'role-page', component: RolePage },
      { path: 'permission-page', component: PermissionPage },
      { path: 'file-page', component: FilePageComponent },
      { path: 'procurement-acceptance-page', component: ProcurementAcceptancePage },
      { path: 'procurement-expert-page', component: ProcurementExpertPage },
      { path: 'procurement-implementation-page', component: ProcurementImplementationPage },
      { path: 'procurement-plan-page', component: ProcurementPlanPage },
      { path: 'procurement-plan-type-page', component: ProcurementPlanTypePage },
      { path: 'procurement-review-page', component: ProcurementReviewPage },
      { path: 'llm-model-page', component: LLMModelPage },
      { path: 'llm-procurement-analysis-page', component: LLMProcurementAnalysisPage },
      { path: 'llm-token-usage-page', component: LLMTokenUsagePage },
      { path: 'procurement-fraud-risk-page', component: ProcurementFraudRiskPage },
      { path: 'assets-page', component: AssetsPage },
      { path: 'asset-type-page', component: AssetTypePage },
      { path: 'asset-category-page', component: AssetCategoryPage },
      { path: 'financial-export-page', component: FinancialExportComponent },
      { path: 'supply-chain-risk-analysis-page', component: SupplyChainRiskAnalysisComponent }
    ]
  },
];
