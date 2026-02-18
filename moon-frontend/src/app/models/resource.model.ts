export interface Resource {
  resource_id: string;
  name: string;
  description: string;
  sensitive_flag: boolean;
  created_at: string;
  updated_at: string;
}

export interface ResourceListItem {
  resource_id: string;
  name: string;
  description: string;
  sensitive_flag: boolean;
  created_at: string;
}

export interface ResourceCreateRequest {
  name: string;
  description: string;
}

export interface ResourceUpdateRequest {
  resource_id: string;
  name: string;
  description: string;
}
