package pdf_generator

import (
	"bytes"
	"fmt"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
	"io"
)

type PDFGenerate interface {
	Build(req *gotenberg.HTMLRequest) ([]byte, error)
	BuildRequest(r models.Record) (*gotenberg.HTMLRequest, error)
}

type PDFGenerator struct {
	Store           *AssetStore
	GotenbergClient gotenberg.Client
}

// Builds PDF and returns as bytes.
func (g *PDFGenerator) Build(req *gotenberg.HTMLRequest) ([]byte, error) {
	res, err := g.GotenbergClient.Post(req)
	if err != nil {
		return nil, err
	}

	all, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return all, nil
}

// Builds request for gotenberg with given record.
func (g *PDFGenerator) BuildRequest(r models.Record) (*gotenberg.HTMLRequest, error) {
	htmlContent := new(bytes.Buffer)
	g.Store.Template.Execute(htmlContent, r)

	qr, err := r.GenerateQRCode()
	if err != nil {
		return nil, err
	}

	index, err := gotenberg.NewDocumentFromString(fmt.Sprintf("%s.pdf", r.UniqueReference), htmlContent.String())
	if err != nil {
		return nil, err
	}

	img1, err := gotenberg.NewDocumentFromBytes("1920px-Emblem_of_Uzbekistan.svg.png", g.Store.StateLogo)
	if err != nil {
		return nil, err
	}

	img2, err := gotenberg.NewDocumentFromBytes("bg.jpg", g.Store.Background)
	if err != nil {
		return nil, err
	}

	img3, err := gotenberg.NewDocumentFromBytes("hse_logo.png", g.Store.HseLogo)
	if err != nil {
		return nil, err
	}

	img4, err := gotenberg.NewDocumentFromBytes("kase.png", g.Store.CompanySignature)
	if err != nil {
		return nil, err
	}

	img5, err := gotenberg.NewDocumentFromBytes("qr_code.png", qr)
	if err != nil {
		return nil, err
	}

	css, err := gotenberg.NewDocumentFromBytes("styles.css", g.Store.Styles)
	if err != nil {
		return nil, err
	}

	req := gotenberg.NewHTMLRequest(index)
	req.Landscape(true)
	req.PaperSize(gotenberg.A4)
	req.Assets(img1, img2, img3, img4, img5, css)
	req.Margins(gotenberg.NoMargins)

	return req, nil
}
