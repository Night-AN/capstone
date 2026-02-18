import { Injectable } from '@angular/core';

@Injectable({
  providedIn: 'root',
})
export class StorageService {
  private readonly prefix = 'moon_';

  set(key: string, value: any): void {
    try {
      const serialized = JSON.stringify(value);
      localStorage.setItem(this.prefix + key, serialized);
    } catch (e) {
      console.error(`Storage set error: ${e}`);
    }
  }

  get<T>(key: string): T | null {
    try {
      const serialized = localStorage.getItem(this.prefix + key);
      return serialized ? JSON.parse(serialized) : null;
    } catch (e) {
      console.error('Storage get error:', e);
      return null;
    }
  }

  remove(key: string): void {
    localStorage.removeItem(this.prefix + key);
  }

  clear(): void {
    const keysToRemove: string[] = [];
    for (let i = 0; i < localStorage.length; i++) {
      const key = localStorage.key(i);
      if (key?.startsWith(this.prefix)) {
        keysToRemove.push(key);
      }
    }
    keysToRemove.forEach((key) => localStorage.removeItem(key));
  }

  clearAll():void{
    localStorage.clear()
  }

  setSession(key: string, value: any): void {
    try {
      sessionStorage.setItem(this.prefix + key, JSON.stringify(value));
    } catch (e) {
      console.error('SessionStorage set error:', e);
    }
  }

  getSession<T>(key: string): T | null {
    try {
      const data = sessionStorage.getItem(this.prefix + key);
      return data ? JSON.parse(data) : null;
    } catch (e) {
      console.error('SessionStorage get error:', e);
      return null;
    }
  }

  removeSession(key: string): void {
    sessionStorage.removeItem(this.prefix + key);
  }
}
