export interface Organization {
  organization_id: string;
  organization_name: string;
  organization_code: string;
  organization_description: string;
  organization_flag: string;
  created_at: string;
  updated_at: string;
}

export interface OrganizationListItem {
  organization_id: string;
  organization_name: string;
  organization_code: string;
  organization_flag: string;
  created_at: string;
}

export interface OrganizationCreateRequest {
  organization_name: string;
  organization_code: string;
  organization_description: string;
  organization_flag: string;
  parent_id?: string | null;
}

export interface OrganizationUpdateRequest {
  organization_id: string;
  organization_name: string;
  organization_code: string;
  organization_description: string;
  organization_flag: string;
  parent_id?: string | null;
}
