import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import {
  ModelConfig,
  ModelConfigCreateRequest,
  ModelConfigUpdateRequest,
  ModelConfigDeleteRequest,
  ModelConfigListResponse,
  ModelConfigGetResponse,
  PromptTemplate,
  PromptTemplateCreateRequest,
  PromptTemplateUpdateRequest,
  PromptTemplateDeleteRequest,
  PromptTemplateListResponse,
  PromptTemplateGetResponse,
  APICallLog,
  APICallLogListRequest,
  APICallLogListResponse,
  APICallLogGetResponse,
  AssetClassifyRequest,
  AssetClassification,
  AssetClassificationApproveRequest,
  RiskAssessRequest,
  RiskAssessment,
  RecommendationGenerateRequest,
  SecurityRecommendation,
  RecommendationFeedbackRequest
} from '../../models/ai.model';

@Injectable({
  providedIn: 'root'
})
export class AIService {
  private apiUrl = '/api/v1/ai';

  constructor(private http: HttpClient) { }

  getModelConfigs(): Observable<ModelConfigListResponse> {
    return this.http.get<ModelConfigListResponse>(`${this.apiUrl}/model-config/list`, {
      timeout: 5000
    });
  }

  getModelConfig(configId: string): Observable<ModelConfigGetResponse> {
    return this.http.get<ModelConfigGetResponse>(`${this.apiUrl}/model-config`, {
      params: { config_id: configId },
      timeout: 5000
    });
  }

  getActiveModelConfig(): Observable<ModelConfigGetResponse> {
    return this.http.get<ModelConfigGetResponse>(`${this.apiUrl}/model-config/active`, {
      timeout: 5000
    });
  }

  createModelConfig(config: ModelConfigCreateRequest): Observable<any> {
    return this.http.post(`${this.apiUrl}/model-config`, config, {
      timeout: 5000
    });
  }

  updateModelConfig(config: ModelConfigUpdateRequest): Observable<any> {
    return this.http.put(`${this.apiUrl}/model-config`, config, {
      timeout: 5000
    });
  }

  deleteModelConfig(configId: string): Observable<any> {
    return this.http.delete(`${this.apiUrl}/model-config`, {
      body: { config_id: configId },
      timeout: 5000
    });
  }

  getPromptTemplates(): Observable<PromptTemplateListResponse> {
    return this.http.get<PromptTemplateListResponse>(`${this.apiUrl}/prompt-template/list`, {
      timeout: 5000
    });
  }

  getPromptTemplate(templateId: string): Observable<PromptTemplateGetResponse> {
    return this.http.get<PromptTemplateGetResponse>(`${this.apiUrl}/prompt-template`, {
      params: { template_id: templateId },
      timeout: 5000
    });
  }

  getPromptTemplateByType(templateType: string): Observable<PromptTemplateGetResponse> {
    return this.http.get<PromptTemplateGetResponse>(`${this.apiUrl}/prompt-template/type`, {
      params: { template_type: templateType },
      timeout: 5000
    });
  }

  createPromptTemplate(template: PromptTemplateCreateRequest): Observable<any> {
    return this.http.post(`${this.apiUrl}/prompt-template`, template, {
      timeout: 5000
    });
  }

  updatePromptTemplate(template: PromptTemplateUpdateRequest): Observable<any> {
    return this.http.put(`${this.apiUrl}/prompt-template`, template, {
      timeout: 5000
    });
  }

  deletePromptTemplate(templateId: string): Observable<any> {
    return this.http.delete(`${this.apiUrl}/prompt-template`, {
      body: { template_id: templateId },
      timeout: 5000
    });
  }

  classifyAsset(assetId: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/classify-asset`, { asset_id: assetId }, {
      timeout: 30000
    });
  }

  getClassificationByAssetId(assetId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/classify-asset`, {
      params: { asset_id: assetId },
      timeout: 5000
    });
  }

  approveClassification(request: AssetClassificationApproveRequest): Observable<any> {
    return this.http.put(`${this.apiUrl}/classify-asset/approve`, request, {
      timeout: 5000
    });
  }

  assessRisk(vulnerabilityId: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/assess-risk`, { vulnerability_id: vulnerabilityId }, {
      timeout: 30000
    });
  }

  getAssessmentByVulnerabilityId(vulnerabilityId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/assess-risk`, {
      params: { vulnerability_id: vulnerabilityId },
      timeout: 5000
    });
  }

  generateRecommendation(vulnerabilityId: string): Observable<any> {
    return this.http.post(`${this.apiUrl}/generate-recommendation`, { vulnerability_id: vulnerabilityId }, {
      timeout: 30000
    });
  }

  getRecommendationByVulnerabilityId(vulnerabilityId: string): Observable<any> {
    return this.http.get(`${this.apiUrl}/generate-recommendation`, {
      params: { vulnerability_id: vulnerabilityId },
      timeout: 5000
    });
  }

  submitFeedback(request: RecommendationFeedbackRequest): Observable<any> {
    return this.http.put(`${this.apiUrl}/generate-recommendation/feedback`, request, {
      timeout: 5000
    });
  }

  getAPICallLogs(request: APICallLogListRequest): Observable<APICallLogListResponse> {
    return this.http.get<APICallLogListResponse>(`${this.apiUrl}/logs`, {
      params: {
        limit: request.limit.toString(),
        offset: request.offset.toString(),
        call_type: request.call_type
      },
      timeout: 5000
    });
  }

  getAPICallLog(logId: string): Observable<APICallLogGetResponse> {
    return this.http.get<APICallLogGetResponse>(`${this.apiUrl}/logs/${logId}`, {
      timeout: 5000
    });
  }

  chat(message: string, conversationId?: string, promptTemplateId?: string): Observable<any> {
    return this.http.post<any>(`${this.apiUrl}/chat`, {
      message: message,
      conversation_id: conversationId || null,
      prompt_template_id: promptTemplateId || null
    }, {
      timeout: 60000
    });
  }

  getConversation(conversationId: string): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/chat`, {
      params: { conversation_id: conversationId },
      timeout: 5000
    });
  }

  listConversations(): Observable<any> {
    return this.http.get<any>(`${this.apiUrl}/chat/list`, {
      timeout: 5000
    });
  }

  deleteConversation(conversationId: string): Observable<any> {
    return this.http.delete<any>(`${this.apiUrl}/chat`, {
      body: { conversation_id: conversationId },
      timeout: 5000
    });
  }
}
