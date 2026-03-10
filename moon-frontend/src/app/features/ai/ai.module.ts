import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { MatTableModule } from '@angular/material/table';
import { MatPaginatorModule } from '@angular/material/paginator';
import { MatSortModule } from '@angular/material/sort';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatDialogModule } from '@angular/material/dialog';
import { MatIconModule } from '@angular/material/icon';
import { MatSelectModule } from '@angular/material/select';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';

import { aiRoutes } from './ai-routing.module';

import { ModelConfigManagementComponent } from './pages/model-config-management/model-config-management.component';
import { ModelConfigFormComponent } from './pages/model-config-form/model-config-form.component';
import { PromptTemplateManagementComponent } from './pages/prompt-template-management/prompt-template-management.component';
import { PromptTemplateFormComponent } from './pages/prompt-template-form/prompt-template-form.component';
import { AILogsComponent } from './pages/ai-logs/ai-logs.component';

@NgModule({
  imports: [
    CommonModule,
    RouterModule.forChild(aiRoutes),
    FormsModule,
    ReactiveFormsModule,
    MatTableModule,
    MatPaginatorModule,
    MatSortModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
    MatDialogModule,
    MatIconModule,
    MatSelectModule,
    MatSlideToggleModule,
    MatProgressSpinnerModule,
    ModelConfigManagementComponent,
    ModelConfigFormComponent,
    PromptTemplateManagementComponent,
    PromptTemplateFormComponent,
    AILogsComponent
  ],
  providers: []
})
export class AIModule { }
