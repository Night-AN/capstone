import { organizationRoutes } from './organization-routing.module';

describe('OrganizationModule', () => {
  it('should have routes defined', () => {
    expect(organizationRoutes).toBeDefined();
    expect(Array.isArray(organizationRoutes)).toBe(true);
  });

  it('should have at least one route', () => {
    expect(organizationRoutes.length).toBeGreaterThan(0);
  });
});