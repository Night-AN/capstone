import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_PROCUREMENT_PLANS, GET_PROCUREMENT_PLAN_BY_ID, CREATE_PROCUREMENT_PLAN, UPDATE_PROCUREMENT_PLAN, DELETE_PROCUREMENT_PLAN } from './procurement-plan.graphql';

@Injectable({
  providedIn: 'root'
})
export class ProcurementPlanService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_PLANS,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.procurementPlans));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_PLAN_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_PROCUREMENT_PLAN,
      variables: { input }
    }).pipe(map((res: any) => res.data.createProcurementPlan));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_PROCUREMENT_PLAN,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateProcurementPlan));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_PROCUREMENT_PLAN,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteProcurementPlan));
  }
}
