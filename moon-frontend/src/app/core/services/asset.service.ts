import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, catchError, map, throwError } from 'rxjs';
import { Asset, AssetListItem } from '@models/asset.model';

interface ApiResponse<T> {
  code: string;
  message: string;
  data: T;
}

@Injectable({
  providedIn: 'root'
})
export class AssetService {
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

  // 获取资产列表
  getAssets(): Observable<AssetListItem[]> {
    console.log('Fetching assets from API:', `${this.apiUrl}/assets/list`);
    return this.http.get<any>(`${this.apiUrl}/assets/list`, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Asset API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const assets = response.data?.assets || [];
        console.log('Raw Assets:', assets);
        const transformedAssets = assets.map((a: any) => ({
          asset_id: a.asset_id,
          asset_name: a.asset_name,
          asset_code: a.asset_code,
          asset_type: a.asset_type,
          asset_class: a.asset_class,
          ip_address: a.ip_address,
          status: a.status,
          created_at: a.created_at || new Date().toISOString()
        }));
        console.log('Transformed Assets:', transformedAssets);
        return transformedAssets;
      })
    );
  }

  // 根据ID获取资产详情
  getAssetById(assetId: string): Observable<Asset> {
    return this.http.get<ApiResponse<any>>(`${this.apiUrl}/assets`, {
      params: { asset_id: assetId },
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Asset Detail API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data || {};
        console.log('Asset Detail Data:', data);
        return {
          asset_id: data.asset_id || '',
          asset_name: data.asset_name || '',
          asset_code: data.asset_code || '',
          asset_description: data.asset_description || '',
          organization_id: data.organization_id || '',
          asset_type: data.asset_type || '',
          asset_class: data.asset_class || '',
          manufacturer: data.manufacturer || '',
          model: data.model || '',
          serial_number: data.serial_number || '',
          ip_address: data.ip_address || '',
          mac_address: data.mac_address || '',
          location: data.location || '',
          department: data.department || '',
          owner: data.owner || '',
          contact_info: data.contact_info || '',
          status: data.status || '',
          purchase_date: data.purchase_date || '',
          warranty_end_date: data.warranty_end_date || '',
          value: data.value || '',
          notes: data.notes || '',
          created_at: data.created_at || new Date().toISOString(),
          updated_at: data.updated_at || new Date().toISOString()
        };
      })
    );
  }

  // 创建资产
  createAsset(assetData: any): Observable<Asset> {
    return this.http.post<ApiResponse<any>>(`${this.apiUrl}/assets`, assetData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Create Asset API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          asset_id: data.asset_id,
          asset_name: data.asset_name,
          asset_code: data.asset_code,
          asset_description: data.asset_description,
          organization_id: data.organization_id || '',
          asset_type: data.asset_type,
          asset_class: data.asset_class,
          manufacturer: data.manufacturer,
          model: data.model,
          serial_number: data.serial_number,
          ip_address: data.ip_address,
          mac_address: data.mac_address,
          location: data.location,
          department: data.department,
          owner: data.owner,
          contact_info: data.contact_info,
          status: data.status,
          purchase_date: data.purchase_date,
          warranty_end_date: data.warranty_end_date,
          value: data.value,
          notes: data.notes,
          created_at: data.created_at || new Date().toISOString(),
          updated_at: data.updated_at || new Date().toISOString()
        };
      })
    );
  }

  // 更新资产
  updateAsset(assetData: any): Observable<Asset> {
    return this.http.put<ApiResponse<any>>(`${this.apiUrl}/assets`, assetData, {
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Update Asset API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const data = response.data;
        return {
          asset_id: data.asset_id,
          asset_name: data.asset_name,
          asset_code: data.asset_code,
          asset_description: data.asset_description,
          organization_id: data.organization_id || '',
          asset_type: data.asset_type,
          asset_class: data.asset_class,
          manufacturer: data.manufacturer,
          model: data.model,
          serial_number: data.serial_number,
          ip_address: data.ip_address,
          mac_address: data.mac_address,
          location: data.location,
          department: data.department,
          owner: data.owner,
          contact_info: data.contact_info,
          status: data.status,
          purchase_date: data.purchase_date,
          warranty_end_date: data.warranty_end_date,
          value: data.value,
          notes: data.notes,
          created_at: data.created_at || new Date().toISOString(),
          updated_at: data.updated_at || new Date().toISOString()
        };
      })
    );
  }

  // 删除资产
  deleteAsset(assetId: string): Observable<boolean> {
    return this.http.delete<ApiResponse<any>>(`${this.apiUrl}/assets`, {
      body: { asset_id: assetId },
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Delete Asset API Response:', response);
        return response.code === '200';
      })
    );
  }

  /**
   * 获取指定组织的资产列表
   * @param organizationId 组织ID
   * @returns 组织资产列表
   */
  getAssetsByOrganization(organizationId: string): Observable<AssetListItem[]> {
    return this.http.get<any>(`${this.apiUrl}/assets/organization`, {
      params: { organization_id: organizationId },
      timeout: 5000 // 设置5秒超时
    }).pipe(
      catchError(this.handleError),
      map((response: any) => {
        console.log('Assets by Organization API Response:', response);
        // 转换后端返回的数据结构为前端期望的格式
        const assets = response.data || [];
        const transformedAssets = assets.map((a: any) => ({
          asset_id: a.asset_id,
          asset_name: a.asset_name,
          asset_code: a.asset_code,
          asset_type: a.asset_type,
          asset_class: a.asset_class,
          ip_address: a.ip_address,
          status: a.status,
          created_at: a.created_at || new Date().toISOString()
        }));
        return transformedAssets;
      })
    );
  }
}
