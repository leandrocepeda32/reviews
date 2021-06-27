package rating

import uuid "github.com/satori/go.uuid"

type Rating struct {
	Id uuid.UUID `json:"id"`
	IdArticle string `json:"id_article"`
	ScoreSum float64 `json:"score_sum"`
	ScoreCount int64 `json:"score_count"`
}