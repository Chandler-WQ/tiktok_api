package excel

import (
	"context"

	"github.com/xuri/excelize/v2"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (cli Client) Create(ctx context.Context, filePath string, info [][]string) error {
	f := excelize.NewFile()
	defer func() {
		_ = f.Close()
	}()
	for i, row := range info {
		for j, colCell := range row {
			cellName, err := excelize.CoordinatesToCellName(j+1, i+1)
			if err != nil {
				return err
			}
			f.SetCellStr("Sheet1", cellName, colCell)
		}
	}
	if err := f.SaveAs(filePath); err != nil {
		return err
	}
	return nil
}
