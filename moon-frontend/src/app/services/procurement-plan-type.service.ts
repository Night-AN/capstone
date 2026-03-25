import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_PROCUREMENT_PLAN_TYPES, GET_PROCUREMENT_PLAN_TYPE_BY_ID, CREATE_PROCUREMENT_PLAN_TYPE, UPDATE_PROCUREMENT_PLAN_TYPE, DELETE_PROCUREMENT_PLAN_TYPE } from './procurement-plan-type.graphql';

@Injectable({
  providedIn: 'root'
})
export class ProcurementPlanTypeService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_PLAN_TYPES,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.procurementPlanTypes));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_PLAN_TYPE_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_PROCUREMENT_PLAN_TYPE,
      variables: { input }
    }).pipe(map((res: any) => res.data.createProcurementPlanType));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_PROCUREMENT_PLAN_TYPE,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateProcurementPlanType));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_PROCUREMENT_PLAN_TYPE,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteProcurementPlanType));
  }
}
