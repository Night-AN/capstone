import { assetRoutes } from './asset-routing.module';

describe('AssetModule', () => {
  it('should have routes defined', () => {
    expect(assetRoutes).toBeDefined();
    expect(Array.isArray(assetRoutes)).toBe(true);
  });

  it('should have at least one route', () => {
    expect(assetRoutes.length).toBeGreaterThan(0);
  });
});