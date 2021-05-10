package main

import (
	"github.com/yigitsadic/hsedocumentgenerator/internal/handlers"
	"os"
)

func main() {
	h := handlers.NewHandler(os.Stdin, os.Stdout, nil)

	h.Do()
}
