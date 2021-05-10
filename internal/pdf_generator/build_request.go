package pdf_generator

import (
	"bytes"
	"fmt"
	"github.com/thecodingmachine/gotenberg-go-client/v7"
	"github.com/yigitsadic/hsedocumentgenerator/internal/models"
)

// Builds request for gotenberg with given record.
func BuildRequest(r models.Record, store *AssetStore) (*gotenberg.HTMLRequest, error) {
	htmlContent := new(bytes.Buffer)
	store.Template.Execute(htmlContent, r)

	qr, err := r.GenerateQRCode()
	if err != nil {
		return nil, err
	}

	index, err := gotenberg.NewDocumentFromString(fmt.Sprintf("%s.pdf", r.UniqueReference), htmlContent.String())
	if err != nil {
		return nil, err
	}

	img1, err := gotenberg.NewDocumentFromBytes("1920px-Emblem_of_Uzbekistan.svg.png", store.StateLogo)
	if err != nil {
		return nil, err
	}

	img2, err := gotenberg.NewDocumentFromBytes("bg.jpg", store.Background)
	if err != nil {
		return nil, err
	}

	img3, err := gotenberg.NewDocumentFromBytes("hse_logo.png", store.HseLogo)
	if err != nil {
		return nil, err
	}

	img4, err := gotenberg.NewDocumentFromBytes("kase.png", store.CompanySignature)
	if err != nil {
		return nil, err
	}

	img5, err := gotenberg.NewDocumentFromBytes("qr_code.png", qr)
	if err != nil {
		return nil, err
	}

	css, err := gotenberg.NewDocumentFromBytes("styles.css", store.Styles)
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
