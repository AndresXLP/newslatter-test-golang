package service

import (
	"context"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
)

func (s *service) Create(
	ctx context.Context,
	interest []newsletter.Interest,
) error {
	if err := s.repo.Create(ctx, interest); err != nil {
		return err
	}
	return nil
}
