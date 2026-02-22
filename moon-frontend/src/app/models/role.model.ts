export interface Role {
  role_id: string;
  role_name: string;
  description: string;
  sensitive_flag: boolean;
  created_at: string;
  updated_at: string;
}

export interface RoleCreateRequest {
  role_name: string;
  description: string;
  role_code?: string;
  role_flag?: string;
  sensitive_flag: boolean;
}

export interface RoleUpdateRequest {
  role_id: string;
  role_name: string;
  description: string;
  role_code?: string;
  role_flag?: string;
  sensitive_flag: boolean;
}

export interface RoleDeleteRequest {
  role_id: string;
}

export interface RoleListResponse {
  code: string;
  message: string;
  data: {
    roles: RoleListItem[];
  };
}

export interface RoleGetResponse {
  code: string;
  message: string;
  data: {
    role_id: string;
    role_name: string;
    description: string;
  };
}

export interface RoleListItem {
  role_id: string;
  role_name: string;
  description: string;
  role_code?: string;
  role_flag?: string;
  sensitive_flag: boolean;
  created_at: string;
  updated_at: string;
}