package compressor

import (
	"archive/zip"
	"bytes"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
	"io/ioutil"
	"log"
)

type Zipper struct {
}

// Writes given files to buffer and returns bytes and error
func (z Zipper) WriteAsZip(fileName string, files []models.ReadFile) error {
	buf := new(bytes.Buffer)

	w := zip.NewWriter(buf)

	for _, file := range files {
		f, err := w.Create(file.FileName)
		if err != nil {
			log.Println(err)
		}
		_, err = f.Write(file.Content)
		if err != nil {
			log.Println(err)
		}
	}

	err := w.Close()
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, buf.Bytes(), 0644)
}
