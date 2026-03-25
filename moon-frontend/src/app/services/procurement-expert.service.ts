import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_PROCUREMENT_EXPERTS, GET_PROCUREMENT_EXPERT_BY_ID, CREATE_PROCUREMENT_EXPERT, UPDATE_PROCUREMENT_EXPERT, DELETE_PROCUREMENT_EXPERT } from './procurement-expert.graphql';

@Injectable({
  providedIn: 'root'
})
export class ProcurementExpertService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_EXPERTS,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.procurementExperts));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_EXPERT_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_PROCUREMENT_EXPERT,
      variables: { input }
    }).pipe(map((res: any) => res.data.createProcurementExpert));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_PROCUREMENT_EXPERT,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateProcurementExpert));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_PROCUREMENT_EXPERT,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteProcurementExpert));
  }
}
