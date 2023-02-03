package repository

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	uuid "github.com/google/uuid"
)

func (r *repository) Search(
	ctx context.Context,
	userID uuid.UUID,
	blogID uuid.UUID,
	interests []newsletter.Interest,
	limit int,
	offset int,
) ([]*newsletter.Subscription, error) {
	var dataResponse []*newsletter.Subscription
	dataSubs := &newsletter.Subs
	for _, data := range *dataSubs {
		if userID == data.UserID && blogID == data.BlogID {
			dataResponse = append(dataResponse, data)
		}
	}

	return dataResponse, nil

}
