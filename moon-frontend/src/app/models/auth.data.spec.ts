import { LoginCredentials, AuthResponse, RegisterData } from './auth.data';

describe('Auth Data Models', () => {
  describe('LoginCredentials interface', () => {
    it('should create a valid LoginCredentials object', () => {
      const loginCredentials: LoginCredentials = {
        email: 'test@example.com',
        password: 'password123'
      };

      expect(loginCredentials.email).toBe('test@example.com');
      expect(loginCredentials.password).toBe('password123');
    });

    it('should handle optional fields', () => {
      const minimalLoginCredentials: LoginCredentials = {
        email: 'test@example.com',
        password: ''
      };

      expect(minimalLoginCredentials).toBeTruthy();
      expect(typeof minimalLoginCredentials.email).toBe('string');
      expect(typeof minimalLoginCredentials.password).toBe('string');
    });
  });

  describe('AuthResponse interface', () => {
    it('should create a valid AuthResponse object', () => {
      const authResponse: AuthResponse = {
        token: 'mock-token-123',
        user: {
          UserID: '1',
          Username: 'testuser',
          Organization: 'Test Org',
          Role: ['Admin', 'User']
        }
      };

      expect(authResponse.token).toBe('mock-token-123');
      expect(authResponse.user.UserID).toBe('1');
      expect(authResponse.user.Username).toBe('testuser');
      expect(authResponse.user.Organization).toBe('Test Org');
      expect(authResponse.user.Role).toEqual(['Admin', 'User']);
    });

    it('should handle optional fields', () => {
      const minimalAuthResponse: AuthResponse = {
        token: 'mock-token-123',
        user: {
          UserID: '1',
          Username: 'testuser',
          Organization: '',
          Role: []
        }
      };

      expect(minimalAuthResponse).toBeTruthy();
      expect(typeof minimalAuthResponse.token).toBe('string');
      expect(typeof minimalAuthResponse.user.UserID).toBe('string');
      expect(typeof minimalAuthResponse.user.Username).toBe('string');
      expect(typeof minimalAuthResponse.user.Organization).toBe('string');
      expect(Array.isArray(minimalAuthResponse.user.Role)).toBe(true);
    });
  });

  describe('RegisterData interface', () => {
    it('should create a valid RegisterData object', () => {
      const registerData: RegisterData = {
        nickname: 'testuser',
        fullName: 'Test User',
        email: 'test@example.com',
        password: 'password123'
      };

      expect(registerData.nickname).toBe('testuser');
      expect(registerData.fullName).toBe('Test User');
      expect(registerData.email).toBe('test@example.com');
      expect(registerData.password).toBe('password123');
    });

    it('should handle optional fields', () => {
      const minimalRegisterData: RegisterData = {
        nickname: 'testuser',
        fullName: '',
        email: 'test@example.com',
        password: 'password123'
      };

      expect(minimalRegisterData).toBeTruthy();
      expect(typeof minimalRegisterData.nickname).toBe('string');
      expect(typeof minimalRegisterData.fullName).toBe('string');
      expect(typeof minimalRegisterData.email).toBe('string');
      expect(typeof minimalRegisterData.password).toBe('string');
    });
  });
});