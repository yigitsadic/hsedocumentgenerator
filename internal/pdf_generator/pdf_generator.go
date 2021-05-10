package pdf_generator

import "github.com/thecodingmachine/gotenberg-go-client/v7"

type PDFGenerate interface {
	Build(req *gotenberg.HTMLRequest) ([]byte, error)
}
