package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sam-berry/ecfr-analyzer/server/dao"
	"github.com/sam-berry/ecfr-analyzer/server/ecfrdata"
	"github.com/sam-berry/ecfr-analyzer/server/httpclient"
	"io"
	"strconv"
	"strings"
	"sync"
)

type TitleImportService struct {
	HttpClient *httpclient.ECFRBulkDataClient
	TitleDAO   *dao.TitleDAO
}

func (s *TitleImportService) ImportTitles(ctx context.Context, titlesFilter []string) error {
	logInfo("Start")

	allFiles, err := s.getAllFiles(ctx, titlesFilter)
	if err != nil {
		return fmt.Errorf("failed to get all files, %w", err)
	}

	var wg sync.WaitGroup

	messages := make(chan string)
	defer close(messages)
	successes := make(chan int)
	defer close(successes)
	failures := make(chan int)
	defer close(failures)

	for _, file := range allFiles {
		wg.Add(1)
		go s.processTitleFile(
			ctx,
			file,
			&wg,
			messages,
			successes,
			failures,
		)
	}

	go func() {
		for message := range messages {
			logInfo(message)
		}
	}()

	var successTitles []string
	go func() {
		for it := range successes {
			successTitles = append(successTitles, strconv.Itoa(it))
		}
	}()

	var failedTitles []string
	go func() {
		for it := range failures {
			failedTitles = append(failedTitles, strconv.Itoa(it))
		}
	}()

	wg.Wait()

	logInfo(fmt.Sprintf("Successfully imported titles: %v", strings.Join(successTitles, ", ")))
	logInfo(fmt.Sprintf("Failed to import titles: %v", strings.Join(failedTitles, ", ")))
	logInfo("Complete")
	return nil
}

func (s *TitleImportService) processTitleFile(
	ctx context.Context,
	file ecfrdata.AllFilesItem,
	wg *sync.WaitGroup,
	messages chan<- string,
	successes chan<- int,
	failures chan<- int,
) {
	defer wg.Done()

	titleNumber := file.CFRTitle
	name := fmt.Sprintf("Title %d", titleNumber)

	messages <- fmt.Sprintf("Fetching: %v", name)

	titleFile, err := s.getTitleFile(ctx, file.Link)
	if err != nil {
		messages <- fmt.Sprintf("failed to get title file, %v, %v", name, err)
		failures <- titleNumber
		return
	}

	messages <- fmt.Sprintf("Downloading: %v", name)

	err = s.downloadTitleFile(ctx, name, titleFile.Link)
	if err != nil {
		messages <- fmt.Sprintf("failed to download title file, %v, %v", name, err)
		failures <- titleNumber
		return
	}

	messages <- fmt.Sprintf("Success: %v", name)
	successes <- titleNumber
}

func (s *TitleImportService) getAllFiles(
	ctx context.Context,
	titlesFilter []string,
) ([]ecfrdata.AllFilesItem, error) {
	allFiles, err := s.HttpClient.GetAllFiles(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch all files, %w", err)
	}

	defer allFiles.Body.Close()
	var allFilesResp ecfrdata.AllFilesResponse
	decoder := json.NewDecoder(allFiles.Body)
	if err := decoder.Decode(&allFilesResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal all files response, %w", err)
	}

	filterMap := make(map[string]bool, len(titlesFilter))
	for _, title := range titlesFilter {
		filterMap[title] = true
	}
	hasFilter := len(titlesFilter) > 0

	var finalFiles []ecfrdata.AllFilesItem
	for _, file := range allFilesResp.Files {
		if file.CFRTitle > 0 && (!hasFilter || filterMap[fmt.Sprintf("%d", file.CFRTitle)]) {
			finalFiles = append(finalFiles, file)
		}
	}

	return finalFiles, nil
}

func (s *TitleImportService) getTitleFile(ctx context.Context, url string) (
	*ecfrdata.TitleFileItem,
	error,
) {
	titleFiles, err := s.HttpClient.GetJSON(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch files for title, %v, %w", url, err)
	}

	defer titleFiles.Body.Close()
	var titleFilesResp ecfrdata.TitleFilesResponse
	decoder := json.NewDecoder(titleFiles.Body)
	if err := decoder.Decode(&titleFilesResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal title files response, %w", err)
	}

	var titleXML *ecfrdata.TitleFileItem
	for _, titleFile := range titleFilesResp.Files {
		if titleFile.FileExtension == "xml" {
			titleXML = &titleFile
			break
		}
	}

	return titleXML, nil
}

func (s *TitleImportService) downloadTitleFile(
	ctx context.Context,
	name string,
	url string,
) error {
	resp, err := s.HttpClient.GetXML(ctx, url)
	if err != nil {
		return fmt.Errorf("failed to fetch title XML, %v, %w", url, err)
	}

	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read title content, %w", err)
	}

	err = s.TitleDAO.Insert(ctx, name, content)
	if err != nil {
		return fmt.Errorf("failed to insert title, %w", err)
	}

	return nil
}

func logInfo(message string) {
	log.Info(fmt.Sprintf("Title Import: %v", message))
}
