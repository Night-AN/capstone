import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class OrganizationService {
  private apiUrl = '/api/v1';

  constructor(private http: HttpClient) { }

  getOrganizations(): Observable<any> {
    return this.http.get(`${this.apiUrl}/organizations/list`, {
      timeout: 5000 // 设置5秒超时
    });
  }

  getOrganizationById(organizationId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/organizations/${organizationId}`, {
      timeout: 5000 // 设置5秒超时
    });
  }

  createOrganization(organizationData: any): Observable<any> {
    return this.http.post(`${this.apiUrl}/organizations`, organizationData, {
      timeout: 5000 // 设置5秒超时
    });
  }

  updateOrganization(organizationData: any): Observable<any> {
    return this.http.put(`${this.apiUrl}/organizations`, organizationData, {
      timeout: 5000 // 设置5秒超时
    });
  }

  deleteOrganization(organizationData: any): Observable<any> {
    return this.http.delete(`${this.apiUrl}/organizations`, {
      body: organizationData,
      timeout: 5000 // 设置5秒超时
    });
  }
}
