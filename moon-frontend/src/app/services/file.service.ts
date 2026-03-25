import { Injectable } from '@angular/core';
import { Apollo } from 'apollo-angular';
import { Observable } from 'rxjs';
import { GET_FILES, GET_FILE_BY_ID, CREATE_FILE, UPDATE_FILE, DELETE_FILE } from './file.graphql';

@Injectable({
  providedIn: 'root'
})
export class FileService {

  constructor(private apollo: Apollo) { }

  // 获取文件列表
  getList(where?: any, orderBy?: any, first?: number, after?: string): Observable<any> {
    return this.apollo.query({
      query: GET_FILES,
      variables: { where, orderBy, first, after }
    });
  }

  // 根据ID获取文件
  getById(id: string): Observable<any> {
    return this.apollo.query({
      query: GET_FILE_BY_ID,
      variables: { id }
    });
  }

  // 创建文件
  create(input: {
    fileName: string;
    fileType: string;
    fileSize: number;
  }): Observable<any> {
    return this.apollo.mutate({
      mutation: CREATE_FILE,
      variables: { input }
    });
  }

  // 更新文件
  update(id: string, input: {
    fileName?: string;
    fileType?: string;
    fileSize?: number;
  }): Observable<any> {
    return this.apollo.mutate({
      mutation: UPDATE_FILE,
      variables: { id, input }
    });
  }

  // 删除文件
  delete(id: string): Observable<any> {
    return this.apollo.mutate({
      mutation: DELETE_FILE,
      variables: { id }
    });
  }
}