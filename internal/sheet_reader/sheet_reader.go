package sheet_reader

import "github.com/yigitsadic/hsedocumentgenerator/internal/models"

type SheetClient interface {
	ReadFromSheets() ([]models.Record, error)
}
