import { HttpErrorResponse, HttpInterceptorFn, HttpRequest, HttpResponse } from '@angular/common/http';
import { inject } from '@angular/core';
import { AuthService } from '../services/auth.service';
import { catchError, tap } from 'rxjs';
import { Router } from '@angular/router';

export const authInterceptor: HttpInterceptorFn = (req, next) => {
  const authService = inject(AuthService)
  return next(req)
};

function handleHttpError(
  error:HttpErrorResponse,
  router:Router
):void{
  switch(error.status){
    case 401:
      handleUnauthorized(router)
      break;
  }
}

function handleUnauthorized(router :Router):void{
  router.navigate(['/auth/login'])
}
