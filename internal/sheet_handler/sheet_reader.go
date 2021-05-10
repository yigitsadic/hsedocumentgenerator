package sheet_handler

import "github.com/yigitsadic/hsedocumentgenerator/internal/models"

type SheetClient interface {
	ReadFromSheets() ([]models.Record, error)
	WriteToSheets([]models.Record) error
}
