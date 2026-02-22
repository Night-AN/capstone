import { Permission, PermissionCreateRequest, PermissionUpdateRequest } from './permission.model';

describe('Permission Models', () => {
  describe('Permission interface', () => {
    it('should create a valid Permission object', () => {
      const permission: Permission = {
        permission_id: '1',
        name: 'Test Permission',
        description: 'Test Permission Description',
        sensitive_flag: false,
        created_at: '2023-01-01T00:00:00Z',
        updated_at: '2023-01-01T00:00:00Z'
      };

      expect(permission.permission_id).toBe('1');
      expect(permission.name).toBe('Test Permission');
      expect(permission.description).toBe('Test Permission Description');
      expect(permission.sensitive_flag).toBe(false);
    });

    it('should handle optional fields', () => {
      const minimalPermission: Permission = {
        permission_id: '1',
        name: 'Test Permission',
        description: '',
        sensitive_flag: false,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      };

      expect(minimalPermission).toBeTruthy();
      expect(typeof minimalPermission.permission_id).toBe('string');
      expect(typeof minimalPermission.name).toBe('string');
      expect(typeof minimalPermission.description).toBe('string');
      expect(typeof minimalPermission.sensitive_flag).toBe('boolean');
    });
  });

  describe('PermissionCreateRequest interface', () => {
    it('should create a valid PermissionCreateRequest object', () => {
      const permissionCreateRequest: PermissionCreateRequest = {
        name: 'Test Permission',
        description: 'Test Permission Description'
      };

      expect(permissionCreateRequest.name).toBe('Test Permission');
      expect(permissionCreateRequest.description).toBe('Test Permission Description');
    });

    it('should handle optional fields', () => {
      const minimalPermissionCreateRequest: PermissionCreateRequest = {
        name: 'Test Permission',
        description: ''
      };

      expect(minimalPermissionCreateRequest).toBeTruthy();
      expect(typeof minimalPermissionCreateRequest.name).toBe('string');
      expect(typeof minimalPermissionCreateRequest.description).toBe('string');
    });
  });

  describe('PermissionUpdateRequest interface', () => {
    it('should create a valid PermissionUpdateRequest object', () => {
      const permissionUpdateRequest: PermissionUpdateRequest = {
        permission_id: '1',
        name: 'Updated Test Permission',
        description: 'Updated Test Permission Description'
      };

      expect(permissionUpdateRequest.permission_id).toBe('1');
      expect(permissionUpdateRequest.name).toBe('Updated Test Permission');
      expect(permissionUpdateRequest.description).toBe('Updated Test Permission Description');
    });

    it('should handle optional fields', () => {
      const minimalPermissionUpdateRequest: PermissionUpdateRequest = {
        permission_id: '1',
        name: 'Updated Test Permission',
        description: ''
      };

      expect(minimalPermissionUpdateRequest).toBeTruthy();
      expect(typeof minimalPermissionUpdateRequest.permission_id).toBe('string');
      expect(typeof minimalPermissionUpdateRequest.name).toBe('string');
      expect(typeof minimalPermissionUpdateRequest.description).toBe('string');
    });
  });
});