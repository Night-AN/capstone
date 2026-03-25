import { afterNextRender, Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { LOGIN, REGISTER } from './auth.graphql';

@Injectable({
  providedIn: 'root'
})
export class AuthService {
  constructor(private apollo: Apollo,) { }

  login(email: string, password: string): Observable<string> {
    return this.apollo.mutate({
      mutation: LOGIN,
      variables: { email, password }
    }).pipe(map((res: any) => {
      const token = res.data.login;
      if (typeof window !== 'undefined') {
        localStorage.setItem('token', token);
      }
      return token;
    }));
  }

  register(email: string, password: string): Observable<string> {
    return this.apollo.mutate({
      mutation: REGISTER,
      variables: { email, password }
    }).pipe(map((res: any) => res.data.register));
  }
}
