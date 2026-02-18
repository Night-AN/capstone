import { describe, it, expect, beforeEach, afterEach, vi } from 'vitest';
import { TestBed } from '@angular/core/testing';
import { StorageService } from './storage.service';

describe('StorageService', () => {
  let service: StorageService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(StorageService);
    localStorage.clear();
    sessionStorage.clear();
  });

  afterEach(() => {
    localStorage.clear();
    sessionStorage.clear();
  });

  describe('localStorage', () => {
    it('should store data correctly', () => {
      const testData = { name: 'Zhang San', age: 25 };
      service.set('user', testData);

      const stored = localStorage.getItem('moon_user');
      expect(stored).toBe(JSON.stringify(testData));
    });

    it('should retrieve stored data', () => {
      const testData = { name: 'Zhang San', age: 25 };
      localStorage.setItem('moon_user', JSON.stringify(testData));

      const result = service.get('user');
      expect(result).toEqual(testData);
    });

    it('should return null for non-existent keys', () => {
      const result = service.get('nonexistent');
      expect(result).toBeNull();
    });

    it('should delete data by key', () => {
      localStorage.setItem('moon_user', JSON.stringify({ name: 'Zhang San' }));

      service.remove('user');
      expect(localStorage.getItem('moon_user')).toBeNull();
    });

    it('should clear only prefixed data', () => {
      localStorage.setItem('moon_user1', 'value1');
      localStorage.setItem('moon_user2', 'value2');
      localStorage.setItem('other_key', 'value3');

      service.clear();

      expect(localStorage.length).toBe(1);
      expect(localStorage.getItem('other_key')).toBe('value3');
      expect(localStorage.getItem('moon_user1')).toBeNull();
    });

    it('should handle storage errors gracefully', () => {
      const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {});
      const circular: any = { self: null };
      circular.self = circular;

      service.set('circular', circular);

      expect(consoleSpy).toHaveBeenCalledWith(
        expect.stringContaining('Storage set error: TypeError: Converting circular structure to JSON')
      );

      consoleSpy.mockRestore();
    });

    it('should handle JSON parse errors during retrieval', () => {
      const consoleSpy = vi.spyOn(console, 'error').mockImplementation(() => {});
      const circular: any = { self: null };
      circular.self = circular;

      service.set('circular', circular);

      expect(consoleSpy).toHaveBeenCalledWith(
        expect.stringContaining(
          'Storage set error: TypeError: Converting circular structure to JSON'
        )
      );

      consoleSpy.mockRestore();
    });
  });

  describe('sessionStorage', () => {
    it('should store session data correctly', () => {
      const sessionData = { sessionId: 'abc123' };
      service.setSession('session', sessionData);

      const stored = sessionStorage.getItem('moon_session');
      expect(stored).toBe(JSON.stringify(sessionData));
    });

    it('should retrieve session data', () => {
      const sessionData = { sessionId: 'abc123' };
      sessionStorage.setItem('moon_session', JSON.stringify(sessionData));

      const result = service.getSession('session');
      expect(result).toEqual(sessionData);
    });

    it('should delete session data', () => {
      sessionStorage.setItem('moon_temp', 'value');

      service.removeSession('temp');
      expect(sessionStorage.getItem('moon_temp')).toBeNull();
    });
  });

  describe('key prefixing', () => {
    it('should prepend prefix to all keys', () => {
      service.set('key', 'value');
      expect(localStorage.getItem('moon_key')).toBe('"value"');
      expect(localStorage.getItem('key')).toBeNull();
    });
  });
});