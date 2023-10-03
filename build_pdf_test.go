package utilitywkhtmltopdf_test

import (
	"bytes"
	"guzram/utility/wkhtmltopdf/lib/wkhtmltopdf"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
)

func Test_BuildPDF(t *testing.T) {
	tmpl, err := template.ParseFiles("html/sample_content.html")
	assert.Nil(t, err)

	var parser bytes.Buffer

	// If need to replace any content, do here in the map[string]any
	err = tmpl.Execute(&parser, map[string]any{})
	assert.Nil(t, err)

	e := BuildPDF(parser, "test.pdf")
	assert.Nil(t, e)
}

func BuildPDF(parse bytes.Buffer, outputName string) error {
	pdf := wkhtmltopdf.NewPDFPreparer()
	res := wkhtmltopdf.NewPageReader(&parse)
	res.DisableExternalLinks.Set(false)
	res.EnableLocalFileAccess.Set(true)

	// Set Header And Footer if Any
	res.HeaderHTML.Set("file:///html/sample_header.html")
	res.FooterHTML.Set("file:///html/sample_footer.html")

	pdf.AddPage(res)

	// Set PDF Margin
	pdf.MarginLeft.Set(0)
	pdf.MarginBottom.Set(51)
	pdf.MarginRight.Set(0)

	js, err := pdf.ToJSON()

	if err != nil {
		return err
	}

	pdfFromJson, err1 := wkhtmltopdf.NewPDFGeneratorFromJSON(bytes.NewReader(js))
	if err1 != nil {
		return err
	}

	err = pdfFromJson.Create()
	if err != nil {
		return err
	}

	pdfFromJson.WriteFile(outputName)

	return nil
}
