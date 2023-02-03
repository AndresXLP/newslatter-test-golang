package newsletter

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Get(
		ctx context.Context,
		UserID uuid.UUID,
		BlogID uuid.UUID,
		Interests []Interest,
	) (*Result[*Subscription], error)
	Create(
		ctx context.Context,
		interest []Interest,
	) error
}
