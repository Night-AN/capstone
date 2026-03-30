import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SupplyChainRiskAnalysis } from './supply-chain-risk-analysis';

describe('SupplyChainRiskAnalysis', () => {
  let component: SupplyChainRiskAnalysis;
  let fixture: ComponentFixture<SupplyChainRiskAnalysis>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [SupplyChainRiskAnalysis]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SupplyChainRiskAnalysis);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
