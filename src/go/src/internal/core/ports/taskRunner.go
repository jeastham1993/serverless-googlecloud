package ports

import (
	"context"
	"gcloud-serverless-gym/internal/core/domain"
)

type TaskRunner interface {
	StartHistoryUpdateFor(ctx context.Context, session domain.Session)
}
