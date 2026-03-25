import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_ROLES, GET_ROLE_BY_ID, CREATE_ROLE, UPDATE_ROLE, DELETE_ROLE } from './role.graphql';

@Injectable({
  providedIn: 'root'
})
export class RoleService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_ROLES,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.roles));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_ROLE_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_ROLE,
      variables: { input }
    }).pipe(map((res: any) => res.data.createRole));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_ROLE,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateRole));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_ROLE,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteRole));
  }
}
