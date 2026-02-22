export interface Asset {
  asset_id: string;
  asset_name: string;
  asset_code: string;
  asset_description: string;
  organization_id: string;
  asset_type: string;
  asset_class: string;
  manufacturer: string;
  model: string;
  serial_number: string;
  ip_address: string;
  mac_address: string;
  location: string;
  department: string;
  owner: string;
  contact_info: string;
  status: string;
  purchase_date: string;
  warranty_end_date: string;
  value: string;
  notes: string;
  created_at: string;
  updated_at: string;
}

export interface AssetListItem {
  asset_id: string;
  asset_name: string;
  asset_code: string;
  organization_id: string;
  asset_type: string;
  asset_class: string;
  ip_address: string;
  status: string;
  created_at: string;
}

export interface AssetCreateRequest {
  asset_name: string;
  asset_code: string;
  asset_description: string;
  organization_id: string;
  asset_type: string;
  asset_class: string;
  manufacturer: string;
  model: string;
  serial_number: string;
  ip_address: string;
  mac_address: string;
  location: string;
  department: string;
  owner: string;
  contact_info: string;
  status: string;
  purchase_date: string;
  warranty_end_date: string;
  value: string;
  notes: string;
}

export interface AssetUpdateRequest {
  asset_id: string;
  asset_name: string;
  asset_code: string;
  asset_description: string;
  organization_id: string;
  asset_type: string;
  asset_class: string;
  manufacturer: string;
  model: string;
  serial_number: string;
  ip_address: string;
  mac_address: string;
  location: string;
  department: string;
  owner: string;
  contact_info: string;
  status: string;
  purchase_date: string;
  warranty_end_date: string;
  value: string;
  notes: string;
}
