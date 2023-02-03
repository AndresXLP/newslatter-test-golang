package repository

import (
	"context"
	"fmt"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (r *repository) Create(ctx context.Context, interest []newsletter.Interest) error {
	subs := &newsletter.Subs

	userID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	blogID, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	newSubs := newsletter.Subscription{
		UserID:    userID,
		BlogID:    blogID,
		Interests: interest,
	}

	fmt.Println(newSubs)

	*subs = append(*subs, &newSubs)
	return nil
}
