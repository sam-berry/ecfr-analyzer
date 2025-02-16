package service

import (
	"context"
	"fmt"
	"github.com/sam-berry/ecfr-analyzer/server/dao"
	"github.com/sam-berry/ecfr-analyzer/server/data"
	"sync"
)

type AgencyResult struct {
	Name   string
	Titles []int
}

var MaxConcurrentAgencyLookups = 10

type WordCountService struct {
	AgencyDAO *dao.AgencyDAO
	TitleDAO  *dao.TitleDAO
}

func (s *WordCountService) CountWordsForAgency(
	ctx context.Context,
	slug string,
) (map[string]any, error) {
	agency, err := s.AgencyDAO.FindFullAgencyBySlug(ctx, slug)
	if err != nil {
		return nil, fmt.Errorf("failed to find agency, %v, %w", slug, err)
	}

	agencyResults := make([]*AgencyResult, len(agency.Children)+1)
	agencyResults[0] = s.buildAgencyResult(agency)
	for i, childAgency := range agency.Children {
		agencyResults[i+1] = s.buildAgencyResult(childAgency)
	}

	var messagesWG sync.WaitGroup

	messages := make(chan string)
	messagesWG.Add(1)
	go func() {
		defer messagesWG.Done()
		for message := range messages {
			logInfo(message)
		}
	}()

	var agencyWg sync.WaitGroup
	var mu sync.Mutex
	var totalWordCount int

	throttle := make(chan struct{}, MaxConcurrentAgencyLookups)

	for _, agencyResult := range agencyResults {
		agencyWg.Add(1)
		throttle <- struct{}{}

		go func(agencyResult *AgencyResult) {
			defer agencyWg.Done()
			defer func() { <-throttle }()

			name := agencyResult.Name
			wordCount, err := s.TitleDAO.CountWords(ctx, name, agencyResult.Titles)
			if err != nil {
				messages <- fmt.Sprintf(
					"failed to count words for agency, %v, %v",
					name,
					err,
				)
				return
			}

			mu.Lock()
			totalWordCount += wordCount
			mu.Unlock()

			messages <- fmt.Sprintf(
				"Counted words for agency: %v, count: %d, total: %d",
				name,
				wordCount,
				totalWordCount,
			)
		}(agencyResult)
	}

	agencyWg.Wait()

	close(messages)
	messagesWG.Wait()

	return map[string]any{
		"wordCount": totalWordCount,
	}, nil
}

func (s *WordCountService) buildAgencyResult(agency *data.Agency) *AgencyResult {
	var titles []int
	for _, ref := range agency.CFRReferences {
		titles = append(titles, ref.Title)
	}
	return &AgencyResult{
		Name:   agency.Name,
		Titles: titles,
	}
}
