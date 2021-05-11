package compressor

import "github.com/yigitsadic/hsedocumentgenerator/internal/models"

type ZipWriter interface {
	WriteAsZip(fileName string, files []models.ReadFile) error
}
