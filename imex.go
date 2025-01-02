package imex

import (
	"errors"
	"io"
	"net/http"

	"github.com/ianyulistios/imex/src"
)

type ImexAgent struct {
	FileURL       string
	RawFile       io.ReadCloser
	Header        http.Header
	ContentLength int64
	ErrorData     error
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
	if err != nil {
		i.ErrorData = err
		return i
	}
	if res.StatusCode != 200 {
		i.ErrorData = errors.New("something went wrong with this url: " + i.FileURL + " status code: " + res.Status)
		return i

	}
	i.RawFile = res.Body
	i.Header = res.Header
	i.ContentLength = res.ContentLength
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

	extracted, _, err := src.ReadResponse(i.RawFile)
	if err != nil {
		return dataImage, err
	}

	mimeType := i.Header["Content-Type"]
	if len(mimeType) == 0 {
		imageType = "-"
	} else {
		imageType += "data:" + mimeType[0] + ";base64,"
	}

	for _, mimeData := range customMime {
		if mimeData != "" {
			imageType = "data:image/" + mimeData + ";base64,"
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

func (i *ImexAgent) Close() {
	if i.RawFile != nil {
		i.RawFile.Close()
	}
}
