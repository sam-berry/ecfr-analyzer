package service

import (
	"context"
)

type ECFRImportService struct {
}

func (c *ECFRImportService) GetData(
	ctx context.Context,
) (string, error) {
	return "Success!", nil
}
