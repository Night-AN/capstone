import { Role, RoleCreateRequest, RoleUpdateRequest } from './role.model';

describe('Role Model', () => {
  describe('Role', () => {
    it('should create a valid Role object', () => {
      const role: Role = {
        role_id: '1',
        role_name: 'Admin',
        description: 'Administrator role',
        sensitive_flag: false,
        created_at: '2023-01-01T00:00:00Z',
        updated_at: '2023-01-01T00:00:00Z'
      };

      expect(role.role_id).toBe('1');
      expect(role.role_name).toBe('Admin');
      expect(role.description).toBe('Administrator role');
      expect(role.sensitive_flag).toBe(false);
      expect(role.created_at).toBe('2023-01-01T00:00:00Z');
      expect(role.updated_at).toBe('2023-01-01T00:00:00Z');
    });

    it('should create a Role object with minimal required fields', () => {
      const minimalRole: Role = {
        role_id: '1',
        role_name: 'User',
        description: 'Regular user role',
        sensitive_flag: false,
        created_at: '2023-01-01T00:00:00Z',
        updated_at: '2023-01-01T00:00:00Z'
      };

      expect(minimalRole.role_id).toBe('1');
      expect(minimalRole.role_name).toBe('User');
      expect(minimalRole.description).toBe('Regular user role');
      expect(minimalRole.sensitive_flag).toBe(false);
      expect(minimalRole.created_at).toBe('2023-01-01T00:00:00Z');
      expect(minimalRole.updated_at).toBe('2023-01-01T00:00:00Z');
    });
  });

  describe('RoleCreateRequest', () => {
    it('should create a valid RoleCreateRequest object', () => {
      const roleCreateRequest: RoleCreateRequest = {
        role_name: 'Admin',
        description: 'Administrator role',
        sensitive_flag: false
      };

      expect(roleCreateRequest.role_name).toBe('Admin');
      expect(roleCreateRequest.description).toBe('Administrator role');
      expect(roleCreateRequest.sensitive_flag).toBe(false);
    });

    it('should create a RoleCreateRequest object with minimal required fields', () => {
      const minimalRoleCreateRequest: RoleCreateRequest = {
        role_name: 'User',
        description: 'Regular user role',
        sensitive_flag: false
      };

      expect(minimalRoleCreateRequest.role_name).toBe('User');
      expect(minimalRoleCreateRequest.description).toBe('Regular user role');
      expect(minimalRoleCreateRequest.sensitive_flag).toBe(false);
    });
  });

  describe('RoleUpdateRequest', () => {
    it('should create a valid RoleUpdateRequest object', () => {
      const roleUpdateRequest: RoleUpdateRequest = {
        role_id: '1',
        role_name: 'Admin',
        description: 'Updated Administrator role',
        sensitive_flag: false
      };

      expect(roleUpdateRequest.role_id).toBe('1');
      expect(roleUpdateRequest.role_name).toBe('Admin');
      expect(roleUpdateRequest.description).toBe('Updated Administrator role');
      expect(roleUpdateRequest.sensitive_flag).toBe(false);
    });

    it('should create a RoleUpdateRequest object with minimal required fields', () => {
      const minimalRoleUpdateRequest: RoleUpdateRequest = {
        role_id: '1',
        role_name: 'User',
        description: 'Updated regular user role',
        sensitive_flag: false
      };

      expect(minimalRoleUpdateRequest.role_id).toBe('1');
      expect(minimalRoleUpdateRequest.role_name).toBe('User');
      expect(minimalRoleUpdateRequest.description).toBe('Updated regular user role');
      expect(minimalRoleUpdateRequest.sensitive_flag).toBe(false);
    });
  });
});