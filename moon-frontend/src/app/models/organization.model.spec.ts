import { Organization, OrganizationCreateRequest, OrganizationUpdateRequest } from './organization.model';

describe('Organization Models', () => {
  describe('Organization interface', () => {
    it('should create a valid Organization object', () => {
      const organization: Organization = {
        organization_id: '1',
        organization_name: 'Test Organization',
        organization_code: 'ORG-001',
        organization_description: 'Test Organization Description',
        organization_flag: 'active',
        created_at: '2023-01-01T00:00:00Z',
        updated_at: '2023-01-01T00:00:00Z'
      };

      expect(organization.organization_id).toBe('1');
      expect(organization.organization_name).toBe('Test Organization');
      expect(organization.organization_code).toBe('ORG-001');
      expect(organization.organization_description).toBe('Test Organization Description');
      expect(organization.organization_flag).toBe('active');
    });

    it('should handle optional fields', () => {
      const minimalOrganization: Organization = {
        organization_id: '1',
        organization_name: 'Test Organization',
        organization_code: 'ORG-001',
        organization_description: '',
        organization_flag: '',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      };

      expect(minimalOrganization).toBeTruthy();
      expect(typeof minimalOrganization.organization_id).toBe('string');
      expect(typeof minimalOrganization.organization_name).toBe('string');
      expect(typeof minimalOrganization.organization_code).toBe('string');
      expect(typeof minimalOrganization.organization_description).toBe('string');
      expect(typeof minimalOrganization.organization_flag).toBe('string');
    });
  });

  describe('OrganizationCreateRequest interface', () => {
    it('should create a valid OrganizationCreateRequest object', () => {
      const organizationCreateRequest: OrganizationCreateRequest = {
        organization_name: 'Test Organization',
        organization_code: 'ORG-001',
        organization_description: 'Test Organization Description',
        organization_flag: 'active'
      };

      expect(organizationCreateRequest.organization_name).toBe('Test Organization');
      expect(organizationCreateRequest.organization_code).toBe('ORG-001');
      expect(organizationCreateRequest.organization_description).toBe('Test Organization Description');
      expect(organizationCreateRequest.organization_flag).toBe('active');
    });

    it('should handle optional fields', () => {
      const minimalOrganizationCreateRequest: OrganizationCreateRequest = {
        organization_name: 'Test Organization',
        organization_code: 'ORG-001',
        organization_description: '',
        organization_flag: ''
      };

      expect(minimalOrganizationCreateRequest).toBeTruthy();
      expect(typeof minimalOrganizationCreateRequest.organization_name).toBe('string');
      expect(typeof minimalOrganizationCreateRequest.organization_code).toBe('string');
      expect(typeof minimalOrganizationCreateRequest.organization_description).toBe('string');
      expect(typeof minimalOrganizationCreateRequest.organization_flag).toBe('string');
    });
  });

  describe('OrganizationUpdateRequest interface', () => {
    it('should create a valid OrganizationUpdateRequest object', () => {
      const organizationUpdateRequest: OrganizationUpdateRequest = {
        organization_id: '1',
        organization_name: 'Updated Test Organization',
        organization_code: 'ORG-001',
        organization_description: 'Updated Test Organization Description',
        organization_flag: 'active'
      };

      expect(organizationUpdateRequest.organization_id).toBe('1');
      expect(organizationUpdateRequest.organization_name).toBe('Updated Test Organization');
      expect(organizationUpdateRequest.organization_code).toBe('ORG-001');
      expect(organizationUpdateRequest.organization_description).toBe('Updated Test Organization Description');
      expect(organizationUpdateRequest.organization_flag).toBe('active');
    });

    it('should handle optional fields', () => {
      const minimalOrganizationUpdateRequest: OrganizationUpdateRequest = {
        organization_id: '1',
        organization_name: 'Updated Test Organization',
        organization_code: 'ORG-001',
        organization_description: '',
        organization_flag: ''
      };

      expect(minimalOrganizationUpdateRequest).toBeTruthy();
      expect(typeof minimalOrganizationUpdateRequest.organization_id).toBe('string');
      expect(typeof minimalOrganizationUpdateRequest.organization_name).toBe('string');
      expect(typeof minimalOrganizationUpdateRequest.organization_code).toBe('string');
      expect(typeof minimalOrganizationUpdateRequest.organization_description).toBe('string');
      expect(typeof minimalOrganizationUpdateRequest.organization_flag).toBe('string');
    });
  });
});