export interface Role {
  role_id: string;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
}

export interface RoleListItem {
  role_id: string;
  name: string;
  description: string;
  created_at: string;
}

export interface RoleCreateRequest {
  name: string;
  description: string;
}

export interface RoleUpdateRequest {
  role_id: string;
  name: string;
  description: string;
}
