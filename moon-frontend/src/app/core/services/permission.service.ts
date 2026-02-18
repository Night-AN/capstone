import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, catchError, map, throwError } from 'rxjs';
import { Permission, PermissionListItem, PermissionCreateRequest, PermissionUpdateRequest } from '@models/permission.model';

// API响应接口
interface ApiResponse<T> {
  code: string;
  message: string;
  data: T;
}

// 权限列表响应接口
interface PermissionListResponse {
  Permissions: {
    PermissionID: string;
    PermissionName: string;
    PermissionCode: string;
    SensitiveFlag: boolean;
  }[];
}

@Injectable({
  providedIn: 'root'
})
export class PermissionService {
  private apiUrl = '/api/v1';

  constructor(private http: HttpClient) { }

  // 获取权限列表
  getPermissions(): Observable<PermissionListItem[]> {
    return this.http.get<any>(`${this.apiUrl}/permissions/list`, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Permission API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const permissions = response.data?.permissions || [];
        console.log('Raw Permissions:', permissions);
        const transformedPermissions = permissions.map((p: any) => ({
          permission_id: p.permission_id,
          name: p.permission_name,
          description: p.permission_code, // 使用permission_code作为description
          sensitive_flag: p.sensitive_flag || false, // 添加敏感标志
          created_at: new Date().toISOString() // 使用当前时间作为created_at
        }));
        console.log('Transformed Permissions:', transformedPermissions);
        return transformedPermissions;
      })
    );
  }

  // 根据ID获取权限详情
  getPermissionById(permissionId: string): Observable<Permission> {
    return this.http.get<ApiResponse<any>>(`${this.apiUrl}/permissions/${permissionId}`, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Permission Detail API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          permission_id: data.permission_id,
          name: data.permission_name,
          description: data.permission_code, // 使用permission_code作为description
          sensitive_flag: data.sensitive_flag || false, // 添加敏感标志
          created_at: data.created_at || new Date().toISOString(), // 使用后端返回的created_at或当前时间
          updated_at: data.updated_at || new Date().toISOString() // 使用后端返回的updated_at或当前时间
        };
      })
    );
  }

  // 创建权限
  createPermission(permissionData: PermissionCreateRequest): Observable<Permission> {
    // 转换前端请求数据为后端期望的格式
    const requestData = {
      permission_name: permissionData.name,
      description: permissionData.description,
      permission_code: permissionData.name.toLowerCase().replace(/\s+/g, '_'), // 生成简单的permission_code
      sensitive_flag: false
    };

    return this.http.post<ApiResponse<any>>(`${this.apiUrl}/permissions`, requestData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          permission_id: data.permission_id,
          name: data.permission_name,
          description: data.permission_code, // 使用PermissionCode作为description
          sensitive_flag: data.sensitive_flag || false, // 添加敏感标志
          created_at: new Date().toISOString(), // 使用当前时间作为created_at
          updated_at: new Date().toISOString() // 使用当前时间作为updated_at
        };
      })
    );
  }

  // 更新权限
  updatePermission(permissionData: PermissionUpdateRequest): Observable<Permission> {
    // 转换前端请求数据为后端期望的格式
    const requestData = {
      permission_id: permissionData.permission_id,
      permission_name: permissionData.name,
      description: permissionData.description,
      permission_code: permissionData.name.toLowerCase().replace(/\s+/g, '_'), // 生成简单的permission_code
      sensitive_flag: false
    };

    return this.http.put<ApiResponse<any>>(`${this.apiUrl}/permissions`, requestData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          permission_id: data.permission_id,
          name: data.permission_name,
          description: data.permission_code, // 使用PermissionCode作为description
          sensitive_flag: data.sensitive_flag || false, // 添加敏感标志
          created_at: new Date().toISOString(), // 使用当前时间作为created_at
          updated_at: new Date().toISOString() // 使用当前时间作为updated_at
        };
      })
    );
  }

  // 删除权限
  deletePermission(permissionId: string): Observable<boolean> {
    // 转换前端请求数据为后端期望的格式
    const requestData = {
      permission_id: permissionId
    };

    return this.http.delete<ApiResponse<any>>(`${this.apiUrl}/permissions`, {
      body: requestData,
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        return response.data?.success || false;
      })
    );
  }

  // 错误处理
  private handleError(error: any): Observable<never> {
    let errorMessage = '未知错误';
    // 检查是否是服务器端渲染环境
    const isServer = typeof window === 'undefined';
    
    if (!isServer && error.error && typeof error.error === 'object' && error.error.message) {
      // 客户端错误
      errorMessage = `客户端错误: ${error.error.message}`;
    } else if (error.status) {
      // 服务器错误
      errorMessage = `服务器错误: ${error.status} ${error.message}`;
    } else {
      // 其他错误
      errorMessage = `错误: ${error.message || JSON.stringify(error)}`;
    }
    console.error(errorMessage);
    return throwError(() => new Error(errorMessage));
  }
}
