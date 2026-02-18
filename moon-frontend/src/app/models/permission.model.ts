export interface Permission {
  permission_id: string;
  name: string;
  description: string;
  sensitive_flag: boolean;
  created_at: string;
  updated_at: string;
}

export interface PermissionListItem {
  permission_id: string;
  name: string;
  description: string;
  sensitive_flag: boolean;
  created_at: string;
}

export interface PermissionCreateRequest {
  name: string;
  description: string;
}

export interface PermissionUpdateRequest {
  permission_id: string;
  name: string;
  description: string;
}
