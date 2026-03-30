import { ComponentFixture, TestBed } from '@angular/core/testing';

import { FinancialExport } from './financial-export';

describe('FinancialExport', () => {
  let component: FinancialExport;
  let fixture: ComponentFixture<FinancialExport>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [FinancialExport]
    })
    .compileComponents();

    fixture = TestBed.createComponent(FinancialExport);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
