package tools

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"
	_ "time"

	"github.com/yusuf/mailapp/model"
)

// JSONReader : is a resuable function to help read the information or details submitted and document upload.
func JSONReader(wr http.ResponseWriter, rq *http.Request, subs *model.Subscriber) (*model.Subscriber, error) {
	read := http.MaxBytesReader(wr, rq.Body, int64(1024*1024)*10)
	defer func(io.ReadCloser) {
		err := read.Close()
		if err != nil {
			panic(err)
		}
	}(read)

	err := json.NewDecoder(read).Decode(&subs)
	if err != nil {
		return nil, err
	}
	return subs, nil
}

// JSONWriter: this will help send json response to the client page of this application
func JSONWriter(wr http.ResponseWriter, msg string, statusCode int) error {
	wr.Header().Set("Content-Type", "application/json")
	wr.WriteHeader(statusCode)

	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	_, err = wr.Write(b)
	if err != nil {
		return err
	}
	return nil
}

func ReadForm(wr http.ResponseWriter, rq *http.Request, mail model.MailUpload) (model.MailUpload, error) {
	if err := rq.ParseMultipartForm(10 << 20); err != nil {
		log.Fatal(err)
	}
	form := rq.MultipartForm
	
	mail.DocxName = form.Value["docx_name"][0]
	mail.Date = time.Now()

	file, ok := form.File["docx"]
	if !ok {
		return model.MailUpload{}, fmt.Errorf("unable to get uploaded document")
	}

	if file[0].Filename != "" {
		fileExtension := filepath.Ext(file[0].Filename)
		switch fileExtension {
		case ".doc", "docx", ".txt":
			f, err := file[0].Open()
			if err != nil {
				return model.MailUpload{}, fmt.Errorf("unable to Open uploaded document")
			}
			defer f.Close()

			content, err := ioutil.ReadAll(f)
			if err != nil {
				return model.MailUpload{}, err
			}

			mail.DocxContent = string(content)
		default:
			return model.MailUpload{}, fmt.Errorf("upload document not allow; try .txt .docx or .doc")
		}
	}
	return mail, nil
}

func HTMLRender(wr http.ResponseWriter, rq *http.Request, dt any) error {
	filePath := "./index.html"

	tmp, err := template.ParseFiles(filePath)
	if err != nil {
		return fmt.Errorf("HTMLRender Error: failed to parse file: %v", err)
	}

	err = tmp.Execute(wr, dt)
	if err != nil {
		return fmt.Errorf("HTMLRender Error: failed to execute template: %v", err)
	}

	return nil
}

// func HTMLRender(wr http.ResponseWriter, rq *http.Request, dt any) error {
// 	matchFiles, err := filepath.Glob("./index.html")

// 	if err != nil || err == filepath.ErrBadPattern {
// 		return fmt.Errorf("HTMLRender Error: invalid pattern; cannot find expected file")
// 	}
// 	if len(matchFiles) == 0 {
// 		return fmt.Errorf("HTMLRender Error: no matching file found")
// 	}

// 	if len(matchFiles) > 1 {
// 		return fmt.Errorf("HTMLRender Error: match file is more than one")
// 	}

// 	tmp, err := template.ParseFiles(matchFiles[0])
// 	if err != nil {
// 		return fmt.Errorf("HTMLRender Error: cannot parse file, ${{err}}")
// 	}

// 	buf := new(bytes.Buffer)
// 	err = tmp.Execute(buf, dt)
// 	if err != nil {
// 		return fmt.Errorf("HTMLRender failed to execute templates", err)
// 	}
// 	_, err = buf.WriteTo(buf)
// 	if err != nil {
// 		return fmt.Errorf("HTMLRender failed to Write response", err)
// 	}

// 	return nil
// }
