import { Routes } from '@angular/router';
import { ModelConfigManagementComponent } from './pages/model-config-management/model-config-management.component';
import { ModelConfigFormComponent } from './pages/model-config-form/model-config-form.component';
import { PromptTemplateManagementComponent } from './pages/prompt-template-management/prompt-template-management.component';
import { PromptTemplateFormComponent } from './pages/prompt-template-form/prompt-template-form.component';
import { AILogsComponent } from './pages/ai-logs/ai-logs.component';
import { AIChatComponent } from './pages/ai-chat/ai-chat.component';

export const aiRoutes: Routes = [
  {
    path: '',
    redirectTo: 'chat',
    pathMatch: 'full'
  },
  {
    path: 'chat',
    component: AIChatComponent
  },
  {
    path: 'model-config',
    children: [
      {
        path: '',
        component: ModelConfigManagementComponent
      },
      {
        path: 'create',
        component: ModelConfigFormComponent
      },
      {
        path: 'edit/:id',
        component: ModelConfigFormComponent
      }
    ]
  },
  {
    path: 'prompt-template',
    children: [
      {
        path: '',
        component: PromptTemplateManagementComponent
      },
      {
        path: 'create',
        component: PromptTemplateFormComponent
      },
      {
        path: 'edit/:id',
        component: PromptTemplateFormComponent
      }
    ]
  },
  {
    path: 'logs',
    component: AILogsComponent
  }
];
