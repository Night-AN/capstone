import { Resource, ResourceCreateRequest, ResourceUpdateRequest } from './resource.model';

describe('Resource Models', () => {
  describe('Resource interface', () => {
    it('should create a valid Resource object', () => {
      const resource: Resource = {
        resource_id: '1',
        name: 'Test Resource',
        description: 'Test Resource Description',
        sensitive_flag: false,
        created_at: '2023-01-01T00:00:00Z',
        updated_at: '2023-01-01T00:00:00Z'
      };

      expect(resource.resource_id).toBe('1');
      expect(resource.name).toBe('Test Resource');
      expect(resource.description).toBe('Test Resource Description');
      expect(resource.sensitive_flag).toBe(false);
    });

    it('should handle optional fields', () => {
      const minimalResource: Resource = {
        resource_id: '1',
        name: 'Test Resource',
        description: '',
        sensitive_flag: false,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      };

      expect(minimalResource).toBeTruthy();
      expect(typeof minimalResource.resource_id).toBe('string');
      expect(typeof minimalResource.name).toBe('string');
      expect(typeof minimalResource.description).toBe('string');
      expect(typeof minimalResource.sensitive_flag).toBe('boolean');
    });
  });

  describe('ResourceCreateRequest interface', () => {
    it('should create a valid ResourceCreateRequest object', () => {
      const resourceCreateRequest: ResourceCreateRequest = {
        name: 'Test Resource',
        description: 'Test Resource Description'
      };

      expect(resourceCreateRequest.name).toBe('Test Resource');
      expect(resourceCreateRequest.description).toBe('Test Resource Description');
    });

    it('should handle optional fields', () => {
      const minimalResourceCreateRequest: ResourceCreateRequest = {
        name: 'Test Resource',
        description: ''
      };

      expect(minimalResourceCreateRequest).toBeTruthy();
      expect(typeof minimalResourceCreateRequest.name).toBe('string');
      expect(typeof minimalResourceCreateRequest.description).toBe('string');
    });
  });

  describe('ResourceUpdateRequest interface', () => {
    it('should create a valid ResourceUpdateRequest object', () => {
      const resourceUpdateRequest: ResourceUpdateRequest = {
        resource_id: '1',
        name: 'Updated Test Resource',
        description: 'Updated Test Resource Description'
      };

      expect(resourceUpdateRequest.resource_id).toBe('1');
      expect(resourceUpdateRequest.name).toBe('Updated Test Resource');
      expect(resourceUpdateRequest.description).toBe('Updated Test Resource Description');
    });

    it('should handle optional fields', () => {
      const minimalResourceUpdateRequest: ResourceUpdateRequest = {
        resource_id: '1',
        name: 'Updated Test Resource',
        description: ''
      };

      expect(minimalResourceUpdateRequest).toBeTruthy();
      expect(typeof minimalResourceUpdateRequest.resource_id).toBe('string');
      expect(typeof minimalResourceUpdateRequest.name).toBe('string');
      expect(typeof minimalResourceUpdateRequest.description).toBe('string');
    });
  });
});