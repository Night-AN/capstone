import { User, UserCreateRequest, UserUpdateRequest } from './user.model';

describe('User Models', () => {
  describe('User interface', () => {
    it('should create a valid User object', () => {
      const user: User = {
        user_id: '1',
        nickname: 'testuser',
        full_name: 'Test User',
        email: 'test@example.com',
        created_at: '2023-01-01T00:00:00Z',
        updated_at: '2023-01-01T00:00:00Z'
      };

      expect(user.user_id).toBe('1');
      expect(user.nickname).toBe('testuser');
      expect(user.full_name).toBe('Test User');
      expect(user.email).toBe('test@example.com');
    });

    it('should handle optional fields', () => {
      const minimalUser: User = {
        user_id: '1',
        nickname: 'testuser',
        full_name: '',
        email: 'test@example.com',
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString()
      };

      expect(minimalUser).toBeTruthy();
      expect(typeof minimalUser.user_id).toBe('string');
      expect(typeof minimalUser.nickname).toBe('string');
      expect(typeof minimalUser.email).toBe('string');
    });
  });

  describe('UserCreateRequest interface', () => {
    it('should create a valid UserCreateRequest object', () => {
      const userCreateRequest: UserCreateRequest = {
        nickname: 'testuser',
        full_name: 'Test User',
        email: 'test@example.com',
        password: 'password123'
      };

      expect(userCreateRequest.nickname).toBe('testuser');
      expect(userCreateRequest.full_name).toBe('Test User');
      expect(userCreateRequest.email).toBe('test@example.com');
      expect(userCreateRequest.password).toBe('password123');
    });

    it('should handle optional fields', () => {
      const minimalUserCreateRequest: UserCreateRequest = {
        nickname: 'testuser',
        full_name: '',
        email: 'test@example.com',
        password: 'password123'
      };

      expect(minimalUserCreateRequest).toBeTruthy();
      expect(typeof minimalUserCreateRequest.nickname).toBe('string');
      expect(typeof minimalUserCreateRequest.email).toBe('string');
      expect(typeof minimalUserCreateRequest.password).toBe('string');
    });
  });

  describe('UserUpdateRequest interface', () => {
    it('should create a valid UserUpdateRequest object', () => {
      const userUpdateRequest: UserUpdateRequest = {
        user_id: '1',
        nickname: 'updateduser',
        full_name: 'Updated Test User',
        email: 'test@example.com'
      };

      expect(userUpdateRequest.user_id).toBe('1');
      expect(userUpdateRequest.nickname).toBe('updateduser');
      expect(userUpdateRequest.full_name).toBe('Updated Test User');
      expect(userUpdateRequest.email).toBe('test@example.com');
    });

    it('should handle optional fields', () => {
      const minimalUserUpdateRequest: UserUpdateRequest = {
        user_id: '1',
        nickname: 'updateduser',
        full_name: '',
        email: 'test@example.com'
      };

      expect(minimalUserUpdateRequest).toBeTruthy();
      expect(typeof minimalUserUpdateRequest.user_id).toBe('string');
      expect(typeof minimalUserUpdateRequest.nickname).toBe('string');
      expect(typeof minimalUserUpdateRequest.email).toBe('string');
    });
  });
});