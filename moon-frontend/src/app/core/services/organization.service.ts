import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, catchError, map, throwError } from 'rxjs';
import { Organization, OrganizationListItem } from '@models/organization.model';

interface ApiResponse<T> {
  code: string;
  message: string;
  data: T;
}

@Injectable({
  providedIn: 'root'
})
export class OrganizationService {
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

  getOrganizations(): Observable<OrganizationListItem[]> {
    console.log('Fetching organizations from API:', `${this.apiUrl}/organizations/list`);
    return this.http.get<any>(`${this.apiUrl}/organizations/list`, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Organizations API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const organizations = response.data?.Organizations || response.data?.organizations || [];
        console.log('Raw Organizations:', organizations);
        const transformedOrganizations = organizations.map((org: any) => ({
          organization_id: org.OrganizationID || org.organization_id,
          organization_name: org.OrganizationName || org.organization_name,
          organization_code: org.OrganizationCode || org.organization_code,
          organization_flag: org.OrganizationFlag || org.organization_flag,
          created_at: org.CreatedAt || org.created_at
        }));
        console.log('Transformed Organizations:', transformedOrganizations);
        return transformedOrganizations;
      })
    );
  }

  getOrganizationById(organizationId: string): Observable<Organization> {
    return this.http.get<ApiResponse<any>>(`${this.apiUrl}/organizations`, {
      params: { organization_id: organizationId },
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Organization Detail API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          organization_id: data.OrganizationID || data.organization_id,
          organization_name: data.OrganizationName || data.organization_name,
          organization_code: data.OrganizationCode || data.organization_code,
          organization_description: data.OrganizationDescription || data.organization_description || '',
          organization_flag: data.OrganizationFlag || data.organization_flag,
          created_at: data.CreatedAt || data.created_at || new Date().toISOString(),
          updated_at: data.UpdatedAt || data.updated_at || new Date().toISOString()
        };
      })
    );
  }

  createOrganization(organizationData: any): Observable<Organization> {
    // 转换请求数据格式为后端期望的格式
    const requestData = {
      OrganizationName: organizationData.organization_name,
      OrganizationCode: organizationData.organization_code,
      OrganizationDescription: organizationData.organization_description,
      OrganizationFlag: organizationData.organization_flag,
      ParentID: organizationData.parent_id
    };

    return this.http.post<ApiResponse<any>>(`${this.apiUrl}/organizations`, requestData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Create Organization API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          organization_id: data.organization_id || data.OrganizationID,
          organization_name: data.organization_name || data.OrganizationName,
          organization_code: data.organization_code || data.OrganizationCode,
          organization_description: data.organization_description || data.OrganizationDescription || '',
          organization_flag: data.organization_flag || data.OrganizationFlag,
          created_at: data.created_at || data.CreatedAt || new Date().toISOString(),
          updated_at: data.updated_at || data.UpdatedAt || new Date().toISOString()
        };
      })
    );
  }

  updateOrganization(organizationData: any): Observable<Organization> {
    // 转换请求数据格式为后端期望的格式
    const requestData = {
      OrganizationID: organizationData.organization_id,
      OrganizationName: organizationData.organization_name,
      OrganizationCode: organizationData.organization_code,
      OrganizationDescription: organizationData.organization_description,
      OrganizationFlag: organizationData.organization_flag,
      ParentID: organizationData.parent_id
    };

    return this.http.put<ApiResponse<any>>(`${this.apiUrl}/organizations`, requestData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Update Organization API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          organization_id: data.organization_id || data.OrganizationID,
          organization_name: data.organization_name || data.OrganizationName,
          organization_code: data.organization_code || data.OrganizationCode,
          organization_description: data.organization_description || data.OrganizationDescription || '',
          organization_flag: data.organization_flag || data.OrganizationFlag,
          created_at: data.created_at || data.CreatedAt || new Date().toISOString(),
          updated_at: data.updated_at || data.UpdatedAt || new Date().toISOString()
        };
      })
    );
  }

  deleteOrganization(organizationData: any): Observable<boolean> {
    // 转换请求数据格式为后端期望的格式
    const requestData = {
      OrganizationID: organizationData.organization_id
    };

    return this.http.delete<ApiResponse<any>>(`${this.apiUrl}/organizations`, {
      body: requestData,
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Delete Organization API Response:', response);
        return response.code === '200';
      })
    );
  }

  /**
   * 分配角色给组织
   * @param organizationId 组织ID
   * @param roleId 角色ID
   * @returns 分配结果
   */
  assignRoleToOrganization(organizationId: string, roleId: string): Observable<any> {
    // 转换请求数据格式为后端期望的格式
    const requestData = {
      OrganizationID: organizationId,
      RoleID: roleId
    };

    return this.http.post(`${this.apiUrl}/organizations/assign-role`, requestData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError)
    );
  }

  /**
   * 从组织移除角色
   * @param organizationId 组织ID
   * @param roleId 角色ID
   * @returns 移除结果
   */
  removeRoleFromOrganization(organizationId: string, roleId: string): Observable<any> {
    // 转换请求数据格式为后端期望的格式
    const requestData = {
      OrganizationID: organizationId,
      RoleID: roleId
    };

    return this.http.post(`${this.apiUrl}/organizations/remove-role`, requestData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError)
    );
  }

  /**
   * 获取组织的角色列表
   * @param organizationId 组织ID
   * @returns 组织角色列表
   */
  getOrganizationRoles(organizationId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/organizations/roles`, {
      params: { organization_id: organizationId },
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError)
    );
  }

  /**
   * 移动组织
   * @param organizationId 组织ID
   * @param newParentId 新父组织ID（可选）
   * @returns 移动结果
   */
  moveOrganization(organizationId: string, newParentId?: string): Observable<any> {
    // 转换请求数据格式为后端期望的格式
    const request = {
      OrganizationID: organizationId,
      NewParentID: newParentId || null
    };
    return this.http.post(`${this.apiUrl}/organizations/move`, request, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError)
    );
  }

  /**
   * 获取组织的用户列表
   * @param organizationId 组织ID
   * @returns 组织用户列表
   */
  getOrganizationUsers(organizationId: string): Observable<any[]> {
    return this.http.get<any>(`${this.apiUrl}/organizations/users`, {
      params: { organization_id: organizationId },
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Organization Users API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const users = response.data || [];
        console.log('Raw Organization Users:', users);
        const transformedUsers = users.map((user: any) => ({
          user_id: user.UserID || user.user_id,
          nickname: user.Nickname || user.nickname,
          full_name: user.FullName || user.full_name,
          email: user.Email || user.email
        }));
        console.log('Transformed Organization Users:', transformedUsers);
        return transformedUsers;
      })
    );
  }

  /**
   * 获取组织树
   * @param rootOrganizationCode 根组织代码
   * @returns 组织树
   */
  getOrganizationTree(rootOrganizationCode: string = ''): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/organizations/tree`, {
      params: { root_organization_code: rootOrganizationCode },
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Organization Tree API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        return response.data;
      })
    );
  }
}
