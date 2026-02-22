import { organizationRoutes } from './organization-routing.module';

describe('OrganizationRoutingModule', () => {
  it('should have routes defined', () => {
    expect(organizationRoutes).toBeDefined();
    expect(Array.isArray(organizationRoutes)).toBe(true);
    expect(organizationRoutes.length).toBeGreaterThan(0);
  });

  it('should have a route for the management page', () => {
    const managementRoute = organizationRoutes.find(route => route.path === '');
    expect(managementRoute).toBeDefined();
    expect(managementRoute?.component).toBeDefined();
  });

  it('should have a route for the detail page', () => {
    const detailRoute = organizationRoutes.find(route => route.path === 'detail/:id');
    expect(detailRoute).toBeDefined();
    expect(detailRoute?.component).toBeDefined();
  });

  it('should have a route for the create page', () => {
    const createRoute = organizationRoutes.find(route => route.path === 'create');
    expect(createRoute).toBeDefined();
    expect(createRoute?.component).toBeDefined();
  });

  it('should have a route for the edit page', () => {
    const editRoute = organizationRoutes.find(route => route.path === 'edit/:id');
    expect(editRoute).toBeDefined();
    expect(editRoute?.component).toBeDefined();
  });
});