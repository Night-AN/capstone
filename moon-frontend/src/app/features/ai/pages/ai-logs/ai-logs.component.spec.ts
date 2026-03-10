import { ComponentFixture, TestBed } from '@angular/core/testing';
import { AILogsComponent } from './ai-logs.component';

describe('AILogsComponent', () => {
  let component: AILogsComponent;
  let fixture: ComponentFixture<AILogsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AILogsComponent]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AILogsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('should return correct status class', () => {
    expect(component.getStatusClass(200)).toBe('text-green-600');
    expect(component.getStatusClass(400)).toBe('text-red-600');
    expect(component.getStatusClass(300)).toBe('text-yellow-600');
  });

  it('should return correct success class', () => {
    expect(component.getSuccessClass(true)).toBe('text-green-600');
    expect(component.getSuccessClass(false)).toBe('text-red-600');
  });
});
