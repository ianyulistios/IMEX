# IMEX - File Extractor Simple Library.
![IMEX](https://i.im.ge/2021/08/26/QMnlzm.jpg)

IMEX - File Extractor Simple Library.

The project can handle some of files with several type.

Supported Files Right Now:

Images:
- JPEG -> OUTPUT BASE64
- JPG -> OUTPUT BASE64
- PNG -> OUTPUT BASE64

PDF: Coming Soon

Upload To CDN: Coming Soon

# Installation

```
$ go get github.com/ianyulistios/imex
```
# Why Imex?

With Imex, you can download file by URL:

- **Make Instance**
```
instance := imex.InitImax(dummyURL)
```

- **Download The File**
```
response := instance.DownloadFile()

fmt.Println(response.RawFile)
```
The Raw File that already downloaded can be consumed as ```io.ReadCloser```.

- **Convert to image base64 formated with auto mimes.**
```
image, err := instance.DownloadFile().ToImage()
```

# LICENSE
IMEX is MIT License.