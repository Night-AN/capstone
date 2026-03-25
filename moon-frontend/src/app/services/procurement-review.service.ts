import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_PROCUREMENT_REVIEWS, GET_PROCUREMENT_REVIEW_BY_ID, CREATE_PROCUREMENT_REVIEW, UPDATE_PROCUREMENT_REVIEW, DELETE_PROCUREMENT_REVIEW } from './procurement-review.graphql';

@Injectable({
  providedIn: 'root'
})
export class ProcurementReviewService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_REVIEWS,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.procurementReviews));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_PROCUREMENT_REVIEW_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_PROCUREMENT_REVIEW,
      variables: { input }
    }).pipe(map((res: any) => res.data.createProcurementReview));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_PROCUREMENT_REVIEW,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateProcurementReview));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_PROCUREMENT_REVIEW,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteProcurementReview));
  }
}
