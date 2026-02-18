import { HttpClient } from '@angular/common/http';
import { inject, Injectable } from '@angular/core';
import { StorageService } from './storage.service';
import { Router } from '@angular/router';
import { catchError, Observable, tap, throwError } from 'rxjs';
import { AuthResponse, LoginCredentials, RegisterData } from '../../models/auth.data';

@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly apiURL = '/api/v1';

  private http = inject(HttpClient);
  private route = inject(Router);
  private storage = inject(StorageService);

  public login(credentials: LoginCredentials): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${this.apiURL}/login`, credentials).pipe(
      tap((response) => this.handleAuthSuccess(response)),
      catchError(error=>{
        console.log(error)
        return throwError(() => error);
      })
    );
  }

  public register(data: RegisterData): Observable<AuthResponse> {
    return this.http.post<AuthResponse>(`${this.apiURL}/register`, data).pipe(
      tap((response) => this.handleAuthSuccess(response)),
      catchError(error=>{
        console.log(error)
        return throwError(() => error);
      })
    );
  }

  private handleAuthSuccess(response: AuthResponse): void {
    if (response.token) {
      this.storage.set('token', response.token);
    }
  }

  public logout(): void {
    this.storage.remove('token');
    this.route.navigate(['/auth/login']);
  }

  public isAuthenticated(): boolean {
    return this.storage.get('token') !== null;
  }


}
