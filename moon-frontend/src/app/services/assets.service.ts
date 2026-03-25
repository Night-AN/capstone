import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_ASSETS, GET_ASSET_BY_ID, CREATE_ASSET, UPDATE_ASSET, DELETE_ASSET } from './assets.graphql';

@Injectable({
  providedIn: 'root'
})
export class AssetsService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_ASSETS,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.assets));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_ASSET_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_ASSET,
      variables: { input }
    }).pipe(map((res: any) => res.data.createAsset));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_ASSET,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateAsset));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_ASSET,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteAsset));
  }
}
