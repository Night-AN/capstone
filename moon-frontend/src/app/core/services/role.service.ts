import { Injectable, inject } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable, map } from 'rxjs';
import { Role, RoleCreateRequest, RoleUpdateRequest, RoleDeleteRequest, RoleListResponse, RoleGetResponse } from '@models/role.model';

interface ApiResponse<T> {
  code: string;
  message: string;
  data: T;
}

@Injectable({
  providedIn: 'root'
})
export class RoleService {
  private http = inject(HttpClient);
  private apiUrl = '/api/v1/roles';

  getRoles(): Observable<Role[]> {
    return this.http.get<ApiResponse<any>>(`${this.apiUrl}/list`).pipe(
      map(response => {
        // 处理不同大小写格式的响应数据
        const roles = response.data?.Roles || response.data?.roles || [];
        return roles.map((role: any) => ({
          role_id: role.RoleID || role.role_id,
          role_name: role.RoleName || role.role_name,
          description: role.Description || role.description,
          role_code: role.RoleCode || role.role_code,
          role_flag: role.RoleFlag || role.role_flag,
          sensitive_flag: role.SensitiveFlag || role.sensitive_flag,
          created_at: role.CreatedAt || role.created_at,
          updated_at: role.UpdatedAt || role.updated_at
        }));
      })
    );
  }

  getRoleById(id: string): Observable<Role> {
    return this.http.get<ApiResponse<any>>(this.apiUrl, {
      params: { role_id: id }
    }).pipe(
      map(response => ({
        role_id: response.data.RoleID || response.data.role_id,
        role_name: response.data.RoleName || response.data.role_name,
        description: response.data.Description || response.data.description,
        role_code: response.data.RoleCode || response.data.role_code,
        role_flag: response.data.RoleFlag || response.data.role_flag,
        sensitive_flag: response.data.SensitiveFlag || response.data.sensitive_flag,
        created_at: response.data.CreatedAt || response.data.created_at,
        updated_at: response.data.UpdatedAt || response.data.updated_at
      }))
    );
  }

  createRole(role: RoleCreateRequest): Observable<Role> {
    // 转换请求数据格式为后端期望的格式
    const requestData = {
      RoleName: role.role_name,
      Description: role.description,
      RoleCode: role.role_code,
      RoleFlag: role.role_flag,
      SensitiveFlag: role.sensitive_flag
    };

    return this.http.post<ApiResponse<any>>(this.apiUrl, requestData).pipe(
      map(response => ({
        role_id: response.data.RoleID,
        role_name: response.data.RoleName,
        description: response.data.Description,
        role_code: response.data.RoleCode,
        role_flag: response.data.RoleFlag,
        sensitive_flag: response.data.SensitiveFlag,
        created_at: response.data.CreatedAt,
        updated_at: response.data.UpdatedAt
      }))
    );
  }

  updateRole(role: RoleUpdateRequest): Observable<Role> {
    // 转换请求数据格式为后端期望的格式
    const requestData = {
      RoleID: role.role_id,
      RoleName: role.role_name,
      Description: role.description,
      RoleCode: role.role_code,
      RoleFlag: role.role_flag,
      SensitiveFlag: role.sensitive_flag
    };

    return this.http.put<ApiResponse<any>>(this.apiUrl, requestData).pipe(
      map(response => ({
        role_id: response.data.RoleID,
        role_name: response.data.RoleName,
        description: response.data.Description,
        role_code: response.data.RoleCode,
        role_flag: response.data.RoleFlag,
        sensitive_flag: response.data.SensitiveFlag,
        created_at: response.data.CreatedAt,
        updated_at: response.data.UpdatedAt
      }))
    );
  }

  deleteRole(role: RoleDeleteRequest): Observable<boolean> {
    // 转换请求数据格式为后端期望的格式
    const requestData = {
      RoleID: role.role_id
    };

    return this.http.delete<ApiResponse<any>>(this.apiUrl, {
      body: requestData
    }).pipe(
      map(response => response.code === '200')
    );
  }
}