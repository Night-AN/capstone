import { roleRoutes } from './role-routing.module';

describe('RoleModule', () => {
  it('should have routes defined', () => {
    expect(roleRoutes).toBeDefined();
    expect(Array.isArray(roleRoutes)).toBe(true);
  });

  it('should have at least one route', () => {
    expect(roleRoutes.length).toBeGreaterThan(0);
  });
});