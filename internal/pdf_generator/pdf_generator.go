package pdf_generator

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
	"github.com/yigitsadic/hsedocumentgenerator/internal/translations"
	"io"
	"net/http"
	"time"
)

type PDFGenerate interface {
	Build(req *gotenberg.HTMLRequest) ([]byte, error)
	BuildRequest(r models.Record) (*gotenberg.HTMLRequest, error)
	Ping() error
}

type PDFGenerator struct {
	Store           *AssetStore
	GotenbergClient gotenberg.Client
}

// Checks gotenberg is available
func (g *PDFGenerator) Ping() error {
	c := http.Client{Timeout: 3 * time.Second}
	resp, err := c.Get(fmt.Sprintf("%s/ping", g.GotenbergClient.Hostname))
	if err != nil {
		return err
	}

	if resp.StatusCode == 200 {
		return nil
	}

	return errors.New("unable to connect gotenberg")
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

	trans := translations.TranslateTo(r.EducationHours, r.EducationDate, r.Language)
	t := TemplateDto{
		StateName:        trans.StateName,
		CertificateTitle: trans.Title,
		CompanyName:      r.CompanyName,
		FullName:         r.FullName,
		EducationName:    r.EducationName,
		Content:          trans.Content,
		UniqueReference:  r.UniqueReference,
	}

	g.Store.Template.Execute(htmlContent, t)

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

	img2, err := gotenberg.NewDocumentFromBytes("background.jpg", g.Store.Background)
	if err != nil {
		return nil, err
	}

	img3, err := gotenberg.NewDocumentFromBytes("hse_logo.png", g.Store.HseLogo)
	if err != nil {
		return nil, err
	}

	img4, err := gotenberg.NewDocumentFromBytes("sirketkase.png", g.Store.CompanySignature)
	if err != nil {
		return nil, err
	}

	img5, err := gotenberg.NewDocumentFromBytes("qr_code.png", qr)
	if err != nil {
		return nil, err
	}

	css, err := gotenberg.NewDocumentFromBytes("style.css", g.Store.Styles)
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
