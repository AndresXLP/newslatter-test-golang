package service

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/google/uuid"
)

func (s *service) Get(
	ctx context.Context,
	UserID uuid.UUID,
	BlogID uuid.UUID,
	Interests []newsletter.Interest,
) (*newsletter.Result[*newsletter.Subscription], error) {
	limit := ctx.Value("limit")
	offset := ctx.Value("offset")
	subs, err := s.repo.Search(ctx, UserID, BlogID, Interests, limit.(int), offset.(int))
	if err != nil {
		return nil, err
	}
	result := newsletter.Result[*newsletter.Subscription]{
		Total: len(subs),
		Pages: 0,
		Page: newsletter.Page[*newsletter.Subscription]{
			Number:   0,
			Elements: subs,
		},
	}
	return &result, nil
}
