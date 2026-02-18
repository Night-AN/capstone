export interface Organization {
  organization_id: string;
  name: string;
  description: string;
  created_at: string;
  updated_at: string;
}

export interface OrganizationListItem {
  organization_id: string;
  name: string;
  description: string;
  created_at: string;
}

export interface OrganizationCreateRequest {
  name: string;
  description: string;
}

export interface OrganizationUpdateRequest {
  organization_id: string;
  name: string;
  description: string;
}
