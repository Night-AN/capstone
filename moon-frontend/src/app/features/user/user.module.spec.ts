import { userRoutes } from './user-routing.module';

describe('UserModule', () => {
  it('should have routes defined', () => {
    expect(userRoutes).toBeDefined();
    expect(Array.isArray(userRoutes)).toBe(true);
  });

  it('should have at least one route', () => {
    expect(userRoutes.length).toBeGreaterThan(0);
  });
});