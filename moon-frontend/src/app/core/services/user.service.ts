import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable, map } from 'rxjs';
import {
  User,
  UserCreateRequest,
  UserUpdateRequest,
  UserDeleteRequest,
  UserListResponse
} from '../../models/user.model';

// 通用的API响应接口
export interface ApiResponse<T> {
  code: string;
  message: string;
  data: T;
}

@Injectable({
  providedIn: 'root'
})
export class UserService {
  private apiUrl = '/api/v1';

  constructor(private http: HttpClient) { }

  /**
   * 获取用户列表
   * @returns 用户列表响应
   */
  getUsers(): Observable<any> {
    return this.http.get(`${this.apiUrl}/users/list`, {
      timeout: 5000 // 设置5秒超时
    });
  }

  /**
   * 创建用户
   * @param user 创建用户请求
   * @returns 创建的用户信息
   */
  createUser(user: UserCreateRequest): Observable<any> {
    return this.http.post(`${this.apiUrl}/users`, user, {
      timeout: 5000 // 设置5秒超时
    });
  }

  /**
   * 更新用户
   * @param user 更新用户请求
   * @returns 更新后的用户信息
   */
  updateUser(user: UserUpdateRequest): Observable<any> {
    return this.http.put(`${this.apiUrl}/users`, user, {
      timeout: 5000 // 设置5秒超时
    });
  }

  /**
   * 删除用户
   * @param user 删除用户请求
   * @returns 删除结果
   */
  deleteUser(user: UserDeleteRequest): Observable<any> {
    return this.http.delete(`${this.apiUrl}/users`, {
      body: user,
      timeout: 5000 // 设置5秒超时
    });
  }

  /**
   * 获取用户详情
   * @param userId 用户ID
   * @returns 用户详情响应
   */
  getUserById(userId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/users/${userId}`, {
      timeout: 5000 // 设置5秒超时
    });
  }

  /**
   * 分配角色给用户
   * @param userId 用户ID
   * @param roleId 角色ID
   * @returns 分配结果
   */
  assignRoleToUser(userId: string, roleId: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/users/assign-role`, {
      user_id: userId,
      role_id: roleId
    }, {
      timeout: 5000 // 设置5秒超时
    });
  }

  /**
   * 从用户移除角色
   * @param userId 用户ID
   * @param roleId 角色ID
   * @returns 移除结果
   */
  removeRoleFromUser(userId: string, roleId: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/users/remove-role`, {
      user_id: userId,
      role_id: roleId
    }, {
      timeout: 5000 // 设置5秒超时
    });
  }

  /**
   * 获取用户的角色列表
   * @param userId 用户ID
   * @returns 用户角色列表
   */
  getUserRoles(userId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/users/roles`, {
      params: { user_id: userId },
      timeout: 5000 // 设置5秒超时
    });
  }
}
