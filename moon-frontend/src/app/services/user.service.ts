import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_USERS, GET_USER_BY_ID, CREATE_USER, UPDATE_USER, DELETE_USER } from './user.graphql';

@Injectable({
  providedIn: 'root'
})
export class UserService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_USERS,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.users));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_USER_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_USER,
      variables: { input }
    }).pipe(map((res: any) => res.data.createUser));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_USER,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateUser));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_USER,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteUser));
  }
}
