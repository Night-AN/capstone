import { ComponentFixture, TestBed } from '@angular/core/testing';
import { PromptTemplateManagementComponent } from './prompt-template-management.component';

describe('PromptTemplateManagementComponent', () => {
  let component: PromptTemplateManagementComponent;
  let fixture: ComponentFixture<PromptTemplateManagementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [PromptTemplateManagementComponent]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(PromptTemplateManagementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
