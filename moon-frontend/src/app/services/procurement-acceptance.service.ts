import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_PROCUREMENT_ACCEPTANCES, GET_PROCUREMENT_ACCEPTANCE_BY_ID, CREATE_PROCUREMENT_ACCEPTANCE, UPDATE_PROCUREMENT_ACCEPTANCE, DELETE_PROCUREMENT_ACCEPTANCE } from './procurement-acceptance.graphql';

@Injectable({
  providedIn: 'root'
})
export class ProcurementAcceptanceService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_ACCEPTANCES,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.procurementAcceptances));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_ACCEPTANCE_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_PROCUREMENT_ACCEPTANCE,
      variables: { input }
    }).pipe(map((res: any) => res.data.createProcurementAcceptance));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_PROCUREMENT_ACCEPTANCE,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateProcurementAcceptance));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_PROCUREMENT_ACCEPTANCE,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteProcurementAcceptance));
  }
}
