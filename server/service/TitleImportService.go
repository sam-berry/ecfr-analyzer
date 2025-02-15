package service

import (
	"context"
	"fmt"
	"github.com/sam-berry/ecfr-analyzer/server/dao"
	"github.com/sam-berry/ecfr-analyzer/server/httpclient"
	"io"
)

type TitleImportService struct {
	HttpClient *httpclient.ECFRBulkDataClient
	TitleDAO   *dao.TitleDAO
}

func (s *TitleImportService) ImportTitles(
	ctx context.Context,
) error {
	url := "https://www.govinfo.gov/bulkdata/ECFR/title-5/ECFR-title5.xml"
	resp, err := s.HttpClient.GetXML(ctx, url)
	if err != nil {
		return fmt.Errorf("failed to fetch title, %w", err)
	}

	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read title content, %w", err)
	}

	err = s.TitleDAO.Insert(ctx, "title-5", content)
	if err != nil {
		return fmt.Errorf("failed to insert title, %w", err)
	}

	return nil
}
