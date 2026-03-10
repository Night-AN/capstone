import { ComponentFixture, TestBed } from '@angular/core/testing';
import { PromptTemplateFormComponent } from './prompt-template-form.component';
import { FormBuilder, ReactiveFormsModule } from '@angular/forms';

describe('PromptTemplateFormComponent', () => {
  let component: PromptTemplateFormComponent;
  let fixture: ComponentFixture<PromptTemplateFormComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PromptTemplateFormComponent, ReactiveFormsModule],
      providers: [
        FormBuilder
      ]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PromptTemplateFormComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should have correct initial state for create mode', () => {
    expect(component.isEditMode).toBe(false);
    expect(component.templateId).toBe('');
  });

  it('should be invalid when required fields are empty', () => {
    expect(component.templateForm.valid).toBe(false);
  });

  it('should be valid when required fields are filled', () => {
    component.templateForm.patchValue({
      template_name: 'Test Template',
      template_type: 'asset_classification',
      template_content: 'This is a test template content'
    });
    expect(component.templateForm.valid).toBe(true);
  });
});
