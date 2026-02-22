import { assetRoutes } from './asset-routing.module';

describe('AssetRoutingModule', () => {
  it('should have routes defined', () => {
    expect(assetRoutes).toBeDefined();
    expect(Array.isArray(assetRoutes)).toBe(true);
    expect(assetRoutes.length).toBeGreaterThan(0);
  });

  it('should have a route for the management page', () => {
    const managementRoute = assetRoutes.find(route => route.path === '');
    expect(managementRoute).toBeDefined();
    expect(managementRoute?.component).toBeDefined();
  });

  it('should have a route for the detail page', () => {
    const detailRoute = assetRoutes.find(route => route.path === 'detail/:id');
    expect(detailRoute).toBeDefined();
    expect(detailRoute?.component).toBeDefined();
  });

  it('should have a route for the create page', () => {
    const createRoute = assetRoutes.find(route => route.path === 'create');
    expect(createRoute).toBeDefined();
    expect(createRoute?.component).toBeDefined();
  });

  it('should have a route for the edit page', () => {
    const editRoute = assetRoutes.find(route => route.path === 'edit/:id');
    expect(editRoute).toBeDefined();
    expect(editRoute?.component).toBeDefined();
  });
});