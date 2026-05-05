package services

import (
	"bytes"

	"github.com/ledongthuc/pdf"
)

func ExtractTextFromPDF(filePath string) (string, error) {
	file, reader, err := pdf.Open(filePath)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	var buffer bytes.Buffer

	totalPages := reader.NumPage()
	for pageIndex := 1; pageIndex <= totalPages; pageIndex++ {
		page := reader.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return "", err
		}

		buffer.WriteString(text)
	}

	return buffer.String(), nil
}
