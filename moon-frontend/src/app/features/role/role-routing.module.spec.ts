import { roleRoutes } from './role-routing.module';

describe('RoleRoutingModule', () => {
  it('should have routes defined', () => {
    expect(roleRoutes).toBeDefined();
    expect(Array.isArray(roleRoutes)).toBe(true);
    expect(roleRoutes.length).toBeGreaterThan(0);
  });

  it('should have a route for the management page', () => {
    const managementRoute = roleRoutes.find(route => route.path === '');
    expect(managementRoute).toBeDefined();
    expect(managementRoute?.component).toBeDefined();
  });

  it('should have a route for the detail page', () => {
    const detailRoute = roleRoutes.find(route => route.path === 'detail/:id');
    expect(detailRoute).toBeDefined();
    expect(detailRoute?.component).toBeDefined();
  });

  it('should have a route for the create page', () => {
    const createRoute = roleRoutes.find(route => route.path === 'create');
    expect(createRoute).toBeDefined();
    expect(createRoute?.component).toBeDefined();
  });

  it('should have a route for the edit page', () => {
    const editRoute = roleRoutes.find(route => route.path === 'edit/:id');
    expect(editRoute).toBeDefined();
    expect(editRoute?.component).toBeDefined();
  });
});