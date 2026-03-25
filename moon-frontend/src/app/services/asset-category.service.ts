import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GET_ASSET_CATEGORIES, GET_ASSET_CATEGORY_BY_ID, CREATE_ASSET_CATEGORY, UPDATE_ASSET_CATEGORY, DELETE_ASSET_CATEGORY } from './asset-category.graphql';

@Injectable({
  providedIn: 'root'
})
export class AssetCategoryService {
  constructor(private apollo: Apollo) {}

  getList(where?: any, first?: number, after?: string, before?: string, last?: number): Observable<any> {
    return this.apollo.query({
      query: GET_ASSET_CATEGORIES,
      variables: { where, first, after, before, last }
    }).pipe(map((res: any) => res.data.assetCategories));
  }

  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_ASSET_CATEGORY_BY_ID,
      variables: { id }
    }).pipe(map((res: any) => res.data.node));
  }

  create(input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_ASSET_CATEGORY,
      variables: { input }
    }).pipe(map((res: any) => res.data.createAssetCategory));
  }

  update(id: string, input: any): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_ASSET_CATEGORY,
      variables: { id, input }
    }).pipe(map((res: any) => res.data.updateAssetCategory));
  }

  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_ASSET_CATEGORY,
      variables: { id }
    }).pipe(map((res: any) => res.data.deleteAssetCategory));
  }
}
