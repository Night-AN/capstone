import { HttpErrorResponse, HttpInterceptorFn, HttpRequest } from '@angular/common/http';
import { inject } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { catchError, throwError } from 'rxjs';
import { Router } from '@angular/router';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const authService = inject(AuthService);
  const router = inject(Router);
  
  return next(req).pipe(
    catchError((error: HttpErrorResponse) => {
      handleHttpError(error, router);
      return throwError(() => error);
    })
  );
};

function handleHttpError(
  error: HttpErrorResponse,
  router: Router
): void {
  switch (error.status) {
    case 401:
      handleUnauthorized(router);
      break;
    case 403:
      handleForbidden(router);
      break;
    case 404:
      handleNotFound();
      break;
    case 500:
      handleServerError();
      break;
  }
}

function handleUnauthorized(router: Router): void {
  console.error('Unauthorized access, redirecting to login');
  router.navigate(['/auth/login']);
}

function handleForbidden(router: Router): void {
  console.error('Forbidden access');
  // You could redirect to a 403 page here
}

function handleNotFound(): void {
  console.error('Resource not found');
  // You could redirect to a 404 page here
}

function handleServerError(): void {
  console.error('Server error occurred');
  // You could show a generic error message here
}
