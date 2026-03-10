export interface ModelConfig {
  config_id: string;
  provider_name: string;
  model_name: string;
  api_key?: string;
  api_endpoint: string;
  api_version: string;
  max_tokens: number;
  temperature: number;
  timeout_seconds: number;
  is_active: boolean;
  priority: number;
  created_at: string;
  updated_at: string;
}

export interface ModelConfigCreateRequest {
  provider_name: string;
  model_name: string;
  api_key: string;
  api_endpoint: string;
  api_version: string;
  max_tokens: number;
  temperature: number;
  timeout_seconds: number;
  is_active: boolean;
  priority: number;
}

export interface ModelConfigUpdateRequest {
  config_id: string;
  provider_name: string;
  model_name: string;
  api_key?: string;
  api_endpoint: string;
  api_version: string;
  max_tokens: number;
  temperature: number;
  timeout_seconds: number;
  is_active: boolean;
  priority: number;
}

export interface ModelConfigDeleteRequest {
  config_id: string;
}

export interface ModelConfigListResponse {
  code: string;
  message: string;
  data: {
    configs: ModelConfig[];
  };
}

export interface ModelConfigGetResponse {
  code: string;
  message: string;
  data: ModelConfig;
}

export interface PromptTemplate {
  template_id: string;
  template_name: string;
  template_type: string;
  template_content: string;
  description: string;
  is_active: boolean;
  created_at: string;
  updated_at: string;
}

export interface PromptTemplateCreateRequest {
  template_name: string;
  template_type: string;
  template_content: string;
  description: string;
  is_active: boolean;
}

export interface PromptTemplateUpdateRequest {
  template_id: string;
  template_name: string;
  template_type: string;
  template_content: string;
  description: string;
  is_active: boolean;
}

export interface PromptTemplateDeleteRequest {
  template_id: string;
}

export interface PromptTemplateListResponse {
  code: string;
  message: string;
  data: {
    templates: PromptTemplate[];
  };
}

export interface PromptTemplateGetResponse {
  code: string;
  message: string;
  data: PromptTemplate;
}

export interface APICallLog {
  log_id: string;
  config_id: string;
  call_type: string;
  prompt_tokens: number;
  completion_tokens: number;
  total_tokens: number;
  request_payload?: string;
  response_payload?: string;
  status_code: number;
  error_message?: string;
  latency_ms: number;
  success: boolean;
  created_at: string;
}

export interface APICallLogListRequest {
  limit: number;
  offset: number;
  call_type: string;
}

export interface APICallLogListResponse {
  code: string;
  message: string;
  data: {
    logs: APICallLog[];
  };
}

export interface APICallLogGetResponse {
  code: string;
  message: string;
  data: APICallLog;
}

export interface AssetClassification {
  classification_id: string;
  asset_id: string;
  classification_type: string;
  confidence: number;
  classification_reason: string;
  is_approved: boolean;
  approved_by?: string;
  approved_at?: string;
  created_at: string;
}

export interface AssetClassifyRequest {
  asset_id: string;
}

export interface AssetClassificationApproveRequest {
  classification_id: string;
  approve: boolean;
}

export interface RiskAssessment {
  assessment_id: string;
  vulnerability_id: string;
  risk_score: number;
  risk_level: string;
  assessment_factors: string;
  created_at: string;
}

export interface RiskAssessRequest {
  vulnerability_id: string;
}

export interface SecurityRecommendation {
  recommendation_id: string;
  vulnerability_id: string;
  log_id: string;
  summary: string;
  vulnerability_analysis: string;
  remediation_steps: string;
  references: string;
  feedback?: string;
  created_at: string;
}

export interface RecommendationGenerateRequest {
  vulnerability_id: string;
}

export interface RecommendationFeedbackRequest {
  recommendation_id: string;
  feedback: string;
}
