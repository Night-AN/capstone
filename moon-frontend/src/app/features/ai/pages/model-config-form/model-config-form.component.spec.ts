import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ModelConfigFormComponent } from './model-config-form.component';
import { AIService } from '@core/services/ai.service';
import { NotificationService } from '@shared/service/notification/notification.service';
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';

describe('ModelConfigFormComponent', () => {
  let component: ModelConfigFormComponent;
  let fixture: ComponentFixture<ModelConfigFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ModelConfigFormComponent, ReactiveFormsModule],
      providers: [
        FormBuilder
      ]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ModelConfigFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have correct initial state for create mode', () => {
    expect(component.isEditMode).toBe(false);
    expect(component.configId).toBe('');
  });

  it('should be invalid when required fields are empty', () => {
    expect(component.configForm.valid).toBe(false);
  });

  it('should be valid when required fields are filled', () => {
    component.configForm.patchValue({
      provider_name: 'openai',
      model_name: 'gpt-4o',
      max_tokens: 4096,
      temperature: 0.7,
      timeout_seconds: 30,
      priority: 1
    });
    expect(component.configForm.valid).toBe(true);
  });
});
