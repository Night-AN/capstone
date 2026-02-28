package usecase

import (
	"github.com/google/uuid"
)

type RecommendationFeedbackRequest struct {
	RecommendationID string `json:"recommendation_id"`
	IsUseful        bool   `json:"is_useful"`
	Feedback        string `json:"feedback"`
}

type RecommendationFeedbackResponse struct {
	RecommendationID uuid.UUID `json:"recommendation_id"`
	Success          bool      `json:"success"`
}
