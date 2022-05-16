package imex

import (
	"errors"
	"io"
	"net/http"

	"github.com/ianyulistios/imex/src"
)

type ImexAgent struct {
	FileURL   string
	RawFile   io.ReadCloser
	ErrorData error
}

func InitImax(url string) *ImexAgent {
	return &ImexAgent{
		FileURL: url,
	}
}

func (i *ImexAgent) GetData() *ImexAgent {
	return i
}

func (i *ImexAgent) DownloadFile() *ImexAgent {
	res, err := http.Get(i.FileURL)
	if err != nil && res.StatusCode != 200 {
		i.ErrorData = err
		return i
	}
	i.RawFile = res.Body
	return i
}

func (i *ImexAgent) ToImage(customMime ...string) (string, error) {
	var (
		dataImage string
		imageType string
		errorData error
	)

	if i.ErrorData != nil {
		errorData = i.ErrorData
		return dataImage, errorData
	}

	extracted, mimeType, err := src.ReadResponse(i.RawFile)
	if err != nil {
		return dataImage, err
	}

	switch mimeType {
	case "image/jpeg":
		imageType += "data:image/jpeg;base64,"
	case "image/png":
		imageType += "data:image/png;base64,"
	case "image/jpg":
		imageType += "data:image/jpg;base64,"
	default:
		imageType += "-"
	}

	for _, mimeData := range customMime {
		if mimeData != "" {
			imageType = "data:image/" + mimeData + ";base64"
		}
	}

	if imageType == "-" {
		return "", errors.New("invalid mime type")
	}

	dataImage = imageType + src.ToBase64(extracted)
	return dataImage, errorData
}

func (i *ImexAgent) ToByte() ([]byte, string, error) {
	if i.ErrorData != nil {
		return nil, "", i.ErrorData
	}

	extractedDataBytes, mimeType, err := src.ReadResponse(i.RawFile)

	return extractedDataBytes, mimeType, err
}
