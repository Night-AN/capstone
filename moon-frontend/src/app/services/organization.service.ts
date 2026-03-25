import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_ORGANIZATIONS, GET_ORGANIZATION_BY_ID, CREATE_ORGANIZATION, UPDATE_ORGANIZATION, DELETE_ORGANIZATION } from './organization.graphql';

@Injectable({
  providedIn: 'root'
})
export class OrganizationService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_ORGANIZATIONS,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.organizations));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_ORGANIZATION_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_ORGANIZATION,
      variables: { input }
    }).pipe(map((res: any) => res.data.createOrganization));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_ORGANIZATION,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateOrganization));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_ORGANIZATION,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteOrganization));
  }
}
