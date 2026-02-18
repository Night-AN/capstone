import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class RoleService {
  private apiUrl = '/api/v1';

  constructor(private http: HttpClient) { }

  getRoles(): Observable<any> {
    return this.http.get(`${this.apiUrl}/roles/list`, {
      timeout: 5000 // 设置5秒超时
    });
  }

  getRoleById(roleId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/roles/${roleId}`, {
      timeout: 5000 // 设置5秒超时
    });
  }

  createRole(roleData: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/roles`, roleData, {
      timeout: 5000 // 设置5秒超时
    });
  }

  updateRole(roleData: any): Observable<any> {
    return this.http.put(`${this.apiUrl}/roles`, roleData, {
      timeout: 5000 // 设置5秒超时
    });
  }

  deleteRole(roleData: any): Observable<any> {
    return this.http.delete(`${this.apiUrl}/roles`, {
      body: roleData,
      timeout: 5000 // 设置5秒超时
    });
  }
}
