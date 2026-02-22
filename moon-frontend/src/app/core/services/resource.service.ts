import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, catchError, map, throwError } from 'rxjs';
import { Resource, ResourceListItem } from '@models/resource.model';

interface ApiResponse<T> {
  code: string;
  message: string;
  data: T;
}

@Injectable({
  providedIn: 'root'
})
export class ResourceService {
  private apiUrl = '/api/v1';

  constructor(private http: HttpClient) { }

  private handleError(error: HttpErrorResponse): Observable<never> {
    console.error('API Error:', error);
    let errorMessage = 'An unknown error occurred';
    if (error.error instanceof ErrorEvent) {
      // Client-side error
      errorMessage = `Error: ${error.error.message}`;
    } else {
      // Server-side error
      errorMessage = `Error Code: ${error.status}\nMessage: ${error.message}`;
    }
    return throwError(() => new Error(errorMessage));
  }

  // 获取资源列表
  getResources(): Observable<ResourceListItem[]> {
    console.log('Fetching resources from API:', `${this.apiUrl}/resources/list`);
    return this.http.get<any>(`${this.apiUrl}/resources/list`, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Resource API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const resources = response.data?.resources || [];
        console.log('Raw Resources:', resources);
        const transformedResources = resources.map((r: any) => ({
          resource_id: r.resource_id,
          name: r.resource_name,
          description: r.resource_code, // 使用resource_code作为description
          sensitive_flag: r.sensitive_flag || false, // 添加敏感标志
          created_at: new Date().toISOString() // 使用当前时间作为created_at
        }));
        console.log('Transformed Resources:', transformedResources);
        return transformedResources;
      })
    );
  }

  // 根据ID获取资源详情
  getResourceById(resourceId: string): Observable<Resource> {
    return this.http.get<ApiResponse<any>>(`${this.apiUrl}/resources`, {
      params: { resource_id: resourceId },
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Resource Detail API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          resource_id: data.resource_id,
          name: data.resource_name,
          description: data.resource_code, // 使用resource_code作为description
          sensitive_flag: data.resource_flag === 'true', // 使用resource_flag作为sensitive_flag
          created_at: new Date().toISOString(), // 使用当前时间作为created_at
          updated_at: new Date().toISOString() // 使用当前时间作为updated_at
        };
      })
    );
  }

  // 创建资源
  createResource(resourceData: any): Observable<Resource> {
    return this.http.post<ApiResponse<any>>(`${this.apiUrl}/resources`, resourceData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Create Resource API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          resource_id: data.resource_id,
          name: data.resource_name,
          description: data.resource_code, // 使用resource_code作为description
          sensitive_flag: data.sensitive_flag || false, // 添加敏感标志
          created_at: data.created_at || new Date().toISOString(), // 使用后端返回的created_at或当前时间
          updated_at: data.updated_at || new Date().toISOString() // 使用后端返回的updated_at或当前时间
        };
      })
    );
  }

  // 更新资源
  updateResource(resourceData: any): Observable<Resource> {
    return this.http.put<ApiResponse<any>>(`${this.apiUrl}/resources`, resourceData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Update Resource API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          resource_id: data.resource_id,
          name: data.resource_name,
          description: data.resource_code, // 使用resource_code作为description
          sensitive_flag: data.sensitive_flag || false, // 添加敏感标志
          created_at: data.created_at || new Date().toISOString(), // 使用后端返回的created_at或当前时间
          updated_at: data.updated_at || new Date().toISOString() // 使用后端返回的updated_at或当前时间
        };
      })
    );
  }

  // 删除资源
  deleteResource(resourceId: string): Observable<boolean> {
    return this.http.delete<ApiResponse<any>>(`${this.apiUrl}/resources`, {
      body: { resource_id: resourceId },
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Delete Resource API Response:', response);
        return response.code === '200';
      })
    );
  }
}
