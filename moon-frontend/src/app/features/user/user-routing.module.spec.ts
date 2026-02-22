import { userRoutes } from './user-routing.module';

describe('UserRoutingModule', () => {
  it('should have routes defined', () => {
    expect(userRoutes).toBeDefined();
    expect(Array.isArray(userRoutes)).toBe(true);
    expect(userRoutes.length).toBeGreaterThan(0);
  });

  it('should have a route for the management page', () => {
    const managementRoute = userRoutes.find(route => route.path === '');
    expect(managementRoute).toBeDefined();
    expect(managementRoute?.component).toBeDefined();
  });

  it('should have a route for the detail page', () => {
    const detailRoute = userRoutes.find(route => route.path === 'detail/:id');
    expect(detailRoute).toBeDefined();
    expect(detailRoute?.component).toBeDefined();
  });

  it('should have a route for the create page', () => {
    const createRoute = userRoutes.find(route => route.path === 'create');
    expect(createRoute).toBeDefined();
    expect(createRoute?.component).toBeDefined();
  });

  it('should have a route for the edit page', () => {
    const editRoute = userRoutes.find(route => route.path === 'edit/:id');
    expect(editRoute).toBeDefined();
    expect(editRoute?.component).toBeDefined();
  });
});