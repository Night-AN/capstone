import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_ASSET_TYPES, GET_ASSET_TYPE_BY_ID, CREATE_ASSET_TYPE, UPDATE_ASSET_TYPE, DELETE_ASSET_TYPE } from './asset-type.graphql';

@Injectable({
  providedIn: 'root'
})
export class AssetTypeService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_ASSET_TYPES,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.assetTypes));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_ASSET_TYPE_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_ASSET_TYPE,
      variables: { input }
    }).pipe(map((res: any) => res.data.createAssetType));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_ASSET_TYPE,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateAssetType));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_ASSET_TYPE,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteAssetType));
  }
}
