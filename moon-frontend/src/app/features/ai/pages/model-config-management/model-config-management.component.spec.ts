import { ComponentFixture, TestBed } from '@angular/core/testing';
import { ModelConfigManagementComponent } from './model-config-management.component';

describe('ModelConfigManagementComponent', () => {
  let component: ModelConfigManagementComponent;
  let fixture: ComponentFixture<ModelConfigManagementComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [ModelConfigManagementComponent]
    }).compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(ModelConfigManagementComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
