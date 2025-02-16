package service

import (
	"context"
	"fmt"
	"github.com/sam-berry/ecfr-analyzer/server/dao"
)

type WordCountService struct {
	AgencyDAO *dao.AgencyDAO
	TitleDAO  *dao.TitleDAO
}

func (s *WordCountService) CountWordsForAgency(
	ctx context.Context,
	slug string,
) (map[string]any, error) {
	agency, err := s.AgencyDAO.FindBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("failed to find agency, %v, %w", slug, err)
	}

	wordCount, err := s.TitleDAO.CountWords(ctx, agency.DisplayName)
	if err != nil {
		return nil, fmt.Errorf("failed to count words for agency, %v, %w", agency.DisplayName, err)
	}

	return map[string]any{
		"wordCount": wordCount,
	}, nil
}
