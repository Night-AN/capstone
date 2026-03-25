import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_PERMISSIONS, GET_PERMISSION_BY_ID, CREATE_PERMISSION, UPDATE_PERMISSION, DELETE_PERMISSION } from './permission.graphql';

@Injectable({
  providedIn: 'root'
})
export class PermissionService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_PERMISSIONS,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.permissions));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_PERMISSION_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_PERMISSION,
      variables: { input }
    }).pipe(map((res: any) => res.data.createPermission));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_PERMISSION,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updatePermission));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_PERMISSION,
      variables: { id }
    }).pipe(map((res: any) => res.data.deletePermission));
  }
}
